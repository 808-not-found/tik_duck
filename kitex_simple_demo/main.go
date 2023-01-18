package main

import (
	"log"

	minimal_demo "github.com/808-not-found/tik_duck/kitex_simple_demo/kitex_gen/minimalDemo/addservice"
)

func Sub(a int, b int) int {
	return a - b
}

func main() {
	svr := minimal_demo.NewServer(new(AddServiceImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
