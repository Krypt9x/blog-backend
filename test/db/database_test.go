package test

import (
	"testing"

	"github.com/Krypt9x/blog-backend/internal/database"
)

func TestConnectDB(t *testing.T) {
	dbObj := database.InitDBService{
		DataSource: "postgres://root:secret@localhost:5432/blog",
	}
	db := dbObj.InitDB()
	defer db.Close()
}
