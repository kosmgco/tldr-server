package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kosmgco/tldr/global"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HotResponse struct {
	Data []HotResponseData `json:"data"`
}

type HotResponseData struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Language string `json:"language"`
}

func Hot(ctx *gin.Context) {
	db := global.Config.DB.Get()
	rows, err := db.Query(`select name, platform, language from tldr_hot order by rand() limit 10`)
	if err != nil {
		logrus.Errorf("query failed. err: %s", err)
		ctx.JSON(http.StatusBadGateway, gin.H{})
		return
	}

	var (
		name, platform, language string
	)

	data := []HotResponseData{}
	for rows.Next() {
		if err := rows.Scan(&name, &platform, &language); err != nil {
			logrus.Errorf("scan failed. err: %s", err)
			ctx.JSON(http.StatusBadGateway, gin.H{})
			return
		}
		data = append(data, HotResponseData{
			Name:     name,
			Platform: platform,
			Language: language,
		})
	}
	ctx.JSON(http.StatusOK, HotResponse{Data: data})
}
