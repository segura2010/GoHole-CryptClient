# GoHole CryptClient

GoHole is a DNS server written in Golang with the same idea than the [PiHole](https://pi-hole.net), blocking advertisements's and tracking's domains.

In order to provide a safer way to make DNS queries, GoHole introduces AES encryption, allowing you to encrypt all your DNS queries. To make use of the encryption feeature, you need the GoHole CryptClient, which is basically a "proxy DNS server" that listen for unencrypted queries in your computer and then encrypts these queries to send the encrypted queries to your GoHole server. The server will respond with an encrypted reply that this client will unencryt and send as result to your system.

All your DNS queries will be encrypted and safe! :)

### Installation

1. Clone this repository and rename the folder to `GoHole-CryptClient` if it is not the name.
2. Run the install script `install.sh` to install all the dependencies.
3. Compile using Makefile (`make`). Or run `make install` to install (you should have included your $GOPATH/bin to your $PATH).

Finally, run using the executable for your platform.

### Usage

Generate a new key with your GoHole server and use it to encrypt your queries. Edit the config_example.json file and run the client.

**Tested on Go 1.9.2**
