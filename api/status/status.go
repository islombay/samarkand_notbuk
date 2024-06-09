package status

import (
	"net/http"
)

type Status struct {
	Code    int               `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Count   int64             `json:"count,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
	Error   map[string]string `json:"error,omitempty"`
}

var (
	StatusOk = Status{
		Code:    http.StatusOK,
		Message: "Ok",
	}

	StatusInternal = Status{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}

	StatusBadRequest = Status{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}

	StatusNotFound = Status{
		Code:    http.StatusNotFound,
		Message: "Not found",
	}

	StatusBadPhone = Status{
		Code:    http.StatusBadRequest,
		Message: "Invalid phone number",
		Error: map[string]string{
			"phone_number": "invalid",
		},
	}

	StatusBadID = Status{
		Code:    http.StatusBadRequest,
		Message: "Invalid id",
		Error: map[string]string{
			"id": "invalid",
		},
	}

	StatusAlreadyExists = Status{
		Code:    http.StatusConflict,
		Message: "Already exists",
	}
)

func (s Status) AddError(key, value string) Status {
	s.Error[key] = value
	return s
}

func (s Status) AddData(data interface{}) Status {
	s.Data = data
	return s
}

func (s Status) AddCode(code int) Status {
	s.Code = code
	return s
}

func (s Status) AddCount(count int64) Status {
	s.Count = count
	return s
}
