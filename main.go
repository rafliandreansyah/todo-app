package main

import "todo-app/router"

func main() {
	r := router.Router()
	r.Run(":8080")
}
