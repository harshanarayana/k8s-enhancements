package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"k8s-enhancements/models"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

var sheetClient *http.Client
var sheetService *sheets.Service
var homeConfigPath string

const (
	K8SSheetId      = "1E89GNZRwmnfVerOqWqrmzVII8TQlQ0iiI_FgYKxanGs"
	EnhancementsTab = "Enhancements!A11:M"
)

func getClient(config *oauth2.Config) *http.Client {
	tokFile := strings.Join([]string{homeConfigPath, "token.json"}, string(os.PathSeparator))
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

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
	json.NewEncoder(f).Encode(token)
}

func InitSheetClient(configFile string) {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	homeConfigPath = strings.Join([]string{home, ".k8s-enhancements"}, string(os.PathSeparator))
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	sheetClient = getClient(config)

	sheetService, err = sheets.New(sheetClient)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
}

func GetMyAssignments(username string) []*models.EnhancementRow {
	enhancements := make([]*models.EnhancementRow, 0)
	if resp, err := sheetService.Spreadsheets.Values.Get(K8SSheetId, EnhancementsTab).Do(); err != nil {
		panic(err)
	} else {
		if len(resp.Values) < 1 {
			return enhancements
		} else {
			for _, row := range resp.Values {
				if fmt.Sprintf("%s", row[2]) != username {
					continue
				}
				enh := &models.EnhancementRow{}
				for _, index := range []int{0, 1, 3, 4, 5, 6, 7, 8, 9} {
					reflect.ValueOf(enh).Elem().FieldByName(models.AttributeMapper[index]).SetString(fmt.Sprintf("%s", row[index]))
				}
				enhancements = append(enhancements, enh)
			}
		}
	}
	return enhancements
}
