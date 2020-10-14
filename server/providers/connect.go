package providers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Connect open conntion with postgresql
func Connect() *sql.DB {
	env := &struct {
		DBNAME string
		DBUSER string
		DBPASS string
		DBHOST string
		DBPORT string
	}{
		DBNAME: os.Getenv("DBNAME"),
		DBUSER: os.Getenv("DBUSER"),
		DBPASS: os.Getenv("DBPASS"),
		DBHOST: os.Getenv("DBHOST"),
		DBPORT: os.Getenv("DBPORT"),
	}

	psqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHOST, env.DBPORT, env.DBUSER, env.DBPASS, env.DBNAME)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		log.Println("DB Open: ", err)
	}
	if err = db.Ping(); err != nil {
		log.Println("Ping: ", err)
	}
	return db
}
