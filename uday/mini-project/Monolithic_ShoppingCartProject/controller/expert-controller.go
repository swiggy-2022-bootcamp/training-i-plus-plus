package controller

import (
	"fmt"
	"strconv"

	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete godoc
// @Summary Shows all the expert accounts
// @Description get string by ID
// @Security ApiKeyAuth
// @Produce json
// @Param expertid path string true "userid"
// @Success 200 {object} bool
// @Failure 400 {object} bool
// @Failure 500 {object} bool
// @Router /expert/delete/{expertid} [delete]
func Delete(ctx *gin.Context) int {
	id_s:= ctx.Param("expertid")
	id, _ := primitive.ObjectIDFromHex(id_s)
	return service.Delete(id)
}


func SignUpExpert(username string, skill string, email string, location int)(int, entity.Expert ){
	return service.SignUpExpert(username, skill, email, location)
}

// GetAllExperts godoc
// @Summary Shows all the expert accounts
// @Description get string by ID
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} []entity.Expert
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/getallexperts [get]
func GetAllExperts(ctx *gin.Context) (int,[]entity.Expert){
	return service.GetAllExperts()
}

// WorkDone godoc
// @Summary Update work status of Expert
// @Description Expert completed the work
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param userid path string true "userid"
// @Param expertid path string true "expertid"
// @Success 200 {object} bool
// @Failure 400 {object} bool
// @Failure 500 {object} bool
// @Router /expert/done/{userid}/{expertid} [get]
func WorkDone(ctx *gin.Context)(int,bool) {
	a := ctx.Param("userid")
	b := ctx.Param("expertid")
	a_s, _ := primitive.ObjectIDFromHex(a)
	b_s, _ := primitive.ObjectIDFromHex(b)
	fmt.Println(a_s, b_s)
	return	service.WorkDone(a_s, b_s)
}

// BookEmployee godoc
// @Summary Books an Employee based on skill
// @Description get string by ID
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param skill path string true "Skill"
// @Param userid path string true "User ID"
// @Success 200 {object} entity.Expert
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/get/{skill}/{userid} [get]
func BookEmployee(ctx *gin.Context) (int,entity.Expert) {
	skill := ctx.Param("skill")
	id_s := ctx.Param("userid")
	fmt.Println(skill, id_s)
	id, _ := primitive.ObjectIDFromHex(id_s)
	return service.BookEmployee(skill, id)
}


// AddRating godoc
// @Summary Add Rating
// @Description get string by ID
// @Consume application/x-www-form-urlencoded
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param expertid path string true "Expet ID"
// @Param review body entity.RatingStruct true "Rating and Review"
// @Success 200 {object} bool
// @Failure 400 {object} bool
// @Failure 500 {object} bool
// @Router /expert/addrating/{expertid} [post]
func AddRating(ctx *gin.Context) {
	var review entity.RatingStruct
	ctx.BindJSON(&review)
	id_s := ctx.Param("expertid")
	// id,_ := strconv.ParseInt(id_s,10,8)
	id, _ := primitive.ObjectIDFromHex(id_s)
	service.AddRating(review.Rating, review.Review, id)
}
// GetExperts godoc
// @Summary Get Experts based on skills
// @Description get string by ID
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param skill path string true "Skill"
// @Success 200 {object} []entity.Expert
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/get/{skill} [get]
func GetExperts(ctx *gin.Context) (int,[]entity.Expert) {
	skill := ctx.Param("skill")
	return service.GetExperts(skill)
}

type Temp struct {
	Username string `json:"username"`
	Skill    string `json:"skill"`
	Email    string `json:"email"`
	Location int    `json:"location"`
}

// DirectSignUp godoc
// @Summary Sign Up
// @Description get string by ID
// @Consume application/x-www-form-urlencoded
// @Security ApiKeyAuth
// @Produce json
// @Param expert body Temp true "Expert Details"
// @Success 200 {object} entity.Expert
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/signupexpert [post]
func DirectSignUp(ctx *gin.Context)(int,entity.Expert) {
	var expert Temp;
	ctx.ShouldBindJSON(&expert)
	fmt.Println(expert)
	return service.SignUpExpert(expert.Username, expert.Skill, expert.Email, expert.Location)
}
// GetExpertByID godoc
// @Summary Get Expert by Id
// @Description get string by ID
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param expertid path string true "expertid"
// @Success 200 {object} entity.Expert
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/getexpert/{expertid} [get]
func GetExpertByID(ctx *gin.Context) (int,entity.Expert) {
	id_s:= ctx.Param("expertid")
	id, _ := primitive.ObjectIDFromHex(id_s)
	return service.GetExpertByID(id)
}

// FilterExpert godoc
// @Summary Filters userd based on skill and rating
// @Description get string by ID
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param skill path string true "Skill"
// @Param rating path string true "rating"
// @Success 200 {object} entity.Expert
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/filter/{skill}/{rating} [get]
func FilterExpert(ctx *gin.Context)(int, []entity.Expert) {
	skill := ctx.Param("skill")
	id_s := ctx.Param("rating")
	rating, _ := strconv.Atoi(id_s)
	return service.FilterExpert(skill, rating)
}

// GetWaitingRequest godoc
// @Summary Get the waiting requests of an expert
// @Description get string by ID
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param expertid path string true "Expertid"
// @Success 200 {object} entity.Expert
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/waitingreq/{expertid} [get]
func GetWaitingRequest(ctx *gin.Context)(int, entity.UserExpert ){
	id_s:= ctx.Param("expertid")
	id, _ := primitive.ObjectIDFromHex(id_s)
	return service.GetWaitingRequest(id)
}


// RejectWaitingResult godoc
// @Summary Reject the waiting request of an expert
// @Description get string by ID
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param expertid path string true "Expert Id"
// @Success 200 {object} bool
// @Failure 400 {object} bool
// @Failure 500 {object} bool
// @Router /expert/rejectreq/{expertid} [get]
func RejectWaitingResult(ctx *gin.Context)(int, bool) {
	id_s := ctx.Param("expertid")
	id, _ := primitive.ObjectIDFromHex(id_s)
	return service.RejectWaitingResult(id)
}

// AcceptWaitingRequest godoc
// @Summary Accept the waiting request of an Expert
// @Description get string by ID
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param expertid path string true "Expert Id"
// @Success 200 {object} entity.UserExpert
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/acceptreq/{expertid} [get]
func AcceptWaitingRequest(ctx *gin.Context) (int,entity.UserExpert ){
	id_s:= ctx.Param("expertid")
	fmt.Println(id_s)
	id, _ := primitive.ObjectIDFromHex(id_s)
	return service.AcceptWaitingRequest(id)
}

// CompletedRequest godoc
// @Summary Get all the Completed Requests based on expertid
// @Description get string by ID
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param expertid path string true "Expert Id"
// @Param cost path int true "Cost"
// @Success 200 {object} entity.UserExpert
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/complete/{cost}/{expertid} [get]
func CompletedRequest(ctx *gin.Context)(int, entity.UserExpert) {
	id_s := ctx.Param("expertid")
	id, _ := primitive.ObjectIDFromHex(id_s)

	cost:= ctx.Param("cost")
	cost_int,_:=strconv.Atoi(cost)
 //  cost_int, _ := strconv.ParseInt(cost, 10, 8)
   fmt.Println(cost_int)
	return service.CompletedRequest(id, cost_int)
}

func GetSkills()(int,[]string) {
	return service.GetSkills()
}
