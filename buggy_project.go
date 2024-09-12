package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const errorPrefix = "ERROR: "

var db *sql.DB

func main() {
	db, _ = sql.Open("postgres", "user=postgres dbname=test password=password sslmode=disable")
	err := db.Ping()
	if err != nil {
		log.Fatalln(errorPrefix+"Database connection issue.", err)
	}

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/create", createUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Println(errorPrefix+"Wrong users limit.", err)
		fmt.Fprintf(w, "Wrong users limit")
		return
	}

	if limit < 0 {
		log.Println(errorPrefix+"Wrong users limit =", limit)
		fmt.Fprintf(w, "Wrong users limit")
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		log.Println(errorPrefix+"Wrong users offset.", err)
		fmt.Fprintf(w, "Wrong users offset")
		return
	}

	if offset < 0 {
		log.Println(errorPrefix+"Wrong users offset =", offset)
		fmt.Fprintf(w, "Wrong users offset")
		return
	}

	rows, err := db.Query("SELECT name FROM users LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		log.Println(errorPrefix+"Database issue.", err)
		fmt.Fprintf(w, "Database issue")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Println(errorPrefix+"Wrong user data.", err)
			fmt.Fprintf(w, "Wrong user data")
			return
		}
		fmt.Fprintf(w, "User: %s\n", name)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second) // Simulate a long database operation

	username := r.URL.Query().Get("name")
	_, err := db.Exec("INSERT INTO users (name) VALUES ($1)", username)
	if err != nil {
		errMsg := "Failed to create user: " + username
		log.Println(errorPrefix+errMsg, err)
		fmt.Fprint(w, errMsg)
		return
	}

	fmt.Fprintf(w, "User %s created successfully", username)
}
