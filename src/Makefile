# defining environment variables

TARGETNAME=contact-book-service
REPONAME=contact-book-service
TARGETPATH=$(GOPATH)/bin/$(TARGETNAME)
#SERVICEPATH=$(GOPATH)/src/github.com/kanikapant/$(REPONAME)/services
#SERVICEPACKPATH=$(SERVICEPATH)/continuum
#SERVICEBINPATH=$(SERVICEPACKPATH)/userAccountService
#SERVICECONFIGPATH=$(SERVICEPACKPATH)/config
#SERVICEPACKNAME=continuum.zip
#SERVICEPACKTARGET=continuum
DB_URL="cassandra://localhost:9042/contact_book_service_db"
MIGRATIONS_DIR="../cql/migrations"

#BUILDNUMBER = $(BUILD_NUMBER)
#BUILDCOMMIT = $(GIT_COMMIT)

#LDFLAGBUILDVERSION=-X github.com/kanikapant/contact-book-service/src/services/version.BuildNumber=$(BUILDNUMBER)
#LDFLAGBUILDCOMMIT=-X github.com/ContinuumLLC/contact-book-service/src/services/version.BuildCommitSHA=$(BUILDCOMMIT)

LINTERCOMMAND=gometalinter --vendor --tests --skip=mock --exclude='_gen.go' --exclude='docs.go' --disable=gotype --disable=errcheck --disable=gas --disable=dupl --deadline=1500s --sort=linter
DOCKER_TAR_NAME=continuum_$(TARGETNAME).tar

# global command
.PHONY: all
all: dependencies contact-book-service test cover

.PHONY: windows
windows: dependencies contact-book-service-win test cover

.PHONY: mac
mac: dependencies contact-book-service-mac test cover

.PHONY: linux
linux: dependencies contact-book-service-linux test cover 

.PHONY: local
local: dependencies localbuild test code-quality

#.PHONY: localbuild
#localbuild:
#	go build -ldflags "${LDFLAGBUILDVERSION} ${LDFLAGBUILDCOMMIT}" -o $(TARGETPATH)

.PHONY: contact-book-service
contact-book-service: contact-book-service-win contact-book-service-linux contact-book-service-mac

.PHONY: contact-book-service-win
contact-book-service-win:
	GOOS=windows GOARCH=amd64 go build -o $(GOPATH)/bin/windows_amd64/$(TARGETNAME).exe

.PHONY: contact-book-service-linux
contact-book-service-linux:
	GOOS=linux GOARCH=amd64 go -o $(GOPATH)/bin/linux_amd64/$(TARGETNAME)

.PHONY: contact-book-service-mac
contact-book-service-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(GOPATH)/bin/darwin_amd64/$(TARGETNAME)

.PHONY: dependencies
dependencies:
	echo "Installing dependencies"
	glide install

packages = \
	./config \
	./logger \
	./translation \
	./messaging \
	./models \
	./services/user \
	./services/logon

.PHONY: test
test:
	@$(foreach package,$(packages), \
		set -e; \
		go test -coverprofile $(package)/cover.out -covermode=count $(package);)
		
.PHONY: cover-xml
cover-xml:
	@$(foreach package,$(packages), \
	    gocov convert $(package)/cover.out | gocov-xml > $(package)/coverage.xml;)

.PHONY: cover
cover: test
	echo "mode: count" > cover-all.out
	@$(foreach package,$(packages), \
		tail -n +2 $(package)/cover.out >> cover-all.out;)
	gocover-cobertura < cover-all.out > cover-cobertura.xml

.PHONY: show
show:
	echo "Launching web browser to show overall coverage..."
	go tool cover -html=cover-all.out
	go tool cover -func=cover-all.out

.PHONY: code-quality
code-quality:
	#$(LINTERCOMMAND) --checkstyle ./... > static-analysis.xml

.PHONY: code-quality-print
code-quality-print:
	$(LINTERCOMMAND) ./...

.PHONE: clean
clean:
	echo "Cleaning ..."
	rm -f $(TARGETPATH)
	@$(foreach package,$(packages), \
		rm -f $(package)/cover.out \
		rm -f $(package)/*.log;)
	rm -f cover-all.out
	rm -f cover-cobertura.xml

migrate:
	@migrate -database $(DB_URL) -path $(MIGRATIONS_DIR) up

migrate-down:
	@migrate -database $(DB_URL) -path $(MIGRATIONS_DIR) down