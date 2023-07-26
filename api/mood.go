package api

import (
	"fmt"
	"time"
)

type Mood string
type MoodDate int

const (
	Happy  Mood = "Happy"
	Normal Mood = "Normal"
	Sad    Mood = "Sad"

	Week MoodDate = iota
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

func ParseMoodDate(s string) MoodDate {
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
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type FindMood struct {
	Mood Mood     `json:"mood"`
	Date MoodDate `json:"date"`
}

func (m Mood) ValidateMood() error {
	if m.String() == "" {
		return fmt.Errorf("invalid request")
	}
	return nil
}
