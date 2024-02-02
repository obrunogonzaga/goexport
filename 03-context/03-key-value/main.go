package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "language", "Go")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	if v := ctx.Value("language"); v != nil {
		language := v.(string)
		fmt.Println("Language is set to", language)
		return
	}
	fmt.Println("Language is not set")
}
