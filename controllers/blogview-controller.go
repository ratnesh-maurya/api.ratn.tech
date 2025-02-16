package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ratnesh-maurya/api.ratn.tech/models"
	"github.com/ratnesh-maurya/api.ratn.tech/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IncrementViews increments the view count for a given blog post
func IncrementViews(blogview *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug") // Extract slug from URL

		if slug == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Slug is required"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Increment view count
		filter := bson.M{"slug": slug}
		update := bson.M{"$inc": bson.M{"views": 1}}
		opts := options.Update().SetUpsert(true)

		_, err := blogview.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ApplicationResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to increment blog view",
			})
			return
		}

		c.JSON(http.StatusOK, responses.ApplicationResponse{
			Status:  http.StatusOK,
			Message: "Blog view incremented successfully",
		})
	}
}


// GetViews retrieves the view count for a given blog post
func GetViews(blogview *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var blog models.BlogView
		err := blogview.FindOne(ctx, bson.M{"slug": slug}).Decode(&blog)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, responses.ApplicationResponse{
					Status:  http.StatusNotFound,
					Message: "Blog not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, responses.ApplicationResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to retrieve blog views",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"slug":  blog.Slug,
			"views": blog.Views,
		})
	}
}
