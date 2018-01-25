package config

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

// MyConfig struct
// This is the struct that the config.json must have
type MyConfig struct {
    DNSPort string // listen on port
    GoHoleServer string // GoHole Server IP
    GoHoleServerPort string // GoHole Server Port

    EncryptionKey string // Path to the encryption key file
}

var instance *MyConfig = nil

func CreateInstance(filename string) *MyConfig {
    var err error
    instance, err = loadConfig(filename)
    if err != nil {
        log.Printf("Error loading config file: %s\nUsing default config.", err)
        // use defaults
        instance = &MyConfig{
            DNSPort: "53",
            GoHoleServer: "127.0.0.1",
            GoHoleServerPort: "443",
            EncryptionKey: "enc.key",
        }
    }

    return instance
}

func GetInstance() *MyConfig {
    return instance
}

func GetGoHoleServerAndPort() (string){
    i := GetInstance()
    return i.GoHoleServer + ":" + i.GoHoleServerPort
}

func loadConfig(filename string) (*MyConfig, error){
    var s *MyConfig

    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        return s, err
    }
    // Unmarshal json
    err = json.Unmarshal(bytes, &s)
    return s, err
}
