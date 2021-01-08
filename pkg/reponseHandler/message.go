package reponseHandler

import "github.com/gofiber/fiber/v2"

type reponseMessage struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Status  string      `json:"status,omitempty"`
}

func ReponseMsg(c *fiber.Ctx, code int, status string, msg string, data interface{}) error {
	reponseMsg := &reponseMessage{
		Status:  status,
		Message: msg,
		Data:    data,
	}
	return c.Status(code).JSON(reponseMsg)
}
