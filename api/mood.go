package api

import (
	"fmt"
	"time"
)

type Mood string
type MoodTimeRange int

const (
	Happy  Mood = "Happy"
	Normal Mood = "Normal"
	Sad    Mood = "Sad"

	Week MoodTimeRange = iota
	Month
	Year
	All
)

func (m Mood) String() string {
	switch m {
	case Happy:
		return "Happy"
	case Normal:
		return "Normal"
	case Sad:
		return "Sad"
	default:
		return ""
	}
}

func ParseMoodTimeRange(s string) MoodTimeRange {
	switch s {
	case "Week":
		return Week
	case "Month":
		return Month
	case "Year":
		return Year
	case "All":
		return All
	default:
		return Month
	}
}

type CreateMood struct {
	ID          int       `json:"id"`
	Mood        Mood      `json:"mood"`
	UserId      int       `json:"user_id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type FindMood struct {
	Mood      Mood
	TimeRange MoodTimeRange
	UserId    *int
}

func (m Mood) ValidateMood() error {
	if m.String() == "" {
		return fmt.Errorf("invalid request")
	}
	return nil
}
