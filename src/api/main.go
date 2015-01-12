package api

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

type media struct {
	Id       int64
	SiteName string
	ApiKey   string
}

type dbConfig struct {
	dialect, user, pass, host, port, dbName string
}

func (c *dbConfig) uri() string {
	format := "%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true"
	return fmt.Sprintf(format, c.user, c.pass, c.host, c.port, c.dbName)
}

func init() {
	m := media{
		Id:       1,
		SiteName: "BorrowMoney.com",
		ApiKey:   "abcdefg",
	}

	fmt.Println(m)

	cfg := &dbConfig{
		dialect: "mysql",
		user:    "root",
		pass:    "",
		host:    "127.0.0.1",
		port:    "13306",
		dbName:  "lmp",
	}

	db, err := gorm.Open(cfg.dialect, cfg.uri())
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.DropTableIfExists(&media{})
	db.AutoMigrate(&media{})

	db.Create(&m)

}

func main() {
	goji.Get("/media/:id", handleMedia)
	goji.Serve()
}

func handleMedia(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["id"]

	cfg := &dbConfig{
		dialect: "mysql",
		user:    "root",
		pass:    "",
		host:    "127.0.0.1",
		port:    "13306",
		dbName:  "lmp",
	}

	db, err := gorm.Open(cfg.dialect, cfg.uri())
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var m1 media
	db.First(&m1, id)
	fmt.Println(m1)

	fmt.Fprintf(w, "Hello, %v", m1)

}

func Add(a, b int) int {
	return a + b
}
