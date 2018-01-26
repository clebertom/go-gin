package employers

import (
	"github.com/gin-gonic/gin"
)

// SignUp efetua o cadastro do empregador/usuário da plataforma
func SignUp(c *gin.Context) {
	// efetuando o bind
	var request SignUpRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	response := DoSignUp(request)
	c.JSON(response.Code, response.Data)
}

// HealthCheck retorna o status da api de empregadores
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
