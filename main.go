package main

import (
	"ped/employers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	employers.Register(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
