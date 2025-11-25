#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

mkfile_if_missing() {
  local path="$1"; shift
  local content="$*"
  if [ -e "$path" ]; then
    echo "exists: $path"
  else
    mkdir -p "$(dirname "$path")"
    cat > "$path" <<'EOF'
'"$content"'
EOF
    echo "created: $path"
  fi
}

cd "$ROOT_DIR"

# Directories
mkdir -p cmd/server
mkdir -p internal/core/domain
mkdir -p internal/core/ports
mkdir -p internal/core/services
mkdir -p internal/platform/api/http
mkdir -p internal/platform/repository/postgres
mkdir -p pkg/logger
mkdir -p config
mkdir -p migrations
mkdir -p scripts

# Files (only create if missing)
if [ ! -f cmd/server/main.go ]; then
cat > cmd/server/main.go <<'EOF'
package main

import (
    "log"
    "net/http"
)

func main() {
    // TODO: wire dependencies (repos -> services -> handlers) and register routes
    log.Println("ScheduleAPI starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
EOF
echo "created: cmd/server/main.go"
else
echo "exists: cmd/server/main.go"
fi

# Domain
if [ ! -f internal/core/domain/dashboard.go ]; then
cat > internal/core/domain/dashboard.go <<'EOF'
package domain

import "github.com/google/uuid"

type Dashboard struct {
    ID    uuid.UUID `json:"id"`
    Title string    `json:"title"`
}
EOF
echo "created: internal/core/domain/dashboard.go"
else
echo "exists: internal/core/domain/dashboard.go"
fi

if [ ! -f internal/core/domain/state.go ]; then
cat > internal/core/domain/state.go <<'EOF'
package domain

import "github.com/google/uuid"

type State struct {
    ID          uuid.UUID `json:"id"`
    DashboardID uuid.UUID `json:"dashboard_id"`
    Name        string    `json:"name"`
    Position    int       `json:"position"`
}
EOF
echo "created: internal/core/domain/state.go"
else
echo "exists: internal/core/domain/state.go"
fi

if [ ! -f internal/core/domain/task.go ]; then
cat > internal/core/domain/task.go <<'EOF'
package domain

import (
    "time"

    "github.com/google/uuid"
)

type Task struct {
    ID          uuid.UUID  `json:"id"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    DashboardID uuid.UUID  `json:"dashboard_id"`
    StateID     uuid.UUID  `json:"state_id"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
EOF
echo "created: internal/core/domain/task.go"
else
echo "exists: internal/core/domain/task.go"
fi

# Ports (interfaces)
if [ ! -f internal/core/ports/repository.go ]; then
cat > internal/core/ports/repository.go <<'EOF'
package ports

import (
    "context"

    "github.com/google/uuid"
    "github.com/gj/oschedule/internal/core/domain"
)

// Outbound ports: interfaces the core depends on
type DashboardRepository interface {
    Create(ctx context.Context, d domain.Dashboard) error
    GetByID(ctx context.Context, id uuid.UUID) (domain.Dashboard, error)
    // List, Update, Delete...
}

type TaskRepository interface {
    Create(ctx context.Context, t domain.Task) error
    GetByID(ctx context.Context, id uuid.UUID) (domain.Task, error)
    // List, Update, Delete...
}
EOF
echo "created: internal/core/ports/repository.go"
else
echo "exists: internal/core/ports/repository.go"
fi

if [ ! -f internal/core/ports/service.go ]; then
cat > internal/core/ports/service.go <<'EOF'
package ports

import (
    "context"

    "github.com/google/uuid"
    "github.com/gj/oschedule/internal/core/domain"
)

// Inbound ports: interfaces that delivery layer calls
type DashboardService interface {
    CreateDashboard(ctx context.Context, title string) (domain.Dashboard, error)
    GetDashboard(ctx context.Context, id uuid.UUID) (domain.Dashboard, error)
}

type TaskService interface {
    CreateTask(ctx context.Context, t domain.Task) (domain.Task, error)
    GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error)
}
EOF
echo "created: internal/core/ports/service.go"
else
echo "exists: internal/core/ports/service.go"
fi

# Services (implementations)
if [ ! -f internal/core/services/dashboard_service.go ]; then
cat > internal/core/services/dashboard_service.go <<'EOF'
package services

import (
    "context"

    "github.com/google/uuid"
    "github.com/gj/oschedule/internal/core/domain"
    "github.com/gj/oschedule/internal/core/ports"
)

type dashboardService struct {
    repo ports.DashboardRepository
}

func NewDashboardService(r ports.DashboardRepository) ports.DashboardService {
    return &dashboardService{repo: r}
}

func (s *dashboardService) CreateDashboard(ctx context.Context, title string) (domain.Dashboard, error) {
    d := domain.Dashboard{
        ID:    uuid.New(),
        Title: title,
    }
    if err := s.repo.Create(ctx, d); err != nil {
        return domain.Dashboard{}, err
    }
    return d, nil
}

func (s *dashboardService) GetDashboard(ctx context.Context, id uuid.UUID) (domain.Dashboard, error) {
    return s.repo.GetByID(ctx, id)
}
EOF
echo "created: internal/core/services/dashboard_service.go"
else
echo "exists: internal/core/services/dashboard_service.go"
fi

if [ ! -f internal/core/services/task_service.go ]; then
cat > internal/core/services/task_service.go <<'EOF'
package services

import (
    "context"

    "github.com/google/uuid"
    "github.com/gj/oschedule/internal/core/domain"
    "github.com/gj/oschedule/internal/core/ports"
)

type taskService struct {
    repo ports.TaskRepository
}

func NewTaskService(r ports.TaskRepository) ports.TaskService {
    return &taskService{repo: r}
}

func (s *taskService) CreateTask(ctx context.Context, t domain.Task) (domain.Task, error) {
    if t.ID == uuid.Nil {
        t.ID = uuid.New()
    }
    if err := s.repo.Create(ctx, t); err != nil {
        return domain.Task{}, err
    }
    return t, nil
}

func (s *taskService) GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error) {
    return s.repo.GetByID(ctx, id)
}
EOF
echo "created: internal/core/services/task_service.go"
else
echo "exists: internal/core/services/task_service.go"
fi

# HTTP Handlers
if [ ! -f internal/platform/api/http/dashboard_handler.go ]; then
cat > internal/platform/api/http/dashboard_handler.go <<'EOF'
package http

import (
    "encoding/json"
    "net/http"

    "github.com/gj/oschedule/internal/core/ports"
)

type DashboardHandler struct {
    svc ports.DashboardService
}

func NewDashboardHandler(s ports.DashboardService) *DashboardHandler {
    return &DashboardHandler{svc: s}
}

func (h *DashboardHandler) Create(w http.ResponseWriter, r *http.Request) {
    var payload struct {
        Title string `json:"title"`
    }
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "invalid payload", http.StatusBadRequest)
        return
    }
    d, err := h.svc.CreateDashboard(r.Context(), payload.Title)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(d)
}
EOF
echo "created: internal/platform/api/http/dashboard_handler.go"
else
echo "exists: internal/platform/api/http/dashboard_handler.go"
fi

if [ ! -f internal/platform/api/http/task_handler.go ]; then
cat > internal/platform/api/http/task_handler.go <<'EOF'
package http

import (
    "net/http"
)

type TaskHandler struct {
    // TODO: inject TaskService
}

func NewTaskHandler() *TaskHandler {
    return &TaskHandler{}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}
EOF
echo "created: internal/platform/api/http/task_handler.go"
else
echo "exists: internal/platform/api/http/task_handler.go"
fi

# Postgres repository stubs
if [ ! -f internal/platform/repository/postgres/dashboard_repository.go ]; then
cat > internal/platform/repository/postgres/dashboard_repository.go <<'EOF'
package postgres

import (
    "context"
    "database/sql"

    "github.com/gj/oschedule/internal/core/domain"
    "github.com/gj/oschedule/internal/core/ports"
)

type dashboardRepo struct {
    db *sql.DB
}

func NewDashboardRepository(db *sql.DB) ports.DashboardRepository {
    return &dashboardRepo{db: db}
}

func (r *dashboardRepo) Create(ctx context.Context, d domain.Dashboard) error {
    // TODO: implement SQL insert
    return nil
}

func (r *dashboardRepo) GetByID(ctx context.Context, id uuid.UUID) (domain.Dashboard, error) {
    // TODO: implement SQL select
    return domain.Dashboard{}, nil
}
EOF
echo "created: internal/platform/repository/postgres/dashboard_repository.go"
else
echo "exists: internal/platform/repository/postgres/dashboard_repository.go"
fi

# Logger
if [ ! -f pkg/logger/logger.go ]; then
cat > pkg/logger/logger.go <<'EOF'
package logger

import "log"

func Info(msg string) {
    log.Println("[INFO]", msg)
}
EOF
echo "created: pkg/logger/logger.go"
else
echo "exists: pkg/logger/logger.go"
fi

# Config and migrations
if [ ! -f config/config.yaml ]; then
cat > config/config.yaml <<'EOF'
server:
  port: 8080

database:
  dsn: "postgres://user:pass@localhost:5432/scheduleapi?sslmode=disable"
EOF
echo "created: config/config.yaml"
else
echo "exists: config/config.yaml"
fi

if [ ! -f migrations/0001_init.sql ]; then
cat > migrations/0001_init.sql <<'EOF'
-- initial schema
CREATE TABLE IF NOT EXISTS dashboards (
  id UUID PRIMARY KEY,
  title TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS states (
  id UUID PRIMARY KEY,
  dashboard_id UUID NOT NULL REFERENCES dashboards(id),
  name TEXT NOT NULL,
  position INT NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
  id UUID PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT,
  dashboard_id UUID NOT NULL REFERENCES dashboards(id),
  state_id UUID NOT NULL REFERENCES states(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE
);
EOF
echo "created: migrations/0001_init.sql"
else
echo "exists: migrations/0001_init.sql"
fi

# go.mod (create if missing)
if [ ! -f go.mod ]; then
cat > go.mod <<'EOF'
module github.com/jgonzalez/scheduleapi

go 1.20

require (
    github.com/google/uuid v1.4.1
)
EOF
echo "created: go.mod"
else
echo "exists: go.mod"
fi

# Show top-level tree
echo
echo "Structure created/verified under: $ROOT_DIR"
echo
find . -maxdepth 3 -print