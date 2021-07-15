// Will authenticate with Twitch via OAuth Client Credentials Flow
// https://dev.twitch.tv/docs/authentication/getting-tokens-oauth/#oauth-client-credentials-flow
// POST https://id.twitch.tv/oauth2/token
//     ?client_id=<your client ID>
//     &client_secret=<your client secret>
//     &grant_type=client_credentials
//     &scope=<space-separated list of scopes>

package main

import (
	"net/http"
)


func requestAuth() {

}

func fetchAuthToken(client_id string, client_secret string, grant_type string, scopes []string) {
	log("fetchAuthToken - fetching oAuth token")
	
	resp, err := http.Get("")
	
	if err != nil {
		// handle error
		log("fetchAuthToken - error response")
	}
	
	defer resp.Body.Close()

	//body, err := io.ReadAll(resp.Body)

}