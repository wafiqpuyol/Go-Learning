package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wafiqpuyol/Go-Learning/api/internal/types"
	"github.com/wafiqpuyol/Go-Learning/api/internal/utils/response"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studentDetails := types.Student{}
		if err := json.NewDecoder(r.Body).Decode(&studentDetails); err != nil && errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}

		// input validation
		validate = validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(studentDetails); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}

		// business logic
		response.WriteJson(w, http.StatusOK, studentDetails)
	}
}

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "OK")
	}
}
