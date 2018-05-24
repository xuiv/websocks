package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/xuiv/websocks/core"
)

var (
	mode       = flag.String("mode", "cert", "mode: server, client, cert")
	listenAddr = flag.String("listen", ":1080", "local listening port")

	ecdsa = flag.Bool("ecdsa", true, "generate ecdsa key and cert(P-256)")
	hosts = flag.String("hosts", "", "certificate hosts")

	webpath  = flag.String("path", "/", "server.com/path, like web, start with '/'")
	tls      = flag.Bool("tls", true, "enable built-in tls")
	certPath = flag.String("cert", "websocks.cer", "tls cert path")
	keyPath  = flag.String("key", "websocks.key", "tls key path")
	proxy    = flag.String("proxy", "", "reverse proxy url, leave blank to disable")

	serverURL    = flag.String("server", "wss://localhost:8080", "server url")
	serverName   = flag.String("name", "", "fake server name for tls client hello, leave blank to disable")
	insecureCert = flag.Bool("insecure", true, "InsecureSkipVerify: true")
)

func gencert() (err error) {
	var key, cert []byte
	lhosts := strings.Split(*hosts, " ")
	if *ecdsa {
		key, cert, err = core.GenP256(lhosts)
		fmt.Println("Generated ecdsa P-256 key and cert")
	} else {
		key, cert, err = core.GenRSA2048(lhosts)
		fmt.Println("Generated rsa 2048 key and cert")
	}

	err = ioutil.WriteFile("websocks.key", key, 0600)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("websocks.cer", cert, 0600)
	if err != nil {
		return
	}
	return
}

func doserver() (err error) {
	server := core.Server{
		Pattern:    *webpath,
		ListenAddr: *listenAddr,
		TLS:        *tls,
		CertPath:   *certPath,
		KeyPath:    *keyPath,
		Proxy:      *proxy,
		CreatedAt:  time.Now(),
	}

	err = server.Listen()
	if err != nil {
		return
	}

	return
}

func doclient() (err error) {
	u, err := url.Parse(*serverURL)
	if err != nil {
		return
	}

	lAddr, err := net.ResolveTCPAddr("tcp", *listenAddr)
	if err != nil {
		return
	}

	local := core.Client{
		ListenAddr:   lAddr,
		URL:          u,
		ServerName:   *serverName,
		InsecureCert: *insecureCert,
	}

	err = local.Listen()
	if err != nil {
		return
	}

	return nil
}

func main() {
	flag.Parse()
	switch *mode {
	case "cert":
		gencert()
		fmt.Println("run server: websocks -mode server -listen :8080")
		fmt.Println("run client: websocks -mode client -listen :1080 -server wss://youdomain.com")
	case "server":
		doserver()
	case "client":
		doclient()
	}
}
