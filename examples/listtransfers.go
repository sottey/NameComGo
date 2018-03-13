package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func ListTransfersExample() {
    fmt.Println("Executing ListTransfers() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Make the call
    nameResponse := namecomgo.ListTransfers(c)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.ListTransfersResponse
            
            // Hydrate the response object with the response data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in ListTransfers(): ", jsonErr)
            
            if len(ret.Transfers) == 0 {
                fmt.Printf("No transfers in account\n")
            } else {
            // Display results
                for _, transfer := range ret.Transfers {
                    fmt.Printf("\nDomainName:%v\nEmail:%v\nStatus:%v\n\n", transfer.DomainName, transfer.Email, transfer.Status)
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