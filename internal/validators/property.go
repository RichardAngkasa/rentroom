package validators

import (
	"errors"
	"strconv"
)

// validators/property.go
func ParsePropertyID(vars map[string]string) (uint, error) {
	idStr, ok := vars["id"]
	if !ok {
		return 0, errors.New("missing property id")
	}
	parsed, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid property id")
	}
	return uint(parsed), nil
}
