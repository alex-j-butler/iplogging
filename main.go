package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/vattle/sqlboiler/queries/qm"

	"alex-j-butler.com/ip-logging/config"
	"alex-j-butler.com/ip-logging/models"
	"alex-j-butler.com/loghandler"
)

var DB *sql.DB

func main() {
	// Create the loghandler server.
	logs, err := loghandler.Dial(config.Conf.LogServer.LogAddress, config.Conf.LogServer.LogPort)
	if err != nil {
		log.Println("LogHandler error:", err)
		os.Exit(1)
	}

	// Connect to the PostgreSQL database.
	DB, err = sql.Open("postgres", config.Conf.Database.DSN)
	if err != nil {
		log.Println("Database error:", err)
		os.Exit(1)
	}

	// Ping the database to make sure we're properly connected.
	if err := DB.Ping(); err != nil {
		log.Println("Database error:", err)
		os.Exit(1)
	}

	// Add the ConnectEvent handler.
	logs.AddHandler(OnConnect)
}

func OnConnect(lh *loghandler.LogHandler, event *loghandler.ConnectEvent) {
	// Check that the connected player isn't a bot.
	if event.SteamID != "BOT" {
		dbUser, err := models.Users(DB, qm.Where("steam_id=?", event.SteamID)).One()
		if err != nil {
			// Insert a new record.
			var newUser models.User
			newUser.SteamID = event.SteamID
			newUser.Name = event.Username
			newUser.LastConnected = null.TimeFrom(time.Now())

			err = newUser.Insert(DB)

			if err != nil {
				log.Println("Database error:", err)
				return
			}

			dbUser = &newUser
		}

		// Update the user LastConnected timestamp.
		dbUser.LastConnected = null.TimeFrom(time.Now())
		dbUser.Update(DB, "last_connected")

		dbIP, err := models.IPAddress(DB, qm.Where("ipaddress=?")).One()
		if err != nil {
			// Insert a new record.
			var newUser models.User
			newUser.SteamID = event.SteamID
			newUser.Name = event.Username
			newUser.LastConnected = null.TimeFrom(time.Now())

			err = newUser.Insert(DB)

			if err != nil {
				log.Println("Database error:", err)
				return
			}

			dbUser = &newUser
		}
	}
}
