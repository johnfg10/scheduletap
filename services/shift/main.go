package main

import (
	"io"
	"strconv"

	"github.com/johnfg10/scheduletap/internal/shiftmodels"
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

			var shift shiftmodels.Shift
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

		err = client.Create(ctx.Request().Context(), &shift_models.Shift{ID: strconv.FormatUint(id, 10), Name: "Test"})
		if err != nil {
			app.Logger().Error(err)
			ctx.JSON(iris.Map{"message": "failed"})
			return
		}

		ctx.JSON(iris.Map{"message": "sucess"})
	})

	app.Run(iris.Addr(":8080"))
}
