package models

type ProjectConfig struct {
	Prompt       string `json:"prompt"`
	BackendPath  string `json:"backend_path"`
	FrontendPath string `json:"frontend_path"`
	BackendPort  string `json:"backend_port"`
}

type Architecture struct {
	BackendFiles  []FileSpec `json:"backend_files"`
	FrontendFiles []FileSpec `json:"frontend_files"`
}

type FileSpec struct {
	Path        string `json:"path"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

type LLMRequest struct {
	Prompt      string   `json:"prompt"`
	MaxTokens   int      `json:"max_tokens"`
	Temperature float64  `json:"temperature"`
	Stop        []string `json:"stop,omitempty"`
}

type LLMResponse struct {
	Text      string `json:"text"`
	Content   string `json:"content,omitempty"`
	Completed bool   `json:"completed"`
	Model     string `json:"model,omitempty"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Details string `json:"details,omitempty"`
}

type CreateProjectRequest struct {
	Prompt       string `json:"prompt"`
	BackendPath  string `json:"backend_path"`
	FrontendPath string `json:"frontend_path"`
	BackendPort  string `json:"backend_port"`
}

type CreateProjectResponse struct {
	Message      string `json:"message"`
	BackendPath  string `json:"backend_path"`
	FrontendPath string `json:"frontend_path"`
	BackendPort  string `json:"backend_port"`
	Status       string `json:"status"`
}

type HealthCheckResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type TodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoResponse struct {
	Todo    *Todo   `json:"todo,omitempty"`
	Todos   []*Todo `json:"todos,omitempty"`
	Message string  `json:"message,omitempty"`
	Total   int     `json:"total,omitempty"`
}

type AgentType string

const (
	AgentArchitecture AgentType = "architecture"
	AgentBackend      AgentType = "backend"
	AgentFrontend     AgentType = "frontend"
	AgentTests        AgentType = "tests"
)

type AgentTask struct {
	Type      AgentType   `json:"type"`
	Config    interface{} `json:"config"`
	Priority  int         `json:"priority"`
	CreatedAt string      `json:"created_at"`
}

type AgentResult struct {
	Type        AgentType `json:"type"`
	Success     bool      `json:"success"`
	Error       string    `json:"error,omitempty"`
	Files       []string  `json:"files,omitempty"`
	Duration    int64     `json:"duration_ms"`
	CompletedAt string    `json:"completed_at"`
}
