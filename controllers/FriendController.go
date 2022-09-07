package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rizqirenaldy/api-go-gin-gorm/helpers"
	"github.com/rizqirenaldy/api-go-gin-gorm/models"
)

type RequestInput struct {
	Requestor string `json:"requestor" binding:"required"`
	To        string `json:"to" binding:"required"`
}

type Request struct {
	Requestor string `json:"requestor"`
	To        string `json:"to"`
	Status    string `json:"status"`
}

type RequestGetData struct {
	Email string `json:"email" binding:"required"`
}

type RequestFriendList struct {
	Requestor string `json:"requestor"`
	Status    string `json:"status"`
}

type FriendListData struct {
	To string `json:"to"`
}

type FriendListCommonInput struct {
	Friends []string `json:"friends"`
}

type BlockInput struct {
	Requestor string `json:"requestor" binding:"required"`
	Block     string `json:"block" binding:"required"`
}

func RequestFriend(c *gin.Context) {

	var input RequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Payload invalid!"})
		return
	}

	var requestBlock models.Request
	if err := models.DB.Table("requests").Where(&Request{Requestor: input.Requestor, To: input.To, Status: "block"}).Find(&requestBlock).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Blocked"})
		return
	}

	var request models.Request
	if err := models.DB.Table("requests").Where(&Request{Requestor: input.Requestor, To: input.To}).Find(&request).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Duplicate Request"})
		return
	}

	requestCreate := models.Request{Requestor: input.Requestor, To: input.To, Status: "pending"}
	models.DB.Create(&requestCreate)

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func RequestFriendStatus(c *gin.Context) {

	var input RequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Payload invalid!"})
		return
	}

	status := c.Param("status")

	statusList := []string{"accept", "reject"}

	if !helpers.Contains(statusList, status) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Status " + status + " not handled"})
		return
	}

	var request models.Request
	if err := models.DB.Where(&Request{Requestor: input.Requestor, To: input.To}).First(&request).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Data Request empty"})
		return
	}

	updateData := models.DB.Table("requests").Where(&Request{Requestor: input.Requestor, To: input.To}).Update(Request{Requestor: input.Requestor, To: input.To, Status: status})

	if updateData == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func RequestFriendData(c *gin.Context) {

	var input RequestGetData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Payload invalid!"})
		return
	}

	var request []RequestFriendList
	models.DB.Table("requests").Where(&Request{To: input.Email}).Find(&request)

	c.JSON(http.StatusOK, gin.H{"requests": request})
}

func FriendList(c *gin.Context) {
	var input RequestGetData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Payload invalid!"})
		return
	}

	var request []Request
	models.DB.Table("requests").Where(&Request{To: input.Email, Status: "accept"}).Find(&request)

	var friendList []string
	for _, v := range request {
		friendList = append(friendList, v.Requestor)
	}

	c.JSON(http.StatusOK, gin.H{"requests": friendList})
}

func FriendListCommon(c *gin.Context) {
	var input FriendListCommonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Payload invalid!"})
		return
	}

	if helpers.LengthofArray(input.Friends) > 2 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Payload friend only with 2 params"})
		return
	}

	userA := input.Friends[0]
	userB := input.Friends[1]

	var dataUserA []Request
	models.DB.Table("requests").Where(&Request{To: userA, Status: "accept"}).Find(&dataUserA)

	var dataUserB []Request
	models.DB.Table("requests").Where(&Request{To: userB, Status: "accept"}).Find(&dataUserB)

	var arrayA []string
	for _, v := range dataUserA {
		if !helpers.Contains(input.Friends, v.Requestor) {
			arrayA = append(arrayA, v.Requestor)
		}
	}

	var arrayB []string
	for _, v := range dataUserB {
		if !helpers.Contains(input.Friends, v.Requestor) && helpers.Contains(arrayA, v.Requestor) {
			arrayB = append(arrayB, v.Requestor)
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "friends": arrayB, "count": helpers.LengthofArray(arrayB)})
}

func BlockFriend(c *gin.Context) {
	var input BlockInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Payload invalid!"})
		return
	}

	var request models.Request
	if err := models.DB.Where(&Request{Requestor: input.Requestor, To: input.Block}).First(&request).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Data is empty"})
		return
	}

	models.DB.Table("requests").Where(&Request{Requestor: input.Requestor, To: input.Block}).Update(Request{Status: "block"})

	c.JSON(http.StatusOK, gin.H{"success": true})
}
