package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/Germanchrystan/GeekStore/api/cmd/server/routes"
)

//=============================================================//
func main() {
	envs, err := godotenv.Read("./../../.env")
	fmt.Print(envs)
	if err != nil {
		fmt.Println("ERROR")
		throwServerError(err)
	}

	dbUrl := envs["DATABASE_URL"]
	fmt.Println(dbUrl)
	if dbUrl == "" {
		throwServerError(err)
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		throwServerError(err)
	}
	r := gin.Default()
	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}

//=============================================================//
func throwServerError(err error) {
	panic(errors.New(err.Error()))
}

//=============================================================//
