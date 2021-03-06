package buffalo

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/buffalo"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gobuffalo/buffalo"
	"net/http"
	"os"
)

func NewBuffaloHandler() http.Handler {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	examplePlugin := example.NewExample()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJson(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(bu); err != nil {
		panic(err)
	}

	bu.GET("/admin", func(ctx buffalo.Context) error {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return datamodel.GetContent()
		})
		return nil
	})

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	return bu
}
