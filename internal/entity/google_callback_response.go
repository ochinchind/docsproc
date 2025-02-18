package entity

type GoogleCallbackResponse struct {
	Id            string `json:"id"                  example:"105482576905442285932"`
	Email         string `json:"email"               example:"abc@example.com"`
	VerifiedEmail bool   `json:"verified_email"      example:"true"`
	Name          string `json:"name"                example:"John Doe"`
	GivenName     string `json:"given_name"          example:"John"`
	FamilyName    string `json:"family_name"         example:"Doe"`
	Picture       string `json:"picture"             example:"https://example.com/picture.jpg"`
}
