# ☕ coffee-chooser

[![Build](https://github.com/rflpazini/coffee-chooser/actions/workflows/build-app.yml/badge.svg)](https://github.com/rflpazini/coffee-chooser/actions/workflows/build-app.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rflpazini/coffee-chooser)](https://goreportcard.com/report/github.com/rflpazini/coffee-chooser)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?logo=go)

> The smart way to discover your perfect brew. Stop staring at your coffee equipment wondering what to make - let Coffee Chooser decide for you!

## 📋 Table of Contents

- [Key Features](#-key-features)
- [Quick Start](#-quick-start)
- [Installation](#-installation)
- [Usage](#-usage)
- [API Documentation](#-api-documentation)
- [Configuration](#-configuration)
- [Why Coffee Chooser?](#-why-coffee-chooser)
- [Contributing](#-contributing)
- [License](#-license)
- [Acknowledgements](#-acknowledgements)

## ✨ Key Features

- **Smart Recommendations** - Get brewing method suggestions based on time, mood, and preferences
- **Brew Library** - Access detailed information about popular brewing methods
- **Equipment Matcher** - Find brewing methods that work with the equipment you already own
- **Taste Profile Builder** - Track your preferences to improve future recommendations
- **Containerized Architecture** - Run anywhere with Docker, no complex setup required

## 🚀 Quick Start

Want to find your perfect brew right away? Here's the quickest path to coffee enlightenment:

```bash
# Clone repository and start application
git clone https://github.com/rflpazini/coffee-chooser.git
cd coffee-chooser
docker compose up --build
```

Then visit `http://localhost:8080` and discover your next amazing cup of coffee!

## 📦 Installation

### Prerequisites

- [Docker](https://www.docker.com/) and Docker Compose
- [Go](https://golang.org/) 1.19+ (only needed for development)
- [MongoDB](https://www.mongodb.com/) (handled by Docker automatically)

### Docker Installation (Recommended)

The simplest way to get up and running:

```bash
# Clone the repository
git clone https://github.com/rflpazini/coffee-chooser.git

# Navigate to project directory
cd coffee-chooser

# Build and start the containers
docker compose up --build
```

### Manual Installation (For Development)

If you prefer to run the application without Docker:

```bash
# Clone the repository
git clone https://github.com/rflpazini/coffee-chooser.git

# Navigate to project directory
cd coffee-chooser

# Install dependencies
go mod download

# Make sure MongoDB is running (installation varies by system)
# Example for macOS with Homebrew:
brew services start mongodb-community

# Run the application
go run cmd/server/main.go
```

## 💻 Usage

### Basic Usage

The application exposes a RESTful API that you can interact with via curl or any HTTP client:

```bash
# Get a random brewing method suggestion
curl http://localhost:8080/api/v1/suggestions/random

# Get a brewing method based on available time (in minutes)
curl http://localhost:8080/api/v1/suggestions?time=5

# Get brewing methods for specific equipment
curl http://localhost:8080/api/v1/suggestions?equipment=chemex
```

### Web Interface

A simple web interface is available at `http://localhost:8080` after starting the application.

## 📘 API Documentation

### Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/v1/methods` | GET | List all brewing methods |
| `/api/v1/methods/{id}` | GET | Get details about a specific brewing method |
| `/api/v1/suggestions/random` | GET | Get a random brewing method suggestion |
| `/api/v1/suggestions` | GET | Get suggestions based on query parameters |
| `/api/v1/profile` | POST | Save user preferences |

For complete API documentation, visit `http://localhost:8080/swagger/` after starting the application.

## ⚙️ Configuration

Coffee Chooser can be customized through environment variables or a `.env` file:

```bash
# MongoDB connection settings
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=coffee_chooser

# Server settings
PORT=8080
LOG_LEVEL=info

# Feature flags
ENABLE_TASTE_PROFILING=true
```

## 🤔 Why Coffee Chooser?

Ever stood in front of your coffee gear in the morning, too tired to decide how to brew your coffee? We've been there!

Coffee Chooser was born from a simple idea: make the decision for me. What started as a fun weekend project became a helpful tool that coffee enthusiasts worldwide use to discover new brewing techniques and get consistent recommendations.

Unlike generic recipe apps, Coffee Chooser understands the nuances of coffee brewing and helps you make the most of the equipment you already own. Think of it as the barista friend who always knows what coffee you should make next.

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/amazing-feature`)
3. Commit your Changes (`git commit -m 'Add some amazing feature'`)
4. Push to the Branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

## 👏 Acknowledgements

- [The Specialty Coffee Association](https://sca.coffee/) for brewing standards
- [James Hoffmann](https://www.youtube.com/channel/UCMb0O2CdPBNi-QqPk5T3gsQ) for brewing method insights
- All the coffee lovers who contributed suggestions and feedback

---

*Made with ☕ by coffee enthusiasts for coffee enthusiasts.*
