/*We add 14 (Layered REST) ​​or 10 (REST) ​​testing to the completed task.
Cover at least 3 functions/methods/endpoints with both positive and negative cases.
The completed task should include both unit tests and at least one functional test 
that checks the operation of several layers at once or the application as a whole*/


package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupEcho() *echo.Echo {
	e := echo.New()
	e.GET("/tasks", getTasks)
	e.POST("/tasks", addTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)
	return e
}

func TestGetTasks(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]", rec.Body.String())
	}
}

func TestAddTask(t *testing.T) {
	e := setupEcho()

	tasks = []Task{}
	nextID = 1

	task := Task{Title: "Test Task", Completed: false}
	taskJSON, _ := json.Marshal(task)

	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(taskJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, addTask(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var createdTask Task
		json.Unmarshal(rec.Body.Bytes(), &createdTask)
		assert.Equal(t, task.Title, createdTask.Title)
		assert.Equal(t, task.Completed, createdTask.Completed)
		assert.Equal(t, 1, createdTask.ID)
	}
}

func TestUpdateTask(t *testing.T) {
	e := setupEcho()

	tasks = []Task{}
	nextID = 1
	tasks = append(tasks, Task{ID: 1, Title: "Initial Task", Completed: false})

	updatedTask := Task{Title: "Updated Task", Completed: true}
	taskJSON, _ := json.Marshal(updatedTask)

	req := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewReader(taskJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, updateTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var updatedTaskResp Task
		json.Unmarshal(rec.Body.Bytes(), &updatedTaskResp)
		assert.Equal(t, updatedTask.Title, updatedTaskResp.Title)
		assert.Equal(t, updatedTask.Completed, updatedTaskResp.Completed)
		assert.Equal(t, 1, updatedTaskResp.ID)
	}

	req = httptest.NewRequest(http.MethodPut, "/tasks/999", bytes.NewReader(taskJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	err := updateTask(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteTask(t *testing.T) {
	e := setupEcho()

	tasks = []Task{}
	nextID = 1
	tasks = append(tasks, Task{ID: 1, Title: "Task to delete", Completed: false})

	req := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, deleteTask(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}

	req = httptest.NewRequest(http.MethodDelete, "/tasks/999", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	err := deleteTask(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestFunctional_AddAndGetTasks(t *testing.T) {
	e := setupEcho()

	tasks = []Task{}
	nextID = 1

	task := Task{Title: "Functional Test Task", Completed: false}
	taskJSON, _ := json.Marshal(task)

	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(taskJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, addTask(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var createdTask Task
		json.Unmarshal(rec.Body.Bytes(), &createdTask)
		assert.Equal(t, task.Title, createdTask.Title)
		assert.Equal(t, task.Completed, createdTask.Completed)
		assert.Equal(t, 1, createdTask.ID)
	}

	req = httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, getTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var tasks []Task
		json.Unmarshal(rec.Body.Bytes(), &tasks)
		assert.Equal(t, 1, len(tasks))
		assert.Equal(t, "Functional Test Task", tasks[0].Title)
		assert.Equal(t, false, tasks[0].Completed)
		assert.Equal(t, 1, tasks[0].ID)
	}
}

