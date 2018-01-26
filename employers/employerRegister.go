package employers

import (
	"github.com/gin-gonic/gin"
)

// Register registra as apis de empregadores
func Register(engine *gin.Engine) {
	group := engine.Group("/v1/employers")
	{
		group.GET("/healthcheck", HealthCheck)
		group.POST("/signup", SignUp)
	}
}
