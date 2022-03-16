GO_BUILD ?= go build
RM.      := rm -rf

.PHONY: binaries clean install

build: build/gopen-amd64 build/gopen-linux-arm64 build/gopen-darwin-amd64 build/gopen-windows-386.exe build/gopen-freebsd-amd64

clean:
	@$(RM) gopen-amd64 gopen-linux-arm64 gopen-darwin-amd64 gopen-windows-386.exe gopen-freebsd-amd64

build/gopen-amd64:
	GOOS=darwin GOARCH=amd64 $(GO_BUILD)  -o $@ .

build/gopen-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GO_BUILD) -o $@ .

build/gopen-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o $@ .

build/gopen-windows-386.exe:
	GOOS=windows GOARCH=386 $(GO_BUILD) -o $@ .

build/gopen-freebsd-amd64:
	GOOS=freebsd GOARCH=amd64 $(GO_BUILD) -o $@ .

artifact: build
	@find build/ -name 'gopen*' -exec tar zcvf {}.tar.gz {} \;
