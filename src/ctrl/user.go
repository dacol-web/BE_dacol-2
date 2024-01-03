package ctrl

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/Hy-Iam-Noval/dacol-2/src/DB"
	"github.com/Hy-Iam-Noval/dacol-2/src/helpers"
	"github.com/Hy-Iam-Noval/dacol-2/src/validation"
	"github.com/gofiber/fiber/v2"
)

func Login(c Ctx) error {
	req := new(DB.User)
	user := new(DB.User)

	IsError(c.BodyParser(req))

	DB.
		Select("user", "*", fmt.Sprintf(`email = "%s"`, req.Email)).
		Single().
		Scan(&user.Id, &user.Email, &user.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil || user.Email == "" {
		return c.
			Status(helpers.UnAuth).
			JSON(ErrRes("Email or Password wrong"))
	}

	return c.SendString(helpers.GenerateToken(*user))
}

func Register(c Ctx) error {
	req := new(DB.User)

	IsError(c.BodyParser(req))

	if err := validation.IsValid(*req); err != nil {
		return c.Status(helpers.Invalid).JSON(err)
	}

	passCrypt, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.ErrBadGateway.Code, err.Error())
	}

	DB.
		Create("user", DB.User{Email: req.Email, Password: string(passCrypt)}).
		Exec()

	return nil
}

func ParseToken(c Ctx) error {
	return SendDatas(c, helpers.ParseTokenUser(c.Get(helpers.UserKey)))
}
