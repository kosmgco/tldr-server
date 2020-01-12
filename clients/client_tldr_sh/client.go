package client_tldr_sh

import (
	"github.com/kosmgco/tools"
)

type IClientTldrSH interface {
	GetIndex(GetIndexInput) (*GetIndexOutput, error)
}

type ClientTldrSH struct {
	tools.Client
}

func (c *ClientTldrSH) GetIndex(input GetIndexInput) (output *GetIndexOutput, err error) {
	err = c.SetURL("/assets/index.json").SetMethod(tools.HTTP_METHOD__GET).Do().Into(&output)
	return
}
