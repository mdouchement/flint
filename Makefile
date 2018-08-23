.PHONY: flint install clean re dir re-all all test release
.PHONY: darwin-386 darwin-amd64 linux-arm linux-arm64 linux-386 linux-amd64
.PHONY: windows-386 windows-amd64 freebsd-386 freebsd-amd64

NAME = flint
DIST_DIR = dist
REPO="github.com/astrocorp42/flint"
VERSION := $(shell cat version/version.go| grep "\sVersion" | cut -d '"' -f2)

define checksums
	echo $$(openssl sha512 $(1) | cut -d " " -f2) $$(echo $(1) | cut -d "/" -f2) >> $(2)/sha512sum$(3)
endef

define build_for_os_arch
	mkdir -p $(DIST_DIR)/$(1)-$(2)/$(VERSION)
	GOOS=$(1) GOARCH=$(2) go build \
		 -ldflags "-X $(REPO)/version.UTCBuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` \
		 -X $(REPO)/version.GitCommit=`git rev-parse HEAD` \
		 -X $(REPO)/version.GoVersion=`go version | cut -d' ' -f 3 | cut -c3-`" \
		 -o $(DIST_DIR)/$(1)-$(2)/$(VERSION)/$(NAME)$(3)
	@# binary checksums
	$(call checksums,$(DIST_DIR)/$(1)-$(2)/$(VERSION)/$(NAME)$(3),$(DIST_DIR)/$(1)-$(2)/$(VERSION),.txt)

	zip -j $(DIST_DIR)/$(NAME)-$(VERSION)-$(1)-$(2).zip $(DIST_DIR)/$(1)-$(2)/$(VERSION)/$(NAME)$(3) \
		$(DIST_DIR)/$(1)-$(2)/$(VERSION)/sha512sum.txt

	@#archive checksums
	$(call checksums,$(DIST_DIR)/$(NAME)-$(VERSION)-$(1)-$(2).zip,dist,s.txt)
endef


$(NAME): dir
	dep ensure
	go build \
		 -ldflags "-X $(REPO)/version.UTCBuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` \
		 -X $(REPO)/version.GitCommit=`git rev-parse HEAD` \
		 -X $(REPO)/version.GoVersion=`go version | cut -d' ' -f 3 | cut -c3-`" \
		 -o $(DIST_DIR)/$(NAME)

test:
	go vet $(go list ./... | grep -v /vendor/)
	go test -v -race ./...

install:
	go install

clean:
	rm -rf $(NAME) $(DIST_DIR)

re: clean $(NAME)

dir:
	mkdir -p $(DIST_DIR)

re-all: clean all

all: $(NAME) darwin-386 darwin-amd64 linux-arm linux-arm64 linux-386 linux-amd64 windows-386 windows-amd64 freebsd-386 freebsd-amd64

release: clean
	git commit -m "Releasing v$(VERSION)"
	git push
	git tag -a v$(VERSION) -m "Release v$(VERSION)"
	git push origin v$(VERSION)

darwin-386:
	$(call build_for_os_arch,darwin,386,)

darwin-amd64:
	$(call build_for_os_arch,darwin,amd64,)

linux-arm:
	$(call build_for_os_arch,linux,arm,)

linux-arm64:
	$(call build_for_os_arch,linux,arm64,)

linux-386:
	$(call build_for_os_arch,linux,386,)

linux-amd64:
	$(call build_for_os_arch,linux,amd64,)

windows-386:
	$(call build_for_os_arch,windows,386,.exe)

windows-amd64:
	$(call build_for_os_arch,windows,amd64,.exe)

freebsd-386:
	$(call build_for_os_arch,freebsd,386,)

freebsd-amd64:
	$(call build_for_os_arch,freebsd,amd64,)
