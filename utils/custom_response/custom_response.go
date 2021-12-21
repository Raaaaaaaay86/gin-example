package custom_response

type CustomResponse map[string]interface{}

func Build(data interface{}, message string) CustomResponse {
	response := make(CustomResponse)
	response["data"] = data
	response["message"] = message
	return response
}
