package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ong-gtp/go-chat/pkg/errors"

	"github.com/ong-gtp/go-chat/pkg/domain/requests"
	"github.com/ong-gtp/go-chat/pkg/services"
	"github.com/ong-gtp/go-chat/pkg/utils"
)

type AuthController struct {
	authServ services.AuthService
}

func (a *AuthController) RegisterService(s services.AuthService) {
	a.authServ = s
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	lP := requests.LoginPayload{}
	err := utils.ParseBody(r, &lP)
	if err != nil {
		utils.ErrResponse(errors.ErrInRequestMarshaling, w)
		return
	}

	res, err := a.authServ.Login(lP.Email, lP.Password)
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	res.User.Password = ""
	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}

func (a *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	lP := requests.SignUpPayload{}
	err := utils.ParseBody(r, &lP)
	if err != nil {
		utils.ErrResponse(errors.ErrInRequestMarshaling, w)
		return
	}

	res, err := a.authServ.SignUp(lP.Email, lP.UserName, lP.Password)
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	res.User.Password = ""
	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}
