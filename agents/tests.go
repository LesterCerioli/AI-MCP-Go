package agents

import (
	"backend/llama"
	"os"
)

type TestsAgent struct {
	llamaClient *llama.Client
}

func NewTestsAgent() *TestsAgent {
	return &TestsAgent{
		llamaClient: llama.NewClient("http://localhost:8080"),
	}
}

func (t *TestsAgent) Generate(config *ProjectConfig) error {

	testContent := `package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
)

func TestGetTodos(t *testing.T) {
	app := fiber.New()
	app.Get("/api/todos", GetTodos)
	
	req := httptest.NewRequest("GET", "/api/todos", nil)
	resp, _ := app.Test(req)
	
	if resp.StatusCode != 200 {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}
}

func TestCreateTodo(t *testing.T) {
	app := fiber.New()
	app.Post("/api/todos", CreateTodo)
	
	todo := Todo{ID: "1", Title: "Test", Completed: false}
	body, _ := json.Marshal(todo)
	
	req := httptest.NewRequest("POST", "/api/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	
	resp, _ := app.Test(req)
	
	if resp.StatusCode != 201 {
		t.Errorf("Expected 201, got %d", resp.StatusCode)
	}
}`

	os.MkdirAll(config.BackendPath+"/handlers", 0755)
	os.WriteFile(config.BackendPath+"/handlers/todo_test.go", []byte(testContent), 0644)

	return nil
}
