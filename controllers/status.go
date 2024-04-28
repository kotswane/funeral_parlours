package controllers

import (
	"funeral_parlour/services/status"
	"github.com/gin-gonic/gin"
)

type Status struct {
	svc status.StatusService
}

func NewStatusController(svc status.StatusService) Status {
	return Status{svc: svc}
}

func (s *Status) AddStatus(c *gin.Context) {

}

func (s *Status) UpdateStatus(c *gin.Context) {

}

func (s *Status) FindAllStatus(c *gin.Context) {

}

func (s *Status) DeleteStatus(c *gin.Context) {

}

func (s *Status) FindStatus(c *gin.Context) {

}
