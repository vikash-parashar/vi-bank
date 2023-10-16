package storage

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func CreateTablesFromFile(fileName, connStr string) error {
	// Database connection setup
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Read SQL file
	sqlFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	// Split SQL statements
	sqlStatements := strings.Split(string(sqlFile), ";")

	// Execute SQL statements
	for _, sqlStatement := range sqlStatements {
		trimmedStatement := strings.TrimSpace(sqlStatement)
		if trimmedStatement == "" {
			continue
		}

		_, err := db.Exec(trimmedStatement)
		if err != nil {
			return err
		}
	}

	fmt.Println("Tables created successfully")
	return nil
}
