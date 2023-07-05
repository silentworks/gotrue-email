CHECK_FILES?=./...
FLAGS?=-ldflags "-X github.com/supabase/gotrue/internal/utilities.Version=`git describe --tags`" -buildvcs=false
VERSION = alpha

build: deps ## Build the binary.
	CGO_ENABLED=0 go build $(FLAGS)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(FLAGS) -o gotrue-email-arm64

dev-deps: ## Install developer dependencies
	@go install github.com/gobuffalo/pop/soda@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

deps: ## Install dependencies.
	@go mod download
	@go mod verify

vet: # Vet the code
	go vet $(CHECK_FILES)

sec: dev-deps # Check for security vulnerabilities
	gosec -quiet -exclude-generated $(CHECK_FILES)
	gosec -quiet -tests -exclude-generated -exclude=G104 $(CHECK_FILES)

docker-build:
	docker build --rm -t gotrue-email:$(VERSION) .

docker-run:
	docker run -d -p 8080:3000 --name gotrue-email-$(VERSION) gotrue-email:$(VERSION)

format:
	gofmt -s -w .
