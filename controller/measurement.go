package controller

import (
	"github.com/Asad2730/anstAPI/connect"
	models "github.com/Asad2730/anstAPI/model"
	"github.com/gin-gonic/gin"
)

func CreateMeasurement(c *gin.Context) {
	var measurement models.Measurement
	err := c.ShouldBindJSON(&measurement)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if err := connect.Db.Create(&measurement).Error; err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(201, measurement)
}

func UpdateMeasurement(c *gin.Context) {
	orderNo := c.Param("orderNo")
	var measurement models.Measurement

	if err := connect.Db.First(&measurement, orderNo).Error; err != nil {
		c.JSON(404, err.Error())
		return
	}

	err := c.ShouldBindJSON(&measurement)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if err := connect.Db.Save(&measurement).Error; err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, "updated!")
}
