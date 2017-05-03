package helpers

import (
	"fmt"
	"strconv"
	"time"
)

type Timestamp struct {
	Natural interface{} `json:"natural"`
	Unix    interface{} `json:"unix"`
}

func ParseDate(s string, formats []string) (Timestamp, error) {
	input, err := strconv.Atoi(s)
	var result Timestamp

	if err != nil {
		// Handle natural date
		// Go throug array of supported formats...
		for _, f := range formats {
			d, _ := time.Parse(f, s)
			// ...return if found matching format
			if !d.IsZero() {
				result.Natural = fmt.Sprintf("%s %d, %d", d.Month().String(), d.Day(), d.Year())
				result.Unix = d.Unix()
				return result, nil
			}
		}
	} else {
		// Handle unix date
		d := time.Unix(int64(input), 0)
		result.Natural = fmt.Sprintf("%s %d, %d", d.Month().String(), d.Day(), d.Year())
		result.Unix = input
	}
	return result, nil
}
