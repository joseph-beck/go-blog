package app

import (
	"log"
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joseph-beck/chit-chat/api/internal/cache"
	"github.com/joseph-beck/chit-chat/api/internal/database"
	"github.com/joseph-beck/chit-chat/api/internal/models"
)

type Config struct {
	Port string
}

func Default() Config {
	return Config{
		Port: ":8080",
	}
}

type App struct {
	Fiber  *fiber.App
	Redis  *cache.Cache
	Store  *database.Store
	Config Config
}

func New(a ...Config) App {
	c := Default()

	if len(a) > 0 {
		c = a[0]
	}

	f := fiber.New()
	r := cache.New()
	s := database.New(database.SQLiteConn())

	return App{
		Fiber:  f,
		Redis:  r,
		Store:  s,
		Config: c,
	}
}

func (a *App) Run() {
	err := a.Store.AutoMigrate(models.User{})
	if err != nil {
		panic(err)
	}

	a.Fiber.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))

	a.Fiber.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/api/v1/swagger.json",
		Path:     "docs",
		Title:    "chit chat docs",
	}))

	a.Fiber.Use(logger.New())

	log.Fatalln(a.Fiber.Listen(a.Config.Port))
}

func (a *App) Close() {

}
