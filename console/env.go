package console

import (
	`github.com/chaodoing/app/assets`
	`github.com/gookit/goutil/fsutil`
	`github.com/urfave/cli`
	`os`
	`strings`
)

var Env = cli.Command{
	Name:        "env",
	ShortName:   "e",
	Usage:       "环境变量配置",
	Description: "生成环境变量文件.env",
	Category:    "Frame",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "dir,d",
			Usage:       "工作目录",
			Required:    false,
			Value:       "",
			Destination: nil,
		},
	},
	Action: func(c *cli.Context) (err error) {
		if !strings.EqualFold(c.String("dir"), "") && fsutil.PathExist(os.ExpandEnv(c.String("dir"))) {
			err = os.Setenv("DIR", os.ExpandEnv(c.String("dir")))
			if err != nil {
				return
			}
		} else {
			err = os.Setenv("DIR", os.ExpandEnv("${PWD}"))
			if err != nil {
				return
			}
		}
		var f = "${DIR}/.env"
		data, err := assets.Asset("service/env")
		if err != nil {
			return err
		}
		_, err = fsutil.PutContents(os.ExpandEnv(f), string(data))
		if err != nil {
			return err
		}
		return nil
	},
}
