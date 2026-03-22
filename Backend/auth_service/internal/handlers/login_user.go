package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/internal/entity/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) LoginUser(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("req: ", req)

	resp, err := h.s.LoginService(c.Request.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, custom_err.ErrInvalidArguments):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errors.Is(err, custom_err.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, custom_err.ErrBuildingQuery):
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, resp)
}
