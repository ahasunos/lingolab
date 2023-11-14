package controllers

import (
	"log"
	"net/http"

	"github.com/ahasunos/lingolab/go/ginandgorm/initializers"
	"github.com/ahasunos/lingolab/go/ginandgorm/models"
	"github.com/gin-gonic/gin"
)

func GetCharacters(c *gin.Context) {
	var chars []models.Character

	// func (*gorm.DB).Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	// Find finds all records matching given conditions conds
	initializers.DB.Find(&chars)
	c.JSON(http.StatusOK, chars)
}

func AddCharacters(c *gin.Context) {
	// character := models.Character{
	// 	Name:     "Michael Scott",
	// 	RealName: "Steve Carell",
	// 	Role:     "Regional Manager",
	// }

	var character models.Character

	c.Bind(&character)

	result := initializers.DB.Create(&character)

	// func (*gorm.DB).Create(value interface{}) (tx *gorm.DB)
	// Create inserts value, returning the inserted data's primary key in value's id
	if result.Error != nil {
		log.Fatal("Unable to create character in DB")
		c.Status(http.StatusInternalServerError)
		return
	}

	// type H map[string]any
	// func (gin.H).MarshalXML(e *xml.Encoder, start xml.StartElement) error
	// H is a shortcut for map[string]any
	c.JSON(http.StatusOK, gin.H{
		"character": character,
	})
}
