package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Hy-Iam-Noval/dacol-2/src/DB"
	"github.com/Hy-Iam-Noval/dacol-2/src/helpers"
	"github.com/stretchr/testify/assert"
)

var newU *[]byte = new([]byte)

func TestRegister(t *testing.T) {
	newUser := DB.User{
		Email:     "o@gmail.com",
		Password:  "21345678",
		Password2: "21345678",
	}
	*newU, _ = json.Marshal(newUser)

	req := POST("/register", *newU)
	R.Test(req)

	success := DB.User{}
	DB.
		Select("user", "*",
			fmt.Sprintf(`email = "%s"`, newUser.Email)).
		Single().
		Scan(&success.Id, &success.Email, &success.Password)

	assert.NotEqual(t, success, DB.User{})
}

func TestFailLogin(t *testing.T) {
	req := POST("/login", nil)
	res, err := R.Test(req)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, helpers.UnAuth, res.StatusCode)
}

func TestLogin(t *testing.T) {
	req := POST("/login", *newU)
	sendJSON(req)
	res, _ := R.Test(req)

	assert.Equal(t, 200, res.StatusCode)
	*token = string(readBody(res))
}

func TestAuth(t *testing.T) {
	req := GET("/auth/user")
	res, _ := R.Test(req)

	assert.Equal(t, 200, res.StatusCode)
}
