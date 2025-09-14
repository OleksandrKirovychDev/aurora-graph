package main

import (
	"aurora-graph/account/config"
	"aurora-graph/account/internal"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tinrab/retry"
)

func main() {
	if err := godotenv.Load(); err != nil {
    	log.Println("No .env file found, relying on system env")
	}
	config.Init()

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
	port, _ := strconv.Atoi(config.GRPCPort)
	log.Printf("Account service running on port %v", port)
	service := internal.NewAccountService(repository)
	log.Fatal(internal.ListenGRPC(service, port))
}