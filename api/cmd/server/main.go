package main

import (
	"database/sql"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Germanchrystan/GeekStore/api/cmd/server/routes"
)

//=============================================================//
func main() {
	_ = godotenv.Load()
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		throwServerError()
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		throwServerError()
	}
	r := gin.Default()
	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}

//=============================================================//
func throwServerError() {
	panic(errors.New("Server Connection Error"))
}

//=============================================================//
