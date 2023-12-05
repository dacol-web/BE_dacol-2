package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Hy-Iam-Noval/dacol-2/src/DB"
	"github.com/stretchr/testify/assert"
)

// Get /auth/product
func TestFailAdd(t *testing.T) {
	req := POST("/auth/product_add", nil)
	res, err := R.Test(req)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 200, res.StatusCode, "msg : "+string(readBody(res)))
}

func TestAdd(t *testing.T) {
	datas, _ := json.Marshal([]DB.Product{newProduct})
	req := POST("/auth/product_add", datas)
	_, err := R.Test(req)
	if err != nil {
		panic(err)
	}

	data := new(DB.Product)
	DB.
		Select("product", "*", fmt.Sprintf(`name = "%s"`, newProduct.Name)).
		Single().
		Scan(&data.Id, &data.Id_user, &data.Name, &data.Qty, &data.Price, &data.Img, &data.Descript)

	assert.NotEmpty(t, data)
}

func TestFind(t *testing.T) {
	var (
		data, newData DB.Product
	)
	json.Unmarshal(*datas, &data)

	DB.
		Select("product", "*", fmt.Sprintf(`name = "%s"`, data.Name)).
		Single().
		Scan(&newData.Id, &newData.Id_user, &newData.Name, &newData.Img, &newData.Qty, &newData.Price, &newData.Descript)

	req := GET(fmt.Sprintf("/auth/product/%d/", newData.Id))
	res, _ := R.Test(req)

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, data.Name, newData.Name)
	assert.Equal(t, data.Img, newData.Img)
}

func TestSuccessDelete(t *testing.T) {
	var (
		data DB.Product
	)
	json.Unmarshal(*datas, &data)

	DB.
		Select("product", "*", fmt.Sprintf(`name = "%s"`, data.Name)).
		Single().
		Scan(&data.Id, &data.Id_user, &data.Name, &data.Img, &data.Qty, &data.Price, &data.Descript)

	assert.Equal(t, DB.Product{}, data)
}
