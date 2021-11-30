package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	_mysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

func initDB(db string, clientCert string, clientKey string, serverCA string, serverName string) *gorm.DB {
	rootCertPool := x509.NewCertPool()
	if ok := rootCertPool.AppendCertsFromPEM([]byte(strings.ReplaceAll(serverCA, "\\n", "\n"))); !ok {
		log.Fatal("Failed to append PEM.")
	}

	certs, err := tls.X509KeyPair([]byte(strings.ReplaceAll(clientCert, "\\n", "\n")), []byte(strings.ReplaceAll(clientKey, "\\n", "\n")))
	if err != nil {
		log.Fatal(err)
	}
	clientX509 := []tls.Certificate{certs}

	err = _mysql.RegisterTLSConfig(serverName, &tls.Config{
		RootCAs:      rootCertPool,
		Certificates: clientX509,
		ServerName: serverName,
		Time: nil,
		InsecureSkipVerify: true,
		VerifyConnection: func(state tls.ConnectionState) error {
			commonName := state.PeerCertificates[0].Subject.CommonName
			if commonName != state.ServerName {
				return fmt.Errorf("invalid certificate name %q, expected %q", commonName, state.ServerName)
			}
			opts := x509.VerifyOptions{
				Roots:         rootCertPool,
				Intermediates: x509.NewCertPool(),
			}
			for _, cert := range state.PeerCertificates[1:] {
				opts.Intermediates.AddCert(cert)
			}
			_, err = state.PeerCertificates[0].Verify(opts)
			return err
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	database, err := gorm.Open(mysql.Open(db), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return database
}