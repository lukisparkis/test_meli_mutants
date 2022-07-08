package stats

import (
	"github.com/gofiber/fiber/v2"
	"meli_test/internal/pkg/services/stats"
	"meli_test/pkg/services"
)

// GetAllChecksPrevious godoc
// @Summary Get all checks previous
// @Description Get all stats previous DNA checks
// @ID GetAllChecksPrevious
// @Accept  json
// @Produce  json
// @Tags stats
// @Success 200 {object} models.Stat "Return stats"
// @Failure 400 {string} error
// @Failure 500 {string} error
// @Router /stats [get]
func GetAllChecksPrevious(c *fiber.Ctx) error {
	data, err := stats.GetAllChecksPrevious()
	if err != nil {
		status, msg := services.DBErrorHandler("Getting all checks previous", err)
		c.SendStatus(status)
		c.JSON(msg)
		return nil
	}
	c.JSON(data)
	c.SendStatus(fiber.StatusOK)
	return nil
}
