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
)
