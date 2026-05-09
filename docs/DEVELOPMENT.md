# Development Setup

## Setup Guides for Different Operating Systems

### macOS

```bash
# Install Homebrew if not installed
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install dependencies
brew install go@1.23
brew install postgresql
brew install redis
brew install docker
brew install kubectl

# Setup Go
export PATH=$PATH:/usr/local/go/bin

# Start services
brew services start postgresql
brew services start redis
```

### Ubuntu/Debian

```bash
# Update package manager
sudo apt update

# Install Go
wget https://go.dev/dl/go1.23.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Install dependencies
sudo apt install -y postgresql postgresql-contrib redis-server docker.io kubectl

# Start services
sudo systemctl start postgresql
sudo systemctl start redis-server
sudo systemctl start docker
```

### Windows (PowerShell)

```powershell
# Install Chocolatey if not installed
Set-ExecutionPolicy Bypass -Scope Process -Force
[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# Install dependencies
choco install golang postgresql redis docker-desktop kubernetes-cli

# Start services via Services app or docker-desktop
```

### With Docker

```bash
# Install Docker Desktop
# Then use docker-compose for all services
make docker-up
```

## IDE Setup

### VS Code Extensions

- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [PostgreSQL](https://marketplace.visualstudio.com/items?itemName=ms-ossdata.vscode-postgresql)
- [Docker](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)
- [Kubernetes](https://marketplace.visualstudio.com/items?itemName=ms-kubernetes-tools.vscode-kubernetes-tools)

### GoLand Configuration

- Enable Go modules support
- Set Go SDK to version 1.23
- Enable code formatting and linting

## Git Workflow

### Initial Setup

```bash
git config --global user.name "Your Name"
git config --global user.email "your@email.com"
```

### Feature Development

```bash
# Create feature branch
git checkout -b feature/your-feature

# Make changes
git add .
git commit -m "feat: description"

# Push to fork
git push origin feature/your-feature

# Create Pull Request on GitHub
```

## Pre-commit Hooks

Setup automatic checks before committing:

```bash
# Install pre-commit
pip install pre-commit

# Setup hooks
pre-commit install

# Run manually
pre-commit run --all-files
```

## Environment Setup

```bash
# Copy environment file
cp .env.example .env

# Edit .env with your settings
# For local dev, defaults usually work
```

## Testing Workflow

```bash
# Run tests frequently
make test

# Watch mode
go test -v ./... -watch

# Coverage
make coverage

# Benchmarks
go test -bench=. ./...
```

---

For more details, see [README.md](../README.md)
