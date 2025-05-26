package misc

import "strconv"

const (
	KB float64 = 1024
	MB float64 = KB * 1024
	GB float64 = MB * 1024
	TB float64 = GB * 1024
)

// Formats a number of bytes into a string. E.g., 1234567 would become 
// "1.23 MB".
func ByteCountToString(byteCount int) string {
	quantity := float64(byteCount)
	units := "B"
	if quantity > TB {
		quantity /= TB
		units = "TB"
	} else if quantity > GB {
		quantity /= GB
		units = "GB"
	} else if quantity > MB {
		quantity /= MB
		units = "MB"
	} else if quantity > KB {
		quantity /= KB
		units = "KB"
	}

	if units == "B" {
		return strconv.FormatInt(int64(byteCount), 10) + " " + units
	}

	return strconv.FormatFloat(quantity, 'f', 2, 64) + " " + units
}
