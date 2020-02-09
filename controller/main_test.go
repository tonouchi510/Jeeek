package controller_test

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/tonouchi510/Jeeek/factory"
	"goa.design/goa/v3/middleware"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	jeeek "github.com/tonouchi510/Jeeek/controller"
	"github.com/tonouchi510/Jeeek/gen/activity"
	"github.com/tonouchi510/Jeeek/gen/admin"
	activitysvr "github.com/tonouchi510/Jeeek/gen/http/activity/server"
	adminsvr "github.com/tonouchi510/Jeeek/gen/http/admin/server"
	usersvr "github.com/tonouchi510/Jeeek/gen/http/user/server"
	"github.com/tonouchi510/Jeeek/gen/user"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
)

var testServeMux *http.ServeMux
var fsClient *firestore.Client
var authClient *auth.Client
var testToken string
var adminToken string
var ctx context.Context

var FirebaseCredentials = os.Getenv("FIREBASE_CREDENTIALS")
var TestUserID = os.Getenv("TEST_USER_ID")
var TestAdminPassWord = os.Getenv("ADMIN_PASSWORD")

func TestMain(m *testing.M) {
	// os.Exit は他の defer を気にせずプロセスを殺す
	// 普通に `defer os.Exit(exitStatus)` とするとその時点での exitStatus (= 0) で
	// 実行されるため、無名関数で実行時に取得するようにする
	var exitStatus int
	defer func() {
		os.Exit(exitStatus)
	}()

	projectID := "jeeek-dev"
	ctx = context.Background()

	// Setup logger and external client.
	var (
		logger *log.Logger
		err error
	)
	{
		logger = log.New(os.Stderr, "[Jeeek] ", log.Ltime)
		_, authClient = InitFirebaseAuth(ctx)
		fsClient, err = firestore.NewClient(ctx, projectID)
		if err != nil {
			logger.Fatalf("Failed to create firestore client: %v", err)
		}
		defer fsClient.Close()
	}

	// setup firebase token for auth
	generalSuite := factory.CreateGeneralSuite(ctx, authClient, TestUserID)
	testToken = generalSuite.TestToken
	adminToken = generalSuite.AdminToken

	// Initialize the services.
	var (
		adminSvc admin.Service
		userSvc user.Service
		activitySvc activity.Service
	)
	{
		adminSvc = jeeek.NewAdmin(logger, authClient)
		userSvc = jeeek.NewUser(logger, authClient)
		activitySvc = jeeek.NewActivity(logger, authClient, fsClient)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		adminEndpoints *admin.Endpoints
		userEndpoints *user.Endpoints
		activityEndpoints *activity.Endpoints
	)
	{
		adminEndpoints = admin.NewEndpoints(adminSvc)
		userEndpoints = user.NewEndpoints(userSvc)
		activityEndpoints = activity.NewEndpoints(activitySvc)
	}
	handleHTTPServer(adminEndpoints, userEndpoints, activityEndpoints, logger)

	exitStatus = m.Run()
}

func handleHTTPServer(adminEndpoints *admin.Endpoints, userEndpoints *user.Endpoints, activityEndpoints *activity.Endpoints, logger *log.Logger) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		adminServer *adminsvr.Server
		userServer *usersvr.Server
		activityServer *activitysvr.Server
	)
	{
		eh := errorHandler(logger)
		adminServer = adminsvr.New(adminEndpoints, mux, dec, enc, eh)
		userServer = usersvr.New(userEndpoints, mux, dec, enc, eh)
		activityServer = activitysvr.New(activityEndpoints, mux, dec, enc, eh)
	}
	// Configure the mux.
	adminsvr.Mount(mux, adminServer)
	usersvr.Mount(mux, userServer)
	activitysvr.Mount(mux, activityServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}
	http.Handle("/", handler)
	testServeMux = http.DefaultServeMux
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}

func InitFirebaseAuth(ctx context.Context) (app *firebase.App, client *auth.Client) {
	opt := option.WithCredentialsFile(FirebaseCredentials)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Get an auth client from the firebase.App
	client, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return app, client
}

type requestGenerator func() *http.Request
type testCondition func(rr *httptest.ResponseRecorder, t *testing.T)
type testCase struct {
	description string
	req         requestGenerator
	status      int
	condition   testCondition
}

func runTestCondition(idx int, test testCase, t *testing.T) {
	t.Logf("case %d: %s", idx, test.description)
	req := test.req()
	res := httptest.NewRecorder()
	testServeMux.ServeHTTP(res, req)
	if res.Code != test.status {
		t.Errorf("invalid response status: excepted: %d, got %d", test.status, res.Code)
		t.Logf("%s", res.Body.Bytes())
	}
	test.condition(res, t)
}
