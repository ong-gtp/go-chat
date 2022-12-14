package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ong-gtp/go-chat/pkg/domain/requests"
	"github.com/ong-gtp/go-chat/pkg/errors"

	"github.com/ong-gtp/go-chat/pkg/services"
	"github.com/ong-gtp/go-chat/pkg/utils"
)

type ChatController struct {
	chatService services.ChatService
}

func (c *ChatController) RegisterService(s services.ChatService) {
	c.chatService = s
}

func (c *ChatController) ChatRooms(w http.ResponseWriter, r *http.Request) {

	res, err := c.chatService.ChatRooms()
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}

func (c *ChatController) Create(w http.ResponseWriter, r *http.Request) {
	cP := requests.ChatRoomCreatePayload{}
	err := utils.ParseBody(r, &cP)
	if err != nil {
		utils.ErrResponse(errors.ErrInRequestMarshaling, w)
		return
	}

	res, err := c.chatService.CreateChatRoom(cP.Name)
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}

func (c *ChatController) ChatRoomMessages(w http.ResponseWriter, r *http.Request) {
	cmP := requests.ChatRoomMessagesPayload{}
	err := utils.ParseBody(r, &cmP)
	if err != nil {
		utils.ErrResponse(errors.ErrInRequestMarshaling, w)
		return
	}

	res, err := c.chatService.ChatRoomMessages(cmP.RoomId)
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}
