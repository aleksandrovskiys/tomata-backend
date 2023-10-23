package googleoauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"tomata-backend/authentication"
	"tomata-backend/interfaces"
)

func GetIDToken(code string, redirectURI string) (string, error) {
	googleAuthURI := os.Getenv("GOOGLE_AUTH_URL") + "/token"
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OPEN_ID_CLIENT_SECRET")
	fmt.Printf("\nCalling %s\n", googleAuthURI)
	response, err := http.PostForm(googleAuthURI, url.Values{
		"code":          {code},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"redirect_uri":  {redirectURI},
		"grant_type":    {"authorization_code"},
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		respBody := interfaces.GoogleOpenIdTokenReponseSchema{}
		err := json.NewDecoder(response.Body).Decode(&respBody)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return "", err
		}
		return string(respBody.IdToken), nil
	} else {
		body, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return "", err
		}
		fmt.Printf("Error: %v\n", string(body))
		return "", nil
	}

}

func GetUserInfo(tokenData interfaces.GoogleOpenIDParametersSchema, redirectURI string) (interfaces.GoogleLoginDataSchema, error) {

	if !authentication.ValidateAntiForgeryToken(tokenData.State) {
		return interfaces.GoogleLoginDataSchema{}, errors.New("invalid CSRF token")
	}

	idToken, err := GetIDToken(tokenData.Code, redirectURI)

	if err != nil {
		return interfaces.GoogleLoginDataSchema{}, errors.New("invalid code")
	}
	userInfo, err := authentication.GetUserInfoFromToken(idToken)

	if err != nil {
		return interfaces.GoogleLoginDataSchema{}, errors.New("invalid token")
	}

	return userInfo, nil
}
