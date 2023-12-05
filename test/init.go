package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/Hy-Iam-Noval/dacol-2/src"
	"github.com/Hy-Iam-Noval/dacol-2/src/DB"
)

var (
	token      *string    = new(string)
	R          src.App    = src.Route()
	datas      *[]byte    = new([]byte)
	newProduct DB.Product = DB.Product{
		Name:  "baru wdadw d wd dadw",
		Qty:   10,
		Price: 20,
		Img:   "[]",
	}
)

func sendJSON(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func setUser(req *http.Request) {
	req.Header.Set("user", "Breaker "+*token)
}

func readBody(res *http.Response) []byte {
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return bytes
}

func buffer(b []byte) *bytes.Buffer {
	return bytes.NewBuffer(b)
}

func POST(path string, b []byte) *http.Request {
	req, _ := http.NewRequest("POST", path, buffer(b))
	sendJSON(req)
	setUser(req)
	return req
}

func GET(path string) *http.Request {
	req, _ := http.NewRequest("GET", path, nil)
	setUser(req)
	return req
}

func getProduct() DB.Product {
	var data DB.Product

	DB.
		Select("product", "*", fmt.Sprintf(`name = "%s"`, newProduct.Name)).
		Single().
		Scan(&data.Id, &data.Id_user, &data.Name, &data.Img, &data.Qty, &data.Price, &data.Descript)
	return data
}
