package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiolucena1/database"
	"github.com/sergiolucena1/models"
	"github.com/sergiolucena1/services"
)

func Login(c *gin.Context){
	db := database.GetDatabase()

	var p models.Login
	err:= c.ShouldBindJSON(&p)
	if err != nil{
		c.JSON(400, gin.H{
			"error": "Não consigo pegar o JSON: " + err.Error(),
		})
		return
	}

	var user models.User
	dbError := db.Where("email = ?", p.Email).First(&user).Error
	if dbError != nil{
		c.JSON(400, gin.H{
			"error": "Não consigo achar o usuário: ",
		})
		return
	}

	//conferir se a senha que recebemos é a mesma que esta no banco

	//se for diferente...
	if user.Password != services.SHA256Encoder(p.Password){
		c.JSON(400, gin.H{
			"error": "As credenciais estão inválidas: ",
		})
		return
	}

	//se for igual e passou desses testes ...
	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil{
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
