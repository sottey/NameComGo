package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func UpdateRecordExample() {
    fmt.Println("Executing Hello() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    record := namecomgo.CreateRecordObject(335560, "any4.org", namecomgo.RandString(6), "A", "10.0.0.12", 300)
    
    // Make the call
    nameResponse := namecomgo.UpdateRecord(c, record)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.Record
            
            // Hydrate the response object with the response data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in UpdateRecord(): ", jsonErr)
            
            // Display member variables from the response object
            fmt.Printf("\nID: %v\nDomainName:%v\nHost:%v\nFqdn:%v\nType:%v\nAnswer:%v\nTTL:%v\nPriority:%v\n\n", ret.ID,ret.DomainName, ret.Host, ret.Fqdn, ret.RecordType, ret.Answer, ret.TTL,ret.Priority)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }

}