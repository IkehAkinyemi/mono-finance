package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/IkehAkinyemi/mono-finance/api"
	db "github.com/IkehAkinyemi/mono-finance/db/sqlc"
	"github.com/IkehAkinyemi/mono-finance/doc/swagger"
	"github.com/IkehAkinyemi/mono-finance/gapi"
	"github.com/IkehAkinyemi/mono-finance/pb"
	"github.com/IkehAkinyemi/mono-finance/utils"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
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

	runDBMigrations(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGRPCServer(config, store)
}

func runDBMigrations(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatalf("cannot create a new migrate instance: %v", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrateup: %v", err)
	}
	
	log.Println("db migrated successfully")
}

func runGRPCServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create monofinance server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMonoFinanceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatalf("cannot start listener: %v", err)
	}

	log.Printf("Starting gRPC server, port = %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server: %v", err)
	}
}

func runGatewayServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create monofinance server: ", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jsonOpts := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOpts)
	pb.RegisterMonoFinanceHandlerServer(ctx, grpcMux, server)

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	fs := http.FileServer(http.FS(swagger.StatisFile))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatalf("cannot start listener: %v", err)
	}

	log.Printf("Starting gRPC Gateway server, port = %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatalf("cannot start gRPC server: %v", err)
	}
}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	if err := server.Start(fmt.Sprint(config.HTTPServerAddress)); err != nil {
		log.Fatalf("error occur starting server: %v", err)
	}
}
