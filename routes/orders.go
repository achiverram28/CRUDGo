package routes

import (
	"context"

	"log"
	"net/http"
	"time"

	"restaurantserver/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var validate = validator.New()

var orderCollection *mongo.Collection = OpenCollection(Client, "orders")

func AddOrder(c *gin.Context){

	var ctx , cancel = context.WithTimeout(context.Background(),100*time.Second)

	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		log.Fatal(err)
		return
	}

	validationErr := validate.Struct(order)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":validationErr.Error(),
		})
		log.Fatal(validationErr)
		return
	}

	order.ID = primitive.NewObjectID()

	result, insertErr := orderCollection.InsertOne(ctx,order)
	if insertErr != nil {
		message := "Could not create order"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":message,
		})
		log.Fatal(insertErr)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)
}

func GetOrders(c *gin.Context){
	var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)

	var orders []bson.M

	cursor, err := orderCollection.Find(ctx,bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		log.Fatal(err)
		return
	}

	if err = cursor.All(ctx,&orders); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		log.Fatal(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK,orders)

}

func DeleteOrder(c *gin.Context){
	 
	orderID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(orderID)

	var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)

	result, err := orderCollection.DeleteOne(ctx,bson.M{"_id":docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		log.Fatal(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.DeletedCount)

}

func UpdateOrder(c *gin.Context){

	orderID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(orderID)

	var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)

	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
       log.Fatal(err)
	   return
	}

	validationErr := validate.Struct(order)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":validationErr.Error(),
		})
		log.Fatal(validationErr)
		return
	}

	result, err := orderCollection.ReplaceOne(ctx,bson.M{"_id":docID},bson.D{
		
		{Key: "category",Value: order.Category},
		{Key: "dish",Value: order.Dish},
		{Key: "quantity",Value: order.Quantity},
		{Key: "tablenumber",Value: order.TableNumber},
		{Key: "servername",Value: order.ServerName},
		{Key: "price",Value: order.Price},

	})

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		log.Fatal(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.ModifiedCount)

}