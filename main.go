package main

import (
	"log"
	"strconv"

	"encoding/json"

	"net/http"

	"fmt"

	"github.com/fabwi987/iRec/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//Env is an environmental struct
type Env struct {
	db models.Datastore
}

func main() {

	//db, err := models.NewDatabase("root:trustno1@/test")
	db, err := models.NewDatabase("fabwi987:trustno1@tcp(aauzfgep8gdf8x.cyuihstloqfq.eu-central-1.rds.amazonaws.com:3306)/irec01")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	router := gin.Default()
	router.GET("/users", env.GetUsersEndpoint)
	router.GET("/user/:id", env.GetUserEndpoint)
	router.GET("/positions", env.GetPositionsEndpoint)
	router.GET("/position/:id", env.GetPositionEndpoint)
	router.GET("/recommendations", env.GetRecommendationsEndpoint)
	router.GET("/recommendation/:id", env.GetRecommendationEndpoint)
	//router.GET("/recommendations/pos/:id", env.GetPositionRecommendationsEndpoint)
	//router.GET("/recommendations/usr/:id", env.GetUserRecommendationsEndpoint)

	router.POST("/user", env.CreateUserEndpoint)
	router.Run(":8080")

}

//GetUsersEndpoint return all users
func (env *Env) GetUsersEndpoint(c *gin.Context) {

	users, err := env.db.GetUsers()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(users)
}

//GetUserEndpoint return a single user from it's id
func (env *Env) GetUserEndpoint(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	usr, err := env.db.GetUser(intid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(usr)
}

//GetPositionsEndpoint return all positions
func (env *Env) GetPositionsEndpoint(c *gin.Context) {
	users, err := env.db.GetPositions()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(users)
}

//GetPositionEndpoint returns a single position fomr it's id
func (env *Env) GetPositionEndpoint(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	usr, err := env.db.GetPosition(intid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(usr)
}

//GetRecommendationsEndpoint returns all recommendations
func (env *Env) GetRecommendationsEndpoint(c *gin.Context) {
	users, err := env.db.GetRecommendations()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(users)
}

//GetRecommendationEndpoint returns a single recommendation from it's id
func (env *Env) GetRecommendationEndpoint(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	usr, err := env.db.GetRecommendation(intid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(usr)
}

//CreateUserEndpoint creates a user and sends it to the database
func (env *Env) CreateUserEndpoint(c *gin.Context) {
	intid, err := strconv.Atoi(c.PostForm("Type"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	NewUser := models.User{
		Id:    0,
		Type:  intid,
		Name:  c.PostForm("Name"),
		Mail:  c.PostForm("Mail"),
		Phone: c.PostForm("Phone"),
	}

	userid, err := env.db.CreateUser(&NewUser)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	fmt.Println(userid)
}

//CreatePositionEndpoint creates a position and sends it to the database
func (env *Env) CreatePositionEndpoint(c *gin.Context) {
	userid, err := strconv.Atoi(c.PostForm("Userid"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	usr, err := env.db.GetUser(userid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	NewPosition := models.Position{
		Id:     0,
		Userid: usr,
		Title:  c.PostForm("Title"),
		Body:   c.PostForm("Body"),
		Reward: c.PostForm("Reward"),
	}

	err = env.db.CreatePosition(&NewPosition)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

}

/**
//GetPositionRecommendationsEndpoint returns all recommendations from a position id
func (env *Env) GetPositionRecommendationsEndpoint(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	rec, err := env.db.GetPositionRecommendations(intid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(rec)
}

//GetUserRecommendationsEndpoint returns all recommendations from a user id
func (env *Env) GetUserRecommendationsEndpoint(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	rec, err := env.db.GetUserRecommendations(intid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(rec)
}*/
