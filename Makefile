BINARY := hugoArticleEncryptor
MODULE := github.com/hotjuicew/hugoArticleEncryptor

BUILD_DIR     := build
BUILD_FLAGS   := -v

CGO_ENABLED := 0
GO111MODULE := on

GO_BUILD = GO111MODULE=$(GO111MODULE) CGO_ENABLED=$(CGO_ENABLED) \
	go build $(BUILD_FLAGS) -trimpath

UNIX_ARCH_LIST = \
	darwin-amd64 \
	darwin-arm64 \
	linux-amd64 \
	linux-arm64

WINDOWS_ARCH_LIST = \
	windows-amd64 \

all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 windows-amd64

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY)-$@

darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY)-$@

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY)-$@

linux-arm64:
	GOARCH=arm64 GOOS=linux $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY)-$@

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY)-$@.exe

clean:
	rm -rf build/