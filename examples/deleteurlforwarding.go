package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func DeleteURLForwardingExample() {
    fmt.Println("Executing Hello() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any2.org"
    host := "www.any1.org"
    
    // <ake the call
    nameResponse := namecomgo.DeleteURLForwarding(c, domainName, host)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
        	var ret namecomgo.URLForwarding
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in DeleteURLForwardings(): ", jsonErr)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}