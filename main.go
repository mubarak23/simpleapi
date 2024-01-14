package main 
import (
	"github.com/labstack/echo/v4"
	"simplefitapi/handlers"
	"simplefitapi/dbconnection"
)

func main() {
	 e := echo.New()

	 // create our first routes
	 e.GET("/", handlers.Home)

	 // connect to the db 
	 dbconnection.InitDb()
	 e.Logger.Fatal(e.Start(":8085"))
}