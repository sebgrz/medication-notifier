package handler

import (
	"medication-notifier/data"
	"medication-notifier/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) ListMedications(ctx *gin.Context) {
	var clientDataAny, exists = ctx.Get(utils.CLIENT_INFO_CONTEXT_CONST)
	if !exists {
		logErrorAndAbort(ctx, "list_medication failed, clientData is empty")
		return
	}
	clientData := clientDataAny.(utils.ClientInfo)
	medicaitonList := h.medicationData.FindByUserId(clientData.Id)
	ctx.JSON(http.StatusOK, medicaitonList)
}

func (h *httpHandler) AddMedication(ctx *gin.Context) {
	var req AddMedicationRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logErrorAndAbort(ctx, "add_medication body err: %s", err)
		return
	}

	var clientDataAny, exists = ctx.Get(utils.CLIENT_INFO_CONTEXT_CONST)
	if !exists {
		logErrorAndAbort(ctx, "add_medication failed, clientData is empty")
		return
	}
	clientData := clientDataAny.(utils.ClientInfo)

	medication := data.Medication{
		UserId:    clientData.Id,
		Name:      req.Name,
		Day:       string(req.Day),
		TimeOfDay: string(req.TimeOfDay),
	}

	if err := h.medicationData.Add(medication); err != nil {
		logErrorAndAbort(ctx, "add_medication store medication failed, err: %s", err)
		return
	}

	ctx.JSON(http.StatusOK, AddMedicationResponse{
		Id: medication.Id,
	})
}

func (h *httpHandler) RemoveMedication(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		logErrorAndAbort(ctx, "remove_medication request require 'id' param")
		return
	}

	h.medicationData.RemoveById(id)

	ctx.Status(http.StatusOK)
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

type ReplaceMedicationRequest struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	Day       DayType       `json:"day"`
	TimeOfDay TimeOfDayType `json:"time_of_day"`
}
