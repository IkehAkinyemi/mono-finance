package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/IkehAkinyemi/mono-finance/api"
	db "github.com/IkehAkinyemi/mono-finance/db/sqlc"
	"github.com/IkehAkinyemi/mono-finance/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config file", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)
	if err := server.Start(fmt.Sprint(config.ServerAddress)); err != nil {
		log.Fatalf("error occur starting server: %v", err)
	}
}
