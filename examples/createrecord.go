package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func CreateRecordExample() {
    fmt.Println("Executing CreateRecord() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    record := namecomgo.CreateRecordObject(0, "any4.org", namecomgo.RandString(6), "A", "10.0.0.9", 300)
    
    // Make the call
    nameResponse := namecomgo.CreateRecord(c, record)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
    	    var ret namecomgo.Record
    	    
    	    // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in CreateRecord(): ", jsonErr)
        
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