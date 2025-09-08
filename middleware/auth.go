package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/utils"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		auth_token := c.Cookies("auth_token")
		if auth_token == "" {
			return c.Status(400).JSON(map[string]string{
				"error": "User not found!",
			})
		}

		jwtToken, err := utils.GetJwtToken(auth_token)
		if err != nil {
			fmt.Println(err)
			return fiber.ErrInternalServerError
		}

		if !jwtToken.Valid {
			return c.Status(400).JSON(map[string]string{
				"error": "Invalid Token",
			})
		}

		err = c.Next()
		if err != nil {
			return err
		}
		return nil
	}
}
