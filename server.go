package main

import (
	"fmt"
	"net/http"
	//"runtime"
	"log"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	nw "github.com/labstack/echo/middleware"
)

func okJSON(c *echo.Context, value interface{}) error {
	return c.JSON(http.StatusOK, value)
}

func welcome(c *echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("正在用Go重写...... %v\n", time.Now()))
}

func ping(c *echo.Context) error {
	return okJSON(c, "pong")
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	e := echo.New()

	e.Use(nw.Logger())
	e.Use(nw.Recover())

	e.Get("/", welcome)
	e.Get("/ping", ping)

	//// Start server
	//e.Run(":1323")

	// Get the http.Server
	s := e.Server(":8000")

	// HTTP2 is currently enabled by default in echo.New(). To override TLS handshake errors
	// you will need to override the TLSConfig for the server so it does not attempt to validate
	// the connection using TLS as required by HTTP2
	s.TLSConfig = nil

	// Serve it like a boss
	log.Fatal(gracehttp.Serve(s))
}
