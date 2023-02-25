package controllers

import (
	"fiber_gorm/initializers"
	"fiber_gorm/models"

	"github.com/gofiber/fiber/v2"
)

func PostCreate(c *fiber.Ctx) error {

	var body = new(models.Post)
	c.BodyParser(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		return c.SendStatus(400)

	}

	return c.Status(200).JSON(&fiber.Map{
		"post": post,
	})
}

func PostRead(c *fiber.Ctx) error {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(&fiber.Map{
		"posts": posts,
	})

}

func PostReadOne(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(&fiber.Map{
		"post": post,
	})
}

func PostUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(models.Post)
	c.BodyParser(&body)
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		return c.SendStatus(400)
	}
	result = initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	if result.Error != nil {
		return c.SendStatus(400)
	}
	return c.Status(200).JSON(&fiber.Map{
		"post": post,
	})
}

// Delete is a function to delete post from database
// @Summary Delete from db
// @Description Delete from id
// @Tags post
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /post/{id} [delete]
func PostDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	result := initializers.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		return c.SendStatus(400)
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
	})

}
