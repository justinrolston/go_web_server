package main

import (
  "os"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Index("public/index.html")
  e.Static("/", "public/")
	e.Static("/js", "public/js")
  e.Static("/css", "public/css")

	e.Run(":"+os.Getenv("PORT"))
}
