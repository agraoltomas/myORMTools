package myORMTools

import "encoding/json"

type APIError struct {
	Err     error
	Message string
	Code    int
}

func (e APIError) Error() string {
	return e.Err.Error()
}

func (e APIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		message string `json:"message"`
	}{
		e.Message,
	})
}
