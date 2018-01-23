// Allows for testing the domain logic without having to go through
// the http server
package main

import (
	"context"
	"time"

	"github.com/gilcrest/go-API-template/pkg/domain/appUser"
	"github.com/gilcrest/go-API-template/pkg/env"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Print("Start main")

	// Start Timer
	start := time.Now()

	// Get an empty context with a Cancel function included
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel ctx as soon as main returns.

	// Initializes "environment" struct type
	env, err := env.NewEnv()
	if err != nil {
		log.Fatal().Err(err)
	}

	// Creates a new instance of the appUser.User struct type
	inputUsr := appUser.User{Username: "repoMan",
		MobileID:  "(617) 302-7777",
		Email:     "repoman@alwaysintense.com",
		FirstName: "Otto",
		LastName:  "Maddox"}

	// Create method does validation and then inserts user into db
	tx, err := inputUsr.Create(ctx, env)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Check to ensure that the CreateDate struct field is populated by
	// making sure it's not at it's zero value to ensure that the db
	// transaction was successful before commiting
	if !inputUsr.CreateDate.IsZero() {
		err = tx.Commit()
		if err != nil {
			log.Fatal().Err(err)
		}
	} else {
		err = tx.Rollback()
		log.Fatal().Err(err).Str("Rollback", "CreateDate is nil, rolled back txn")
		if err != nil {
			log.Fatal().Err(err)
		}
	}

	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
