package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func GetTransferExample() {
    fmt.Println("Executing GetTransfer() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any4.org"
    
    // Make the call
    nameResponse := namecomgo.GetTransfer(c, domainName)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create response object
            var ret namecomgo.GetTransferResponse
            
            // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in GetTransfer(): ", jsonErr)
    
            // Display returned data
            fmt.Printf("\nDomainName:%v\nEmail:%v\nStatus:%v\n\n", ret.DomainName, ret.Email, ret.Status)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}