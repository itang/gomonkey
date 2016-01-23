package main

import (
	"fmt"
	"net/http"
	//"runtime"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	nw "github.com/labstack/echo/middleware"
)

var count = 0

func welcome(c *echo.Context) error {
	count++
	return c.String(http.StatusOK, fmt.Sprintf("[%v]Hello, World! %v\n", count, time.Now()))
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	e := echo.New()

	e.Use(nw.Logger())
	e.Use(nw.Recover())

	e.Get("/", welcome)

	//// Start server
	//e.Run(":1323")

	// Get the http.Server
	s := e.Server(":3000")

	// HTTP2 is currently enabled by default in echo.New(). To override TLS handshake errors
	// you will need to override the TLSConfig for the server so it does not attempt to validate
	// the connection using TLS as required by HTTP2
	s.TLSConfig = nil

	// Serve it like a boss
	gracehttp.Serve(s)
}
