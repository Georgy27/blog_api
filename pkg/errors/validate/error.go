package validate

import (
	"github.com/pkg/errors"
	"strings"
)

type ValidationErrors struct {
	Messages []string `json:"error_messages"`
}

func (v *ValidationErrors) addError(message string) {
	v.Messages = append(v.Messages, message)
}

func NewValidationErrors(messages ...string) *ValidationErrors {
	return &ValidationErrors{
		Messages: messages,
	}
}

func (v *ValidationErrors) Error() string {
	//data, err := json.Marshal(v.Messages)
	//if err != nil {
	//	return err.Error()
	//}

	return strings.Join(v.Messages, ", ")
	//return string(data)
}

func IsValidationError(err error) bool {
	var ve *ValidationErrors
	return errors.As(err, &ve)
}
