package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func CancelTransferExample() {
    fmt.Println("Executing CancelTransfer() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domain := "any99.org"

    // Make the call
    nameResponse := namecomgo.CancelTransfer(c, domain)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
    	    var ret namecomgo.Transfer
    	    
    	    // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in CancelTransfer(): ", jsonErr)
        
            // Display member variables from the response object
            fmt.Printf("\nNew Status: %v\n\n", ret.Status)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}