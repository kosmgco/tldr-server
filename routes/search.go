package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kosmgco/tldr/database"
	"github.com/kosmgco/tldr/global"
	"net/http"
	"strings"
)

type SearchRequest struct {
	Query    string `json:"query" form:"query" required:"true"`
	Platform string `json:"platform" form:"platform"`
	Lang     string `json:"lang" form:"lang"`
}

type SearchResponse struct {
	Data []SearchResponseData `json:"data"`
}

type SearchResponseData struct {
	Name     string                 `json:"name"`
	Platform []string               `json:"platform"`
	Language []string               `json:"language"`
	Targets  []SearchResponseTarget `json:"targets"`
}

type SearchResponseTarget struct {
	Os       string `json:"os"`
	Language string `json:"language"`
}

func Search(ctx *gin.Context) {
	var req SearchRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if strings.Trim(req.Query, " ") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := global.Config.DB.Get()

	index := database.Index{}
	indices, err := index.SearchBy(db, database.SearchByParams{
		Name:     req.Query,
		Platform: req.Platform,
		Language: req.Lang,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{})
		return
	}

	data := []SearchResponseData{}
	for _, item := range indices {
		var p, l []string
		var t []SearchResponseTarget
		_ = json.Unmarshal(item.Platform, &p)
		_ = json.Unmarshal(item.Language, &l)
		_ = json.Unmarshal(item.Targets, &t)
		data = append(data, SearchResponseData{
			Name:     item.Name,
			Platform: p,
			Language: l,
			Targets:  t,
		})
	}

	ctx.JSON(http.StatusOK, SearchResponse{Data: data})

}
