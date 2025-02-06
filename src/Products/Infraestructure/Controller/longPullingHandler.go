package controller

import (
	"net/http"
	"strconv"
	applicationnegocio "tienda/src/Products/Application-negocio"
	"time"

	"github.com/gin-gonic/gin"
)

type LongPollingHandler struct {
	GetNewProductsUseCase *applicationnegocio.GetNewProductsUseCase
}

func NewLongPollingHandler(getNewProductsUseCase *applicationnegocio.GetNewProductsUseCase) *LongPollingHandler {
	return &LongPollingHandler{GetNewProductsUseCase: getNewProductsUseCase}
}

// Long Polling Handler
func (h *LongPollingHandler) HandleLong(c *gin.Context) {
	lastProductID, err := strconv.Atoi(c.Query("lastProductID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lastProductID"})
		return
	}

	timeout := time.After(30 * time.Second)
	ticker := time.Tick(5 * time.Second)

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusNoContent, gin.H{"message": "No new products"})
			return
		case <-ticker:
			products, latestID, err := h.GetNewProductsUseCase.Execute(lastProductID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if len(products) > 0 {
				c.JSON(http.StatusOK, gin.H{"products": products, "latestID": latestID})
				return
			}
		}
	}
}
