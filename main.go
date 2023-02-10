package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/IkehAkinyemi/mono-finance/api"
	db "github.com/IkehAkinyemi/mono-finance/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbURI    = "postgresql://mono_finance:Akinyemi@localhost:5432/mono_finance?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbURI)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)
	if err := server.Start(fmt.Sprint(serverAddress)); err != nil {
		log.Fatalf("error occur starting server: %v", err)
	}
}