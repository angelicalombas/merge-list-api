package controllers

import (
	"github.com/angelicalombas/merge-list-api/models"
	"github.com/gin-gonic/gin"
)

type InputLists struct {
	List1 []int `json:"list1"`
	List2 []int `json:"list2"`
}

var savedLists []*models.ListNode

// SaveLists godoc
// @Summary Save lists
// @Description Save two lists provided in the request body
// @Tags Lists
// @Accept json
// @Produce json
// @Param lists body controllers.InputLists true "Lists to save"
// @Success 200 {string} string "Lists saved successfully."
// @Failure 400 {string} string "Unable to read the lists provided."
// @Router /SaveLists [post]
func SaveLists(c *gin.Context) {
	var inputLists InputLists
	if err := c.BindJSON(&inputLists); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if inputLists.List1 == nil || inputLists.List2 == nil {
		c.JSON(400, gin.H{"error": "Missing lists in the request."})
		return
	}

	list1 := models.ConvertToLinkedList(inputLists.List1)
	list2 := models.ConvertToLinkedList(inputLists.List2)

	savedLists = []*models.ListNode{list1, list2}
	c.JSON(200, gin.H{"message": "Lists saved successfully."})
}

// Merge godoc
// @Summary Merge saved lists
// @Description Merge saved lists
// @Tags Lists
// @Produce json
// @Success 200 {array} int "Merged list"
// @Failure 400 {string} string "You must save at least two lists to merge."
// @Router /Merge [get]
func Merge(c *gin.Context) {
	if len(savedLists) == 0 {
		c.JSON(200, []int{})
		return
	} else if len(savedLists) == 1 {
		merged := models.ConvertLinkedListToArray(savedLists[0])
		c.JSON(200, merged)
		return
	}

	merged := models.MergeTwoLists(savedLists[0], savedLists[1])
	result := models.ConvertLinkedListToArray(merged)

	if len(result) == 0 {
		c.JSON(200, []int{})
		return
	}

	c.JSON(200, result)
}
