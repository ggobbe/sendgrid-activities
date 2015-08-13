package api

// SGClient will contain the credentials and default values
type SGClient struct {
	apiUser string
	apiKey  string
}

// NewSendGridClient will return a new SGClient. Used for username, password and API URL
func NewSendGridClient(apiUser, apiKey string) *SGClient {
	Client := &SGClient{
		apiUser: apiUser,
		apiKey:  apiKey,
	}

	return Client
}
