# Albert Espí­n 2019 CELSA,SL
# set variables
BINDIR = ./bin
VERSION = latest
BINARYFILE = book
DOCKERFILE = book
DOCKERREGISTRY = ea3hsp
DOCKERPUSH = $(DOCKERREGISTRY)/$(DOCKERFILE):$(VERSION)
ARCH = amd64
OS = linux

.PHONY: clean

build:
	go build -o $(BINDIR)/$(BINARYFILE) -i cmd/main.go
run:build
	$(BINDIR)/$(DOCKERFILE)
clean:
	rm -f .$(BINDIR)/$(BINARYFILE)
docker:
	env GOARCH=$(ARCH) GOOS=$(OS) go build -o $(BINDIR)/$(BINARYFILE) -i cmd/main.go
	docker build -t $(DOCKERFILE) .
docker-push:
	docker tag $(DOCKERFILE) $(DOCKERPUSH)
	docker push $(DOCKERPUSH)