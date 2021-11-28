package helper

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

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
func EnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value

}
