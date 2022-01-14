package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type PgConnection struct {
	Host            string
	Port            int
	User            string
	Password        string
	DbName          string
	SslMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func CreateDBConnection(pgConn PgConnection) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=%s",
		pgConn.Host, pgConn.Port, pgConn.User, pgConn.Password, pgConn.DbName, pgConn.SslMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(pgConn.MaxOpenConns)
	db.SetMaxIdleConns(pgConn.MaxIdleConns)
	db.SetConnMaxLifetime(pgConn.ConnMaxLifetime)

	log.Println("Connected to database: ", pgConn.Host, "at port:", pgConn.Port, "db_name:", pgConn.DbName, "successfully!")

	return db, nil
}
