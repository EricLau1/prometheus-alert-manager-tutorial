package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prometheus-alert-manager-tutorial/api/httpext"
	"prometheus-alert-manager-tutorial/api/services"
	"prometheus-alert-manager-tutorial/api/types"
	"strconv"
)

type todosHandler struct {
	todosService services.TodosService
}

func RegisterTodos(router *gin.Engine, todosService services.TodosService) {
	h := &todosHandler{todosService: todosService}
	router.POST("/todos", h.PostTodo)
	router.GET("/todos", h.GetTodos)
	router.PUT("/todos/:id", h.PutTodo)
	router.GET("/todos/:id", h.GetTodo)
	router.DELETE("/todos/:id", h.DeleteTodo)
}

// PostTodo     godoc
// @Summary     Store a new Todo
// @Description Takes a Todo JSON and store in DB. Return saved JSON.
// @Tags        Todos
// @Produce     json
// @Param       Todo   body     types.Todo true "Todo JSON"
// @Success     201    {object} types.Todo
// @Failure     422    {object} httpext.JsonError
// @Router      /todos [post]
func (h *todosHandler) PostTodo(ctx *gin.Context) {
	todo := new(types.Todo)
	err := ctx.Bind(todo)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, httpext.JsonError{Error: err.Error()})
		return
	}
	todo, err = h.todosService.CreateTodo(ctx.Request.Context(), todo)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, httpext.JsonError{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, todo)
}

// PutTodo      godoc
// @Summary     Update a Todo
// @Description Takes a Todo JSON and update in DB by id. Return saved JSON.
// @Tags        Todos
// @Produce     json
// @Param       id     path string true "Todo identifier"
// @Param       Todo   body     types.Todo true "Todo JSON"
// @Success     200    {object} types.Todo
// @Failure     400    {object} httpext.JsonError
// @Failure     422    {object} httpext.JsonError
// @Router      /todos [put]
func (h *todosHandler) PutTodo(ctx *gin.Context) {
	todo := new(types.Todo)
	err := ctx.Bind(todo)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, httpext.JsonError{Error: err.Error()})
		return
	}
	todo.ID, err = strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpext.JsonError{Error: err.Error()})
	}
	todo, err = h.todosService.UpdateTodo(ctx.Request.Context(), todo)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, httpext.JsonError{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

// GetTodo      godoc
// @Summary     Get a Todo
// @Description Get a Todo by id.
// @Tags        Todos
// @Produce     json
// @Param       id          path string true "Todo identifier"
// @Success     200         {object} types.Todo
// @Failure     404         {object} httpext.JsonError
// @Failure     422         {object} httpext.JsonError
// @Router      /todos/{id} [get]
func (h *todosHandler) GetTodo(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpext.JsonError{Error: err.Error()})
	}
	todo, err := h.todosService.GetTodo(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, httpext.JsonError{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

// GetTodos     godoc
// @Summary     Get a Todo list
// @Description Get a Todo list.
// @Tags        Todos
// @Produce     json
// @Success     200    {array} types.Todo
// @Failure     500    {object} httpext.JsonError
// @Router      /todos [get]
func (h *todosHandler) GetTodos(ctx *gin.Context) {
	todo, err := h.todosService.GetTodos(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpext.JsonError{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

// DeleteTodo   godoc
// @Summary     Delete a Todo
// @Description Delete a Todo by id.
// @Tags        Todos
// @Produce     json
// @Param       id          path string true "Todo identifier"
// @Success     200         {object} types.Todo
// @Failure     400         {object} httpext.JsonError
// @Failure     404         {object} httpext.JsonError
// @Router      /todos/{id} [delete]
func (h *todosHandler) DeleteTodo(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpext.JsonError{Error: err.Error()})
	}
	err = h.todosService.DeleteTodo(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, httpext.JsonError{Error: err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
