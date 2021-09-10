package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=vagrant dbname=test-dipta")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Read migrations from a folder
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/postgres",
	}

	results, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(results)

	/*results, err = migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(results)*/

	status, err := migrate.GetMigrationRecords(db, "postgres")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, stat := range status {
		fmt.Println(*stat)
	}
}
