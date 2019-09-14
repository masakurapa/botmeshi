OS := linux
ARCH := amd64

build:
	rm -rf built/
	make build_event
	make build_interactive
	make build_search

build_event:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o built/event cmd/event/main.go
	zip -j built/event.zip built/event

build_interactive:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o built/interactive cmd/interactive/main.go
	zip -j  built/interactive.zip built/interactive

build_search:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o built/invoke-search cmd/search/main.go
	zip -j built/invoke-search.zip built/invoke-search
