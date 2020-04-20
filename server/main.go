package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	post struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		ZipAddress  string `json:"zip_address"`
		Address     string `json:"address"`
		HasOpenbath bool   `json:"has_openbath"`
		HasSauna    bool   `json:"has_sauna"`
		Mountains   string `json:"mountains"`
	}
)

var (
	posts = map[int]*post{}
	seq   = 1
)

//----------
// Handlers
//----------
func index(c echo.Context) error {
	db, err := sqlConnect()
	var postList = []post{}
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
		db.Raw("SELECT * FROM posts ORDER BY id ASC LIMIT 10").Scan(&postList)
		// db.Limit(10).Find(&postList)
		return c.JSON(http.StatusOK, postList)
	}
}

func createPost(c echo.Context) error {
	p := &post{
		ID: seq,
	}
	if err := c.Bind(p); err != nil {
		return err
	}
	posts[p.ID] = p
	seq++
	return c.JSON(http.StatusCreated, p)
}

func getPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := sqlConnect()
	var p post
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
		db.Raw("SELECT * FROM posts WHERE id = ? ORDER BY id DESC", id).Scan(&p)
		//db.Where("id = ?", id).Find(&p)
		return c.JSON(http.StatusOK, p)
	}
}

func updatePost(c echo.Context) error {
	p := new(post)
	if err := c.Bind(p); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	posts[id].Name = p.Name
	return c.JSON(http.StatusOK, posts[id])
}

func deletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(posts, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()
	db, err := sqlConnect()
	var h post
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
		db.First(&h, 1)
		sample_json, _ := json.Marshal(h)
		fmt.Printf("[+] %s\n", string(sample_json))
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/posts", index)
	e.POST("/api/posts", createPost)
	e.GET("/api/posts/:id", getPost)
	e.PUT("/api/posts/:id", updatePost)
	e.DELETE("/api/posts/:id", deletePost)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "moritomo"
	PASS := "moritomo_mountain"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "mountain_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
