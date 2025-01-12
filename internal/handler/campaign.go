package handler

import (
	"net/http"

	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/response"
	"api-service/internal/retriever"
	"api-service/internal/types"

	"github.com/gin-gonic/gin"
)

type ICampaignHandler interface {
	Create(c *gin.Context)
	Query(c *gin.Context)
}

type campaignHandler struct {
	retriever retriever.CampaignRetriever
}

func NewCampaignHandler() ICampaignHandler {
	return &campaignHandler{
		retriever: retriever.NewCampaignRetriever(
			model.GetDb(false),
			&cache.Cache{}),
	}
}

func (h *campaignHandler) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *campaignHandler) Query(c *gin.Context) {
	form := &types.CampaignQueryReqest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// db handle campaign query

	res, err := h.retriever.Query(c, *form)

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "campaign query success",
	}, res)
}
