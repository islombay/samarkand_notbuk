package status

import (
	"net/http"

	"github.com/minio/minio-go/v7"
)

type Status struct {
	Code    int               `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Count   int64             `json:"count,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
	Error   map[string]string `json:"error,omitempty"`

	FileObject     *minio.Object    `json:"-"`
	FileObjectInfo minio.ObjectInfo `json:"-"`
}

// type File struct {
// 	name        string
// 	contentType string
// 	size        int
// }

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
			"phone_number": string(ErrInvalid),
		},
	}

	StatusBadID = Status{
		Code:    http.StatusBadRequest,
		Message: "Invalid id",
		Error: map[string]string{
			"id": string(ErrInvalid),
		},
	}

	StatusAlreadyExists = Status{
		Code:    http.StatusConflict,
		Message: "Already exists",
	}
)

func (s Status) AddError(key string, value StatusError) Status {
	if s.Error == nil {
		s.Error = map[string]string{}
	}
	s.Error[key] = string(value)
	return s
}

func (s Status) AddData(data interface{}) Status {
	s.Data = data
	return s
}

func (s Status) AddDataMap(key, val string) Status {
	if s.Data == nil {
		s.Data = map[string]interface{}{}
	}
	s.Data.(map[string]interface{})[key] = val
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

func (s Status) AddFileObject(obj *minio.Object) Status {
	s.FileObject = obj
	return s
}

func (s Status) AddFileObjectInfo(obj minio.ObjectInfo) Status {
	s.FileObjectInfo = obj
	return s
}

type StatusError string

var (
	ErrInvalid  = StatusError("invalid")
	ErrNotFound = StatusError("not_found")
)
