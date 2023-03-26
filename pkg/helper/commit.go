package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx, resultErr error) {
	if resultErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalln(rollbackErr)
		}
		log.Fatalln(resultErr)

	}

	if err := tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
