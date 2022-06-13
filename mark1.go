package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//Student type
type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	//inserting a rows
	insert(Student{101, "Dinesh"})
	insert(Student{102, "Dinesh"})

	//updating the student by id
	updateById(Student{101, "Dinesh Krishnan"})

	//select all students
	results := selectAll()

	//iterating a results
	for results.Next() {
		var student Student
		results.Scan(&student.Id, &student.Name)
		fmt.Println(student.Id, student.Name)
	}

	//select customer by id
	result := selectById(101)
	var student Student
	result.Scan(&student.Id, &student.Name)
	fmt.Println(student.Id, student.Name)

	//delete a customer by id
	delete(102)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Name": "Hello World",
		})
	})
	r.Run()
}

//function to get a database connection
func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Prabhat@2022@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error! Getting connection...")
	}
	return db
}

//function to insert a row in customer table
func insert(student Student) {
	db := connect()
	insert, err := db.Query("INSERT INTO student VALUES (?, ?)", student.Id, student.Name)
	if err != nil {
		fmt.Println("Error! Inserting records...")
	}
	defer insert.Close()
	defer db.Close()
}

//function to select all records from student table
func selectAll() *sql.Rows {
	db := connect()
	results, err := db.Query("SELECT * FROM student")
	if err != nil {
		fmt.Println("Error! Getting records...")
	}
	defer db.Close()
	return results
}

//function to select a student record from table by student id
func selectById(id int) *sql.Row {
	db := connect()
	result := db.QueryRow("SELECT * FROM student WHERE id=?", id)
	defer db.Close()
	return result
}

//function to update a studPent record by student id
func updateById(student Student) {
	db := connect()
	db.QueryRow("UPDATE student SET name=? WHERE id=?", student.Name, student.Id)

}

//function to delete a student by student id
func delete(id int) {
	db := connect()
	db.QueryRow("DELETE FROM student WHERE id=?", id)
}
