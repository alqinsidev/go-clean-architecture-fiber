package middlewares

import (
	"fmt"

	"alqinsidev.io/go-clean-architecture/src/app"
	"github.com/gofiber/fiber/v2"
)

func RouteLogger(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		uri := c.Path()
		method := c.Method()
		apiUri := fmt.Sprintf(`[%s] %s`, method, uri)
		logger := app.GetLogger()

		additionalInfo := make(map[string]interface{})
		additionalInfo["http_method"] = method
		additionalInfo["http_uri"] = uri
		additionalInfo["remote_addr"] = c.IP()
		additionalInfo["host_name"] = c.Hostname()
		logger.LogSuccessWithAdditionalInfo("router", apiUri, "success", additionalInfo)

		return c.Next()
	}
}
