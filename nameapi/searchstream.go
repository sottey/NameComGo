package nameapi
    
import (
    "encoding/json"
)

/// Search Models are in objects.go

// Search Methods
func SearchStream(c Configuration, keyword string, tldFilter []string, timeout int32, promoCode string) NameResponse {
    myCall := new(SearchRequest)
    myCall.url = c.BaseURL + "/domains:searchStream"
    myCall.config = c
    myCall.method = "POST"
    myCall.Keyword = keyword
    myCall.TLDFilter = tldFilter
    myCall.Timeout = timeout
    myCall.PromoCode = promoCode
    
    body, err := json.Marshal(myCall)
    Check("", err)

    return Execute(myCall, body)
}
