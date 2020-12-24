package v1

import (
	"qa_go/serializer"
	"strconv"

	v1 "qa_go/service/v1/answer"

	"qa_go/api"

	"github.com/gin-gonic/gin"
)
//添加回答
func AddAnswer(c *gin.Context) {
	// qid 所属问题id
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	// 解析参数
	var service v1.AddAnswerService
	err = c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}

	user := api.CurrentUser(c)
	var uid uint
	if user==nil{
		uid=0
	}else{
		uid=user.ID
	}
	res := service.AddAnswer(user, uint(qid),uid)
	c.JSON(200, res)
}

func FindAnswer(c *gin.Context) {
	// qid 所属问题id
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	// aid 回答id
	aid, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	user := api.CurrentUser(c)
	var uid uint
	if user==nil{
		uid=0
	}else{
		uid=user.ID
	}
	res := v1.FindOneAnswer(uint(qid), uint(aid),uid)
	c.JSON(200, res)
}
