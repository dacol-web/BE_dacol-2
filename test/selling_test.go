package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/Hy-Iam-Noval/dacol-2/src/DB"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSellingEmpty(t *testing.T) {
	req := GET("/auth/selling")
	res, _ := R.Test(req)

	sellingRes, _ := json.Marshal(fiber.Map{
		"datas": []interface{}{},
	})
	assert.JSONEq(t, string(sellingRes), string(readBody(res)))
}

func TestSellingAdd(t *testing.T) {
	data := getProduct()

	apiReq, _ := json.Marshal([]struct {
		Id  int `json:"id"`
		Qty int `json:"qty"`
	}{{Id: 1, Qty: 3}, {Id: int(data.Id), Qty: 2}})

	dataSelling, _ := json.Marshal(DB.Selling{
		Product_list: string(apiReq),
		Date:         time.Now().GoString(),
	})

	req := POST("/auth/selling_add", dataSelling)
	res, _ := R.Test(req)

	assert.Equal(t, 200, res.StatusCode, string(readBody(res)))
}

func TestAllSellingNotEmpty(t *testing.T) {
	req := GET("/auth/selling/")
	res, _ := R.Test(req)

	assert.NotEqual(t, `{"datas":[]}`, string(readBody(res)), string(readBody(res)))
}

func TestSellingNotEmpty(t *testing.T) {
	data := getProduct()
	req := GET(fmt.Sprintf("/auth/selling/%d", data.Id))
	res, _ := R.Test(req)

	assert.NotEqual(t, `{"datas":[]}`, string(readBody(res)), string(readBody(res)))
}
