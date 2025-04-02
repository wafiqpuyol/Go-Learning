package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/wafiqpuyol/Go-Learning/api/internal/types"
)

func WriteJson(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(data)
}

func ValidationError(errs validator.ValidationErrors) types.Response {
	var errorMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errorMsgs = append(errorMsgs, fmt.Sprintf("%s is required", err.Field()))
		default:
			errorMsgs = append(errorMsgs, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}
	return types.Response{
		Status: "error",
		Error:  strings.Join(errorMsgs, ", "),
	}
}
