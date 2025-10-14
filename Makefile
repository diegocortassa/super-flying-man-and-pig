MAIN_NAME=super-flying-man-and-pig
GO_PACKAGE=github.com/diegocortassa/$(MAIN_NAME)

# Version information
VERSION=$(shell cat VERSION)
RELEASE=1
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Build variables
GO=go
GOROOT=$(shell bash -c "go env GOROOT")
LDFLAGS="-X $(GO_PACKAGE)/version.Version=$(VERSION) \
         -X $(GO_PACKAGE)/version.Commit=$(COMMIT) \
         -X $(GO_PACKAGE)/version.BuildTime=$(BUILD_TIME)"
LDFLAGS_WIN=-X=$(GO_PACKAGE)/version.Version=$(VERSION),-X=$(GO_PACKAGE)/internal/version.Commit=$(COMMIT),-X=$(GO_PACKAGE)/internal/version.BuildTime=$(BUILD_TIME)

# Platform-specific variables
LINUX_AMD64_BINARY=$(MAIN_NAME)
LINUX_AMD64_DIR=$(MAIN_NAME)-v$(VERSION)-linux-amd64
WINDOWS_AMD64_BINARY=$(MAIN_NAME).exe
WINDOWS_AMD64_DIR=$(MAIN_NAME)-v$(VERSION)-windows-amd64
WASM_BINARY=$(MAIN_NAME).wasm
WASM_DIR=$(MAIN_NAME)-v$(VERSION)-wasm

# Build all platforms
.PHONY: build
build: build-linux build-windows build-wasm

# Build for Linux
.PHONY: build-linux
build-linux:
	mkdir -p dist/$(LINUX_AMD64_DIR)
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags $(LDFLAGS) -o dist/$(LINUX_AMD64_DIR)/$(LINUX_AMD64_BINARY) .
	cp README.md LICENSE dist/$(LINUX_AMD64_DIR)/
	cd dist && tar czf $(LINUX_AMD64_DIR).tar.gz $(LINUX_AMD64_DIR)

# Build for Windows
.PHONY: build-windows
build-windows:
	go-winres simply --product-version $(VERSION).0 --file-version $(VERSION).0 --file-description "An old style shoot'em up written in Go" --product-name "Super Flying Man And Pig" --copyright "Diego Cortassa" --original-filename "$(WINDOWS_BINARY)" --icon assets/Icon.png
	mkdir -p dist/$(WINDOWS_AMD64_DIR)
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags $(LDFLAGS) -o dist/$(WINDOWS_AMD64_DIR)/$(WINDOWS_AMD64_BINARY) .
	cp README.md LICENSE dist/$(WINDOWS_AMD64_DIR)/
	cd dist && 7z a -bd -r $(WINDOWS_AMD64_DIR).zip $(WINDOWS_AMD64_DIR)

# Build for WEB
.PHONY: build-wasm
build-wasm:
	mkdir -p dist/$(WASM_DIR)
	env GOOS=js GOARCH=wasm $(GO) build -ldflags $(LDFLAGS) -o dist/$(WASM_DIR)/$(WASM_BINARY) .
	cp README.md LICENSE assets/index.html dist/$(WASM_DIR)/
	cp $(GOROOT)/lib/wasm/wasm_exec.js dist/$(WASM_DIR)/
	cd dist && tar czf $(WASM_DIR).tar.gz $(WASM_DIR)

# Print version
.PHONY: version
version:
	@echo $(VERSION)

# Bump version (patch by default)
.PHONY:
version-bump: 
	@current_version=`cat VERSION`; \
	major=`echo $$current_version | cut -d. -f1`; \
	minor=`echo $$current_version | cut -d. -f2`; \
	patch=`echo $$current_version | cut -d. -f3`; \
	new_minor=$$((minor + 1)); \
	new_version="$$major.$$new_minor.$$patch"; \
	echo $$new_version > VERSION; \
	echo "Version bumped from $$current_version to $$new_version"
	$(MAKE) update-toml

# Create a new version tag
.PHONY: tag
tag: version
	git add VERSION
	@if git diff --quiet --cached -- VERSION; then \
		echo "VERSION up to date, tagging"; \
		git tag -a v$(VERSION) -m "Version $(VERSION)"; \
		echo "Tagged, now push to GitHub: git push origin v$(VERSION)"; \
	else \
		echo "VERSION need to be committed first"; \
	fi

PHONY: clean
clean:
	# go clean ;
	rm -rf dist
	rm -f *.syso

PHONY: run-debug
run-debug:
	go run . --debug ;

PHONY: run
run:
	go run . ;

PHONY: run-web
run-web:
	go run github.com/hajimehoshi/wasmserve@latest .
	open http://localhost:8080
