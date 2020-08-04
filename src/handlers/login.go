package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"portfolio_api/src/database"
	"portfolio_api/src/entities"
	"portfolio_api/src/utils"
)

func LoginHandlers(r *gin.Engine) {
	loginRoute := r.Group("/login")
	{
		loginRoute.POST("/", login)
	}

	r.POST("/logout", Logout)
}

func login(c *gin.Context) {
	var user entities.User
	var socialmedia []*entities.SocialMedia

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	var password = user.Password

	if err := database.Db.Select("id, name, email, password, created_at").Where("email = ?", user.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, "Email or password is incorrect")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(*password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, "Email or password is incorrect")
		return
	}

	database.Db.Select("id, name, url, icon").Find(&socialmedia, "user_id = ?", user.ID)
	user.SocialMedia = socialmedia
	user.Password = nil

	token, err := utils.CreateToken(user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := utils.CreateAuth(user, token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)

}

func Logout(c *gin.Context) {
	au, err := utils.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := utils.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
