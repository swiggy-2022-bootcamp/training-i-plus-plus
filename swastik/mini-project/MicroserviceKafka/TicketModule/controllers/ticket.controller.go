package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/swastiksahoo153/MicroserviceKafka/TicketModule/models"
	"github.com/swastiksahoo153/MicroserviceKafka/TicketModule/services"
	"fmt"
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
	err := tc.TicketService.CreateTicket(&ticket)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}

func (tc *TicketController) GetTicket(ctx *gin.Context) {
	pnr_number := ctx.Param("pnr_number")
	fmt.Println("pnr Number:: ", pnr_number)
	ticket, err := tc.TicketService.GetTicket(&pnr_number)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  ticket)
}

func (tc *TicketController) GetAll(ctx *gin.Context) {
	tickets, err := tc.TicketService.GetAll()
	// fmt.Println("tickets: ", tickets, "errors: ", err.Error())
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
	err := tc.TicketService.UpdateTicket(&ticket)
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
	ticketroute.GET("/get/:pnr_number", tc.GetTicket)
	ticketroute.GET("/getall", tc.GetAll)
	ticketroute.PATCH("/update", tc.UpdateTicket)
	ticketroute.DELETE("/delete/:pnr_number", tc.DeleteTicket)
}
