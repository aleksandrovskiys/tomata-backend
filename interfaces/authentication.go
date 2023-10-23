package interfaces

type LoginDataSchema struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GoogleOpenIDParametersSchema struct {
	State string `json:"state" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

type GoogleOpenIdTokenReponseSchema struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type GoogleLoginDataSchema struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	GoogleID string `json:"google_id" binding:"required"`
}
