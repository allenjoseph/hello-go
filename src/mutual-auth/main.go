package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
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

	// 1. Cargar certificado y clave privada del cliente
	clientCert, err := tls.LoadX509KeyPair(dir+"/src/mutual-auth/certs/client.crt", dir+"/src/mutual-auth/certs/client.key")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Cargar certificado del servidor
	caServerCert, err := ioutil.ReadFile(dir + "/src/mutual-auth/certs/server.crt")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caServerCert)

	// 3. Configurar TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}
	tlsConfig.BuildNameToCertificate()

	// 4. Crear un cliente HTTPS y realizar la llamada
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: tlsConfig},
	}
	resp, err := client.Get("https://localhost/sample")
	if err != nil {
		fmt.Println(err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(contents))
}
