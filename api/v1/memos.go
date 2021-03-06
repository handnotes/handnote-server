package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/handnotes/handnote-server/models"
	"github.com/handnotes/handnote-server/pkg/util"
	"net/http"
	"strconv"
)

// UpdateMemoForm 更新备忘录/便笺表单
type UpdateMemoForm struct {
	Title    string `form:"title" json:"title"`
	Content  string `form:"content" json:"content" binding:"required"`
	Archived int    `form:"archived" json:"archived"`
}

// ListMemo 备忘录/便笺列表
func ListMemo(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(*util.Claims)
	memos := models.GetMemoList(user.ID)
	c.JSON(http.StatusOK, gin.H{"data": memos})
}

// UpdateMemo 创建备忘录/便笺
func UpdateMemo(c *gin.Context) {
	var request UpdateMemoForm
	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID错误"})
		return
	}
	user := c.MustGet(gin.AuthUserKey).(util.Claims)
	memo := models.Memo{
		ID:       uint(id),
		UserID:   user.ID,
		Title:    request.Title,
		Content:  request.Content,
		Archived: request.Archived,
	}
	if err := models.SaveMemo(&memo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建备忘录/便笺失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": memo})
}
