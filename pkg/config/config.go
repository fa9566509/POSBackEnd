package config

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("posbackend.db"), &gorm.Config{})
	if err != nil {
		panic("Cannot initiate database")
	}
	return db, err
}

func GetListner() net.Listener {
	cer, err := tls.LoadX509KeyPair("github.com/fa9566509/POSBackEnd/pkg/config/certs/ssl.cert", "github.com/fa9566509/POSBackEnd/pkg/config/certs/ssl.key")
	if err != nil {
		fmt.Println("Error while reading cert files")
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	ln, err := tls.Listen("tcp", ":9010", config)
	if err != nil {
		panic(err)
	}

	return ln
}
