package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"status": "ok"})
	})

	app.Get("/hc", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"status": "hc ok"})
	})

	app.Listen(":9999")
}
