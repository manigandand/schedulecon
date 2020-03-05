package store

import (
	"hotstar/types"
	"time"
)

// Slots ...
var Slots = make(map[string]*types.Slots, 0)

// many day ..

// var Slots = make(map[string]types.Schedular, 0)

// Init ...
func Init() {
	time := time.Now().Local()
	// config

	Slots[time.String()] = &types.Slots{
		Morning: &types.SlotDetails{
			Slots:            make([]*types.Schedular, 0),
			TotalMins:        4 * 60,
			AvailableMins:    4 * 60,
			StartTime:        9,
			EndTime:          13,
			IsSlotsAvailable: true,
		},
		LunchBreak: &types.SlotBreak{
			StartTime: 13,
			EndTime:   14,
		},
		Afternoon: &types.SlotDetails{
			Slots:            make([]*types.Schedular, 0),
			TotalMins:        2 * 60,
			AvailableMins:    2 * 60,
			StartTime:        14,
			EndTime:          16,
			IsSlotsAvailable: true,
		},
		TeaBreak: &types.SlotBreak{
			StartTime: 16,
			EndTime:   16.30,
		},
		Evening: &types.SlotDetails{
			Slots:            make([]*types.Schedular, 0),
			TotalMins:        90,
			AvailableMins:    90,
			StartTime:        16.30,
			EndTime:          18,
			IsSlotsAvailable: true,
		},
	}

	process(input)
}

func process(input []*types.Schedular) {
	for _, talk := range input {

	}
}

var input = []*types.Schedular{
	&types.Schedular{
		Title:    "Title 1",
		Author:   "Author 1",
		Duration: 60,
	},
	&types.Schedular{
		Title:    "Title 2",
		Author:   "Author 2",
		Duration: 120,
	},
	&types.Schedular{
		Title:    "Title 3",
		Author:   "Author 3",
		Duration: 180,
	},
	&types.Schedular{
		Title:    "Title 4",
		Author:   "Author 4",
		Duration: 60,
	},
	&types.Schedular{
		Title:    "Title 5",
		Author:   "Author 5",
		Duration: 30,
	},
	&types.Schedular{
		Title:    "Title 6",
		Author:   "Author 6",
		Duration: 10,
	},
}
