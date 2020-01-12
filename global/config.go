package global

import (
	"fmt"
	"github.com/kosmgco/tldr/clients/client_api_github"
	"github.com/kosmgco/tldr/clients/client_tldr_sh"
	"github.com/kosmgco/tools"
	"os"
	"time"
)

var Config = struct {
	GinApp          *tools.GinApp
	ClientTldrSH    *client_tldr_sh.ClientTldrSH
	ClientApiGithub *client_api_github.ClientApiGithub
	DB              *tools.MySQL
}{
	GinApp: &tools.GinApp{
		Port: 4324,
	},
	ClientTldrSH: &client_tldr_sh.ClientTldrSH{
		Client: tools.Client{
			Host:    "https://tldr.sh",
			Timeout: time.Second * 300,
		},
	},
	ClientApiGithub: &client_api_github.ClientApiGithub{
		Client: tools.Client{
			Host: fmt.Sprintf(
				"https://%s:%s@api.github.com",
				os.Getenv("GITHUB_USER"),
				os.Getenv("GITHUB_PASS"),
			),
			Timeout: time.Second * 300,
		},
	},
	DB: &tools.MySQL{
		User:     "root",
		Database: "ooops",
		Host:     "172.16.8.72",
		Password: "root",
		Port:     "3306",
	},
}
