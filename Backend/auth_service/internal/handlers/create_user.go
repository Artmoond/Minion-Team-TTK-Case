package handlers

import (
	"errors"
	"net/http"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	res, err := h.s.CreateUser(c.Request.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, custom_err.ErrInvalidArguments):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		case errors.Is(err, custom_err.ErrUserIsExist):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

		case errors.Is(err, custom_err.ErrBuildingQuery):
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		case errors.Is(err, custom_err.ErrCreateUser):
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	c.JSON(http.StatusCreated, res)
}
