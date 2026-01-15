package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	dbPath := flag.String("db", "./data.db", "Path to the database file")
	flag.Parse()

	query := strings.Join(flag.Args(), " ")
	if query == "" {
		fmt.Println("Usage: go run cmd/sqlite/main.go <SQL QUERY>")
		return
	}

	db, err := sql.Open("sqlite", *dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if strings.HasPrefix(strings.ToLower(query), "select") {
		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		cols, err := rows.Columns()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(strings.Join(cols, " | "))
		fmt.Println(strings.Repeat("-", 20))

		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		for rows.Next() {
			if err := rows.Scan(valuePtrs...); err != nil {
				log.Fatal(err)
			}

			var rowStrings []string
			for _, val := range values {
				if b, ok := val.([]byte); ok {
					rowStrings = append(rowStrings, string(b))
				} else {
					rowStrings = append(rowStrings, fmt.Sprintf("%v", val))
				}
			}
			fmt.Println(strings.Join(rowStrings, " | "))
		}
	} else {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Command executed successfully.")
	}
}
