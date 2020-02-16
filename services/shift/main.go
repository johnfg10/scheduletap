package main

import (
	"io"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/sony/sonyflake"
	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/mongodocstore"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		return 232, nil
	}
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/test", func(ctx iris.Context) {
		client, err := docstore.OpenCollection(ctx.Request().Context(), "mongo://test/test?id_field=ID")
		if err != nil {
			app.Logger().Error(err)
			ctx.StatusCode(500)
			ctx.JSON(iris.Map{"message": "failed"})
			return
		}

		defer client.Close()

		iter := client.Query().Where("Name", "=", "Test").Get(ctx.Request().Context())

		for {
			var shift Shift
			err := iter.Next(ctx.Request().Context(), &shift)
			if err == io.EOF {
				app.Logger().Debug("end of file")
				ctx.StatusCode(404)
				ctx.JSON(iris.Map{"message": "not found"})
				return
			} else if err != nil {
				app.Logger().Error(err)
				ctx.StatusCode(500)
				ctx.JSON(iris.Map{"message": "failed"})
				return
			} else {
				app.Logger().Debug(shift)
				ctx.JSON(shift)
				return
			}
		}
	})

	app.Handle("GET", "/", func(ctx iris.Context) {
		client, err := docstore.OpenCollection(ctx.Request().Context(), "mongo://test/test?id_field=ID")
		if err != nil {
			app.Logger().Error(err)
			ctx.StatusCode(500)
			ctx.JSON(iris.Map{"message": "failed"})
			return
		}
		defer client.Close()
		id, err := sf.NextID()
		if err != nil {
			app.Logger().Error(err)
			ctx.JSON(iris.Map{"message": "failed"})
			return
		}

		err = client.Create(ctx.Request().Context(), &Shift{ID: strconv.FormatUint(id, 10), Name: "Test"})
		if err != nil {
			app.Logger().Error(err)
			ctx.JSON(iris.Map{"message": "failed"})
			return
		}

		ctx.JSON(iris.Map{"message": "sucess"})
	})

	app.Run(iris.Addr(":8080"))
}

type OfficeLocation struct {
	Address string `json:"address"`
}

type Company struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Shift s
type Shift struct {
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	Location         OfficeLocation `json:"location"`
	StartTime        time.Time      `json:"start_time"`
	Duration         time.Duration  `json:"duration"`
	DocstoreRevision interface{}    `json:"-"`
}

type Position struct {
	ID               string        `json:"id"`
	Company          Company       `json:"company"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	StartTime        time.Time     `json:"start_time"`
	Duration         time.Duration `json:"duration"`
	DocstoreRevision interface{}   `json:"-"`
}
