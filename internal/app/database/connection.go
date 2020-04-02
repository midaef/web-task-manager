package connection

import (
	"log"

	"github.com/jackc/pgx"
)

func NewConnection() (conn *pgx.Conn) {
	var runtimeParams map[string]string
	runtimeParams = make(map[string]string)
	runtimeParams["application_name"] = "todo"
	connConfig := pgx.ConnConfig{
		User:              "postgres",
		Password:          "1212ru",
		Host:              "localhost",
		Port:              65001,
		Database:          "tasks",
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
		RuntimeParams:     runtimeParams,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		log.Fatal("Connection error", err)
	}
	return conn
}
