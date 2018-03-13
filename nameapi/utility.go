package nameapi

import (
    "strings"
    "fmt"
    "log"
    "time"
    "bytes"
    "net/http"
    "io/ioutil"
    "math/rand"
)

// Utility Methods
func Execute(myCall IRequest, body []byte) NameResponse {
    client := &http.Client{}
    req, newReqErr := http.NewRequest(myCall.GetMethod(), myCall.GetUrl(), bytes.NewBuffer(body))
    Check("new Request Error: ", newReqErr)
    
    config := myCall.GetConfig()
    
    req.SetBasicAuth(config.Username, config.Token)
    
    if (config.Debug) { fmt.Printf("\n\nRequest:\n", req) }
    
    resp, doErr := client.Do(req)
    Check("client Do Error: ", doErr)
    
    defer resp.Body.Close()
    bodyText, readErr := ioutil.ReadAll(resp.Body)
    Check("Read error: ", readErr)
    
    if (config.Debug) { fmt.Printf("\n\nResult:\n", resp.Body)}
	
	var ret NameResponse
	ret.StatusCode = resp.StatusCode
	ret.Body = string(bodyText)
	return ret
}

// Error verifier
func Check(msg string, err error) {
    if err != nil {
	    fmt.Printf("%s%s\n", msg,err)
        log.Fatal(err)
	}
}

// Format request for printing
func formatRequest(r *http.Request) string {
    // Create return string
    var request []string
    
    // Add the request string
    url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
    request = append(request, url)
    
    // Add the host
    request = append(request, fmt.Sprintf("Host: %v", r.Host))
    
    // Loop through headers
    for name, headers := range r.Header {
        for _, h := range headers {
            request = append(request, fmt.Sprintf("%v: %v", name, h))
        }
    }
    
    // If this is a POST, add post data
    if r.Method == "POST" {
        r.ParseForm()
        request = append(request, "\n")
        request = append(request, r.Form.Encode())
    }
    
    // Return the request as a string
    return strings.Join(request, "\n")
}


const alphabet = "abcdefghijklmnopqrstuvwxyz"
func RandString(n int) string {
    rand.Seed(time.Now().UTC().UnixNano())
    
    b := make([]byte, n)
    for i := range b {
        b[i] = alphabet[rand.Intn(len(alphabet))]
    }
    return string(b)
}