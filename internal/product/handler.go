package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(r *gin.RouterGroup) {
	r.POST("", h.create)
	r.GET("", h.getAll)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

func (h *Handler) create(c *gin.Context) {
	var req Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (h *Handler) getAll(c *gin.Context) {
	categories, _ := h.service.GetAll()
	c.JSON(http.StatusOK, categories)
}

func (h *Handler) getByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = uint(id)
	h.service.Update(&req)
	c.JSON(http.StatusOK, req)
}

func (h *Handler) delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.service.Delete(uint(id))
	c.Status(http.StatusNoContent)
}
