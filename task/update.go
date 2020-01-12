package task

import (
	"encoding/json"
	"github.com/kosmgco/tldr/clients/client_api_github"
	"github.com/kosmgco/tldr/clients/client_tldr_sh"
	"github.com/kosmgco/tldr/global"
	"github.com/sirupsen/logrus"
)

func Run() {
	output, err := global.Config.ClientTldrSH.GetIndex(client_tldr_sh.GetIndexInput{})
	if err != nil {
		logrus.Errorf("request err: %s", err)
		return
	}

	db := global.Config.DB.Get()
	for _, item := range output.Commands {
		p, _ := json.Marshal(item.Platform)
		t, _ := json.Marshal(item.Targets)
		l, _ := json.Marshal(item.Language)
		_, err := db.Exec(`insert into tldr_index(name, platform, language, targets) values (?, ?, ?, ?) on duplicate key update platform = ?, language = ?, targets = ?`, item.Name, string(p), string(l), string(t), string(p), string(l), string(t))
		if err != nil {
			logrus.Errorf("insert err: %s", err)
			continue
		}
		for _, target := range item.Targets {
			req := client_api_github.GetContentInput{}
			if target.Language == "en" {
				req.Language = ""
			} else {
				req.Language = "." + target.Language
			}
			req.Platform = target.Os
			req.Fn = item.Name
			content, err := global.Config.ClientApiGithub.GetContent(req)
			if err != nil {
				logrus.Errorf("get content failed. err: %s", err)
				continue
			}
			_, err = db.Exec(`insert into tldr_content(name, platform, language, content) values (?, ?, ?, ?) on duplicate key update content = ?`, item.Name, target.Os, target.Language, content.Content, content.Content)
			if err != nil {
				logrus.Errorf("insert content failed. err: %s", err)
				continue
			}
			_, _ = db.Exec(`insert into tldr_hot(name, platform, language) values (?, ?, ?)`, item.Name, target.Os, target.Language)

		}
	}
	logrus.Info("end")
}
