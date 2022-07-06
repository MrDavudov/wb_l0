package main

import (
	"database/sql"
	"github.com/MrDavudov/LessonPosgresql/models"
	"github.com/MrDavudov/LessonPosgresql/service"
	"github.com/gorilla/mux"
)

var db *sql.DB

var dbConf = models.Config {
	Host:     "localhost",
	Port:     "5432",
	Username: "test",
	Password: "admin",
	DBName:   "postgres",
	SSLMode:  "disable",
}

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	db = service.DBInit(dbConf)
	service.Db = db
}