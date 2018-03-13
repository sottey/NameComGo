package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestSetContacts(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    domainName := "any1.org"
    var contacts nameapi.ContactsWrapper
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
 
    nameResponse := nameapi.SetContacts(c,domainName, contacts)
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.Domain
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in SetContacts(): ", jsonErr)
        
        if ret.Contacts.Registrant.LastName != "Bowbemma" {
            t.Errorf("Incorrect SetContacts response! Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}