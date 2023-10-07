package middleware

import (
	"fmt"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

func RegisterWhiteListingMiddleware(whiteListFilePath string) func(c *fiber.Ctx) error {
	whitelist, err := loadWhitelistFromFile(whiteListFilePath)
	if err != nil {
		fmt.Println("Error loading whitelist: ", err)
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	return func(c *fiber.Ctx) error {
		clientIP := c.IP()
		if len(whitelist) > utils.IntZero && !slices.Contains(whitelist, clientIP) {
			fmt.Println("Blacklisted IP:", clientIP)
			return fiber.ErrForbidden
		}
		err := c.Next()
		return err
	}
}

func loadWhitelistFromFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ips := strings.Split(string(content), "\n")

	var whitelist []string
	for _, ip := range ips {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			whitelist = append(whitelist, ip)
		}
	}

	return whitelist, nil
}
