package types

import (
	"errors"
)

// SchedularReq ... in mintes
type SchedularReq struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Duration int    `json:"duration"`
}

// Schedular ... talks
type Schedular struct {
	Title     string `json:"title,omitempty"`
	Author    string `json:"author,omitempty"`
	Duration  int    `json:"duration,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
}

// SlotDetails ...
type SlotDetails struct {
	Slots            []*Schedular `json:"slots"`
	TotalMins        int          `json:"total_mins"`
	AvailableMins    int          `json:"available_mins"`
	StartTime        int          `json:"start_time,omitempty"`
	EndTime          int          `json:"end_time,omitempty"`
	IsSlotsAvailable bool         `json:"is_slots_available,omitempty"`
}

// SlotBreak ...
type SlotBreak struct {
	StartTime float32 `json:"start_time,omitempty"`
	EndTime   float32 `json:"end_time,omitempty"`
}

// Slots ...
type Slots struct {
	Morning    *SlotDetails `json:"morning,omitempty"`
	LunchBreak *SlotBreak   `json:"lunch_break,omitempty"`
	Afternoon  *SlotDetails `json:"afternoon,omitempty"`
	TeaBreak   *SlotBreak   `json:"tea_break,omitempty"`
	Evening    *SlotDetails `json:"evening,omitempty"`
}

// AssignTalk ...
func (s *Slots) AssignTalk(talk *Schedular) error {
	// check morning first
	// check afternoon first
	// check evening first
	// check morning first

	return errors.New("no slots available")
}

// IsFree ...
func (s *SlotDetails) IsFree(duration int) error {
	if !s.IsSlotsAvailable {
		return errors.New("no slots available")
	}
	// check the map

	return errors.New("no slots available")
}
