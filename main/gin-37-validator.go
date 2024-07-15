package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// Booking 包含绑定和验证的数据。
type Booking struct {
	CheckIn  time.Time `form:"check_in" json:"check_in" binding:"required,bookableDate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" json:"check_out" binding:"required,gtfield=CheckIn,bookableDate" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	engine := gin.Default()
	defer engine.Run()

	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		fmt.Println("RegisterValidation bookableDate")
		err := validate.RegisterValidation("bookableDate", bookableDate)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	engine.GET("/bookable", func(c *gin.Context) {
		var booking Booking
		if err := c.ShouldBind(&booking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": booking})
		}
	})

}
