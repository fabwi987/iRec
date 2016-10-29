package main

import (
	"log"
	"strconv"

	"encoding/json"

	"net/http"

	"github.com/fabwi987/iRec/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	db models.Datastore
}

func main() {

	db, err := models.NewDatabase("root:trustno1@/test")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	router := gin.Default()
	router.GET("/users", env.GetUsersEndpoint)
	router.GET("/user/:id", env.GetUserEndpoint)
	router.GET("/positions", env.GetPositionsEndpoint)
	router.GET("/position/:id", env.GetPositionEndpoint)
	router.GET("/recommendation", env.GetRecommendationsEndpoint)
	router.GET("/recommendation/:id", env.GetRecommendationEndpoint)
	router.Run(":3000")

}

func (env *Env) GetUsersEndpoint(c *gin.Context) {
	users, err := env.db.GetUsers()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(users)
}

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

func (env *Env) GetPositionsEndpoint(c *gin.Context) {
	users, err := env.db.GetPositions()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(users)
}

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

func (env *Env) GetRecommendationsEndpoint(c *gin.Context) {
	users, err := env.db.GetRecommendations()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	json.NewEncoder(c.Writer).Encode(users)
}

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
