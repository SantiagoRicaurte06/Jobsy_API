package main

import (
	_ "Api_Josby/Catalogo_API/routers"
	"fmt"
	"time"

	"github.com/joho/godotenv"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error cargando archivo .env")
	}

	//recargar la api.confi despues de cargar .env
	err = beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	pgUser, _ := beego.AppConfig.String("PGuser")
	pgPass, _ := beego.AppConfig.String("PGpass")
	pgHost, _ := beego.AppConfig.String("PGhost")
	pgPort, _ := beego.AppConfig.String("PGport")
	pgDatabase, _ := beego.AppConfig.String("PGdatabase")
	pgSchem, _ := beego.AppConfig.String("PGschema")

	fmt.Printf("PostgreSQL connection string: postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=UTC&search_path=%s\n", pgUser, pgPass, pgHost, pgPort, pgDatabase, pgSchem)

	// Las columnas son "timestamp without time zone" y lib/pq las entrega como UTC.
	// Con la hora local el ORM restaba el desfase en cada escritura y las fechas
	// retrocedian un dia por cada PUT; lectura, escritura y sesion van en UTC.
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterDataBase(
		"default",
		"postgres",
		"postgres://"+
			pgUser+":"+
			pgPass+"@"+
			pgHost+":"+
			pgPort+"/"+
			pgDatabase+
			"?sslmode=disable&timezone=UTC&search_path="+
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
	beego.Run()
}
