package test

import (
	"testing"

	"github.com/Hy-Iam-Noval/dacol-2/src/DB"
	"github.com/stretchr/testify/assert"
)

func TestQ(t *testing.T) {
	data := new(DB.Product)
	DB.
		Select("product", "*", "id = "+"40").
		Single().
		Scan(&data.Id, &data.Id_user, &data.Name, &data.Qty, &data.Price, &data.Img, &data.Descript)
	assert.NotEqual(t, DB.Product{}, data)
}
