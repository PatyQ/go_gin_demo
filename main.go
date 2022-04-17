package main

import "go_gin_demo/router"

func init() {
	router.Init()
}

func main() {

	router.Run()
}
