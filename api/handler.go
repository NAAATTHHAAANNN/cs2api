package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LandingPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func GetAllSkins(c *gin.Context) {
	skins, err := getAllSkinsJson()
	if err != nil {
		log.Printf("Error fetching skins: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't serve skins!"})
		return
	}

	c.JSON(http.StatusOK, skins)
}
func GetSkinById(c *gin.Context) {
	idParam := c.Query("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id parameter is required!"})
		return
	}

	Id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id parameter!"})
		return
	}

	skin, err := getSkinByIdJson(Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't find skin with that id!"})
		return
	}

	c.JSON(http.StatusOK, skin)
}

func GetSkinByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name parameter is required!"})
		return
	}

	skin, err := getSkinByNameJson(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't find skin with that name!"})
		return
	}
	c.JSON(http.StatusOK, skin)
}
func GetCollectionByName(c *gin.Context) {

}
func GetCollections(c *gin.Context) {

}
