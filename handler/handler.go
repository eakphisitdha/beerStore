package handler

import (
	"beerstore/model"
	"beerstore/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Get(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Handler struct {
	s service.IService
}

func NewHandler(s service.IService) IHandler {
	return &Handler{s: s}
}

func (h *Handler) Get(c *gin.Context) {
	var req model.GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	beer := h.s.GetBeers(req)
	if beer == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found the " + req.Name})
		return
	}
	c.JSON(http.StatusOK, beer)
}

func (h *Handler) Add(c *gin.Context) {
	var req model.AddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.AddBeer(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully"})

}

func (h *Handler) Update(c *gin.Context) {
	var req model.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	req.ID = id
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.s.UpdateBeer(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully"})
}

func (h *Handler) Delete(c *gin.Context) {
	var req model.DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.s.DeleteBeer(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully"})

}
