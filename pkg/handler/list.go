package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kefirchick13/todo-app"
)


type statusOk struct{
	Status string `json:"status"`
}



func (h *Handler) createList(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}
	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil{
	NewErrorResponse(c, http.StatusBadRequest, err.Error())
	return	
	}
	listId,err :=h.services.TodoList.Create(userId, input)
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"listId": listId,
	})
}


func (h *Handler) getAllLists(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}
	lists, err := h.services.TodoList.GetAll(userId) 
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, lists)

}



func (h *Handler) getListById(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	
	if(err != nil){
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetListById(userId, listId)
	
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)

}
func (h *Handler) updateList(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	
	if(err != nil){
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil{
	NewErrorResponse(c, http.StatusBadRequest, err.Error())
	return	
	}

	err = h.services.UpdateListById(userId,listId, input)

	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, statusOk{
		Status: "ok",
	})


}
func (h *Handler) deleteListById(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	
	if(err != nil){
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.DeleteListById(userId, listId)

	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, "Лист не найден")
		return
	}
	c.JSON(http.StatusOK, statusOk{
		Status: "ok",
	})

}
