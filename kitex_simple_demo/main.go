package main

import (
	"fmt"
	"log"

	minimal_demo "github.com/808-not-found/tik_duck/kitex_simple_demo/kitex_gen/minimal_demo/addservice"
)

func main() {
	svr := minimal_demo.NewServer(new(AddServiceImpl))
	fmt.Println("123")
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
