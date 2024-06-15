package main

import (
	"github.com/arielamaral/myidm/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
