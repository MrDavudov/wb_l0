package main

import (
	"github.com/MrDavudov/LessonPosgresql/models"
	"github.com/MrDavudov/LessonPosgresql/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

var dbConf = models.Config{
	Host:     "localhost",
	Port:     "5432",
	Username: "test",
	Password: "test",
	DBName:   "wb_test",
	SSLMode:  "disable",
}

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	db = service.DBInit(dbConf)
	db.SetMaxOpenConns(100)
	service.Db = db

}