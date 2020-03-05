package store

import (
	"hotstar/types"
	"time"
)

func populateDefaultSlots() {
	time := time.Now().Local()
	default1 := getDefaultSlots()
	Slots[time] = &default1

	default2 := getDefaultSlots()
	Slots[time.AddDate(0, 0, 1)] = &default2

	default3 := getDefaultSlots()
	Slots[time.AddDate(0, 0, 2)] = &default3
}

func getDefaultSlots() types.Slots {
	return types.Slots{
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
}
