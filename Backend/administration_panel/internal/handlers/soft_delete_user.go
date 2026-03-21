package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
	"github.com/gin-gonic/gin"
)

func (h *handlers) DeleteUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)

	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user id"})
		return
	}

	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err = h.service.DeleteUser(c.Request.Context(), &models.DeleteUserRequest{
		Token:      tokenString,
		ToDeleteId: userID,
	})
	if err != nil {
		switch {
		case errors.Is(err, custom_err.ErrTokenInvalid):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
