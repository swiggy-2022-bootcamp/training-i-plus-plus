package controller

import (
	"sync"
	"strconv"
	"context"
	pb "github.com/Udaysonu/SwiggyGoLangProject/pb"
	grpc "google.golang.org/grpc"
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	log "github.com/Udaysonu/SwiggyGoLangProject/config"
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
	id, err:= primitive.ObjectIDFromHex(id_s)
	if err!=nil{
		log.Info.Println(err)
	}
	log.Info.Println("Delete request of Expert Id",id_s)
	return service.Delete(id)
}


func SignUpExpert(username string, skill string, email string, location int)(int, entity.Expert ){
	log.Info.Println("Post Signing up Expert",username,skill,email,location)
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
	log.Info.Println("Get All Experts")
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
	log.Info.Println("Completed the work, Userid: ",a,"Expertid: ",b)
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
	id, _ := primitive.ObjectIDFromHex(id_s)
	log.Info.Println("Booking Employee Userid: ",id_s,"Skill: ",skill)
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
	log.Info.Println("Adding Rating Expertid: ",id_s,"Review: ",review)
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
var wg=sync.WaitGroup{}

func insert(msg string){
	log.Channel<-msg;
	wg.Done()
}
// @Router /expert/getbyskill/{skill} [get]
func GetExperts(ctx *gin.Context) (int,[]entity.Expert) {
	skill := ctx.Param("skill")
	log.Info.Println("Get Experts with Skills: ",skill)
	
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
	log.Info.Println("Direct SignUp with Details: ",expert)
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
//	id, _ := primitive.ObjectIDFromHex(id_s)
	conn,_:=grpc.Dial("localhost:8082",grpc.WithInsecure())
	defer conn.Close()
	c:=pb.NewServiceClient(conn)
	log.Info.Println("Get Expert details by Id: ",id_s);
 	return 200,ExpertService(c,id_s)
}
func ExpertService(c pb.ServiceClient,id string)entity.Expert{
	nameRequest:=pb.ExpertRequest{
		Expertid:&pb.ExpertId{
			Id:id,
		},
	}
	res,_:=c.ExpertService(context.Background(),&nameRequest)
 
	id_s,_:=primitive.ObjectIDFromHex(res.Expert.Id)
	ratingStruct:=[]entity.RatingStruct{}
	for _,val:=range res.Expert.Ratingstruct{
		ratingStruct=append(ratingStruct,entity.RatingStruct{int(val.Rating),val.Review})
	}
	return entity.Expert{id_s,res.Expert.Username,res.Expert.Skill,res.Expert.Email,res.Expert.Isavailable,int(res.Expert.Served),float64(res.Expert.Rating),int(res.Expert.Location),ratingStruct}
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
	log.Info.Println("Filter Experts based on Skill: ",skill, "Rating: ",rating)
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
	log.Info.Println("Get Waiting Requests of Expert: ",id_s)
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
	log.Info.Println("Reject Waiting Request with ExpertId: ",id_s)
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
	id, _ := primitive.ObjectIDFromHex(id_s)
	log.Info.Println("Accept Waiting Request of Expert: ",id_s)
	wg.Add(1)
	message:="Your request has been accepted by userd with ID "+id_s+"\n\n Regards \n OnlineShoppingMart.com"
	go insert(message)
	wg.Wait()
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
	return service.CompletedRequest(id, cost_int)
}

// CompletedRequest godoc
// @Summary Get all the Completed Requests based on expertid
// @Description get string by ID
// @Accept json
// @Produce json
// @Success 200 {object} []string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /expert/services [get]
func GetSkills()(int,[]string) {
	return service.GetSkills()
}
