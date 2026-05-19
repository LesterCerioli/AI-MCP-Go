
# 🚀 MCP Orchestrator - AI-Powered Project Generator

The **MCP Orchestrator** is an intelligent system that generates complete software projects using AI (LLaMA). It orchestrates multiple specialized agents to create a backend in Go with Fiber and a frontend in Next.js with Styled Components, in isolated repositories.

## ✨ Features

- 🤖 **Intelligent Orchestration**: Coordinates specialized agents using AI
- 📦 **Isolated Repositories**: Backend and frontend in separate directories
- 🎯 **Multiple Agents**:
  - Architecture Agent - Defines project structure
  - Backend Agent - Generates Go code with Fiber
  - Frontend Agent - Creates Next.js with Styled Components
  - Tests Agent - Generates automated tests
- 🐳 **LLaMA Integration**: Runs AI model locally via Docker
- ⚡ **High Performance**: Written in Go with Fiber, concurrent support
- 🎨 **Styled Components**: CSS-in-JS with TypeScript
- 📱 **App Router**: Next.js 14 with server components

## 📋 Prerequisites

- **Go** 1.21 or higher
- **Docker** (to run LLaMA)
- **Node.js** 18+ (to test the generated frontend)

## 🛠️ Installation

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/mcp-orchestrator.git
cd mcp-orchestrator
```

### 2. Install Go dependencies

```bash
go mod tidy
```

### 3. Set up the LLaMA container

```bash
# Build the Docker image
docker build -f Dockerfile.llama -t llama-server .

# Run the container
docker run -p 8080:8080 llama-server
```

LLaMA will be available at `http://localhost:8080`

## 🚀 How to Use

### CLI Mode (Recommended)

```bash
go run main.go cli
```

The assistant will ask for:
1. 📝 Project description
2. 📁 Backend path
3. 📁 Frontend path
4. 🔌 Backend port

### HTTP Server Mode

```bash
go run main.go
```

Make a POST request:

```bash
curl -X POST http://localhost:3000/api/create \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "Task management system with complete CRUD",
    "backend_path": "/home/user/task-api",
    "frontend_path": "/home/user/task-app",
    "backend_port": "8080"
  }'
```

## 📁 Project Structure

```
mcp-orchestrator/
├── main.go                 # Main orchestrator (Fiber)
├── models/
│   └── types.go           # Data types and structures
├── agents/
│   ├── architecture.go     # Architecture agent
│   ├── backend.go         # Backend agent (Go/Fiber)
│   ├── frontend.go        # Frontend agent (Next.js)
│   └── tests.go           # Tests agent
├── llama/
│   └── client.go          # HTTP client for LLaMA
├── Dockerfile.llama       # LLaMA container
├── go.mod                 # Go dependencies
└── README.md
```

## 🎯 Example Generated Project

### Backend (Go with Fiber)

```go
// main.go
package main

import (
    "github.com/gofiber/fiber/v2"
    // ...
)

func main() {
    app := fiber.New()
    app.Get("/api/todos", handlers.GetTodos)
    app.Post("/api/todos", handlers.CreateTodo)
    app.Listen(":8080")
}
```

### Frontend (Next.js with Styled Components)

```tsx
// components/TodoList/index.tsx
import { ListContainer } from './styles';

export function TodoList({ todos }) {
  return (
    <ListContainer>
      {todos.map(todo => <TodoItem key={todo.id} todo={todo} />)}
    </ListContainer>
  );
}
```

```ts
// components/TodoList/styles.ts
import styled from 'styled-components';

export const ListContainer = styled.div`
  display: flex;
  flex-direction: column;
  gap: 1rem;
`;
```

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `LLAMA_API_URL` | `http://localhost:8080` | LLaMA API URL |
| `MCP_PORT` | `3000` | Orchestrator HTTP port |

### Customization

To modify generated templates, edit the agents in `agents/`:

- `backend.go`: Backend Go template
- `frontend.go`: Next.js frontend template
- `tests.go`: Test templates

## 🧪 Testing the Generated Project

### Backend

```bash
cd /path/to/backend
go mod tidy
go run main.go
```

Access: `http://localhost:8080/health`

### Frontend

```bash
cd /path/to/frontend
npm install
npm run dev
```

Access: `http://localhost:3000`

## 🐛 Troubleshooting

### Error: "LLaMA API not available"

**Solution**: Check if the Docker container is running:

```bash
docker ps | grep llama-server
docker logs <container-id>
```

### Error: "Permission denied" when creating directories

**Solution**: Run with proper permissions or use a writable directory:

```bash
sudo chmod 755 /path/to/project
```

### Error: "illegal label declaration"

**Solution**: This error has been fixed in the current version. Make sure you're using the latest code.

## 📊 Performance

- **Backend Go**: ~50ms response time in production mode
- **Frontend Next.js**: Build time ~30s (depends on project size)
- **LLaMA (Docker)**: ~2-5s per generation (depends on model)

## 🏗️ Architecture

```
[User] → [MCP Orchestrator] → [Architecture Agent] → [LLaMA]
                    ↓                      ↓
            [Backend Agent]        [Frontend Agent]
                    ↓                      ↓
            [Go Code]             [Next.js + TS]
                    ↓                      ↓
            [Backend Repo]        [Frontend Repo]
```

## 🔄 Workflow

1. User provides prompt and paths
2. MCP validates inputs
3. Architecture Agent queries LLaMA
4. Specialized agents generate code
5. System writes files to specified paths
6. READMEs and instructions are generated

## 🤝 Contributing

1. Fork the project
2. Create your branch: `git checkout -b feature/new-feature`
3. Commit: `git commit -m 'Add new feature'`
4. Push: `git push origin feature/new-feature`
5. Open a Pull Request

## 📄 License

MIT License - see the [LICENSE](LICENSE) file for details

## 🙏 Acknowledgments

- [Go Fiber](https://gofiber.io/) - Web framework
- [Next.js](https://nextjs.org/) - React framework
- [Styled Components](https://styled-components.com/) - CSS-in-JS
- [LLaMA](https://github.com/facebookresearch/llama) - AI model
- [Ollama](https://ollama.ai/) - Model management

## 📞 Support

- Open an issue on GitHub
- Email: cerioli728@gmail.com
- Documentation: [docs.mcp-orchestrator.com](https://docs.mcp-orchestrator.com)

## 🎉 Roadmap

- [ ] Support for more databases (PostgreSQL, MongoDB)
- [ ] E2E test generation with Cypress
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] Automatic Vercel deployment
- [ ] Web interface for the orchestrator
- [ ] Support for other AI models (GPT4All, Mistral)

---

