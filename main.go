package main

import (
	"fmt"
	"golang_udemy/todo_app_heroku/app/controllers"
	"golang_udemy/todo_app_heroku/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	/*
		user, _ := models.GetUserByEmail("test@example.com")
		fmt.Println(user)

		session, err := user.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(session)

		valid, _ := session.CheckSession()
		fmt.Println(valid)
	*/

}
