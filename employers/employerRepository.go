package employers

import "ped/results"

// DoSignUp efetua o registro do usuário no banco de dados
func DoSignUp(request SignUpRequest) (result results.BaseResult) {
	result = results.OkResult(SignUpResponse{Name: "Crebs"})
	result = results.NotFoundResult("Registro não encontrado!")
	return
}
