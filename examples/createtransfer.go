package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func CreateTransferExample() {
    fmt.Println("Executing CreateTransfer() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domain := "any99.org"
    authcode := "j5Y4U3f4S0N"
    purchasePrice := 12.99
    privacy := true
    promoCode := "seano"
    
    
    // Make the call
    nameResponse := namecomgo.CreateTransfer(c, domain, authcode, promoCode, privacy, purchasePrice)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
    	    var ret namecomgo.CreateTransferResponse
    	    
    	    // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in CreateTransfer(): ", jsonErr)
        
            // Display member variables from the response object
            fmt.Printf("\nID: %v\nTotalPaid:%v\n\n", ret.OrderID,ret.TotalPaid)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}