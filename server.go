package main

import (
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Define a route to serve image entries as JSON
	app.Get("/imagelist", func(c *fiber.Ctx) error {
		return c.JSON(imageEntries)
	})

	// Define a route to update images manually
	app.Get("/imageupdate", func(c *fiber.Ctx) error {
		updateImageEntries()
		return c.SendStatus(200)
	})

	// Define a route to serve images
	app.Get("/image/:iid", func(c *fiber.Ctx) error {
		iid, err := c.ParamsInt("iid", -1)
		if err != nil || iid < 0 || iid > len(imageEntries) {
			return c.Status(http.StatusNotFound).SendString("Image not found")
		}
		entry := imageEntries[iid]
		return c.SendFile(filepath.Join(IMAGE_DIR, entry.Folder, entry.Name))
	})

	// Define a route to serve static files
	app.Static("/", STATIC_DIR)

	// Start the web server
	go func() {
		log.Fatal(app.Listen(":3000"))
	}()

	// Update image entries every 1 hour
	go func() {
		for {
			updateImageEntries()
			time.Sleep(time.Hour)
		}
	}()

	// Block the main goroutine
	select {}
}
