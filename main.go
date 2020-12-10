package main

import (
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"microblog/database"
	"microblog/handler/server"
	"os"
	"os/signal"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("error in main")

			os.Exit(1)
		}
	}()

	// connection to the database.
	db := database.New()
	if err := db.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	conn := &database.Data{
		DB: db.DB,
	}

	//Versioning the database
	err = database.VersionedDB(db, false)
	if err != nil {
		log.Fatal(err)
	}

	DaemonPort := os.Getenv("DAEMON_PORT")
	serv := server.NewApplication(DaemonPort, conn)

	// start the server.
	go serv.Start()

	// Wait for an in interrupt.
	// If you ask about <- look here https://tour.golang.org/concurrency/2
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown.
	_ = serv.Close()
	_ = database.Close()
}
