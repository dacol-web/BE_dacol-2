package ctrl

import (
	"github.com/gofiber/fiber/v2"
)

const (
	UserKey = "user"
)

type (
	Ctx = *fiber.Ctx

	DataReq[T interface{}] struct {
		Data T `json:"data"`
	}
)

func Acceptable(c Ctx) error {
	// accept json
	c.Accepts("application/json")
	c.Accepts("multipart/form-data")

	return c.Next()
}

func FiberBadReq(err error) error {
	return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
}

func ErrRes(data interface{}) fiber.Map {
	return fiber.Map{"errors": data}
}

func SendDatas(c Ctx, data interface{}) error {
	return c.JSON(fiber.Map{"datas": data})
}

func IsError(err error) {
	if err != nil {
		panic(err)
	}
}
