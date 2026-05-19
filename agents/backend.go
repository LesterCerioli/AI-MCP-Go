package agents

import (
	"backend/llama"
	"fmt"
	"os"
)

type ProjectConfig struct {
	Prompt       string
	BackendPath  string
	FrontendPath string
	BackendPort  string
}

type BackendAgent struct {
	llamaClient *llama.Client
}

func NewBackendAgent() *BackendAgent {
	return &BackendAgent{
		llamaClient: llama.NewClient("http://localhost:8080"),
	}
}

func (b *BackendAgent) Generate(config *ProjectConfig) error {

	if err := b.generateMainGo(config); err != nil {
		return err
	}

	if err := b.generateHandlers(config); err != nil {
		return err
	}

	if err := b.generateModels(config); err != nil {
		return err
	}

	if err := b.generateGoMod(config); err != nil {
		return err
	}

	return nil
}

func (b *BackendAgent) generateMainGo(config *ProjectConfig) error {
	prompt := fmt.Sprintf(`Gere um arquivo main.go para um servidor HTTP em Go usando Fiber.
Requisitos:
- Porta: %s
- Deve incluir rotas: /health, /api/todos (GET, POST, PUT, DELETE)
- Deve usar in-memory storage para todos
- Incluir middleware de logging e CORS

Responda APENAS com o código Go completo, sem explicações.`, config.BackendPort)

	content, err := b.llamaClient.Generate(prompt, 800)
	if err != nil {

		content = b.getDefaultMainGo(config.BackendPort)
	}

	return os.WriteFile(config.BackendPath+"/main.go", []byte(content), 0644)
}

func (b *BackendAgent) generateHandlers(config *ProjectConfig) error {
	prompt := `Gere um handler para todos em Go com Fiber.
Inclua funções: GetTodos, CreateTodo, UpdateTodo, DeleteTodo.
Use uma struct Todo com ID, Title, Completed.
Responda APENAS com o código Go.`

	content, err := b.llamaClient.Generate(prompt, 600)
	if err != nil {
		content = `package handlers

import (
	"github.com/gofiber/fiber/v2"
	"sync"
)

type Todo struct {
	ID        string ` + "`json:\"id\"`" + `
	Title     string ` + "`json:\"title\"`" + `
	Completed bool   ` + "`json:\"completed\"`" + `
}

var (
	todos  = make(map[string]Todo)
	mu     sync.RWMutex
)

func GetTodos(c *fiber.Ctx) error {
	mu.RLock()
	defer mu.RUnlock()
	
	list := make([]Todo, 0, len(todos))
	for _, todo := range todos {
		list = append(list, todo)
	}
	return c.JSON(list)
}

func CreateTodo(c *fiber.Ctx) error {
	var todo Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	
	mu.Lock()
	defer mu.Unlock()
	todos[todo.ID] = todo
	
	return c.Status(201).JSON(todo)
}
`
	}

	return os.WriteFile(config.BackendPath+"/handlers/todo.go", []byte(content), 0644)
}

func (b *BackendAgent) generateModels(config *ProjectConfig) error {
	modelContent := `package models

type Todo struct {
	ID        string ` + "`json:\"id\"`" + `
	Title     string ` + "`json:\"title\"`" + `
	Completed bool   ` + "`json:\"completed\"`" + `
}

type ErrorResponse struct {
	Error string ` + "`json:\"error\"`" + `
}
`
	return os.WriteFile(config.BackendPath+"/models/todo.go", []byte(modelContent), 0644)
}

func (b *BackendAgent) generateGoMod(config *ProjectConfig) error {
	content := `module backend

go 1.26

require github.com/gofiber/fiber/v2 v2.52.0

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
`
	return os.WriteFile(config.BackendPath+"/go.mod", []byte(content), 0644)
}

func (b *BackendAgent) getDefaultMainGo(port string) string {
	return fmt.Sprintf(`package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	
	app.Use(logger.New())
	app.Use(cors.New())
	
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	
	app.Get("/api/todos", handlers.GetTodos)
	app.Post("/api/todos", handlers.CreateTodo)
	
	log.Printf("🚀 Server running on :%s", "%s")
	log.Fatal(app.Listen(":" + "%s"))
}`, port, port)
}
