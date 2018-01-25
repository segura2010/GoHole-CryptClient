package main

import (
    "flag"
    "os"
    "os/exec"
    "fmt"

    "GoHole-CryptClient/config"
    "GoHole-CryptClient/dnsserver"
    "GoHole-CryptClient/encryption"
)

/* Update version number on each release:
    Given a version number x.y.z, increment the:

    x - major release
    y - minor release
    z - build number
*/
const BINARY_VERSION = "1.0.0"
var Commit string
var CompilationDate string

func showVersionInfo(){
    fmt.Println("----------------------------------------")
    fmt.Printf("GoHole-CryptClient v%s\nCommit: %s\nCompilation date: %s\n", BINARY_VERSION, Commit, CompilationDate)
    fmt.Println("----------------------------------------")
}

func main(){

    // Command line options
    port := flag.String("p", "", "Set proxy DNS server port")
    cfgFile := flag.String("c", "./config.json", "Config file")
    version := flag.Bool("v", false, "Show current version")

    // option to start the DNS server
    startDNS := flag.Bool("s", false, "Start proxy DNS server")
    // option to stop the DNS server
    stopDNS := flag.Bool("stop", false, "Stop proxy DNS server")

    
    flag.Parse()

    config.CreateInstance(*cfgFile)
    if *port != ""{
        config.GetInstance().DNSPort = *port
    }

    encryption.CreateInstance()
    encryption.ImportKeyFromFile(config.GetInstance().EncryptionKey)

    if *version{
        showVersionInfo()
    }

    if *startDNS{
        dnsserver.ListenAndServe()
    }
    if *stopDNS{
        exec.Command("killall", os.Args[0]).Run()
    }

}
