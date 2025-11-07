package validators

import (
	"errors"
	"strconv"
)

func ParseCountryID(countryStr string) (uint, error) {
	if countryStr == "" {
		return 0, nil
	}
	parsed, err := strconv.ParseUint(countryStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid country id")
	}
	return uint(parsed), nil
}
