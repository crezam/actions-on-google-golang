package actions_on_google_golang

type ServiceResponse struct {
	Speech string `json:"speech"`
	DisplayText string `json:"displayText"`
	Data struct {
		Google struct {
			ExpectUserResponse bool `json:"expect_user_response"`
			IsSsml bool `json:"is_ssml"`
			PermissionsRequest struct {
				OptContext string `json:"opt_context"`
				Permissions []string `json:"permissions"`
			} `json:"permissions_request"`
		} `json:"google"`
	} `json:"data"`
	ContextOut []struct {
		Name string `json:"name"`
		Lifespan int `json:"lifespan"`
		Parameters struct {
			City string `json:"city"`
		} `json:"parameters"`
	} `json:"contextOut"`

	//TODO
	//add contextOut
}
