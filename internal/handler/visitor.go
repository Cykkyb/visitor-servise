package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
	"time"
	"visitor/internal/entity"
)

func (h *Handler) addUsersHandler(c *gin.Context) {
	var (
		err  error
		user entity.User
	)

	h.LogRequest("add user", c)

	defer func() {
		c.Next()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusCreated, user)
	}()

	if err = c.BindJSON(&user); err != nil {
		h.log.Error(err.Error())
		return
	}

	h.log.Info("parsed user",
		slog.Any("data", user),
	)
	if err = h.services.Visitor.CreateUser(&user); err != nil {
		h.log.Error(err.Error())
		return
	}
	h.log.Info("add user",
		slog.Any("data", user),
	)
}

type updateUserRequest struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
}

func (h *Handler) updateUserHandler(c *gin.Context) {
	var (
		err  error
		req  updateUserRequest
		user entity.User
	)

	h.LogRequest("update user", c)

	defer func() {
		c.Next()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, user)
	}()

	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		err = fmt.Errorf("invalid user id")
		return
	}

	if err = c.BindJSON(&req); err != nil {
		h.log.Error(err.Error())
		err = fmt.Errorf("invalid request body")
		return
	}

	user, err = h.services.GetUser(id)
	if err != nil {
		h.log.Error(err.Error())
		err = fmt.Errorf("user not found")
		return
	}

	user.Name = req.Name
	user.Surname = req.Surname
	user.Email = req.Email
	user.Phone = req.Phone
	user.Code = req.CountryCode

	if err = h.services.Visitor.UpdateUser(user); err != nil {
		h.log.Error(err.Error())
		err = fmt.Errorf("can't update user")
		return
	}

	h.log.Info("update user", user)
}

func (h *Handler) deleteUserHandler(c *gin.Context) {
	var (
		err error
	)

	h.LogRequest("delete user", c)

	defer func() {
		c.Next()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})
	}()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		err = fmt.Errorf("invalid user id")
		return
	}

	if err = h.services.Visitor.DeleteUser(id); err != nil {
		h.log.Error(err.Error())
		err = fmt.Errorf("can't delete user")
		return
	}

	h.log.Info("delete user",
		slog.Int("id", id),
	)
}

func (h *Handler) getUserHandler(c *gin.Context) {
	var (
		err   error
		users entity.User
	)

	h.LogRequest("get user", c)

	defer func() {
		c.Next()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, users)
	}()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id == 0 {
		h.log.Error("id parameter is required")
		err = fmt.Errorf("id parameter is required")
		return
	}

	users, err = h.services.Visitor.GetUser(id)
	if err != nil {
		h.log.Error(err.Error())
		err = fmt.Errorf("can't get user")
		return
	}
}

func (h *Handler) LogRequest(message string, c *gin.Context) {
	h.log.Info("Request:"+message,
		slog.String("ip", c.ClientIP()),
		slog.String("time", time.Now().Format("02-01-2006 15:04:05")),
		slog.String("method", c.Request.Method),
	)
}
