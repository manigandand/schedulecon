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

// Talk ... talks
type Talk struct {
	ID        uint   `json:"id"`
	Title     string `json:"title,omitempty"`
	Author    string `json:"author,omitempty"`
	Duration  int    `json:"duration,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
}

// SlotDetails ...
type SlotDetails struct {
	Slots         []*Talk `json:"slots"`
	TotalMins     int     `json:"total_mins"`
	AvailableMins int     `json:"available_mins"`
	StartTime     float32 `json:"start_time,omitempty"`
	EndTime       float32 `json:"end_time,omitempty"`
}

// SlotBreak ...
type SlotBreak struct {
	StartTime float32 `json:"start_time,omitempty"`
	EndTime   float32 `json:"end_time,omitempty"`
	Duration  int     `json:"duration"`
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
// check morning first
// then afternoon first
// and then evening first
func (s *Slots) AssignTalk(talk *Talk) error {
	switch {
	case s.Morning.IsFree(talk.Duration):
		s.Morning.AssignTalk(talk)
		return nil
	case s.Afternoon.IsFree(talk.Duration):
		s.Afternoon.AssignTalk(talk)
		return nil
	case s.Evening.IsFree(talk.Duration):
		s.Evening.AssignTalk(talk)
		return nil
	}

	// if s.Morning.IsFree(talk.Duration) {
	// 	s.Morning.AssignTalk(talk)
	// 	return nil
	// }
	// if s.Afternoon.IsFree(talk.Duration) {
	// 	s.Afternoon.AssignTalk(talk)
	// 	return nil
	// }
	// if s.Evening.IsFree(talk.Duration) {
	// 	s.Evening.AssignTalk(talk)
	// 	return nil
	// }

	return errors.New("no slots available")
}

// IsFree ...
func (s *SlotDetails) IsFree(duration int) bool {
	if s.AvailableMins < duration {
		return false
	}

	return true
}

// AssignTalk ...
// TODO: set start time end time
func (s *SlotDetails) AssignTalk(talk *Talk) {
	s.Slots = append(s.Slots, talk)
	s.updateAvailableTime(talk.Duration)
}

// UpdateAvailableTime ...
func (s *SlotDetails) updateAvailableTime(duration int) {
	s.AvailableMins = s.AvailableMins - duration
}
