package main

import (
	"aurora-graph/account/config"
	"aurora-graph/account/internal"
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/tinrab/retry"
)

func main() {
	var repository internal.Repository

	retry.ForeverSleep(2 * time.Second, func(_ int) (err error){
		db, err := sql.Open("postgres", config.DatabaseURL)
		if err != nil {
			log.Fatal(err)
		}

		repository, err = internal.NewPostgresRepository(db)
		if err != nil {
    		log.Fatal(err)
		}

		return
	})

	defer repository.Close()
	log.Println("Account service running")
	service := internal.NewAccountService(repository)
	port, _ := strconv.Atoi(config.GRPCPort)
	log.Fatal(internal.ListenGRPC(service, port))
}