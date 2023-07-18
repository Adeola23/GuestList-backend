package main

import (
	_ "fmt"
	_ "log"
	_ "net/http"

	"github.com/getground/tech-tasks/backend/database"
	_ "github.com/getground/tech-tasks/backend/database"
	"github.com/getground/tech-tasks/backend/router"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	database.Setup()
	router.Setup()
	database.Setup()
	router.Router.Run(":3000")

}


