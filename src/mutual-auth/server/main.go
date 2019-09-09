package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 1. Configurar handler
	http.HandleFunc("/sample", func(res http.ResponseWriter, req *http.Request) {
		defer log.Printf(req.URL.Path)
		res.Write([]byte("Hello World.. with mutual authentication"))
	})

	// 2. Carga de certificado cliente
	caClientCert, err := ioutil.ReadFile(dir + "/src/mutual-auth/certs/client.crt")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caClientCert)

	// 3. Configurar TLS
	tlsConfig := &tls.Config{
		ClientCAs:  certPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	// 4. Instanciar y arrancar servidor HTTPS
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}
	errServer := server.ListenAndServeTLS(dir+"/src/mutual-auth/certs/server.crt", dir+"/src/mutual-auth/certs/server.key")
	if errServer != nil {
		log.Fatal(errServer)
	}
}
