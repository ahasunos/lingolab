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

// Payload Example!
// {
// 	"Name": "Kevin Malone",
// 	"RealName": "Brian Baumgartner",
//     "Role": "Accountant"
// }

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

func GetCharacterByID(c *gin.Context) {
	// get id of the url
	id := c.Param("id")

	var character models.Character

	result := initializers.DB.First(&character, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Character not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"character": character,
	})
}

func UpdateCharacters(c *gin.Context) {
	// get id of the url
	id := c.Param("id")

	// fetch the char from the body
	var character, charinDB models.Character
	c.Bind(&character)

	// find the char from the db
	result := initializers.DB.First(&charinDB, id)

	// return not found if not available in db
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Character not found to update",
		})
		return
	} else {
		// update the char

		// syntaxt from gorm docs:
		// Update attributes with `struct`, will only update non-zero fields
		// db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
		// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

		updateResult := initializers.DB.Model(&charinDB).Updates(models.Character{
			Name:     character.Name,
			RealName: character.RealName,
			Role:     character.Role,
		})

		if updateResult.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "unable to update",
			})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"character updated": charinDB,
	})
}

func DeleteCharacterByID(c *gin.Context) {
	// get id of the url
	id := c.Param("id")

	var character models.Character

	result := initializers.DB.Delete(&character, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Character not deleted",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"character deleted": character,
	})
}

func DeleteAllCharacters(c *gin.Context) {

	var chars []models.Character

	// Find finds all characters
	result := initializers.DB.Find(&chars)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch characters for deletion",
		})
		return
	}

	// delete all
	result = initializers.DB.Delete(&chars)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete characters",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All characters deleted successfully",
	})

}
