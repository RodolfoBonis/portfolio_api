package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio_api/src/database"
	"portfolio_api/src/entities"
	"portfolio_api/src/utils"
)

func PostsHandlers(r *gin.Engine) {
	posts := r.Group("/posts")
	{
		posts.GET("/", getAllPosts)
		posts.GET("/:id", getPostById)
		posts.POST("/", utils.TokenAuthMiddleware(), postCreate)
		posts.PUT("/:id", utils.TokenAuthMiddleware(), postUpdate)
		posts.PATCH("/claps/:id", clapsPost)
		posts.DELETE("/:id", utils.TokenAuthMiddleware(), postDelete)
	}
}

func clapsPost(c *gin.Context) {
	id := c.Params.ByName("id")
	var post entities.Post

	if err := database.Db.Where("id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	c.BindJSON(&post)

	if err := database.Db.Save(post).Error; err != nil {
		c.JSON(http.StatusNotModified, gin.H{"message": err})
	}


	c.JSON(http.StatusOK, post)

}

func postDelete(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var post entities.Post
	database.Db.Where("id = ?", id).Delete(&post)
	data := fmt.Sprintf("id %v/", id)
	database.Db.Exec(
		"UPDATE tags_posts SET deleted_at = now() "+
			"WHERE deleted_at IS NULL AND "+
			"post_id = ?", id)
	database.Db.Exec(
		"UPDATE comments SET deleted_at = now() "+
			"WHERE deleted_at IS NULL AND "+
			"post_id = ?", id)

	c.JSON(200, gin.H{data: "deleted"})
}

func getPostById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var post entities.Post

	var user entities.User

	if err := database.Db.Find(&post, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	} else {
		var tags []entities.Tag
		database.Db.Model(&post).Select("name, email").Association("User").Find(&user)

		database.Db.Raw(
			"SELECT t.* FROM tags_posts tp  "+
				"LEFT JOIN tags t ON t.id = tp.tag_id "+
				"WHERE t.deleted_at IS NULL AND tp.deleted_at IS NULL AND tp.post_id = ?", post.ID).Scan(&tags)

		post.User = &user

		post.Tags = &tags

		c.JSON(http.StatusOK, post)
	}
}

func postUpdate(c *gin.Context) {
	var post entities.Post

	id := c.Params.ByName("id")

	if err := database.Db.Where("id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	c.BindJSON(&post)

	if err := database.Db.Update("like", post.Like).Error; err != nil {
		c.JSON(http.StatusNotModified, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, post)
}

func postCreate(c *gin.Context) {
	var post entities.Post
	var tagPost entities.TagsPosts

	c.BindJSON(&post)

	err := database.Db.Create(&post).Error

	for _, tag := range *post.Tags {
		tagPost.TagID = tag.ID
		tagPost.PostID = post.ID

		database.Db.Create(&tagPost)
	}

	if err != nil {
		c.JSON(http.StatusNotModified, map[string]string{"message": "Not created"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func getAllPosts(c *gin.Context) {
	var posts []entities.Post
	var user entities.User

	if err := database.Db.Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	} else {
		for i, _ := range posts {
			var tags []entities.Tag
			database.Db.Model(&posts[i]).Select("name, email").Association("User").Find(&user)

			database.Db.Raw(
				"SELECT t.* FROM tags_posts tp  "+
					"LEFT JOIN tags t ON t.id = tp.tag_id "+
					"WHERE t.deleted_at IS NULL AND tp.deleted_at IS NULL AND tp.post_id = ?", posts[i].ID).Scan(&tags)

			posts[i].User = &user

			posts[i].Tags = &tags

		}

		c.JSON(http.StatusOK, posts)
	}
}
