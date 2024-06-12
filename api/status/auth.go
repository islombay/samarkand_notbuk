package status

import "net/http"

var (
	StatusUnauthorized = Status{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}

	StatusTokenResponse = Status{
		Code: http.StatusOK,
	}

	StatusForbidden = Status{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
	}

	StatusUserCreated = Status{
		Code:    http.StatusCreated,
		Message: "Created",
	}
)
