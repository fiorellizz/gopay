package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/fiorellizz/gopay/internal/application"
)

type OrderHandler struct {
	createOrderUseCase *application.CreateOrderUseCase
	getOrderByIDUseCase *application.GetOrderByIDUseCase
	listOrdersUseCase   *application.ListOrdersUseCase
}

func NewOrderHandler(
	createUC *application.CreateOrderUseCase,
	getByIDUC *application.GetOrderByIDUseCase,
	listUC *application.ListOrdersUseCase,
) *OrderHandler {
	return &OrderHandler{
		createOrderUseCase: createUC,
		getOrderByIDUseCase: getByIDUC,
		listOrdersUseCase:   listUC,
	}
}

type createOrderRequest struct {
	Amount float64 `json:"amount"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {

	var req createOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	output, err := h.createOrderUseCase.Execute(
		c.Request.Context(),
		application.CreateOrderInput{
			Amount: req.Amount,
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {

	orders, err := h.listOrdersUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	order, err := h.getOrderByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "order not found",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}