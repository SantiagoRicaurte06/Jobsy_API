package main

import (
	_ "Jobsy_API/Core_api/routers"
	"fmt"
	
	"github.com/joho/godotenv"
	
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lib/pq"
)

func main() {
	err:= godotenv.Load()
	if err !=nil {
		panic("Error cargando archivo .env")
	}

	//recargar la api.confi despues de cargar .env
	err = beego.LoadAppConfig("ini", "conf/app.conf")
	if err !=nil {
		panic(err)
	}

	pgUser, _ := beego.AppConfig.String("PGuser")
	pgPass, _ := beego.AppConfig.String("PGpass")
	pgHost, _ := beego.AppConfig.String("PGhost")
	pgPort, _ := beego.AppConfig.String("PGport")
	pgDatabase, _ := beego.AppConfig.String("PGdatabase")
	pgSchem, _ := beego.AppConfig.String("PGschema")

	fmt.Printf("PostgreSQL connection string: postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s\n", pgUser, pgPass, pgHost, pgPort, pgDatabase, pgSchem)

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

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
