package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateRequest(request interface{}) error {
	err := validate.Struct(request)
	return err
}

func ParseRequestBody(c *fiber.Ctx, request interface{}) error {
	err := c.BodyParser(request)
	return err
}

func ParseAndValidateRequest(c *fiber.Ctx, request interface{}) error {
	err := ParseRequestBody(c, request)
	if err != nil {
		return fmt.Errorf("error while parsing request: %w", err)
	}

	err = ValidateRequest(request)
	if err != nil {
		return fmt.Errorf("error while validating request: %w", err)
	}

	return nil
}
