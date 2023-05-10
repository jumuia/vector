package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jumuia/vector"
)

type Config struct {
}

func Run(cfg *Config) error {
	// engine := html.New("./views", ".html")
	engine := html.NewFileSystem(http.FS(vector.ViewsFS), ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "views/layouts/default",
	})

	phdlr := NewPageHandler()

	app.Get("/", phdlr.ShowIndexPage)
	app.Get("/about", phdlr.ShowAboutPage)

	return app.Listen(":3000")
}
