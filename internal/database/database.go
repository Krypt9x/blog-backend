package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type (
	DatabaseConfig interface {
		SetupDB(dataSource string) *sql.DB
	}

	DatabaseConfigPostgres struct {
	}
)

func (Configdb *DatabaseConfigPostgres) SetupDB(dataSource string) *sql.DB {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Panicln("database failed to connect")
	}
	log.Println("database connected")
	return db
}

type (
	InitDBContract interface {
		InitDB()
	}

	InitDBService struct {
		DataSource string
	}
)

func (db *InitDBService) InitDB() *sql.DB {
	dbObj := DatabaseConfigPostgres{}
	return dbObj.SetupDB(db.DataSource)
}
