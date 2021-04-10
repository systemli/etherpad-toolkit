package helper

import "strings"

// SortPads will put the padIds into a string map organized by their suffixes
func SortPads(padIds []string) map[string][]string {
	sorted := make(map[string][]string)

	for _, pad := range padIds {
		if strings.HasSuffix(pad, "-keep") {
			sorted["keep"] = append(sorted["keep"], pad)
		} else if strings.HasSuffix(pad, "-temp") {
			sorted["temp"] = append(sorted["temp"], pad)
		} else {
			sorted["none"] = append(sorted["none"], pad)
		}
	}

	return sorted
}
