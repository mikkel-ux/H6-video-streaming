package routes

import (
	scalargo "github.com/bdpiprava/scalar-go"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/api.yaml", func(c *gin.Context) {
		c.File("./api/api.yaml")
	})

	r.GET("/docs", func(c *gin.Context) {
		html, err := scalargo.NewV2(
			scalargo.WithSpecURL("/api.yaml"),
		)
		if err != nil {
			c.String(500, "Error generating Swagger UI: %v", err)
			return
		}
		c.Data(200, "text/html; charset=utf-8", []byte(html))
	})
}
