package api

import (
	"hotstar/pkg/errors"
	"hotstar/pkg/respond"
	"hotstar/store"
	"hotstar/types"
	"hotstar/utils"
	"net/http"
)

func getAllScheduledTalksHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	respond.OK(w, store.Slots)
	return nil
}

// creates a new city
func scheduleTalkHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	var input types.TalkReq

	if err := utils.Decode(r, &input); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}

	talks, err := store.ScheduleSlot(&types.Talk{
		Title:    input.Title,
		Author:   input.Author,
		Duration: input.Duration,
	})
	if err != nil {
		return err
	}

	respond.Created(w, talks)
	return nil
}
