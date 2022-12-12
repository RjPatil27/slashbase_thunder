YARN ?= yarn
GOPATH ?= $(HOME)/go
STUFFBIN ?= $(GOPATH)/bin/stuffbin

BIN := slashbase
STATIC := web

.PHONY: build
build: $(BIN)

$(STUFFBIN):
	go install github.com/knadh/stuffbin/...

$(BIN): $(shell find . -type f -name "*.go")
	go build --o ${BIN} -trimpath -ldflags="-X 'main.Build=production'"

.PHONY: dist
dist: $(STUFFBIN) build-web build pack-bin


.PHONY: build-web
build-web: 
	cd frontend; yarn build; mv out ../web 


.PHONY: pack-bin
pack-bin: $(BIN) $(STUFFBIN)
	$(STUFFBIN) -a stuff -in ${BIN} -out ${BIN} ${STATIC}