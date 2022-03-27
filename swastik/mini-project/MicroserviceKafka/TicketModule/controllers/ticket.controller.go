package controllers

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swastiksahoo153/ticket-module/models"
	"github.com/swastiksahoo153/ticket-module/services"
)

type TicketController struct{
	TicketService services.TicketService
}

func New(ticketService services.TicketService) TicketController{
	return TicketController{
		TicketService: ticketService,
	}
}

func (tc *TicketController) CreateTicket(ctx *gin.Context) {
	var ticket models.Ticket
	if err := ctx.ShouldBindJSON(&ticket); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}
	err := tc.TicketService.CreateTicket(&train)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}

func (tc *TicketController) GetTicket(ctx *gin.Context) {
	pnrnumber := ctx.Param("pnr_number")
	ticket, err := uc.TicketService.GetTicket(&pnrnumber)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  ticket)
}

func (tc *TicketController) GetAll(ctx *gin.Context) {
	tickets, err := tc.TicketService.GetAll()
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tickets)
}

func (tc *TicketController) UpdateTicket(ctx *gin.Context) {
	var ticket models.Ticket
	if err := ctx.ShouldBindJSON(&ticket); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}
	err := tc.TicketService.UpdateTicket(&train)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}

func (tc *TicketController) DeleteTicket(ctx *gin.Context) {
	pnrnumber := ctx.Param("pnr_number")
	err := tc.TicketService.DeleteTicket(&pnrnumber)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}

func (tc *TicketController) RegisterTicketRoutes(rg *gin.RouterGroup) {
	ticketroute := rg.Group("/ticket")
	ticketroute.POST("/create", tc.CreateTicket)
	ticketroute.GET("/get/:name", tc.GetTicket)
	ticketroute.GET("/getall", tc.GetAll)
	ticketroute.PATCH("/update", tc.UpdateTicket)
	ticketroute.DELETE("/delete/:name", tc.DeleteTicket)
}
