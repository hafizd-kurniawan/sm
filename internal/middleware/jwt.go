package middleware

import (
	"boilerplate/pkg/exception"
	"boilerplate/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RoleAuthMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		init := exception.InitException(c, initData.Conf, initData.Log)

		authorizationHeader := c.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Invalid token format", "", nil)
		}

		accessToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		claims, err := utils.CheckAccessToken(init.Conf, accessToken)
		if err != nil {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Token expired or invalid", "", nil)
		}

		// Ambil role_name dari klaim JWT
		roleName, ok := claims["role_name"].(string)
		if !ok {
			// Jika klaim role_name tidak ada di dalam token
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Role information is missing from token", "", nil)
		}

		// Pemeriksaan apakah peran pengguna ada dalam daftar peran yang diizinkan
		isAllowed := false
		for _, allowedRole := range allowedRoles {
			if roleName == allowedRole {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Forbidden: you don't have the required role", "", nil)
		}

		// Menyimpan info pengguna di locals untuk digunakan oleh handler selanjutnya
		if employeeName, ok := claims["employee_name"].(string); ok {
			c.Locals("employee_name", employeeName)
		}
		c.Locals("role_name", roleName)

		return c.Next()
	}
}
