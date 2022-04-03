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


// @Summary Book Ticket 
// @Description To book ticket.
// @Tags Ticket
// @Schemes
// @Accept json
// @Produce json
// @Param	ticket	body	models.Ticket  true  "Ticket structure"
// @Success	201  {string} 	success
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /ticket/book [POST]
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


// @Summary Get Ticket
// @Description To get ticket details.
// @Tags Ticket
// @Schemes
// @Accept json
// @Param pnr_number path string true "PNR Number"
// @Produce json
// @Success	200  {object} 	models.Ticket
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /ticket/get/{pnr_number} [GET]
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


// @Summary Get all Ticket details
// @Description To get every ticket detail.
// @Tags Ticket
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	models.Ticket
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /ticket/getall [GET]
func (tc *TicketController) GetAll(ctx *gin.Context) {
	tickets, err := tc.TicketService.GetAll()
	// fmt.Println("tickets: ", tickets, "errors: ", err.Error())
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tickets)
}


// @Summary Update Ticket
// @Description To update ticket details.
// @Tags Ticket
// @Schemes
// @Accept json
// @Param pnr_number path string true "PNR Number"
// @Produce json
// @Success	200  {string} 	success
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /ticket/update/{pnr_number} [PATCH]
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


// @Summary Delete Ticket
// @Description To remove a particular ticket.
// @Tags Ticket
// @Schemes
// @Accept json
// @Param pnr_number path string true "PNR Number"
// @Produce json
// @Success	200  {string} 	success
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /ticket/delete/{pnr_number} [DELETE]
func (tc *TicketController) DeleteTicket(ctx *gin.Context) {
	pnrnumber := ctx.Param("pnr_number")
	err := tc.TicketService.DeleteTicket(&pnrnumber)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}


//Routes
func (tc *TicketController) RegisterTicketRoutes(rg *gin.RouterGroup) {
	ticketroute := rg.Group("")
	ticketroute.POST("/ticket/book", tc.CreateTicket)
	ticketroute.GET("/ticket/get/:pnr_number", tc.GetTicket)
	ticketroute.GET("/ticket/getall", tc.GetAll)
	ticketroute.PATCH("/ticket/update", tc.UpdateTicket)
	ticketroute.DELETE("/ticket/delete/:pnr_number", tc.DeleteTicket)
}
