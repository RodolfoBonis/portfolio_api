package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"portfolio_api/src/database"
	"portfolio_api/src/entities"
)

func UsersHandlers(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", fetchUsers)
		users.POST("/", createUser)
		users.PUT("/:id", updateUser)
	}
}

func createUser(c *gin.Context) {
	var user entities.User

	c.Bind(&user)

	dp := database.Db.Where("email = ?", user.Email).First(&user).RowsAffected

	if dp > 0 {
		c.JSON(http.StatusConflict, map[string]interface{}{"message": "user already exists"})
		return
	}

	err := database.Db.FirstOrCreate(&user, &entities.User{Email: user.Email}).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func fetchUsers(c *gin.Context) {
	var users []*entities.User
	var socialmedia []*entities.SocialMedia
	if err := database.Db.Select("id, name, email, created_at").Find(&users).Order("created_at ASC").Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	for i, _ := range users {
		database.Db.Select("id, name, url, icon").Find(&socialmedia, "user_id = ?", users[i].ID)
		users[i].SocialMedia = socialmedia
	}

	c.JSON(http.StatusOK, users)
}

func updateUser(c *gin.Context) {
	var user entities.User
	var socialmedia entities.SocialMedia
	var listSocialmedia []*entities.SocialMedia

	id := c.Params.ByName("id")

	if err := database.Db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.BindJSON(&user)

	for i, _ := range user.SocialMedia {
		socialmedia = *user.SocialMedia[i]

		sId, _ := uuid.FromString(id)
		socialmedia.UserID = &sId

		if socialmedia.ID == nil {
			sId := uuid.NewV4()
			socialmedia.ID = &sId

			database.Db.Create(&socialmedia)
		}

		database.Db.Save(&socialmedia)

	}

	user.SocialMedia = nil

	database.Db.Save(&user)

	database.Db.Select("id, name, email, created_at").First(&user)

	database.Db.Select("id, name, url, icon").Find(&listSocialmedia, "user_id = ?", user.ID)
	user.SocialMedia = listSocialmedia

	c.JSON(http.StatusOK, user)

}
