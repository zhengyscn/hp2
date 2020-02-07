package controllers

import (
	"encoding/json"
	"errors"
	"hp2/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (bc *BaseController) success(c *gin.Context, data interface{}, message string) {
	c.JSON(200, gin.H{
		"code":    0,
		"data":    data,
		"message": message,
	})
}

func (bc *BaseController) successPaginator(c *gin.Context, data interface{}, count int) {
	c.JSON(200, gin.H{
		"code":    0,
		"data":    data,
		"message": "",
		"count":   count,
	})
}

func (bc *BaseController) fail(c *gin.Context, message string) {
	c.JSON(200, gin.H{
		"code":    -1,
		"data":    nil,
		"message": message,
	})
}

func (bc *BaseController) checkAndResponse(c *gin.Context, err error) {
	if err != nil {
		bc.fail(c, err.Error())
	} else {
		bc.success(c, "", "OK")
	}
}

func (bc *BaseController) GetBodyData(c *gin.Context) (map[string]interface{}, error) {
	m := make(map[string]interface{}, 5)
	cCp := c.Copy()
	rawDataByte, _ := cCp.GetRawData()
	if len(rawDataByte) == 0 {
		return m, errors.New("body is empty")
	}
	if err := json.Unmarshal(rawDataByte, &m); err != nil {
		return m, err
	}
	return m, nil
}

func (bc *BaseController) GetPaginator(c *gin.Context) (int, int, interface{}) {
	cCp := c.Copy()
	page, _ := strconv.Atoi(cCp.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(cCp.DefaultQuery("page_size", "15"))

	uriParams := cCp.Request.URL.String() // /api/v1/users?name=xxx&password=123&page=2&page_size=3
	if strings.Contains(uriParams, "?") {
		mWhereAll := utils.SplitUriStrToMapTmp(uriParams, "?", "&", "=")
		//单条件全局查询
		if mWhere, ok := mWhereAll["default"]; ok {
			// 保留1个
			return page, pageSize, mWhere
		} else {
			// 多条件查询
			delete(mWhereAll, "page_size")
			delete(mWhereAll, "page")
			delete(mWhereAll, "default")
			return page, pageSize, mWhereAll
		}
	} else {
		return page, pageSize, ""
	}
}
