package stringUtils

import (
	"log"
	"strconv"
)

func ParseInt(num string) int {
	parsed, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}
