package main

import (
	"fmt"
	"go-grabber/news"
	"go-grabber/router"
)

func main() {
	fmt.Println("Hello go!")
	r := router.New()
	a := news.NewArchive()

	go a.Serve()
	r.Run()
}
