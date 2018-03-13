package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func GetRecordExample() {
    fmt.Println("Executing GetRecord() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any4.org"
    var id int32 = 335560
    
    // Make the call
    nameResponse := namecomgo.GetRecord(c, domainName, id)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create response object
            var ret namecomgo.Record
            
            // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in GetRecord(): ", jsonErr)
    
            // Display returned data
            fmt.Printf("\nID: %v\nDomainName:%v\nHost:%v\nFqdn:%v\nType:%v\nAnswer:%v\nTTL:%v\nPriority:%v\n\n", ret.ID,ret.DomainName, ret.Host, ret.Fqdn, ret.RecordType, ret.Answer, ret.TTL,ret.Priority)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}