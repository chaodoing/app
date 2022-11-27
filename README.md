# iris+gorm 基本配置整合

```go
package main

import (
	`github.com/chaodoing/app/boot`
	`github.com/chaodoing/app/models`
	`github.com/chaodoing/app/o`
	`github.com/gookit/goutil/envutil`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	`log`
	`os`
)

var (
	// ENV     = "production"
	ENV     = "development"
	APP     = "app"
	VERSION = "v1.0.0"
)

func main() {
	var err error
	err = os.Setenv("APP", envutil.Getenv("APP", APP))
	if err != nil {
		panic(err)
		return
	}
	err = os.Setenv("ENV", envutil.Getenv("ENV", ENV))
	if err != nil {
		panic(err)
		return
	}
	err = os.Setenv("VERSION", envutil.Getenv("VERSION", VERSION))
	if err != nil {
		panic(err)
		return
	}
	
	err = os.Setenv("DIR", os.ExpandEnv("${PWD}"))
	if err != nil {
		panic(err)
	}
	err = boot.Boot("${DIR}/config/app.xml").Handle(func(app *iris.Application) {
		app.Get(`/`, hero.Handler(func(ctx iris.Context, container boot.Container) {
			auth, err := container.Auth()
			if err != nil {
				o.O(ctx, 1, err.Error(), nil)
				return
			}
			db, err := container.MySQL()
			if err != nil {
				log.Fatal(err)
				return
			}
			data, err := models.CategoriesMgr(db.Order("`sort` ASC")).Gets()
			if err != nil {
				o.O(ctx, 1, err.Error(), nil)
				return
			}
			err = auth.Create(ctx, data)
			if err != nil {
				o.O(ctx, 1, err.Error(), nil)
				return
			}
			o.O(ctx, 0, "Ok", data)
		}))
	}).Run()
	if err != nil {
		log.Fatal(err)
	}
}
```

## 编译脚本

```makefile
APP=app
ENV=production
VERSION=v1.0.0
DATE=`date +%FT%T%z`
EXECUTE=/usr/local/go/bin/go
GOARCH=`go env GOARCH`

ifeq ($(shell uname), Darwin)
	PRINT=@echo
else
	PRINT=@echo -e
endif

define begin
	$(PRINT) "\t->当前系统: \t\033[1;36m[ $(1) ]\033[0m\t";
	$(PRINT) "\t->系统平台: \t\033[1;36m[ $(2) ]\033[0m\t";
	$(PRINT) "\t->正在编译: \t\033[1;36m[ ... ]\033[0m\t";
endef

define finish
	$(PRINT) "\t->编译程序:\t\033[1;32m[ OK ]\033[0m";
endef

define compile
	@GO111MODULE=on GOOS=$(1) GOARCH=$(2) $(EXECUTE) build -o build/$(1)/$(2)/$(3)/bin/$(3) -ldflags "-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.APP=$(APP)" ./
endef

default:
	$(PRINT) "你可以使用: \n\t make all \t 编译所有平台 \n\t make linux \t 来编译 linux 平台数据 \n\t make window \t 编译 window 平台执行文件 \n\t make darwin \t 编译 MacOS 平台执行文件"

darwin:
	$(call begin,darwin,$(GOARCH))
	$(call compile,darwin,$(GOARCH),$(APP_NAME))
	$(call finish)

linux:
	$(call begin,linux,386)
	$(call compile,linux,386,$(APP_NAME))
	$(call finish)
	$(call begin,linux,amd64,$(APP_NAME))
	$(call compile,linux,amd64,$(APP_NAME))
	$(call finish)

window:
	$(call begin,windows,386)
	$(call compile,windows,386,$(APP_NAME).exe)
	$(call finish)
	$(call begin,windows,amd64)
	$(call compile,windows,amd64,$(APP_NAME).exe)
	$(call finish)

.PHONY:
	$(PRINT) "编译所有平台可执行程序\t\033[1;32m[ 请等待 ]\033[0m";

all: .PHONY darwin linux window
	$(PRINT) "所有平台可执行程序编译\t\033[1;32m[ 完成 ]\033[0m";

clear:
	@rm -rf build

install:
	@cp build/darwin/$(APP_NAME) `go env GOPATH`/bin/
```