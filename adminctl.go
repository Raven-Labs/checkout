package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Initialise the default admin account
func initAdmin() {
	insertStatement := `insert into "Admins"("Username", "Password") values($1, $2)`

	// Hash password to store in DB
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(config.DefaultAdmin.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash default admin's password.")
	}

	// Insert into DB
	_, err = db.Exec(insertStatement, "admin", pwdHash)
	if err != nil {
		log.Fatal("Failed to insert default admin creds into DB.", err)
	}
}

// Authenticate a login attempt
func authAdmin(username, password string) int {
	rows, err := db.Query(`SELECT "Username" FROM "Admins"`)
	if err != nil {
		log.Fatal("Failed to query database for usernames")
	}

	err = rows.Scan(&username)
	if err != nil {
		return 401
	}
	return 200

}
