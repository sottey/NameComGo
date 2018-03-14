package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func GetAuthCodeForDomainExample() {
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // test data
    domainName := "any99.org"
    
    // Create a response object
    var ret namecomgo.GetAuthCodeForDomainResponse
    
    // Make the call
    nameResponse := namecomgo.GetAuthCodeForDomain(c, domainName)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)

    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
            // Hydrate the response object with the response data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in GetAuthCodeForDomain(): ", jsonErr)
            
            // Display member variables from the response object
            fmt.Printf("\nAuth code for domain: '%v' is '%v'\n", domainName, ret.AuthCode)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}
