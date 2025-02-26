package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/DMAproject25/auth-service/internal/entity"
	"github.com/DMAproject25/auth-service/internal/usecase"
	"github.com/DMAproject25/auth-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	u usecase.Users
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, logger logger.Interface, userUseCase usecase.Users) {
	r := userRoutes{userUseCase, logger}

	handler.GET("/users", r.doGetAllUsers)
	handler.GET("/users/:id", r.doGetUserByID)
	handler.POST("/user", r.doSaveUser)
}

// @Summary     Get all toos
// @Description Get all todos
// @ID          get-all-todos
// @Tags  	    todos
// @Accept      json
// @Success     200
// @Failure     500
// @Produce     json
// @Router      /todos/ [get]
func (t *userRoutes) doGetAllUsers(ctx *gin.Context) {
	users, err := t.u.Users(ctx.Request.Context())

	if err != nil {
		t.l.Error(err, "http - v1 - doGetAllAdmins")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

// @Summary     Get todo by id
// @Description Get todo by id
// @ID          get-todo-by-id
// @Tags  	    todos
// @Param       id   path      int  true  "Todo ID"
// @Accept      json
// @Success     200
// @Failure     400
// @Failure     404
// @Failure     500
// @Produce     json
// @Router      /todos/{id} [get]
// @Security    BearerAuth
func (t *userRoutes) doGetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")

		return
	}

	user, err := t.u.UserByID(ctx.Request.Context(), uint64(userID))

	if errors.Is(err, entity.ErrTodoNotFound) {
		errorResponse(ctx, http.StatusNotFound, "todo not found")

		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}

type doSaveUserRequest struct {
	User string `json:"email" binding:"required"`
}

// @Summary     Create todo
// @Description Create todo
// @ID          create-todo
// @Tags  	    todos
// @Param 			request body doSaveTodoRequest true "query params"
// @Accept      json
// @Success     200
// @Failure     500
// @Produce     json
// @Router      /todos [post]
func (t *userRoutes) doSaveUser(ctx *gin.Context) {
	var request doSaveUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		t.l.Error(err, "http - v1 - doSaveTodo")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")

		return
	}

	if err := t.u.SaveUser(ctx.Request.Context(), request.User); err != nil {
		t.l.Error(err, "http - v1 - doSaveTodo")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")

		return
	}

	ctx.JSON(http.StatusOK, nil)
}
