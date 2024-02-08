package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joseph-beck/chit-chat/api/internal/database"
	"github.com/joseph-beck/chit-chat/api/internal/models"
)

func ListPost(s *database.Store) fiber.Handler {
	type Response struct {
		Posts      []models.Post `json:"posts"`
		PostsCount int           `json:"postsCount"`
	}

	return func(c *fiber.Ctx) error {
		var p []models.Post
		err := s.DB.Table("posts").Preload("Author").Find(&p).Error
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		r := Response{
			Posts:      p,
			PostsCount: len(p),
		}

		return c.Status(http.StatusOK).JSON(r)
	}
}

func GetPost(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		b, err := c.ParamsInt("blog")
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}
		a, err := c.ParamsInt("author")
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		var p models.Post
		err = s.DB.Table("posts").Preload("Author").Preload("Blog").Find(&p).Where("blog_refer = ?", b).Where("author_refer = ?", a).Error
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		return c.Status(http.StatusOK).JSON(p)
	}
}
