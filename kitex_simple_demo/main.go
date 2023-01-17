package main

import (
	minimal_demo "github.com/808-not-found/tik_duck/kitex_simple_demo/kitex_gen/minimal_demo/addservice"
	"log"
)

func main() {
	svr := minimal_demo.NewServer(new(AddServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
