package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joseph-beck/chit-chat/api/internal/database"
	"github.com/joseph-beck/chit-chat/api/internal/models"
)

func ListBlog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func GetBlog(s *database.Store) fiber.Handler {
	type Response struct {
		Blog  models.Blog
		Posts []models.Post
	}

	return func(c *fiber.Ctx) error {
		b, err := c.ParamsInt("blog")
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		r := Response{}
		r.Blog = models.Blog{Model: models.Model{ID: uint(b)}}

		err = s.Get(&r.Blog, "blogs")
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		err = s.DB.Table("posts").Preload("Author").Find(&r.Posts).Where("blog_refer = ?", b).Error
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		return c.Status(http.StatusOK).JSON(r)
	}
}
