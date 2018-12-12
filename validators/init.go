package validators

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

func Init() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("validateTitle", ValidateTitle); err != nil {
			return err
		}
	}
	return nil
}
