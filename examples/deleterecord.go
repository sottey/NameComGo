package examples

import (
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func DeleteRecordExample() {
    fmt.Println("Executing DeleteRecord() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any4.org"
    var id int32 = 335553
    
    // Make the call
    nameResponse := namecomgo.DeleteRecord(c, domainName, id)

    if nameResponse.StatusCode != 200 {
        fmt.Printf("\nNon 200 Status Code, Status: %v, Message: %v\n", nameResponse.StatusCode, nameResponse.Body)
    } else {
        fmt.Printf("\nSucessfully deleted record\n")
    }
}