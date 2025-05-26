# 🧭 Proxite

> A lightweight reverse proxy for serving multiple SPA frontends with API proxying, configured entirely via environment
> variables.

[![Go](https://img.shields.io/badge/Go-1.24-blue)](https://golang.org)
[![License](https://img.shields.io/github/license/yourname/proxite)](LICENSE)

**Proxite** is a zero-config, single-binary reverse proxy and static file server designed to support modern frontend
development workflows.

It helps teams serve multiple **Single Page Applications (SPAs)** with associated API proxies — with no configuration
files — just pure environment variables.

---

## ✨ Features

- 🧩 **Multiple SPAs** served from different roots
- 🔁 **API proxying** per SPA via prefix match
- ⚙️ Configured entirely with `ENV` variables (no YAML/JSON)
- 📦 Single static binary, Docker-friendly
- 🔒 Graceful fallback (404 pages, SPA routing support)
- 📊 **Prometheus metrics** endpoint (`/metrics`)
- 🧵 Optional in-memory stats tracking
- 🚀 Suitable for **dev**, **preview**, or **lightweight production**

---

## 📦 Installation

### Option 1: Download binary

```bash
curl -L https://github.com/yourname/proxite/releases/download/v0.1.0/proxite-linux-amd64 -o proxite
chmod +x proxite
./proxite
```

### Option 2: Docker

```bash
docker run -p 9876:9876 \
  -e PROXITE_SPA_PROXIES_PROJECT1_ROOT=/project1 \
  -e PROXITE_SPA_PROXIES_PROJECT1_SPA_PATH=./dist/project1 \
  -e PROXITE_SPA_PROXIES_PROJECT1_PROXY_1_PATH_PREFIX=/api \
  -e PROXITE_SPA_PROXIES_PROJECT1_PROXY_1_TARGET=http://localhost:3001 \
  yourname/proxite

```

### 🚀 Quick Start

```text
export PROXITE_PORT=9876
export PROXITE_SPA_PROXIES_PROJECT1_ROOT=/project1
export PROXITE_SPA_PROXIES_PROJECT1_SPA_PATH=./dist/project1
export PROXITE_SPA_PROXIES_PROJECT1_PROXY_1_PATH_PREFIX=/api
export PROXITE_SPA_PROXIES_PROJECT1_PROXY_1_TARGET=http://localhost:3001

./proxite
```