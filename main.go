package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type ProjectConfig struct {
	BackendPath  string
	FrontendPath string
	BackendPort  string
	Prompt       string
}

type CreateProjectRequest struct {
	Prompt       string `json:"prompt"`
	BackendPath  string `json:"backend_path"`
	FrontendPath string `json:"frontend_path"`
	BackendPort  string `json:"backend_port"`
}

func main() {
	
	if len(os.Args) > 1 && os.Args[1] == "cli" {
		runCLI()
		return
	}

	
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(logger.New())
	app.Use(cors.New())

	app.Post("/api/create", createProject)
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	log.Println("🚀 MCP Orchestrator running on :3000")
	log.Fatal(app.Listen(":3000"))
}

func runCLI() {
	fmt.Println("🤖 MCP Orchestrator - Project generator with AI")
	fmt.Println("================================================")
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("📝 Describe project: ")
	prompt, _ := reader.ReadString('\n')
	prompt = strings.TrimSpace(prompt)
	
	fmt.Print("📁 Path do BACKEND (ex: /home/user/my-backend): ")
	backendPath, _ := reader.ReadString('\n')
	backendPath = strings.TrimSpace(backendPath)
	
	fmt.Print("📁 Path do FRONTEND (ex: /home/user/my-app): ")
	frontendPath, _ := reader.ReadString('\n')
	frontendPath = strings.TrimSpace(frontendPath)
	
	fmt.Print("🔌 Porta do backend (ex: 8080): ")
	backendPort, _ := reader.ReadString('\n')
	backendPort = strings.TrimSpace(backendPort)
	
	if backendPort == "" {
		backendPort = "8080"
	}
	
	config := &ProjectConfig{
		Prompt:       prompt,
		BackendPath:  backendPath,
		FrontendPath: frontendPath,
		BackendPort:  backendPort,
	}
	
	if err := orchestrateProject(config); err != nil {
		log.Fatalf("❌ Error: %v", err)
	}
	
	fmt.Println("\n✅ Projeto criado com sucesso!")
}

func createProject(c *fiber.Ctx) error {
	var req CreateProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	
	config := &ProjectConfig{
		Prompt:       req.Prompt,
		BackendPath:  req.BackendPath,
		FrontendPath: req.FrontendPath,
		BackendPort:  req.BackendPort,
	}
	
	if err := orchestrateProject(config); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(fiber.Map{
		"message":      "Project built successfully",
		"backend_path": config.BackendPath,
		"frontend_path": config.FrontendPath,
	})
}

func orchestrateProject(config *ProjectConfig) error {
	log.Printf("🎯 Starting orchestration to: %s", config.Prompt)
	
	
	log.Println("🏗️  Architecture Agents: definin structure...")
	archAgent := NewArchitectureAgent()
	if err := archAgent.DefineStructure(config); err != nil {
		return fmt.Errorf("arquitetura falhou: %w", err)
	}
	
	
	if err := createDirectories(config); err != nil {
		return fmt.Errorf("directories creation failed: %w", err)
	}
	
	
	log.Println("🔧 Backend Agent: generating Go code...")
	backendAgent := NewBackendAgent()
	if err := backendAgent.Generate(config); err != nil {
		return fmt.Errorf("backend falhou: %w", err)
	}
	
	
	log.Println("🎨 Frontend Agent: Next js generated...")
	frontendAgent := NewFrontendAgent()
	if err := frontendAgent.Generate(config); err != nil {
		return fmt.Errorf("frontend falhou: %w", err)
	}
	
	
	log.Println("🧪 Test Agent generating tests...")
	testsAgent := NewTestsAgent()
	if err := testsAgent.Generate(config); err != nil {
		log.Printf("⚠️ No tests (no critical): %v", err)
	}
	
	
	generateReadmes(config)
	
	log.Println("✅ Orquestração concluída!")
	return nil
}

func createDirectories(config *ProjectConfig) error {
	dirs := []string{
		config.BackendPath,
		config.BackendPath + "/handlers",
		config.BackendPath + "/models",
		config.FrontendPath,
		config.FrontendPath + "/app",
		config.FrontendPath + "/components",
		config.FrontendPath + "/styles",
		config.FrontendPath + "/services",
	}
	
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func generateReadmes(config *ProjectConfig) {
	backendReadme := fmt.Sprintf(`# Backend API}
	
## Descrição
Gerado automaticamente pelo MCP Orchestrator

## Setup

cd %s
go mod init backend
go get github.com/gofiber/fiber/v2
go run main.go</document>
