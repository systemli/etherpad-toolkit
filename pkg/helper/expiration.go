package helper

import (
	"errors"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const DefaultSuffix = "default"

type PadExpiration map[string]time.Duration

// ParsePadExpiration splits a string with format "default:30d,temp:24h,keep:365d" and returns a PadExpiration type.
// The key "default:<duration>" is mandatory in the input string.
func ParsePadExpiration(s string) (PadExpiration, error) {
	exp := make(PadExpiration)

	if s == "" {
		return exp, errors.New("input string is empty")
	}

	for _, str := range strings.Split(s, ",") {
		split := strings.Split(str, ":")
		if len(split) != 2 {
			log.WithField("string", str).Error("string is not valid")
			continue
		}
		duration, err := time.ParseDuration(split[1])
		if err != nil {
			log.WithError(err).WithField("duration", split[1]).Error("unable to parse the duration")
			continue
		}

		exp[split[0]] = duration
	}

	if _, ok := exp[DefaultSuffix]; !ok {
		return exp, errors.New("missing default expiration duration")
	}

	return exp, nil
}

// GetDuration tries to get the Duration by pad name, returns the default duration if no suffix matches.
func (pe *PadExpiration) GetDuration(pad string) time.Duration {
	for suffix, duration := range *pe {
		if strings.HasSuffix(pad, fmt.Sprintf("-%s", suffix)) {
			return -duration
		}
	}

	return -(*pe)[DefaultSuffix]
}

// GroupPadsByExpiration sorts pads for the given expiration and returns a map with string keys and string slices.
func GroupPadsByExpiration(pads []string, expiration PadExpiration) map[string][]string {
	var suffixes []string
	for suffix := range expiration {
		if suffix == DefaultSuffix {
			continue
		}
		suffixes = append(suffixes, suffix)
	}

	return GroupPadsBySuffixes(pads, suffixes)
}

// GroupPadsBySuffixes sorts pads for the given suffixes and returns a map with string keys and string slices.
func GroupPadsBySuffixes(pads, suffixes []string) map[string][]string {
	sorted := make(map[string][]string)

	for _, pad := range pads {
		found := false
		for _, suffix := range suffixes {
			if strings.HasSuffix(pad, fmt.Sprintf("-%s", suffix)) {
				sorted[suffix] = append(sorted[suffix], pad)
				found = true
			}
		}
		if !found {
			sorted[DefaultSuffix] = append(sorted[DefaultSuffix], pad)
		}
	}

	return sorted
}
