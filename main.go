package main

import (
	"fmt"
	"net/http"
	"adrequestauctionsystem/app"
	auctioneerController "adrequestauctionsystem/app/auctioneer"
	bidderController "adrequestauctionsystem/app/bidder"
	"adrequestauctionsystem/config"
	"adrequestauctionsystem/internal/command"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Ad Request Auction Project")
	})

	e.POST("/create/auction", auctioneerController.HandleCreateAuction)
	e.GET("/all/auctions", auctioneerController.HandleGetAllAuctionsList)

	e.POST("/create/bidder", bidderController.HandleCreateBidder)
	e.GET("/all/bidders", bidderController.HandleGetAllBiddersList)

	e.POST("/bid/round/start", auctioneerController.HandleBidRound)

	// Server
	app.SetCommands()
	command.RunApp()
	appConfig := config.GetConfig()
	fmt.Println("Starting server at ", appConfig.Port)
	e.Run(standard.New(fmt.Sprintf(":%d", appConfig.Port)))
}
