package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func ListRecordsExample() {
    fmt.Println("Executing ListRecords() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any2.org"
    
    // Make the call
    nameResponse := namecomgo.ListRecords(c, domainName)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.ListRecordsResponse
            
            // Hydrate the response object with the response data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in ListRecords(): ", jsonErr)
        
            // Display results
            for _, record := range ret.Records {
                fmt.Printf("\nID: %v\nDomainName:%v\nHost:%v\nFqdn:%v\nType:%v\nAnswer:%v\nTTL:%v\nPriority:%v\n\n", record.ID,record.DomainName, record.Host, record.Fqdn, record.RecordType, record.Answer, record.TTL,record.Priority)
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}