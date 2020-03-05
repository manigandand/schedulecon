package store

import (
	"fmt"
	"hotstar/types"
	"time"
)

// Slots ... N day wise
var Slots = make(map[time.Time]*types.Slots, 0)

// Track wise
// var Slots = make(map[string]types.Talk, 0)

// Init ...
func Init() {
	time := time.Now().Local()
	// config

	Slots[time] = &types.Slots{
		Morning: &types.SlotDetails{
			Slots:         make([]*types.Talk, 0),
			TotalMins:     4 * 60,
			AvailableMins: 4 * 60,
			StartTime:     9,
			EndTime:       13,
		},
		LunchBreak: &types.SlotBreak{
			StartTime: 13,
			EndTime:   14,
			Duration:  60,
		},
		Afternoon: &types.SlotDetails{
			Slots:         make([]*types.Talk, 0),
			TotalMins:     2 * 60,
			AvailableMins: 2 * 60,
			StartTime:     14,
			EndTime:       16,
		},
		TeaBreak: &types.SlotBreak{
			StartTime: 16,
			EndTime:   16.30,
			Duration:  30,
		},
		Evening: &types.SlotDetails{
			Slots:         make([]*types.Talk, 0),
			TotalMins:     90,
			AvailableMins: 90,
			StartTime:     16.30,
			EndTime:       18,
		},
	}

	process(input)
}

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
