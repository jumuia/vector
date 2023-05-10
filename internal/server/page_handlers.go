package server

import "github.com/gofiber/fiber/v2"

type PageHandler interface {
	ShowIndexPage(c *fiber.Ctx) error
	ShowAboutPage(c *fiber.Ctx) error
}

func NewPageHandler() PageHandler {
	return &pageHandler{}
}

type pageHandler struct {
}

// ShowAboutPage implements PageHandler
func (*pageHandler) ShowAboutPage(c *fiber.Ctx) error {
	return c.Render("views/about", fiber.Map{
		"Title": "Welcome to vector!",
	})
}

// ShowIndexPage implements PageHandler
func (*pageHandler) ShowIndexPage(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{
		"Title": "Hello, World!",
	})
}
