package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strings"
	// "net/http"
)

type userRef struct {
	gorm.Model
	From   string
	To     string
	PostID int
}

var postID = 1

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=sasuke dbname=point_one password=Wb922149@...S sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	db.LogMode(true)
	db.AutoMigrate(&userRef{})

	r := gin.Default()
	r.GET("/suggest", func(c *gin.Context) {
		user := c.Query("user")
		users := []string{}

		postIDs := []int{}
		rows, err := db.Raw("select distinct user_refs.post_id from user_refs where user_refs.to = ?", user).Rows()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			rows.Scan(&id)
			postIDs = append(postIDs, id)
		}

		rows, err = db.Raw("select distinct user_refs.to from user_refs where post_id in (?) and user_refs.to != ?", postIDs, user).Rows()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var user string
			rows.Scan(&user)
			users = append(users, user)
		}

		froms := []string{}
		rows, err = db.Raw("select distinct user_refs.from from user_refs where user_refs.to = ?", user).Rows()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var from string
			rows.Scan(&from)
			froms = append(froms, from)
		}

		rows, err = db.Raw("select distinct user_refs.to from user_refs where user_refs.from in (?) and user_refs.to != ?", froms, user).Rows()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var user string
			rows.Scan(&user)
			dup := false
			for _, u := range users {
				if u == user {
					dup = true
					break
				}
			}

			if !dup {
				users = append(users, user)
			}
		}

		c.JSON(200, gin.H{"suggest": users})
	})

	r.POST("/postWeibo", func(c *gin.Context) {
		fromUser := c.PostForm("from")
		to := c.PostForm("to")
		toUsers := strings.Split(to, ",")
		fmt.Println(toUsers)

		for _, toUser := range toUsers {
			ref := userRef{From: fromUser, To: toUser, PostID: postID}
			db.Create(&ref)
		}

		c.JSON(200, gin.H{
			"message": "ok",
			"id":      postID,
		})
		postID++
	})

	r.Run()
}
