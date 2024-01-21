package main

import (
	route "GoMon/api"
	"GoMon/middlewares"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logFile *os.File
var logger = logrus.New()

func init() {

	logFile, err := os.OpenFile("GoMon.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatal("Error handling log File: ", err)
	}

	logger.SetOutput(logFile)
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Log error if .env file does not exist
	if err := godotenv.Load("envs/.env"); err != nil {
		logger.Fatal("No .env file found")
	}

	//Set mode in init function
	gin.SetMode(os.Getenv("GIN_MODE"))

}
func main() {

	defer logFile.Close()

	//Gin setup
	app := gin.Default()

	//custom middleware
	app.Use(middlewares.AddHeaders())
	// Use a custom middleware to log to the file using Logrus
	app.Use(middlewares.LoggerWithLogrus(logger))

	//Get the router
	route.Setup(app, logger)

	//hppt server Config
	server := &http.Server{
		Addr:    ":" + os.Getenv("SERVERPORT"),
		Handler: app,
	}

	//Channel to capture interrupt
	quit := make(chan os.Signal)

	//Notify the quit channel
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//another Goroutine to handle interrupt and shutdown gracefully
	go func() {

		//Starting Server
		fmt.Printf("Starting Server on : %s\n", os.Getenv("SERVERPORT"))
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logger.Error("Server closed under request : ", err)
			} else {
				logger.Error("Server closed unexpect  : ", err)
			}
		}

	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-quit
	logger.Warn("Trying to Shutdown Server !")

	//Context with a timeout of 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Server shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("Error while Server Shutdown: %s\n", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logger.Info("Timeout :5 sec")
	}

	logger.Info("Exiting Server !")

}
