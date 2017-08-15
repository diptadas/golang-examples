package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"os/exec"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func getClient(fileName string) (*http.Client, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		return nil, fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}

	tok, err := getTokenFromWeb(config)
	if err != nil {
		return nil, err
	}

	return config.Client(context.Background(), tok), nil
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	err := exec.Command("xdg-open", authURL).Start()
	if err != nil {
		return nil, err
	}

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, fmt.Errorf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve token from web %v", err)
	}
	return tok, nil
}

func main() {
	client, err := getClient("/home/dipta/Downloads/client_secret.json")
	if err != nil {
		panic(err)
	}

	srv, err := sheets.New(client)
	if err != nil {
		panic(fmt.Errorf("Unable to retrieve Sheets Client %v", err))
	}

	spreadsheetId := "1TKojDv8-vNT7AK-re9ShMVh3qauYkuiu_skhyvYqN7o"
	readRange := "Sheet1!B2:B2"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		panic(fmt.Errorf("Unable to retrieve data from sheet. %v", err))
	}

	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			fmt.Println(row)
		}
	} else {
		fmt.Print("No data found.")
	}
}
