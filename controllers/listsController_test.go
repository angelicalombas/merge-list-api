package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/angelicalombas/merge-list-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func createTestList(vals []int) *models.ListNode {
	var head, current *models.ListNode
	for _, val := range vals {
		node := &models.ListNode{Val: val}
		if head == nil {
			head = node
			current = node
		} else {
			current.Next = node
			current = node
		}
	}
	return head
}

func TestSaveLists(t *testing.T) {
	router := SetupDasRotasDeTeste()

	router.POST("/SaveLists", SaveLists)

	// Test case: SaveLists - Successful case
	requestBody, _ := json.Marshal(gin.H{"list1": []int{1, 2, 4}, "list2": []int{1, 3, 4}})

	req, _ := http.NewRequest("POST", "/SaveLists", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// Test case: SaveLists - Error in JSON parsing
	invalidJSONReq, _ := http.NewRequest("POST", "/SaveLists", bytes.NewBuffer([]byte("invalid JSON")))
	invalidJSONReq.Header.Set("Content-Type", "application/json")
	invalidJSONResp := httptest.NewRecorder()

	router.ServeHTTP(invalidJSONResp, invalidJSONReq)

	assert.Equal(t, 400, invalidJSONResp.Code)

	// Test case: SaveLists - Missing list1
	requestBody, _ = json.Marshal(gin.H{"list2": []int{1, 3, 4}})

	req, _ = http.NewRequest("POST", "/SaveLists", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)

	// Test case: SaveLists - Missing list2
	requestBody, _ = json.Marshal(gin.H{"list1": []int{1, 2, 4}})

	req, _ = http.NewRequest("POST", "/SaveLists", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)

	// Test case: SaveLists - No lists
	requestBody, _ = json.Marshal(gin.H{})

	req, _ = http.NewRequest("POST", "/SaveLists", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestMerge(t *testing.T) {
	router := SetupDasRotasDeTeste()

	router.GET("/Merge", Merge)

	// Test case: Merge - Successful case
	list1 := createTestList([]int{1, 2, 4})
	list2 := createTestList([]int{1, 3, 4})

	expectedMerged := "[1,1,2,3,4,4]"

	req, _ := http.NewRequest("GET", "/Merge", nil)
	resp := httptest.NewRecorder()

	savedLists = []*models.ListNode{list1, list2}

	router.ServeHTTP(resp, req)

	respostaBody, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, expectedMerged, string(respostaBody))

	// Test case: Merge - Empty lists
	reqEmptyLists, _ := http.NewRequest("GET", "/Merge", nil)
	respEmptyLists := httptest.NewRecorder()

	savedLists = []*models.ListNode{}

	router.ServeHTTP(respEmptyLists, reqEmptyLists)

	respostaBodyEmptyLists, _ := io.ReadAll(respEmptyLists.Body)

	assert.Equal(t, 200, respEmptyLists.Code)
	assert.Equal(t, "[]", string(respostaBodyEmptyLists))

}
