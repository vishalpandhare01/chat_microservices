package funcation

import (
	"strconv"
)

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// Handle error as needed, maybe return a default value
		return 0 // default fallback value if parsing fails
	}
	return i
}

func Pagination(pageStr string, limitStr string) (int, int) {
	// Parse page and limit (parse as int, not int64)
	pageInt := parseInt(pageStr)
	limitInt := parseInt(limitStr)

	// Add default values if page or limit is missing or invalid
	if pageInt < 1 {
		pageInt = 1
	}
	if limitInt < 1 {
		limitInt = 4
	}

	// Calculate offset
	offset := (pageInt - 1) * limitInt
	return offset, limitInt
}
