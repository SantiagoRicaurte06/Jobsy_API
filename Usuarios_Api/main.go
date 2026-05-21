package main

import (
	_ "JOBSY_API/routers"
	"fmt"

	"github.com/beego/beego/v2/server/web/filter/cors"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main() {

    err := godotenv.Load()
    if err != nil {
        panic("Error cargando el archivo .env: " + err.Error())
    }

	//regargar app.config despues de cargar .env
	err = beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

    User, _    := beego.AppConfig.String("PGuser")
    password, _ := beego.AppConfig.String("PGpassword")
    host,    _ := beego.AppConfig.String("PGhost")
    port,     _ := beego.AppConfig.String("PGport")
    database, _ := beego.AppConfig.String("PGdatabase")
    schema, _ := beego.AppConfig.String("PGschema")

	fmt.Printf("postgresSqL connection string: postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path="+"%s\n", User, password, host, port, database, schema)


// Register the database connection
orm.RegisterDataBase("default", "postgres", "postgres://"+User+":"+password+"@"+host+":"+port+"/"+database+"?sslmode=disable&search_path="+schema)

// config swagger	
	if  beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

// config cors
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: 		[]string{"*"},
		AllowMethods: 		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: 		[]string{"Origin","x-requested-with","content-type","accept","origin","authorization","x-csrftoken",},
		ExposeHeaders:   	[]string{"Content-Length"},
		AllowCredentials: 	true,
	}))

// run the server
	beego.Run()
}