package console

import (
	"github.com/chaodoing/app/o"
	"github.com/urfave/cli"
)

var file string

var Config = cli.Command{
	Name:        "config",
	ShortName:   "c",
	Usage:       "生成配置文件",
	Description: "生成配置文件 -f [文件名称].xml",
	Category:    "Frame",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "file,f",
			Usage:       "配置文件输出位置",
			Required:    true,
			Value:       "./config.xml",
			Destination: &file,
		},
	},
	Action: func(c *cli.Context) (err error) {
		return o.SaveXML(Configuration, file)
	},
}
