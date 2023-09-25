package controllers

import (
	"context"
	"main/backend/configs"
	"main/backend/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var albumCollection = configs.GetCollection(configs.ConnectDB(), "albums")
var validate = validator.New()

func CreateAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var album models.Album
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&album); err != nil {
			c.JSON(http.StatusBadRequest, "Request error"+err.Error())
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&album); validationErr != nil {
			c.JSON(http.StatusBadRequest, "Error, missing field"+validationErr.Error())
			return
		}

		newAlbum := models.Album{
			Id:     primitive.NewObjectID(),
			Title:  album.Title,
			Artist: album.Artist,
			Price:  album.Price,
		}

		//insert the album into the database
		result, insertErr := albumCollection.InsertOne(ctx, newAlbum)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, "Error inserting album"+insertErr.Error())
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"album":  newAlbum,
			"result": result,
		})
	}
}

func GetAnAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		albumID := c.Param("albumId")
		var model models.Album
		defer cancel()

		//find the album by ID
		objID, _ := primitive.ObjectIDFromHex(albumID)
		err := albumCollection.FindOne(
			ctx, bson.M{"id": objID}).Decode(&model)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error retrieving album \n"+err.Error())
			return
		}

		c.JSON(http.StatusOK, model)
	}
}

func UpdateAnAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		albumID := c.Param("albumId")
		var album models.Album
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(albumID)

		//validate the request body
		if err := c.BindJSON(&album); err != nil {
			c.JSON(http.StatusBadRequest, "Request error \n"+err.Error())
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&album); validationErr != nil {
			c.JSON(http.StatusBadRequest, "Error missing field \n"+validationErr.Error())
			return
		}

		//update the album in the database
		update := bson.M{
			"Title":  album.Title,
			"Artist": album.Artist,
			"Price":  album.Price,
		}

		result, err := albumCollection.UpdateOne(ctx,
			bson.M{"id": objId},
			bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error updating album \n"+err.Error())
			return
		}

		var UpdatedAlbum models.Album
		if result.ModifiedCount == 1 {
			err := albumCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&UpdatedAlbum)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Error retrieving album \n"+err.Error())
				return
			}
		}
		c.JSON(http.StatusOK, UpdatedAlbum)
	}
}

func DeleteAnAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		albumId := c.Param("albumId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(albumId)

		result, err := albumCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, "DB Error deleting album \n"+err.Error())
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, "Album with specified ID not found! \n")
			return
		}

		c.JSON(http.StatusOK, "Album "+objId.String()+" successfully deleted! \n")
	}
}

func GetAllAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var albums []models.Album
		defer cancel()

		//find all albums
		results, err := albumCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		//reading from the db more optimally
		defer results.Close(ctx)

		//loop through the results
		for results.Next(ctx) {
			var singleAlbum models.Album
			if err := results.Decode(&singleAlbum); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
			}

			albums = append(albums, singleAlbum)
		}

		c.JSON(http.StatusOK, albums)
	}
}
