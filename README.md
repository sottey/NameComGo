# NameComGo

Simple Go language wrapper for the [name.com v4 API](https://www.name.com/api-docs)

## Getting Started

All the calls are in separate go files in the nameapi directory.

API wide models are in the ./nameapi/objects.go file

Misc utility methods are in the ./nameapi/utility.go file

All unit tests are in separate files in the ./nameapi/tests directory

./main.go contains calls for all the example methods which are contained in the examples directory

### Prerequisites

You will need a name.com account. Once you have that, log in and go to [this link](https://www.name.com/account/settings/api) to set up your api tokens.

Once you have the tokens, set the following environmental variables (you can change the var names if you like by changing the entries in ./nameapi/config.go)   :
nameuser=[prod user name]
nametoken=[prod token]
nametestuser=[test user name]
nametesttoken=[test token]

### Installing

Either use go install OR place the nameapi folder in your project and import namecomgo/nameapi

## Call Status

| API Call      | Status
| ------------- |:-------------:
|Record Model | ***Complete***
|ListRecords | ***Complete***
|GetRecord | ***Complete***
|CreateRecord | ***Complete***
|UpdateRecord | ***Complete***
|DeleteRecord | ***Complete***
|Dnssec Model | ***Complete***
|ListDNSSECs | ***Complete***
|GetDNSSEC | ***Complete***
|CreateDNSSEC | ***Complete***
|DeleteDNSSEC | ***Complete***
|Domain Model | ***Complete***
|Search Result Model | ***Complete***
|ListDomains | ***Complete***
|GetDomain | ***Complete***
|CreateDomain | ***Complete***
|EnableAutorenew | ***Complete***
|DisableAutorenew | ***Complete***
|RenewDomain | ***Complete***
|GetAuthCodeForDomain | ***Complete***
|SetNameservers | ***Complete***
|LockDomain | ***Complete***
|UnlockDomain | ***Complete***
|CheckAvailability | ***Complete***
|Search | ***Complete***
|Hello | ***Complete***
|Urlforwarding Model | ***Complete***
|ListURLForwardings | ***Complete***
|GetURLForwarding | ***Complete***
|CreateURLForwarding | ***Complete***
|UpdateURLForwarding | ***Complete***
|DeleteURLForwarding | ***Complete***
|Emailforwarding Model | ***Complete***
|ListEmailForwardings | ***Complete***
|GetEmailForwarding | ***Complete***
|CreateEmailForwarding | ***Complete***
|UpdateEmailForwarding | ***Complete***
|DeleteEmailForwarding | ***Complete***
|PurchasePrivacy | ***Complete*** (API issue: No way to determine purchase price)
|SetContacts | ***Complete*** (API Issue: 500: Internal Error if no email)
|Vanitynameserver Model | ***Complete***
|ListVanityNameservers | ***Complete***
|GetVanityNameserver | ***Complete***
|CreateVanityNameserver | ***Complete***
|UpdateVanityNameserver | ***Complete***
|DeleteVanityNameserver | ***Complete***
|Transfer Model | ***Complete***
|ListTransfers | ***Complete***
|GetTransfer | ***Complete***
|CreateTransfer | Incomplete
|CancelTransfer | Incomplete
|SearchStream | Incomplete


### Examples

All examples are listed in ./main.go. Each example is in the examples directory, with the method name as it's filename. Note that some calls require you to change the arguments to appropriate values for your environment.

## Running the tests

In the ./nameapi/tests directory simply run go test

## Built With

* [Name.com](http://www.name.com/) - Their excellent domain API

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors

* **Sean Ottey** - *Initial work* - [sottey](https://github.com/sottey)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Name.com
* The GoLang Community
