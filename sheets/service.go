package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"k8s-enhancements/common"
	"log"
	"os"
	"strings"
)

var sheetSvc *sheets.Service

func GetSheetService() *sheets.Service {
	return sheetSvc
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	_ = json.NewEncoder(f).Encode(token)
}

func CreateSheetServiceWithOAuth() {
	credentialFile := strings.Join([]string{common.GetConfigHome(), "credentials.json"}, string(os.PathSeparator))
	if b, err := ioutil.ReadFile(credentialFile); err != nil {
		panic(err)
	} else {
		if config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsScope); err != nil {
			panic(err)
		} else {
			tokenFile := strings.Join([]string{common.GetConfigHome(), "token.json"}, string(os.PathSeparator))
			if tok, err := tokenFromFile(tokenFile); err != nil {
				tok = getTokenFromWeb(config)
				saveToken(tokenFile, tok)
			} else {
				httpClient := config.Client(context.Background(), tok)
				if srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(httpClient)); err != nil {
					panic(err)
				} else {
					sheetSvc = srv
				}
			}
		}
	}
}

func CreateSheetServiceWithAPIKey(apiKey string) {
	if svc, err := sheets.NewService(context.Background(), option.WithAPIKey(apiKey), option.WithScopes(sheets.SpreadsheetsScope)); err != nil {
		panic(err)
	} else {
		sheetSvc = svc
	}
}
