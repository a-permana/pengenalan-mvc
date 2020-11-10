package main

import (
	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-MVC-Golang-Concept/app/controller"
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian", controller.PageAntrianHandler)
	router.Run(":8080")
}



/* package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type Antrian struct {
	Id     string `json:"id"`
	Status bool   `json:"status"`
}

//TODO: delete after implement db
var data []Antrian

var client *db.Client
var ctx context.Context

func init() {
	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://mvc-plus-golang.firebaseio.com",
	}

	opt := option.WithCredentialsFile("firebase-admin-sdk.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/", getSomething)
	router.POST("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/", GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrian", DeleteAntrianHandler)
	router.Run(":8000")
}

func getSomething(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"body1": "Get Something Success",
	})

	return
}

func AddAntrianHandler(c *gin.Context) {
	flag, err := addAntrian()

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "failed",
		})
		return
	}

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "failed",
		})
		return
	}
}

func addAntrian() (bool, error) {
	_, _, dataAntrian := getAntrian()
	var Id string
	var antrianRef *db.Ref
	ref := client.NewRef("antrian")

	if dataAntrian == nil {
		Id = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	} else {
		Id = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d", len(dataAntrian)))
	}
	antrian := Antrian{
		Id:     Id,
		Status: false,
	}
	if err := antrianRef.Set(ctx, antrian); err != nil {
		log.Fatalln(err)
		return false, err
	}

	return true, nil

}

func GetAntrianHandler(c *gin.Context) {
	flag, err, resp := getAntrian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   resp,
		})

		return
	} else {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "unknown Error",
		})
		return
	}
}

func getAntrian() (bool, error, []map[string]interface{}) {
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	err := ref.Get(ctx, &data)
	if err != nil {
		return false, err, nil
	}
	return true, nil, data
}

func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	err := updateAntrian(idAntrian)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func updateAntrian(idAntrian string) error {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	antrian := Antrian{
		Id:     idAntrian,
		Status: true,
	}

	err := childRef.Set(ctx, antrian)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func DeleteAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("isAntrian")
	err := deleteAntrian(idAntrian)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func deleteAntrian(idAntrian string) error {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])

	err := childRef.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
