package nameapi

/// Hello Models
type HelloRequest struct {
    Url string
    Config Configuration
    Method string
}

type HelloResponse struct {
    ServerName string `json:"serverName,omitempty"`
    MOTD string `json:"motd,omitempty"`
    UserName string `json:"username,omitempty"`
    ServerTime string `json:"serverTime,omitempty"`
}


// Hello Accessors
func (h *HelloRequest) GetConfig() Configuration {
    return h.Config
}
func (h *HelloRequest) GetMethod() string {
    return h.Method
}
func (h *HelloRequest) GetUrl() string {
    return h.Url
}

// Hello methods
func Hello(c Configuration) NameResponse {
    myCall := new(HelloRequest)
    myCall.Url = c.BaseURL + "/hello"
    myCall.Config = c
    myCall.Method = "GET"
    

    return Execute(myCall, nil)
}

