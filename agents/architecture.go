package agents

import (
	"backend/llama"
	"fmt"
	"log"

	"os"
)

type ArchitectureAgent struct {
	llamaClient *llama.Client
}

func NewArchitectureAgent() *ArchitectureAgent {
	return &ArchitectureAgent{
		llamaClient: llama.NewClient("http://localhost:8080"),
	}
}

func (a *ArchitectureAgent) DefineStructure(config interface{}) error {

	projConfig, ok := config.(*ProjectConfig)
	if !ok {
		return fmt.Errorf("invalid config type")
	}

	prompt := fmt.Sprintf(`You are a software architect. For a project with the following characteristics:
"%s"

Generate a directory and file structure for:
1. BACKEND in Go with Fiber (isolated repository)
2. FRONTEND in Next.js with App Router and Styled Components (isolated repository)

Respond ONLY with JSON in the format:
{
  "backend_files": [{"path": "main.go", "description": "..."}],
  "frontend_files": [{"path": "app/page.tsx", "description": "..."}]
}
`, projConfig.Prompt)

	response, err := a.llamaClient.Generate(prompt, 500)
	if err != nil {
		log.Printf("⚠️ LLaMA not available, using default structure")
		return a.createDefaultStructure(projConfig)
	}

	log.Printf("📐 Architecture generated: %s", response)

	os.WriteFile(projConfig.BackendPath+"/.architecture.json", []byte(response), 0644)
	os.WriteFile(projConfig.FrontendPath+"/.architecture.json", []byte(response), 0644)

	return nil
}

func (a *ArchitectureAgent) createDefaultStructure(config *ProjectConfig) error {
	defaultStructure := map[string]interface{}{
		"backend_files": []string{
			"main.go", "handlers/todo.go", "models/todo.go", "go.mod",
		},
		"frontend_files": []string{
			"app/page.tsx", "app/layout.tsx", "components/TodoList/index.tsx",
			"components/TodoList/styles.ts", "services/api.ts",
		},
	}

	log.Printf("📐 Using default structure: %v", defaultStructure)
	return nil
}
