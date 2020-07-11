package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// M is interface
type M map[string]interface{}

func main() {
	api := os.Getenv("API_KEY")

	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/json", func(ctx echo.Context) error {
		data := M{"Message": "Hello", "Counter": 2, "API": api}
		return ctx.JSON(http.StatusOK, data)
	})

	fmt.Println("API_KEY: ", api)

	// os.Setenv("PORT", "8080")
	r.Start(":" + os.Getenv("PORT"))
}
