package rest

import (
	grpc_client "Clean-sweep-solution/internal/api-gateway/controller/grpc"
	"Clean-sweep-solution/internal/genproto/product/pb"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	productClient grpc_client.Client
}

func NewHandler(client grpc_client.Client) *Handler {
	return &Handler{
		productClient: client,
	}
}

// @Summary Create a new product
// @Security ApiKeyAuth
// @Description Create a new product with the provided JSON data
// @Tags products
// @Accept json
// @Produce json
// @Param product body ProductRequest true "Product object that needs to be created"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	reqBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var product pb.ProductRequest
	if err = json.Unmarshal(reqBytes, &product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = h.productClient.CreateProduct(ctx, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}
