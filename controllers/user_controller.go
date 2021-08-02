package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiolucena1/database"
	"github.com/sergiolucena1/models"
	"github.com/sergiolucena1/services"
)

func CreateUser(c *gin.Context){
	db := database.GetDatabase()

	var p models.User

	err:= c.ShouldBindJSON(&p)
	if err != nil{
		c.JSON(400, gin.H{
			"error": "Não consigo pegar o JSON: " + err.Error(),
		})
		return
	}
	//antes de salvar no banco temos que encriptar o password
	p.Password = services.SHA256Encoder(p.Password)


	err = db.Create(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não consigo criar o blogposts: " + err.Error(),
		})
	}
	c.Status(204)
}