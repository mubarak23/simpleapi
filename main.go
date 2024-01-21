package main 
import (
	"github.com/labstack/echo/v4"
	"simplefitapi/handlers"
	"simplefitapi/dbconnection"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	 e := echo.New()

	  // connect to the db 
		dbconnection.InitDb()

		// cors handle 
		...
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"*"},
  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
}))
...


		// custom middleware 
		e.Use(handlers.LogRequest)

		// echo middleware
		e.Use(middleware.Logger())

	 // create our first routes
	 e.GET("/", handlers.Home)

	 // Create a User routes
	e.POST("/users", handlers.CreateUser)
	e.POST("/measurement", handlers.CreateMeasurement)
	e.PUT("/users/:id", handlers.HandleUpdateUser)
	e.PUT("/measurement/:id", handlers.HandleUpdateMeasurement)
	// e.POST("/event", handlers.SubmitEvent)

	// create user measurment
	// e.POST("/measurement", handlers.CreateUserMeasurement)
	
	 e.Logger.Fatal(e.Start(":8085"))
}