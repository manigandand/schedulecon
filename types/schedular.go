package types

import (
	"strings"

	"hotstar/pkg/errors"
)

// TalkReq ... in mintes
type TalkReq struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Duration int    `json:"duration"`
}

// Ok implements the Ok interface, it validates city input
func (t *TalkReq) Ok() error {
	switch {
	case strings.TrimSpace(t.Title) == "":
		return errors.IsRequiredErr("title")
	case strings.TrimSpace(t.Author) == "":
		return errors.IsRequiredErr("author")
	case t.Duration < 10:
		return errors.New("talk duration should not less than 10 mins")
	case t.Duration > 180:
		return errors.New("talk duration should not more than 180 mins")
	}

	return nil
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
	talk.ID = uint(len(s.Slots) + 1)
	s.Slots = append(s.Slots, talk)
	s.updateAvailableTime(talk.Duration)
}

// UpdateAvailableTime ...
func (s *SlotDetails) updateAvailableTime(duration int) {
	s.AvailableMins = s.AvailableMins - duration
}
