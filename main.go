package main 
import (
	"github.com/labstack/echo/v4"
	"simplefitapi/handlers"
	"simplefitapi/dbconnection"
)

func main() {
	 e := echo.New()

	  // connect to the db 
		dbconnection.InitDb()

	 // create our first routes
	 e.GET("/", handlers.Home)

	 // Create a User routes
	e.POST("/users", handlers.CreateUser)
	e.POST("/measurement", handlers.CreateMeasurement)

	// create user measurment
	// e.POST("/measurement", handlers.CreateUserMeasurement)
	
	 e.Logger.Fatal(e.Start(":8085"))
}