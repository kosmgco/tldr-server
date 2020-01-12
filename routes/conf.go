package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kosmgco/tldr/database"
	"github.com/kosmgco/tldr/global"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GetConfRequest struct {
	Platform string `json:"platform" form:"platform"`
	Language string `json:"language" form:"language"`
}

type GetConfResponse struct {
	Platforms []string `json:"platforms"`
	Languages []string `json:"languages"`
}

func GetConf(ctx *gin.Context) {
	var req GetConfRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if req.Platform != "" && req.Language != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := global.Config.DB.Get()
	var platforms, languages []string
	content := database.Content{}
	if req.Platform != "" {
		if v, err := content.GetDistinctLanguageBy(db, req.Platform); err != nil {
			logrus.Errorf("get distinct language failed. platform: %s. err: %s", req.Platform, err)
			ctx.JSON(http.StatusBadGateway, gin.H{})
			return
		} else {
			languages = v
			if z, err := content.GetDistinctPlatformBy(db, ""); err != nil {
				logrus.Errorf("get distinct platform failed. platform: %s. err: %s", req.Platform, err)
				ctx.JSON(http.StatusBadGateway, gin.H{})
				return
			} else {
				platforms = z
			}
		}
	} else if req.Language != "" {
		if v, err := content.GetDistinctPlatformBy(db, req.Language); err != nil {
			logrus.Errorf("get distinct platform failed. language: %s. err: %s", req.Language, err)
			ctx.JSON(http.StatusBadGateway, gin.H{})
			return
		} else {
			platforms = v
			if z, err := content.GetDistinctLanguageBy(db, ""); err != nil {
				logrus.Errorf("get distinct language failed. language: %s. err: %s", req.Language, err)
				ctx.JSON(http.StatusBadGateway, gin.H{})
				return
			} else {
				languages = z
			}
		}
	} else {
		if v, err := content.GetDistinctPlatformBy(db, ""); err != nil {
			logrus.Errorf("get distinct platform failed. err: %s", err)
			ctx.JSON(http.StatusBadGateway, gin.H{})
			return
		} else {
			platforms = v
			if z, err := content.GetDistinctLanguageBy(db, ""); err != nil {
				logrus.Errorf("get distinct language failed. err: %s", err)
				ctx.JSON(http.StatusBadGateway, gin.H{})
				return
			} else {
				languages = z
			}
		}
	}

	ctx.JSON(http.StatusOK, GetConfResponse{
		Platforms: platforms,
		Languages: languages,
	})
}
