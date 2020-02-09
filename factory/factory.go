package factory

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type GeneralSuite struct {
	TestToken   string
	AdminToken  string
}

func InitializeFirestore(ctx context.Context, client *firestore.Client) error {

	// Firestoreの初期化
	url := "https://asia-northeast1-jeeek-dev.cloudfunctions.net/recursiveDelete"
	_, err := http.Get(url)
	if err != nil {
		return err
	}

	// 開発環境用の初期データ投入
	url = "https://asia-northeast1-jeeek-dev.cloudfunctions.net/initializeFirestore"
	_, err = http.Get(url)
	if err != nil {
		return err
	}

	return nil
}

func CreateGeneralSuite(ctx context.Context, client *auth.Client, testUserID string) *GeneralSuite {
	var err error
	testTk, err := client.CustomToken(ctx, testUserID)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}
	testTk, err = signInWithCustomToken(testTk)

	claims := map[string]interface{}{"admin": true}
	adminTk, err := client.CustomTokenWithClaims(ctx, testUserID, claims)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}
	adminTk, err = signInWithCustomToken(adminTk)

	return &GeneralSuite{
		TestToken: testTk,
		AdminToken: adminTk,
	}
}

func signInWithCustomToken(token string) (string, error) {
	req, err := json.Marshal(map[string]interface{}{
		"token":             token,
		"returnSecureToken": true,
	})
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s",
		os.Getenv("FIREBASE_APIKEY"))
	res, err := http.Post(path, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected http status code: %d", res.StatusCode)
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var resBody struct {
		IDToken string `json:"idToken"`
	}
	if err := json.Unmarshal(resp, &resBody); err != nil {
		return "", err
	}
	return resBody.IDToken, err
}
