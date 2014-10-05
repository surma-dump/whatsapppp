package main

import (
	"github.com/jmoiron/sqlx/reflectx"
	// "database/sql"
	"log"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/voxelbrain/goptions"
)

var options = struct {
	DBPath string        `goptions:"-p, --path, description='Path to database folder', obligatory"`
	Help   goptions.Help `goptions:"-h, --help, description='Show this help'"`
}{
	DBPath: "/data/data/com.whatsapp/databases",
}

func main() {
	goptions.ParseAndFail(&options)

	dbpath, err := filepath.Abs(options.DBPath)
	if err != nil {
		log.Fatalf("Could not absolutify path: %s", err)
	}

	db, err := sqlx.Open("sqlite3", "file://"+filepath.Join(dbpath, "msgstore.db"))
	if err != nil {
		log.Fatalf("Could not open database: %s", err)
	}

	rows, err := db.Queryx("SELECT * FROM chat_list ORDER BY sort_timestamp DESC;")
	if err != nil {
		log.Fatalf("Could not query database: %s", err)
	}
	defer rows.Close()
	rows.Mapper = reflectx.NewMapper("sql")
	for rows.Next() {
		res := ChatListItem{}
		err := rows.StructScan(&res)
		if err != nil {
			log.Printf("Could not scan row: %s", err)
			continue
		}
		log.Printf("%#v", res)
	}

	defer db.Close()
}
