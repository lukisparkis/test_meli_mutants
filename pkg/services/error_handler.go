package services

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

// DBErrorHandler - Helper to handle errors from DB
func DBErrorHandler(objType string, err error) (int, map[string]interface{}) {
	if errors.Is(err, mongo.ErrNilDocument) {
		return fiber.StatusNoContent, fiber.Map{"message": fmt.Sprint(objType, " not found.")}
	}
	// Gets the postgres code error
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" {
			return fiber.StatusUnprocessableEntity, fiber.Map{"message": fmt.Sprint("The ", objType, " already exists.")}
		}
		if pqErr.Code == "23503" {
			return fiber.StatusUnprocessableEntity, fiber.Map{"message": fmt.Sprint(pqErr.Constraint, " not found.")}
		}
	}

	return fiber.StatusInternalServerError, fiber.Map{"message": fmt.Sprint("There was an error on the database. Contact with an administrator. Error: ", err.Error())}
}
