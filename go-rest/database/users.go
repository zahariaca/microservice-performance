package database

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
}

type Response struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsers(ctx *gin.Context) {
	log.Println("Before ping to db")
	err := Db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("After ping to db")

	rows, err := Db.Exec("select * from example")

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(400, "Couldn't get users.")
	} else {
		log.Println(rows)
		ctx.JSON(http.StatusOK, "Got users.")
	}
}

func AddUser(ctx *gin.Context) {
	body := User{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	id, _ := uuid.NewRandom()
	alteredUsername := fmt.Sprintf("%s_%s", body.Username, id)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s_%s", body.Password, body.Password)), 12)
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec("insert into example(username,password) values ($1,$2)",
		alteredUsername,
		hashedPassword)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(400, "Couldn't create the new user.")
	} else {
		ctx.JSON(http.StatusOK, Response{
			Username: alteredUsername,
			Password: fmt.Sprintf("%s", hashedPassword),
		})
	}

}
