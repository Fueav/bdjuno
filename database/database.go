package database

import (
	"fmt"

	db "github.com/Fueav/juno/database"
	"github.com/Fueav/juno/database/postgresql"
	"github.com/jmoiron/sqlx"
)

var _ db.Database = &Db{}

// Db represents a PostgreSQL database with expanded features.
// so that it can properly store custom BigDipper-related data.
type Db struct {
	*postgresql.Database
	Sql *sqlx.DB
}

// Builder allows to create a new Db instance implementing the db.Builder type
func Builder(ctx *db.Context) (db.Database, error) {
	database, err := postgresql.Builder(ctx)
	if err != nil {
		return nil, err
	}

	psqlDb, ok := (database).(*postgresql.Database)
	if !ok {
		return nil, fmt.Errorf("invalid configuration database, must be PostgreSQL")
	}

	return &Db{
		Database: psqlDb,
		Sql:      sqlx.NewDb(psqlDb.Sql, "postgresql"),
	}, nil
}

// Cast allows to cast the given db to a Db instance
func Cast(db db.Database) *Db {
	bdDatabase, ok := db.(*Db)
	if !ok {
		panic(fmt.Errorf("given database instance is not a Db"))
	}
	return bdDatabase
}
