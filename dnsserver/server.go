package dnsserver

import (
	"log"
	"net"

    "github.com/miekg/dns"

    "GoHole-CryptClient/config"
    "GoHole-CryptClient/encryption"
)

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	
	queryBytes, err := r.Pack()
	if err != nil{
		log.Printf("*** error: %s\n", err.Error())
		return
	}

	ciphertext, err := encryption.Encrypt(queryBytes)
	if err != nil{
		log.Printf("*** error: %s\n", err.Error())
		return
	}

	// send encrypted query to GoHole server
	conn, err := net.Dial("udp", config.GetGoHoleServerAndPort())
	if err != nil {
		log.Printf("*** error: %s\n", err.Error())
		return
	}
	defer conn.Close()

	conn.Write(ciphertext)

	// get reply
	buffer := make([]byte, 2048)
	conn.Read(buffer)
	
	reply, err := encryption.Decrypt(buffer)
	if err != nil {
		log.Printf("*** error: %s\n", err.Error())
		return
	}

	m.Unpack(reply)
	w.WriteMsg(m)
}

func ListenAndServe(){
	// Start DNS proxy server
	port := config.GetInstance().DNSPort
	server := &dns.Server{Addr: ":" + port, Net: "udp"}
	dns.HandleFunc(".", handleDnsRequest)

	log.Printf("Starting at %s\n", port)

	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start DNS Proxy Server: %s\n ", err.Error())
	}

}