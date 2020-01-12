package client_api_github

import (
	"fmt"
	"github.com/kosmgco/tools"
)

type IClientApiGithub interface {
	GetContent(GetContentInput) (*GetContentOutput, error)
}

type ClientApiGithub struct {
	tools.Client
}

func (c *ClientApiGithub) GetContent(input GetContentInput) (output *GetContentOutput, err error) {
	err = c.
		SetURL(fmt.Sprintf("/repos/tldr-pages/tldr/contents/pages%s/%s/%s.md", input.Language, input.Platform, input.Fn)).
		SetMethod(tools.HTTP_METHOD__GET).Do().Into(&output)
	return
}
