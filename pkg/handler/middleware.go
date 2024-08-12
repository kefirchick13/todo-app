package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (		
	authorizationHeader = "Authorization"
 	userCtx ="userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header:= c.GetHeader(authorizationHeader)
	if header == ""{
		NewErrorResponse(c, http.StatusUnauthorized, "errorAuthHeader")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2{
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1]) 
	if err != nil{
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	
	c.Set(userCtx, userId)

}

func GetUserId(c *gin.Context)(int, error){
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user is not found" )
		return 0, errors.New("error id is not found")
	}
	idInt, ok := id.(int)
	if !ok{
	NewErrorResponse(c, http.StatusInternalServerError, "userId is not valid type" )
		return 0, errors.New("error id is not found")
	}
	return idInt, nil

}