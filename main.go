package main

/*eslint-disable */
import (
	"log"
	"manlogin/controllers"
	db "manlogin/dgclient"
	"manlogin/keynodes"
	"manlogin/models"
	"manlogin/mysql"
	"os"
	"time"

	"github.com/didip/tollbooth/limiter"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"

	"github.com/didip/tollbooth"
	"github.com/iris-contrib/middleware/tollboothic"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/websocket"
)

//cheack user
/*
TODO : move this examples to sampleController
/hi
*/
func main() {
	var PortNum string
	if len(os.Args) > 1 {
		PortNum = os.Args[1]
	} else {
		PortNum = ":9090"
	}
	app := iris.New()

	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	// uncomment it to active gunzip
	app.Use(iris.Gzip)
	app.Use(logger.New())

	globalLocale := i18n.New(i18n.Config{
		Default:      "fa",
		URLParameter: "lang",
		Languages: map[string]string{
			"fa": "./web/locales/locale_fa-Persian.ini",
			"en": "./web/locales/locale_en-US.ini",
		}})
	app.Use(globalLocale)
	// ? )  request limiter, 30 req/second for each ip
	lmt := tollbooth.NewLimiter(30, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	app.Use(tollboothic.LimitHandler(lmt))

	tmp := iris.Handlebars("./web/views", ".handlebars")
	tmp.Reload(true)
	tmp.AddFunc("BaseURL", func() string {
		return keynodes.BaseURL
	})

	tmp.AddFunc("Show", func(s, t string) string {
		return s + " showing " + t
	})

	app.RegisterView(tmp)
	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.RegisterView(iris.HTML("./web/views", ".js"))

	app.StaticWeb("/public", "./web/public")
	app.StaticWeb("/assets", "./web/public")
	app.StaticWeb("/static", "./web/public")

	app.Any("/iris-ws.js", websocket.ClientHandler())
	PCAP := map[string]models.TmpSes{}
	hero.Register(PCAP)

	sessionManager := sessions.New(sessions.Config{
		Cookie:       "b502320222bfe165e6bc37db8ea466c3bad11fad72de1d54bbcfe220bb3c94c8",
		Expires:      365 * 24 * time.Hour,
		AllowReclaim: true,
	})
	// var sessionStarter *sessions.Session
	hero.Register(sessionManager.Start)

	sqlc, err := mysql.NewClient()
	if err != nil {
		log.Panic(err)
	}
	hero.Register(sqlc)
	ws := websocket.New(websocket.Config{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		MaxMessageSize:   1024 * 500,
		EvtMessagePrefix: []byte("Ara:"),
	})
	hero.Register(ws)

	DgClient := db.NewClient()
	hero.Register(DgClient)

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE", "PATCH", "WS"},
		AllowedHeaders:   []string{"Accept", "X-Cat", "X-User", "X-USER", "X-S2SToken", "content-type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Authorization-Token", "Screen"},
		AllowCredentials: true,
	})
	/*
		@ X-User : standard user uid
		@ X-Token : standard user Token
	*/

	mvc.New(app.Party("/file", crs)).Handle(new(controllers.FileController))

	mvc.New(app.Party("/", crs)).Handle(new(controllers.MainController))

	app.Get("/favicon.ico", func(ctx iris.Context) {
		ctx.SendFile("./web/public/favicon.ico", "favicon.ico")
	})

	/*app.Post("/file/{action:path}", func(ctx iris.Context) {
		path := ctx.Params().Get("action")
		//paths := strings.Split(path, "/")
		trueUrl := FServers["1"] + "/" + path
		ctx.Redirect(trueUrl)
	})*/

	app.Post("/ping", func(ctx iris.Context) {
		ctx.Writef("Pong")
	})
	app.Post("/ctx", func(ctx iris.Context) {
		Xuser := ctx.GetHeader("X-User")
		AToken := ctx.GetHeader("Authorization-Token")

		ctx.Writef("X-User : %s \nAuthorization-Token : %s \n", Xuser, AToken)
	})
	app.Get("/latency", func(ctx iris.Context) {
		ctx.StatusCode(200)
	})
	app.Get("/see", func(ctx iris.Context) {
		ctx.Header("Location", "http://yaqoti.ir/?name=hamid")
		ctx.StatusCode(iris.StatusFound)

	})
	app.Get("/hash", func(ctx iris.Context) {
		t, _ := models.GeneratePassword(ctx.URLParam("pass"))
		ctx.Write(t)
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	log.Printf("\n:%s", PortNum)
	app.Run(iris.Addr(PortNum), iris.WithoutServerError(iris.ErrServerClosed))
	//? ) HTTPS:
	//app.Run(iris.AutoTLS(":443", "example.com", "mail@example.com"))

	// start the server (HTTPS) on port 443, this is a blocking func
	/*
		target, _ := url.Parse("https://127.0.1:443")
		go host.NewProxy("127.0.0.1:9090", target).ListenAndServe()
		app.Run(iris.TLS("127.0.0.1:443", "mycert.cert", "mykey.key"))
	*/
}
