package store

import (
	"fmt"
	"hotstar/pkg/errors"
	"hotstar/types"
	"time"
)

// Slots ... N day wise
var Slots = make(map[time.Time]*types.Slots, 0)

// Track wise
// var Slots = make(map[string]types.Talk, 0)

// Init ...
func Init() {
	populateDefaultSlots()

	// Process default inputs
	process(input)
}

// ScheduleSlot assigns the slots for the talk
func ScheduleSlot(talk *types.Talk) (map[time.Time]*types.Slots, *errors.AppError) {
	isSlotAlloted := false
	for _, slot := range Slots {
		err := slot.AssignTalk(talk)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		isSlotAlloted = true
		break
	}
	if !isSlotAlloted {
		return nil, errors.BadRequest("Can't allocate slots for this talk, try agian")
	}

	return Slots, nil
}

// -----------------------------------------------------------------------------
// Defaults
func process(input []*types.Talk) {
	for _, talk := range input {
		for _, slot := range Slots {
			err := slot.AssignTalk(talk)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}
	}
}

var input = []*types.Talk{
	&types.Talk{
		Title:    "Title 1",
		Author:   "Author 1",
		Duration: 60,
	},
	&types.Talk{
		Title:    "Title 2",
		Author:   "Author 2",
		Duration: 120,
	},
	&types.Talk{
		Title:    "Title 3",
		Author:   "Author 3",
		Duration: 180,
	},
	&types.Talk{
		Title:    "Title 4",
		Author:   "Author 4",
		Duration: 60,
	},
	&types.Talk{
		Title:    "Title 5",
		Author:   "Author 5",
		Duration: 30,
	},
	&types.Talk{
		Title:    "Title 6",
		Author:   "Author 6",
		Duration: 10,
	},
}
