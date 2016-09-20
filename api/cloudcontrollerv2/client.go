package cloudcontrollerv2

type Warnings []string

type CloudControllerClient struct {
	cloudControllerURL        string
	cloudControllerAPIVersion string
	authorizationEndpoint     string
	loggregatorEndpoint       string
	dopplerEndpoint           string
	tokenEndpoint             string

	connection        *Connection
	skipSSLValidation bool
}

func NewCloudControllerClient(APIURL string, skipSSLValidation bool) *CloudControllerClient {
	return &CloudControllerClient{
		cloudControllerURL: APIURL,
		skipSSLValidation:  skipSSLValidation,
		connection:         NewConnection(APIURL, skipSSLValidation),
	}
}
