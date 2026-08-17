# Resizem release build
#
# Usage:
#   make release          Build release artifacts for every platform the current host can produce
#   make release-macos    Build + zip the macOS app (universal by default)
#   make release-windows  Build the Windows exe + NSIS installer
#   make release-linux    Build the Linux binary (requires a Linux host)
#   make release-all      Attempt every platform; the Linux step needs a Linux host
#   make frontend         Build the Svelte frontend into frontend/dist
#   make test             Run the Go unit tests
#   make clean            Remove build outputs (keeps node_modules)
#   make distclean        Remove build outputs, frontend/dist and node_modules
#   make version          Print the version this build would use
#
# All artifacts land in build/bin/ (gitignored). Nothing is uploaded anywhere.
#
# Notes
#   - APP_VERSION comes from wails.json "productVersion"; override with: make release APP_VERSION=1.2.0
#   - The app and the build CLI are both pinned to wails v2.14.0, the latest v2 release.
#   - proxy.golang.org is unreachable from this network, so builds default to a
#     mirror. Override on the command line, e.g.  make release GOPROXY=https://proxy.golang.org,direct

APP_NAME    := Resizem
APP_VERSION := $(shell sed -n 's/.*"productVersion"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' wails.json | head -1)

# Overridable on the command line, not by the environment (see note above).
export GOPROXY = https://goproxy.cn,direct

GO        := go
GOBIN     := $(shell $(GO) env GOPATH)/bin
WAILS     := $(or $(shell command -v wails 2>/dev/null), $(GOBIN)/wails)
WAILS_VER ?= v2.14.0

MAC_ARCH     ?= universal
WINDOWS_ARCH ?= amd64
LINUX_ARCH   ?= amd64

# Shared wails build flags: keep committed bindings, skip the frontend
# (built by the `frontend` target) and mod tidy.
WAILS_FLAGS := -skipbindings -m -s

UNAME_S := $(shell uname -s)

.DEFAULT_GOAL := help

.PHONY: help version tools nsis frontend test require-linux restore \
        release release-all release-macos release-windows release-linux \
        clean distclean

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2}'

version: ## Print the version this build would use
	@test -n "$(APP_VERSION)" || { echo "error: could not read productVersion from wails.json" >&2; exit 1; }
	@echo "$(APP_VERSION)"

tools: ## Ensure the wails CLI is installed
	@if command -v wails >/dev/null 2>&1; then \
		echo "wails: $$(command -v wails)"; \
	elif [ -x "$(WAILS)" ]; then \
		echo "wails: $(WAILS)"; \
	else \
		echo "Installing wails $(WAILS_VER)..."; \
		$(GO) install github.com/wailsapp/wails/v2/cmd/wails@$(WAILS_VER); \
	fi

nsis: ## Ensure NSIS (makensis) is available for Windows installers
	@if command -v makensis >/dev/null 2>&1; then \
		echo "NSIS: $$(command -v makensis)"; \
	elif command -v brew >/dev/null 2>&1; then \
		echo "Installing NSIS via Homebrew..."; brew install nsis; \
	elif command -v apt-get >/dev/null 2>&1; then \
		echo "Installing NSIS via apt..."; sudo apt-get update && sudo apt-get install -y nsis; \
	elif command -v choco >/dev/null 2>&1; then \
		echo "Installing NSIS via choco..."; choco install nsis -y; \
	else \
		echo "error: makensis not found and no package manager available." >&2; \
		echo "Install NSIS first: brew install nsis | apt-get install nsis | choco install nsis" >&2; \
		exit 1; \
	fi

frontend: ## Build the Svelte frontend (npm install + npm run build)
	@cd frontend && npm install --no-audit --no-fund
	@cd frontend && npm run build

test: ## Run the Go unit tests
	@$(GO) test ./rimage/... ./rmanager/...

release: clean ## Build release artifacts for the current host
	@case "$(UNAME_S)" in \
		Darwin) $(MAKE) release-macos release-windows; \
			echo; echo "Note: Linux is not cross-buildable from macOS."; \
			echo "      Run 'make release-linux' on a Linux host, or push a v* tag to use .github/workflows/release.yml." ;; \
		Linux)  $(MAKE) release-linux release-windows ;; \
		MINGW*|MSYS*|CYGWIN*) $(MAKE) release-windows ;; \
		*) echo "error: unsupported host $(UNAME_S)" >&2; exit 1 ;; \
	esac
	@$(MAKE) restore
	@echo; echo "Release artifacts are in build/bin/"

release-all: clean ## Attempt every platform (Linux step needs a Linux host)
	$(MAKE) release-macos release-windows
	@echo; echo "=== Linux ==="
	@$(MAKE) release-linux || { echo "Linux skipped: not on a Linux host (see .github/workflows/release.yml)."; true; }
	@$(MAKE) restore

restore: ## Remove the package-lock.json that npm creates during a build
	rm -f frontend/package-lock.json

release-macos: tools frontend ## Build + zip the macOS app ($(MAC_ARCH))
	@test -n "$(APP_VERSION)" || { echo "error: cannot determine version" >&2; exit 1; }
	$(WAILS) build $(WAILS_FLAGS) -platform darwin/$(MAC_ARCH)
	@test -d "build/bin/$(APP_NAME).app" || { echo "error: $(APP_NAME).app was not produced" >&2; exit 1; }
	@cd build/bin && COPYFILE_DISABLE=1 zip -r -q -X "$(APP_NAME)-macOS-$(MAC_ARCH)-$(APP_VERSION).zip" "$(APP_NAME).app"
	@echo "macOS: build/bin/$(APP_NAME)-macOS-$(MAC_ARCH)-$(APP_VERSION).zip"

release-windows: tools frontend nsis ## Build the Windows exe + NSIS installer ($(WINDOWS_ARCH))
	@test -n "$(APP_VERSION)" || { echo "error: cannot determine version" >&2; exit 1; }
	$(WAILS) build $(WAILS_FLAGS) -platform windows/$(WINDOWS_ARCH) -nsis
	@ls build/bin/$(APP_NAME)-*-installer.exe >/dev/null 2>&1 || { echo "error: NSIS installer was not produced" >&2; exit 1; }
	@echo "Windows: build/bin/$(APP_NAME)-$(WINDOWS_ARCH)-installer.exe"
	@if command -v zip >/dev/null 2>&1; then \
		cd build/bin && zip -q -X "$(APP_NAME)-windows-$(WINDOWS_ARCH)-$(APP_VERSION)-portable.zip" "$(APP_NAME).exe" && \
		echo "Windows portable: build/bin/$(APP_NAME)-windows-$(WINDOWS_ARCH)-$(APP_VERSION)-portable.zip"; \
	else \
		echo "Portable zip skipped (zip not available on this host); the installer is the Windows artifact."; \
	fi

require-linux: ## Ensure we are on a Linux host (webkit2gtk system libraries)
	@if [ "$(UNAME_S)" != "Linux" ]; then \
		echo "error: Linux builds must run on a Linux host (webkit2gtk system libraries)." >&2; \
		echo "  Run 'make release-linux' in a Linux VM/container, or push a v* tag to" >&2; \
		echo "  trigger .github/workflows/release.yml, which builds Linux in CI." >&2; \
		exit 1; \
	fi

release-linux: require-linux tools frontend ## Build the Linux binary ($(LINUX_ARCH)); requires a Linux host
	@test -n "$(APP_VERSION)" || { echo "error: cannot determine version" >&2; exit 1; }
	$(WAILS) build $(WAILS_FLAGS) -platform linux/$(LINUX_ARCH)
	@BIN=$$(ls build/bin/$(APP_NAME) build/bin/$(shell echo $(APP_NAME) | tr 'A-Z' 'a-z') 2>/dev/null | head -1); \
	[ -n "$$BIN" ] || { echo "error: linux binary was not produced" >&2; exit 1; }; \
	tar -czf build/bin/$(APP_NAME)-linux-$(LINUX_ARCH)-$(APP_VERSION).tar.gz -C build/bin "$$(basename $$BIN)"
	@echo "Linux: build/bin/$(APP_NAME)-linux-$(LINUX_ARCH)-$(APP_VERSION).tar.gz"

clean: ## Remove build outputs (keeps node_modules and tooling)
	rm -rf build/bin

distclean: clean ## Remove build outputs, frontend/dist and node_modules
	rm -rf frontend/dist frontend/node_modules
