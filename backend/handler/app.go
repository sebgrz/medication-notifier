package handler

import "github.com/gin-gonic/gin"

func (h *httpHandler) ListMedications(ctx *gin.Context) {
}

func (h *httpHandler) AddMedication(ctx *gin.Context) {
	var req AddMedicationRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logErrorAndAbort(ctx, "add_medication body err: %s", err)
		return
	}
}

func (h *httpHandler) RemoveMedication(ctx *gin.Context) {
	var req RemoveMedicationRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logErrorAndAbort(ctx, "remove_medication body err: %s", err)
		return
	}
}

func (h *httpHandler) ReplaceMedication(ctx *gin.Context) {
	var req ReplaceMedicationRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logErrorAndAbort(ctx, "replace_medication body err: %s", err)
		return
	}
}

type DayType string

const (
	MONDAY    DayType = "MO"
	TUESDAY   DayType = "TU"
	WEDNESDAY DayType = "WE"
	THURSDAY  DayType = "TH"
	FRIDAY    DayType = "FR"
	SATURDAY  DayType = "SA"
	SUNDAY    DayType = "SU"
)

type TimeOfDayType string

const (
	MORNING TimeOfDayType = "MOR"
	MIDDAY  TimeOfDayType = "MID"
	EVENING TimeOfDayType = "EVE"
)

type AddMedicationRequest struct {
	Name      string        `json:"name"`
	Day       DayType       `json:"day"`
	TimeOfDay TimeOfDayType `json:"time_of_day"`
}

type AddMedicationResponse struct {
	Id string `json:"id"`
}

type RemoveMedicationRequest struct {
	Id string `json:"id"`
}

type ReplaceMedicationRequest struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	Day       DayType       `json:"day"`
	TimeOfDay TimeOfDayType `json:"time_of_day"`
}
