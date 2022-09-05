package main

import (
	"net/http"

	"github.com/rabocse/playing/sftool"
)

func main() {

	// Getting the credentials for authentication via enviroment variables.
	salesforceInstance, username, password, clientID, clientSecret, SecurityKey := sftool.EnvHandler()

	// Building Salesforce URL for authentication purposes.
	authURL := sftool.BuildURL(salesforceInstance, 1)

	// Parsing the credentials.
	authPayload := sftool.CraftPayload(username, password, clientID, clientSecret, SecurityKey, "auth")

	// Crafting a valid HTTPS request with TLS ignore for authentication.
	authReq := sftool.CraftRequest(http.MethodPost, authURL, "no-token", authPayload)

	// Sending the request and getting a valid server response for authentication.
	authResponse := sftool.SendRequest(authReq)

	// Extracting the access token value from the server response.
	accessToken := sftool.ExtractAuthToken(authResponse)

	// Building the URL to query the data.
	casesURL := sftool.BuildURL(salesforceInstance, 2)

	// Crafting a valid HTTPS request with TLS ignore.
	casesReq := sftool.CraftRequest(http.MethodGet, casesURL, accessToken, nil)

	// Sending the request and getting a valid server response.
	casesResponse := sftool.SendRequest(casesReq)

	// Parsing the JSON response.
	output := sftool.UnmarshalSF(casesResponse)

	// Printing the relevant info from the response.
	sftool.PrettyPrintBacklog(output)

}
