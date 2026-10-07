package encoder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/0xys/claude-code-status-line/internal/model"
)

// RateLimits formats the 5-hour and 7-day windows (and the spend limit, if any)
// as "5h 23% ↻2h13m 7d 41% ↻3d4h". It returns "" when no window is present.
func RateLimits(r *model.RateLimits, now time.Time) string {
	if r == nil {
		return ""
	}
	var parts []string
	if w := r.FiveHour; w != nil {
		parts = append(parts, window("5h", w.UsedPercentage, w.ResetsAt, now))
	}
	if w := r.SevenDay; w != nil {
		parts = append(parts, window("7d", w.UsedPercentage, w.ResetsAt, now))
	}
	if s := r.SpendLimit; s != nil {
		label := window("spend", s.UsedPercentage, s.ResetsAt, now)
		if s.UsedUsd != nil && s.LimitUsd != nil {
			label += fmt.Sprintf(" $%.2f/$%.0f", *s.UsedUsd, *s.LimitUsd)
		}
		parts = append(parts, label)
	}
	return strings.Join(parts, " ")
}

func window(name string, pct float64, resetsAt int64, now time.Time) string {
	return Cyan(fmt.Sprintf("%s %.0f%% ↻%s", name, pct, untilReset(resetsAt, now)))
}

// untilReset renders the time left until resetsAt as "3d4h", "2h13m" or "45m".
func untilReset(resetsAt int64, now time.Time) string {
	d := time.Unix(resetsAt, 0).Sub(now)
	if d <= 0 {
		return "now"
	}
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	switch {
	case days > 0:
		return fmt.Sprintf("%dd%dh", days, hours)
	case hours > 0:
		return fmt.Sprintf("%dh%dm", hours, minutes)
	default:
		return fmt.Sprintf("%dm", minutes)
	}
}

// usageSnapshot is what WriteRateLimits stores, so that other tools can read the
// latest usage without an OAuth token.
type usageSnapshot struct {
	UpdatedAt  int64             `json:"updated_at"` // Unix epoch seconds
	SessionID  string            `json:"session_id,omitempty"`
	RateLimits *model.RateLimits `json:"rate_limits"`
}

// WriteRateLimits atomically writes the latest rate limits to path. It does
// nothing when the input carries no rate limits, so a session that has not yet
// received an API response does not erase the previous snapshot.
func WriteRateLimits(path string, input model.Input, now time.Time) error {
	if input.RateLimits == nil {
		return nil
	}
	data, err := json.MarshalIndent(usageSnapshot{
		UpdatedAt:  now.Unix(),
		SessionID:  input.SessionID,
		RateLimits: input.RateLimits,
	}, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".usage-*.json")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.Write(append(data, '\n')); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}
