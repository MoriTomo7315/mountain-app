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
	mountain struct {
		ID         int     `json:"id"`
		Name       string  `json:"name"`
		NameKana   string  `json:"name_kana"`
		Elevation  float64 `json:"elecation"`
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
		Detail     string  `json:detail`
		Prefecture string  `json:"prefecture"`
	}
)

var (
	mountains = map[int]*mountain{}
	seq       = 1
)

//----------
// Handlers
//----------

func createMountain(c echo.Context) error {
	m := &mountain{
		ID: seq,
	}
	if err := c.Bind(m); err != nil {
		return err
	}
	mountains[m.ID] = m
	seq++
	return c.JSON(http.StatusCreated, m)
}

func getMountain(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := sqlConnect()
	var m mountain
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
		db.Where("id = ?", id).Find(&m)
		return c.JSON(http.StatusOK, m)
	}
}

func updateMountain(c echo.Context) error {
	m := new(mountain)
	if err := c.Bind(m); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	mountains[id].Name = m.Name
	mountains[id].Prefecture = m.Prefecture
	return c.JSON(http.StatusOK, mountains[id])
}

func deleteMountain(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(mountains, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()
	db, err := sqlConnect()
	var m mountain
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
		db.First(&m, 1)
		sample_json, _ := json.Marshal(m)
		fmt.Printf("[+] %s\n", string(sample_json))
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/mountains", createMountain)
	e.GET("/mountains/:id", getMountain)
	e.PUT("/mountains/:id", updateMountain)
	e.DELETE("/mountains/:id", deleteMountain)

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
