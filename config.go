package bluemix

import (
	"net/http"
	"time"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/endpoints"
)

//ServiceName ..
type ServiceName string

const (
	//AccountService ...
	AccountService ServiceName = ServiceName("account")
	//AccountServicev1 ...
	AccountServicev1 ServiceName = ServiceName("accountv1")
	//CisService ...
	CisService ServiceName = ServiceName("cis")
	//ContainerService ...
	ContainerService ServiceName = ServiceName("container")
	//RegistryService ...
	ContainerRegistryService ServiceName = ServiceName("container-registry")
	//GlobalSearchService ...
	GlobalSearchService ServiceName = ServiceName("global-search")
	//GlobalTaggingService ...
	GlobalTaggingService ServiceName = ServiceName("global-tagging")
	//IAMService ...
	IAMService ServiceName = ServiceName("iam")
	//IAMPAPService
	IAMPAPService ServiceName = ServiceName("iampap")
	//IAMUUMService ...
	IAMUUMService ServiceName = ServiceName("iamuum")
	//ICDService ...
	ICDService ServiceName = ServiceName("icd")
	//MccpService ...
	MccpService ServiceName = ServiceName("mccp")
	//resourceManagementService
	ResourceManagementService ServiceName = ServiceName("resource-management")
	//resourceControllerService
	ResourceControllerService ServiceName = ServiceName("resource-controller")
	//resourceCatalogService
	ResourceCatalogrService ServiceName = ServiceName("resource-catalog ")
	//UAAService ...
	UAAService ServiceName = ServiceName("uaa")
	//CSEService
	CseService ServiceName = ServiceName("cse")
)

//Config ...
type Config struct {
	IBMID string

	IBMIDPassword string

	BluemixAPIKey string

	IAMAccessToken  string
	IAMRefreshToken string
	UAAAccessToken  string
	UAARefreshToken string

	//Region is optional. If region is not provided then endpoint must be provided
	Region string
	//ResourceGroupID
	ResourceGroup string
	//Endpoint is optional. If endpoint is not provided then endpoint must be obtained from region via EndpointLocator
	Endpoint *string
	//TokenProviderEndpoint is optional. If endpoint is not provided then endpoint must be obtained from region via EndpointLocator
	TokenProviderEndpoint *string
	EndpointLocator       endpoints.EndpointLocator
	MaxRetries            *int
	RetryDelay            *time.Duration

	HTTPTimeout time.Duration

	Debug bool

	HTTPClient *http.Client

	SSLDisable bool
}

//Copy allows the configuration to be overriden or added
//Typically the endpoints etc
func (c *Config) Copy(mccpgs ...*Config) *Config {
	out := new(Config)
	*out = *c
	if len(mccpgs) == 0 {
		return out
	}
	for _, mergeInput := range mccpgs {
		if mergeInput.Endpoint != nil {
			out.Endpoint = mergeInput.Endpoint
		}
	}
	return out
}

//ValidateConfigForService ...
func (c *Config) ValidateConfigForService(svc ServiceName) error {
	if (c.IBMID == "" || c.IBMIDPassword == "") && c.BluemixAPIKey == "" {
		return bmxerror.New(ErrInsufficientCredentials, "Please check the documentation on how to configure the Bluemix credentials")
	}

	if c.Region == "" && (c.Endpoint == nil || *c.Endpoint == "") {
		return bmxerror.New(ErrInvalidConfigurationCode, "Please provide region or endpoint")
	}
	return nil
}
