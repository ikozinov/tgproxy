package main

import (
	"fmt"

	socks5 "github.com/armon/go-socks5"
	flag "github.com/ogier/pflag"
	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func main() {
	var listen = flag.StringP("listen", "l", "", "Listen on ip_address:port to serve requests")
	var port = flag.IntP("port", "p", 6785, "Listen on port to serve requests")
	var proxyname = flag.String("proxyname", "tgproxy", "Username for proxy auth")
	var proxypasswd = flag.String("proxypasswd", "83azFuckOffRKN!", "User password for proxy auth")

	flag.Parse()

	log.Infoln("Starting free Telegram Proxy...")

	if *listen == "" {
		s := fmt.Sprintf(":%d", *port)
		listen = &s
	}

	// Create a SOCKS5 server
	cred := make(socks5.StaticCredentials)
	cred[*proxyname] = *proxypasswd

	log.Infoln("Use one of the following username/password pair to pass proxy auth")

	for k, v := range cred {
		log.Infof("%s : %s", k, v)
	}

	conf := &socks5.Config{Credentials: cred}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", *listen); err != nil {
		panic(err)
	}

}
