package http

import "github.com/gofiber/fiber/v2"

func Render(c *fiber.Ctx, httpStatusCode int, data interface{}, meta interface{}) error {
	body := responseBuilder(httpStatusCode, data, meta)

	return c.Status(httpStatusCode).JSON(body)
}

func RenderError(c *fiber.Ctx, httpStatusCode int, err error) error {
	body := responseErrorBuilder(httpStatusCode, err)

	return c.Status(httpStatusCode).JSON(body)
}

func responseErrorBuilder(httpStatusCode int, err error) fiber.Map {
	return fiber.Map{
		"status":  false,
		"message": err.Error(),
		"code":    httpStatusCode,
	}
}

func responseBuilder(httpStatusCode int, data interface{}, meta interface{}) fiber.Map {
	return fiber.Map{
		"status":  true,
		"message": "success",
		"code":    httpStatusCode,
		"data":    data,
		"meta":    meta,
	}
}
