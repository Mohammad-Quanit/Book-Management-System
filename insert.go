package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Insert(db *sql.DB) {
	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	id := 0
	err := db.QueryRow(sqlStatement, 25, "quanit.io", "mohd", "quanit").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", id)
}
