package nameapi
    
/// GetAuthCodeForDomain Models
type GetAuthCodeForDomainRequest struct {
    DomainName string `json:"domainName"`
    Url string
    Config Configuration
    Method string
}

type GetAuthCodeForDomainResponse struct {
    AuthCode string `json:"authCode,omitempty"`
}

// GetAuthCodeForDomain Accessors
func (h *GetAuthCodeForDomainRequest) GetConfig() Configuration {
    return h.Config
}

func (h *GetAuthCodeForDomainRequest) GetMethod() string {
    return h.Method
}

func (h *GetAuthCodeForDomainRequest) GetUrl() string {
    return h.Url
}

// GetAuthCodeForDomain methods
func GetAuthCodeForDomain(c Configuration, domainName string) NameResponse {
    myCall := new(HelloRequest)
    
    myCall.Url = c.BaseURL + "/domains/" + domainName + ":getAuthCode"
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}

