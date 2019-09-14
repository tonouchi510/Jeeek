package main

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"flag"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
	"os/signal"
	"sync"


	jeeek "github.com/tonouchi510/Jeeek/controller"
	admin "github.com/tonouchi510/Jeeek/gen/admin"
	user "github.com/tonouchi510/Jeeek/gen/user"
)

func main() {
	// service.
	var dbgF = flag.Bool("debug", false, "Log request and response bodies")
	flag.Parse()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Setup logger and external client.
	var (
		logger *log.Logger
		authClient *auth.Client
	)
	{
		logger = log.New(os.Stderr, "[takin] ", log.Ltime)
		_, authClient = InitFirebaseAuth(ctx)
	}

	// Initialize the services.
	var (
		adminSvc admin.Service
		userSvc user.Service
	)
	{
		adminSvc = jeeek.NewAdmin(logger, authClient)
		userSvc = jeeek.NewUser(logger, authClient)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		adminEndpoints *admin.Endpoints
		userEndpoints *user.Endpoints
	)
	{
		adminEndpoints = admin.NewEndpoints(adminSvc)
		userEndpoints = user.NewEndpoints(userSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Start the servers and send errors (if any) to the error channel.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	host := ":" + port
	handleHTTPServer(ctx, host, adminEndpoints, userEndpoints, &wg, errc, logger, *dbgF)

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

func InitFirebaseAuth(ctx context.Context) (app *firebase.App, client *auth.Client) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
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
