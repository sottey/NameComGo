package nameapi

import (
    "os"
)
// Configuration Model
type Configuration struct {
    Username string
    Token string
    BaseURL string
    Debug bool
}

// GetConfig Method - Sets username, token, url and debug to defaults (debug is true in test, false in prod)

// NOTE: for prod, nameuser and nametoken environment variables must be set
//       for test, nametestuser and nametesttoken environment variables must be set
func GetConfig(mode string) Configuration {
    var ret Configuration
    
    // Anthing other than "prod" results in test env
    if (mode == "prod") {
        // Set defaults for prod
        ret.Username = os.Getenv("nameuser")
        ret.Token = os.Getenv("nametoken")
        ret.BaseURL = "https://api.name.com/v4"
        ret.Debug = false
    } else {
        // Set defaults for test
        ret.Username = os.Getenv("nametestuser")
        ret.Token = os.Getenv("nametesttoken")
        ret.BaseURL = "https://api.dev.name.com/v4"
        ret.Debug = true
    }
    
    return ret
}
