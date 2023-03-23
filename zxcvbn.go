package zxcvbn

import (
	"time"

	"github.com/wneessen/zxcvbn-go/match"
	"github.com/wneessen/zxcvbn-go/matching"
	"github.com/wneessen/zxcvbn-go/scoring"
	"github.com/wneessen/zxcvbn-go/utils/math"
)

// PasswordStrength takes a password, userInputs and optional filters and returns a MinEntropyMatch
func PasswordStrength(password string, userInputs []string, filters ...func(match.Matcher) bool) scoring.MinEntropyMatch {
	start := time.Now()
	matches := matching.Omnimatch(password, userInputs, filters...)
	result := scoring.MinimumEntropyMatchSequence(password, matches)
	end := time.Now()

	calcTime := end.Nanosecond() - start.Nanosecond()
	result.CalcTime = zxcvbnmath.Round(float64(calcTime)*time.Nanosecond.Seconds(), .5, 3)
	return result
}
