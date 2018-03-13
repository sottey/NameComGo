package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func GetDomainExample() {
    fmt.Println("Executing GetDomain() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")

    // Test data
    domainName := "any4.org"
    
    // Make the call
    nameResponse := namecomgo.GetDomain(c, domainName)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.Domain
            
            // Hydrate object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in GetDomain(): ", jsonErr)
            
            fmt.Printf("\nDomain: %v\nPrivacyEnabled:%v\nLocked:%v\nAutoRenew:%v\nExpireDate:%v\nCreateDate:%v\nRenewalPrice:%v\n\n", ret.DomainName,ret.PrivacyEnabled, ret.Locked, ret.AutorenewEnabled, ret.ExpireDate, ret.CreateDate,ret.RenewalPrice)

    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }

}