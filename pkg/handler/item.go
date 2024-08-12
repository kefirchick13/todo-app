package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kefirchick13/todo-app"
)

func (h *Handler) createItem(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	
	if(err != nil){
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input) ; err != nil{
		 NewErrorResponse(c, http.StatusInternalServerError, "invalid JSON")
		 return
	}
	id, err := h.services.TodoItem.CreateItem(userId, listId, input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) getAllItem(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	
	if(err != nil){
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	items, err := h.services.TodoItem.GetAllItems(userId, listId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)



}
func (h *Handler) getItemById(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	item, err := h.services.TodoItem.GetItemById(userId, itemId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)




}
func (h *Handler) updateItem(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	var input todo.UpdateItemInput

	if err := c.BindJSON(&input); err != nil{
	NewErrorResponse(c, http.StatusBadRequest, err.Error())
	return	
	}

	err = h.services.TodoItem.UpdateItemById(userId, itemId, input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, statusOk{
		Status: "ok",
	})
	

}
func (h *Handler) deleteItemById(c *gin.Context) {
	userId,err :=GetUserId(c)
	if err != nil{
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	err = h.services.TodoItem.DeleteItemById(userId, itemId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusOk{
		Status: "ok",
	})


}
