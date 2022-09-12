package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kierquebs/aranguren-piggery-farm-API/cloud"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/setup"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}

	cloud.FirestoreInit()
	database.CCDBCon()

}

const idleTimeout = 5 * time.Second

func main() {
	app := fiber.New()
	setup.Routes(app)
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Listen from a different goroutine
	port := os.Getenv("API_Port")
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
