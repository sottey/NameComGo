package main

import (
    "namecomgo/examples"
)

    // NOTES
    // For these examples, you will need to change the parameters in some methods to
    // reflect your testing or production environment.
    // For example, the GetAuthCodeForDomainExample needs a domain you have registered
    // in your specified environment (test or prod). For ListDomainsExample, you will
    // not get anything back unless you have domains registered.
    
    // Also, to allow these examples to work without a bunch of changing of strings,
    // we are creating random strings for things like records and url forwarding. In
    // production, obviously, you would use your own, real, values.
    // END NOTES

func main() {
    // ////////////////////////////////
    // Simple server "Hello"
    // examples.HelloExample()
    
    // ////////////////////////////////
    // List all domains in your account
    // examples.ListDomainsExample()
    
    // ////////////////////////////////
    // Retrieve transfer auth code for a registered domain in your account
    examples.GetAuthCodeForDomainExample()
    
    // ////////////////////////////////
    // Search for unregistered domains using a keyword
    // examples.SearchExample()
    
    // ////////////////////////////////
    // Check the availability of 1-n domain names (NOTE: Name's API needs domain names
    // to be all lowercase. Issue submitted)
    // examples.CheckAvailabilityExample()
    
    // List all DNS records for a domain you have in the account
    // examples.ListRecordsExample()
    
    // ////////////////////////////////
    // Create a DNS record for a domain you have registered
    // examples.CreateRecordExample()
    
    // ////////////////////////////////
    // Retrieve a specific DNS record for a domain you have registered
    // examples.GetRecordExample()
    
    // ////////////////////////////////
    // Update a DNS record for a domain you have registered
    // examples.UpdateRecordExample()

    // ////////////////////////////////
    // Delete a DNS record for a domain you have registered
    // examples.DeleteRecordExample()
    
    // ////////////////////////////////
    // Enable AutoRenew for a domain you have registered
    // examples.EnableAutoRenewExample()
    
    // ////////////////////////////////
    // Disable AutoRenew for a domain you have registered
    // examples.DisableAutoRenewExample()
    
    // ////////////////////////////////
    // Create a URL forward for a domain you have registered
    // examples.CreateURLForwardingExample()

    // ////////////////////////////////
    //List all forwarding for a domain you have registered
    // examples.ListURLForwardingsExample()
    
    // ////////////////////////////////
    // Get a URL forward for a domain you have registered
    // examples.GetURLForwardingExample()
    
    // ////////////////////////////////
    // Update a URL forward for a domain you have registered
    // examples.UpdateURLForwardingExample()
    
    // ////////////////////////////////
    // Delete a URL forward for a domain you have registered
    // examples.DeleteURLForwardingExample()
    
    // ////////////////////////////////
    // Retrieve details for a domain you have registered
    // examples.GetDomainExample()
    
    // ////////////////////////////////
    // Lock a domain you have registered
    // examples.LockDomainExample()
    
    // ////////////////////////////////
    // Unock a domain you have registered
    // examples.UnlockDomainExample()
    
    // ////////////////////////////////
    // Register a domain
    // examples.CreateDomainExample()
    
    // ////////////////////////////////
    // Purchase privacy for a domain you have registered
    // examples.PurchasePrivacyExample()
    
    // List DNSSEC's added for a domain you have registered
    // examples.ListDNSSECsExample()
    
    // Create a DNSSEC entry for a domain you have registered
    // Note: if exact entry already exists, 500 server error occurrs
    // examples.CreateDNSSECExample()
    
    // Get a specific DNSSEC entry for a domain you have registered
    // examples.GetDNSSECExample()
    
    // Delete a specific DNSSEC entry for a domain you have registered
    // examples.DeleteDNSSECExample()
    
    // Configure Name Servers for a domain you have registered
    // examples.SetNameServersExample()
    
    // Set contact information for a domain you have registered (broken, API issue)
    // examples.SetContactsExample()
    
    // Configure email forwarding for a domain you have registered
    // examples.CreateEmailForwardingExample()
    
    // List all email forwarding for a domain you have registered
    // examples.ListEmailForwardingsExample()
    
    // Get specific email forwarding for a domain you have registered
    // examples.GetEmailForwardingExample()

    // Update specific email forwarding for a domain you have registered
    // examples.UpdateEmailForwardingExample()

    // Delete specific email forwarding for a domain you have registered
    // examples.DeleteEmailForwardingExample()
    
    // List vanity Nameservers for a domain you have registered
    // examples.ListVanityNameserversExample()

    // Create Vanity Server entry for a domain you have registered
    // examples.CreateVanityNameserverExample()
    
    // Get specific Vanity Server entry for a domain you have registered
    // examples.GetVanityNameserverExample()

    // Update specific Vanity Server entry for a domain you have registered
    // examples.UpdateVanityNameserverExample()
    
    // Delete specific Vanity Server entry for a domain you have registered
    // examples.DeleteVanityNameserverExample()
    
    // List all transfers in the account
    // examples.ListTransfersExample()
    
    // Get a specific Transfer in the account
    // examples.GetTransferExample()
    
    // Create a Transfer in the account
    // examples.CreateTransferExample()
}
