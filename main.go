package main 
import (
	"github.com/labstack/echo/v4"
	"simplefitapi/handlers"
)

func main() {
	 e := echo.New()

	 // create our first routes
	 e.GET("/", handlers.Home)
	 e.Logger.Fatal(e.Start(":8085"))
}