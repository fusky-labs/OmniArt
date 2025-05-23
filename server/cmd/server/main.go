package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func customHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Made-With", "Your mom lol")
		return next(c)
	}
}

func main() {
	// Load env
	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		log.Fatal("Error getting current working directory:", cwdErr)
	}

	envPath := filepath.Join(cwd, "..", "..", ".env")
	envErr := godotenv.Load(envPath)
	if envErr != nil {
		log.Fatal("Something fucked up when loading that .env file")
	}

	// customAPIKey := os.Getenv("OM_SECRET_API_TOKEN")
	// allowedCORSUrl := os.Getenv("OM_SERVER_CORS_ALLOWED_DOMAINS")

	// Main server
	e := echo.New()

	e.Use(
		// CORS stuff
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}),
		// Timeout stuff
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: 25 * time.Second,
		}),
		// Redirect trailing slash, cuz those are freaking cringe
		middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
			RedirectCode: http.StatusPermanentRedirect,
		}),
		customHeader,
	)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "It's me hi I'm the problem it's me",
		})
	})

	// `/blacklist`
	e.GET("/blacklist", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string][]string{
			"blacklist": {"There's supposed to be a list of blacklist of websites and stuff, but I'm lazy lol"},
		})
	})

	e.POST("/blacklist", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "This is a POST request to add to the blacklist",
		})
	})

	// Search stuff

	e.Logger.Fatal(e.Start(":4000"))
}
