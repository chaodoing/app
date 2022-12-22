package boot

import (
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	`github.com/kataras/iris/v12/middleware/logger`
	`github.com/kataras/iris/v12/middleware/recover`
	`os`
)

var (
	InitMySQL = true
	InitRedis = true
)

func Boot(conf string, isJson ...bool) Bootstrap {
	var (
		config *Configuration
		err    error
	)
	if len(isJson) > 0 && isJson[0] {
		config, err = JSON(conf)
	} else {
		config, err = XML(conf)
	}
	
	if err != nil {
		panic(err)
	}
	var container = Container{
		conf: config,
	}
	if InitMySQL {
		container.db, err = container.MySQL()
		if err != nil {
			panic(err)
		}
	}
	if InitRedis {
		container.redis, err = container.Redis()
		if err != nil {
			panic(err)
		}
	}
	
	hero.Register(container)
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Query:   true,
		Columns: true,
	}))
	drive, err := container.Log("iris-%Y-%m-%d.log")
	if err != nil {
		panic(err)
	}
	app.Logger().SetOutput(drive)
	if config.Log.Stdout {
		app.Logger().AddOutput(os.Stdout)
	}
	return Bootstrap{
		app:       app,
		container: container,
	}
}
