package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name":"Bruno", "Idade":"37"}`))
	resp, err := c.Post("http://www.coritiba.com.br", "application/json", jsonVar)
	if err != nil {

		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.CopyBuffer(os.Stdout, resp.Body, nil)
	if err != nil {
		return
	}

}
