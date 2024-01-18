package main

import (
	"GoMon/api"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	// Log error if .env file does not exist
	if err := godotenv.Load("envs/.env"); err != nil {
		fmt.Println("No .env file found")
		panic("No .env file found")
	}

	//Set mode in init function
	gin.SetMode(os.Getenv("GIN_MODE"))
}
func main() {

	//Get the router
	var app *gin.Engine = api.Setup()

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
				fmt.Printf("Server closed under request : %s\n", err)
				//panic(fmt.Sprintf("Server closed under request : %s\n", err))
			} else {
				fmt.Printf("Server closed unexpect : %s\n", err)
				//panic(fmt.Sprintf("Server closed unexpect : %s\n", err))
			}
		}

	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-quit
	fmt.Println("Trying to Shutdown Server !")

	//Context with a timeout of 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Server shutdown
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error while Server Shutdown: %s\n", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		fmt.Println("Timeout :5 sec")
	}

	fmt.Println("Exiting Server !")

}
