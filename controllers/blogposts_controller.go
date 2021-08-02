package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiolucena1/database"
	"github.com/sergiolucena1/models"
	"strconv"
)

//Primeiro endpoint
func ShowBlogPost(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id) // convertendo pra inteiro
	if err != nil{
		c.JSON(400,gin.H{
			"error": "ID tem que ser inteiro",
		})

		return
	}
	db := database.GetDatabase()

	var p models.BlogPost
	err = db.First(&p, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não consigo encontrar o produto:" + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

//Segundo endpoint
func CreateBlogPost(c *gin.Context){
	db := database.GetDatabase()

	var p models.BlogPost

	err:= c.ShouldBindJSON(&p)
	if err != nil{
		c.JSON(400, gin.H{
			"error": "Não consigo pegar o JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não consigo criar o blogposts: " + err.Error(),
		})
	}
	c.JSON(200, p)
}

//terceiro endpoint
func ShowBlogPosts(c *gin.Context){

	db := database.GetDatabase()

	var p []models.BlogPost
	err := db.Find(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não consigo listar os blogposts: " + err.Error(),
		})
		return
	}
	c.JSON(200, p)

}

//quarto end point
func UpdateBlogPost(c *gin.Context){
	db := database.GetDatabase()

	var p models.BlogPost

	err:= c.ShouldBindJSON(&p)
	if err != nil{
		c.JSON(400, gin.H{
			"error": "Não consigo pegar o JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode atualizar o blogpost: " + err.Error(),
		})
	}
	c.JSON(200, p)
}

//quinto endpoint
func DeleteBlogPost(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id) // convertendo pra inteiro
	if err != nil{
		c.JSON(400,gin.H{
			"error": "ID tem que ser inteiro",
		})

		return
	}
	db := database.GetDatabase()

	err = db.Delete(&models.BlogPost{},newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode deletar o blogpost: " + err.Error(),
		})
		return
	}
	c.Status(204)
}