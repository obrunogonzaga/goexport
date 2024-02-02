package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Microsecond)
	defer cancel()

	res, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://www.coritiba.com.br", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
