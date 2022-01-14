package main

import (
	"database/sql"
	"go-context/config"
	"go-context/db"
	"go-context/handler"
	"go-context/repository"
	"go-context/usecase"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var (
	dbConn *sql.DB
	conf   *config.Config
)

func init() {
	conf = config.InitConfig()

	db, err := db.CreateDBConnection(db.PgConnection{
		Host:            conf.Database.Pg.Host,
		Port:            conf.Database.Pg.Port,
		User:            conf.Database.Pg.User,
		Password:        conf.Database.Pg.Password,
		DbName:          conf.Database.Pg.Dbname,
		SslMode:         conf.Database.Pg.SslMode,
		MaxOpenConns:    conf.Database.Pg.MaxOpenConnection,
		MaxIdleConns:    conf.Database.Pg.MaxIdleConnection,
		ConnMaxLifetime: conf.Database.Pg.MaxConnectionLifetime,
	})
	if err != nil {
		log.Fatalln("failed connect to database", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("failed ping database", err)
	}
	dbConn = db
}

func main() {
	repo := repository.NewRepository(dbConn)
	usecase := usecase.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	http.HandleFunc("/user", handler.HandlerPrintUser)

	http.ListenAndServe(":8080", nil)
}
