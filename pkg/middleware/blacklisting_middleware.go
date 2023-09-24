package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RegisterBlackListingMiddleware(blacklistFilePath string) (func(c *fiber.Ctx) error) {
	blacklist, err := loadBlacklistFromFile(blacklistFilePath)
	if err != nil {
		fmt.Println("Error loading blacklist: ", err)
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	return func(c *fiber.Ctx) error {
		clientIP := c.IP()
		for _, ip := range blacklist {
			if ip == clientIP {
				fmt.Println("Blacklisted IP:", clientIP)
				return fiber.ErrForbidden
			}
		}
		err := c.Next()
		return err
	}
}

func loadBlacklistFromFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ips := strings.Split(string(content), "\n")

	var blacklist []string
	for _, ip := range ips {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			blacklist = append(blacklist, ip)
		}
	}

	return blacklist, nil
}