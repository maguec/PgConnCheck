package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/kouhin/envflag"
	_ "github.com/lib/pq"
)

func main() {
	var (
		pghost     = flag.String("pghost", "localhost", "Database host")
		pgpassword = flag.String("pgpassword", "", "Database password")
		pgdb       = flag.String("pgdb", "", "Database name")
		pguser     = flag.String("pguser", "postgres", "Database user")
		pgport     = flag.Int("pgport", 5432, "Database master port")
		pgsleep     = flag.Int("pgsleep", 100, "Time to sleep between checks in ms")
	)
	if err := envflag.Parse(); err != nil {
		panic(err)
	}
	is_error := false
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			*pghost, *pgport, *pguser, *pgpassword, *pgdb),
	)
  psqlurl := fmt.Sprintf(
    "postgres://%s:XXXXX@%s:%d/%s",
    *pguser, *pghost, *pgport, *pgdb)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connecting to %s\n", psqlurl)
	for {
		err = db.Ping()
		if !is_error && err != nil {
			is_error = true
			log.Printf("Error: %s %s\n", err, psqlurl)
		}
		if is_error && err == nil {
			is_error = false
			log.Printf("Connected to %s\n", psqlurl)
		}
		time.Sleep(time.Duration(*pgsleep) * time.Millisecond)
	}
}
