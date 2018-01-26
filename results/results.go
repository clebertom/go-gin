package results

// BaseResult - retorna ok
type BaseResult struct {
	Status  bool
	Code    int
	Message string
	Data    interface{}
}

type baseResultMessage struct {
	Message string
}

// OkResult retorna uma estrutura com o resultado Ok
func OkResult(data interface{}) (result BaseResult) {
	return BaseResult{Status: true, Code: 200, Data: data}
}

// BadRequestResult retorna o resultado BadRequest
func BadRequestResult(message string) (result BaseResult) {
	return BaseResult{Status: false, Code: 400, Message: message, Data: baseResultMessage{Message: message}}
}

// NotFoundResult ...
func NotFoundResult(message string) (result BaseResult) {
	return BaseResult{Status: false, Code: 404, Message: message, Data: baseResultMessage{Message: message}}
}

// // WriteResult escreve o resultado no contexto
// func (result *BaseResult) WriteResult(c *gin.Context) {
// 	if result.Status {
// 		c.JSON(result.Code, result.Data)
// 	} else {
// 		c.JSON(result.Code, gin.H{"message": result.Message})
// 	}
// 	return
// }
