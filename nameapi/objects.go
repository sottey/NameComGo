package nameapi

/// Models
type Domain struct {
    DomainName string `json:"domainName"`
    NameServers [] string `"json:nameservers"`
    Contacts Contacts `json:contacts"`
    PrivacyEnabled bool `json:privacyEnabled"`
    Locked bool `json:locked"`
    AutorenewEnabled bool `json:autorenewEnabled"`
    ExpireDate string `json:expireDate"`
    CreateDate string `json:createDate"`
    RenewalPrice float64 `json:renewalPrice"`
}

type PurchaseDetails struct {
    DomainObj Domain `json:"domain"`
    PurchasePrice float64 `json:"purchasePrice"`
    PurchaseType string `json:"purchaseType"`
    Years int32 `json:"years"`
    TLDRequirements map[string]string `json:"tldRequirements"`
    PromoCode string `json:"promoCode"`
}

type ContactsWrapper struct {
    ContactData Contacts `json:"contacts"`
}
type Contacts struct {
    Registrant Contact `json:"registrant"`
    Administrative Contact `json:"admin"`
    Technical Contact `json:"tech"`
    Billing Contact `json:"billing"`
}
type Contact struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    CompanyName string `json:"companyName"`
    Address1 string `json:"address1"`
    Address2 string `json:"address2"`
    City string `json:"city"`
    State string `json:"state"`
    Zip string `json:"zip"`
    Country string `json:"country"`
    Phone string `json:"phone"`
    Fax string `json:"fax"`
    Email string `json:"email"`
    
}
type Record struct {
    ID int32 `json:"id"`
    DomainName string `json:"domainName"`
    Host string `json:"host"`
    Fqdn string `json:"fqdn"`
    RecordType string `json:"type"`
    Answer string `json:"answer"`
    TTL uint32 `json:"ttl"`
    Priority uint32 `json:"priority"`
}

type DNSSEC struct {
    DomainName string `json:"domainName"`
    KeyTag int32 `json:"keyTag"`
    Algorithm int32 `json:"algorithm"`
    DigestType int32 `json:"digestType"`
    Digest string `json:"digest"`
}

type URLForwardingGroup struct {
    URLForwardings []URLForwarding `json:"urlForwarding"`
}

type URLForwarding struct {
    DomainName string `json:"domainName"`
    Host string `json:"host"`
    ForwardsTo string `json:"forwardsTo"`
    ForwardType string `json:"type"`
    Title string `json:"title"`
    Meta string `json:"meta"`
}

type EmailForwarding struct {
    DomainName string `json:"domainName"`
    EmailBox string `json:"emailBox"`
    EmailTo string `json:"emailTo"`
}

type VanityNameserver struct {
    DomainName string `json:"domainName"`
    Hostname string `json:"hostname"`
    IPs []string `json:"ips"`
}

/// Interfaces
type IRequest interface {
    GetUrl() string
    GetConfig() Configuration
    GetMethod() string
}

type NameResponse struct {
    StatusCode int
    Body string
    Message string
}

// Create methods
func CreateDomainObject(domainName string) Domain {
    ret := new(Domain)
    
    ret.DomainName = domainName
    
    return *ret
}

func CreatePuchaseObject(domainName string, years int32, price float64) PurchaseDetails {
    ret := new(PurchaseDetails)
    
    ret.DomainObj = CreateDomainObject(domainName)
    ret.Years = years
    ret.PurchasePrice = price
    
    return *ret
}

func CreateRecordObject(id int32, domainName, host,recordType, answer string, ttl uint32) Record {
    ret := new(Record)
    if id != 0 {
        ret.ID = id
    }
    ret.DomainName = domainName
    ret.Host = host
    ret.RecordType = recordType
    ret.Answer = answer
    ret.TTL = ttl
    
    return *ret
}

func CreateURLForwardingObject(domainName, host, forwardsTo, forwardType string) URLForwarding {
    ret := new(URLForwarding)
    
    ret.DomainName = domainName
    ret.Host = host
    ret.ForwardsTo = forwardsTo
    ret.ForwardType = forwardType
    return *ret
}

func CreateDNSSECObject(domainName string, keyTag, algorithm, digestType int32, digest string) DNSSEC {
    ret := new (DNSSEC)
    
    ret.DomainName = domainName
    ret.KeyTag = keyTag
    ret.Algorithm = algorithm // https://www.iana.org/assignments/dns-sec-alg-numbers/dns-sec-alg-numbers.xhtml
    ret.DigestType = digestType // https://www.iana.org/assignments/ds-rr-types/ds-rr-types.xhtml
    ret.Digest = digest
    
    return *ret
}

func CreateVanityNameserverObject(domainName, hostname string, ips []string) VanityNameserver {
    ret := new (VanityNameserver)
    
    ret.DomainName = domainName
    ret.Hostname = hostname
    ret.IPs = ips

    return *ret
}