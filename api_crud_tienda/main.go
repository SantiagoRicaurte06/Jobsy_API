package main

import (
	_ "api_crud_tienda/routers"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lib/pq"
)

func main() {
	pgUser, _ := beego.AppConfig.String("PGuser")
	pgPass, _ := beego.AppConfig.String("PGpass")
	pgHost, _ := beego.AppConfig.String("PGhost")
	pgPort, _ := beego.AppConfig.String("PGport")
	pgDatabase, _ := beego.AppConfig.String("PGdatabase")
	pgSchem, _ := beego.AppConfig.String("PGschema")

	orm.RegisterDataBase(
		"default",
		"postgres",
		"postgres://"+
			pgUser+":"+
			pgPass+"@"+
			pgHost+":"+
			pgPort+"/"+
			pgDatabase+
			"?sslmode=disable&search_path="+
			pgSchem,
	)
	fmt.Printf("PostgreSQL connection string: postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s\n", pgUser, pgPass, pgHost, pgPort, pgDatabase, pgSchem)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"X-Requested-With",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-CsrfToken",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}
