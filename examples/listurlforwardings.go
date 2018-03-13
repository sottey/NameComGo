package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func ListURLForwardingsExample() {
    fmt.Println("Executing ListURLForwardings() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any2.org"
    // Make the call
    nameResponse := namecomgo.ListURLForwardings(c, domainName)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
            var ret namecomgo.URLForwardingGroup
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in ListURLForwardings(): ", jsonErr)
            if len(ret.URLForwardings) < 1 {
                fmt.Printf("\nObject is empty! Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
            } else {
                for _, forward := range ret.URLForwardings {
                    fmt.Printf("\nDomainName:%v\nForwaradsTo:%v\n\n", forward.Host, forward.ForwardsTo)
                }
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}

