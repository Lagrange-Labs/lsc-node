package utils

import (
	"time"
)

// TimeDuration is a wrapper around time.Duration that allows us to unmarshal in TOML.
type TimeDuration time.Duration

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *TimeDuration) UnmarshalText(text []byte) error {
	parsedDuration, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = TimeDuration(parsedDuration)
	return nil
}
