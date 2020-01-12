package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kosmgco/tldr/global"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GetRequest struct {
	Name     string `json:"name" form:"name"`
	Platform string `json:"platform" form:"platform"`
	Language string `json:"language" form:"language"`
}

type GetResponse struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Language string `json:"language"`
	Content  string `json:"content"`
}

func Get(ctx *gin.Context) {
	var req GetRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := global.Config.DB.Get()

	task, err := db.Begin()
	defer task.Commit()
	if err != nil {
		logrus.Errorf("begin task failed. err: %s", err)
		ctx.JSON(http.StatusBadGateway, gin.H{})
		return
	}

	_, err = task.Exec("insert into tldr_hot(`name`, `platform`, `language`, `num`) values (?, ?, ?, 0) on DUPLICATE KEY update `num` = `num` + 1", req.Name, req.Platform, req.Language)
	if err != nil {
		logrus.Errorf("exec err: %s", err)
		if err := task.Rollback(); err != nil {
			logrus.Errorf("roll back failed. err: %s", err)
			ctx.JSON(http.StatusBadGateway, gin.H{})
			return
		}
	}
	row := task.QueryRow("select `content` from `tldr_content` where `name` = ? and `platform` = ? and `language` = ?", req.Name, req.Platform, req.Language)
	var content string
	if err := row.Scan(&content); err != nil {
		logrus.Errorf("query row failed. err: %s", err)
		ctx.JSON(http.StatusBadGateway, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, GetResponse{
		Name:     req.Name,
		Platform: req.Platform,
		Language: req.Language,
		Content:  content,
	})
}
