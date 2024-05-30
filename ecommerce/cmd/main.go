package main

import (
	"ecommerce/app"
)

func main() {
	app := app.NewApplication()
	app.Init()
	app.Run()
	app.Wait()
	app.Cleanup()
}
