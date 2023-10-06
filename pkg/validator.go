package validator

import (
	"errors"
	"fmt"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
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

func ParseRequestQuery(c *fiber.Ctx, request interface{}) error {
	err := c.QueryParser(request)
	return err
}

func ParseAndValidateRequest(c *fiber.Ctx, request interface{}, parserType types.ParserType) error {
	var err error
	switch parserType {
	case types.BodyParserType:
		err = ParseRequestBody(c, request)
	case types.QueryParserType:
		err = ParseRequestQuery(c, request)
	default:
		err = errors.New("invalid parser type")
	}
	if err != nil {
		return fmt.Errorf("error while parsing request: %w", err)
	}

	err = ValidateRequest(request)
	if err != nil {
		return fmt.Errorf("error while validating request: %w", err)
	}

	return nil
}
