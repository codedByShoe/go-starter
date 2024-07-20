package utils

import (
	"encoding/json"
	"io"
)

func JSONResponse(w io.Writer, statusCode int, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

// IsEmptyString checks if a string is empty.
func IsEmptyString(s string) bool {
	return len(s) == 0
}
