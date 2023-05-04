package systemd

import (
	`github.com/gookit/goutil/envutil`
)

type Systemd struct {
	Description string
	Execute     string
	Directory   string
	Username    string
	Group       string
	Env         string
	Version     string
}

func NewSystemd() (data Systemd, err error) {
	var (
		description description
		execute     execute
		directory   directory
	)
	description, err = Description()
	if err != nil {
		return
	}
	execute, err = Execute()
	if err != nil {
		return
	}
	directory, err = Directory()
	if err != nil {
		return
	}
	data = Systemd{
		Description: string(description),
		Execute:     string(execute),
		Directory:   string(directory),
		Username:    "root",
		Group:       "root",
		Env:         envutil.Getenv("ENV", "production"),
		Version:     envutil.Getenv("VERSION", "v1.0.0"),
	}
	return
}
