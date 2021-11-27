package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{}
	meta.Message = message
	meta.Code = code
	meta.Status = status

	jsonRes := Response{}
	jsonRes.Meta = meta
	jsonRes.Data = data

	return jsonRes

}
func FormatErrorValidation(err error) map[string]interface{} {
	var validationErr []string
	for _, e := range err.(validator.ValidationErrors) {
		validationErr = append(validationErr, e.Error())
	}
	var messageError = map[string]interface{}{
		"errors": validationErr,
	}
	return messageError
}
