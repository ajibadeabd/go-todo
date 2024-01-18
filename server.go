package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos).GET("/todos/:id", getTodo) .POST("/todos", createTodo) .PUT("/todos/:id", updateTodo).DELETE("/todos/:id", deleteTodo)

	router.Run(":8080")
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func createTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

func updateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = Todo{ID: id, Title: updatedTodo.Title}
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
