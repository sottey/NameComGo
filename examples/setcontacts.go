package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func SetContactsExample() {
    fmt.Println("Executing SetContacts() example...\n")

    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any1.org"
    var contacts namecomgo.ContactsWrapper
    contacts.ContactData.Registrant.FirstName = "Emma"
    contacts.ContactData.Registrant.LastName = "Bowbemma"
    contacts.ContactData.Registrant.Address1 = "1234 s 8th"
    contacts.ContactData.Registrant.City = "Seattle"
    contacts.ContactData.Registrant.State = "WA"
    contacts.ContactData.Registrant.Country = "US"
    contacts.ContactData.Registrant.Phone = "206-111-1111"
    contacts.ContactData.Registrant.Email = "test@test.com"
    
    
    // To shorten this, just using the same for all contacts
    contacts.ContactData.Administrative = contacts.ContactData.Registrant
    contacts.ContactData.Technical = contacts.ContactData.Registrant
    contacts.ContactData.Billing = contacts.ContactData.Registrant
    
    // Make the call
    nameResponse := namecomgo.SetContacts(c, domainName, contacts)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.Domain
            
            // Hydrate the response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in SetContacts(): ", jsonErr)
            
            // Display results
            if ret.DomainName != "" {
                fmt.Printf("SetContacts Returned for domain %v\n", ret.DomainName)
            } else {
                fmt.Printf("Empty SetContacts response!")
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}