#)####################################################################################
#   _____ _           _   _                    _   _____            _               #
#  / ____| |         | | | |                  | | |  __ \          | |              #
# | (___ | |__   __ _| |_| |_ ___ _ __ ___  __| | | |__) |___  __ _| |_ __ ___  ___ #
#  \___ \| '_ \ / _` | __| __/ _ \ '__/ _ \/ _` | |  _  // _ \/ _` | | '_ ` _ \/ __|#
#  ____) | | | | (_| | |_| ||  __/ | |  __/ (_| | | | \ \  __/ (_| | | | | | | \__ \#
# |_____/|_| |_|\__,_|\__|\__\___|_|  \___|\__,_| |_|  \_\___|\__,_|_|_| |_| |_|___/#
#####################################################################################

#
# Makefile for building, running, and testing
#

APP_NAME = gameserver-service

# Import dotenv
ifneq (,$(wildcard ../.env))
	include ../.env
	export
endif

# Application versions
BASE_VERSION = $(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')
COMMIT_HASH = $(shell git rev-parse --short HEAD)


# Gets the directory containing the Makefile
ROOT_DIR = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

# Base container registry
SRO_BASE_REGISTRY ?= 779965382548.dkr.ecr.us-east-1.amazonaws.com
SRO_REGISTRY ?= $(SRO_BASE_REGISTRY)/sro

# The registry for this service
REGISTRY = $(SRO_REGISTRY)/$(APP_NAME)
time=$(shell date +%s)

PROTO_DIR=$(ROOT_DIR)/api

PROTO_FILES = "$(PROTO_DIR)/sro/gameserver/dimension.proto" "$(PROTO_DIR)/sro/gameserver/map.proto" "$(PROTO_DIR)/sro/gameserver/connection.proto" "$(PROTO_DIR)/sro/gameserver/data.proto"

MOCK_INTERFACES = $(shell egrep -rl --include="*.go" "type (\w*) interface {" $(ROOT_DIR)/pkg | sed "s/.go$$//")

# Versioning
VERSION=$(BASE_VERSION)
ifeq ($(VERSION),)
	VERSION := 0.0.0
endif

VERSION_PARTS=$(subst ., ,$(VERSION))
MAJOR_VERSION=$(word 1,$(VERSION_PARTS))
MINOR_VERSION=$(word 2,$(VERSION_PARTS))
PATCH_VERSION=$(word 3,$(VERSION_PARTS))

#   _____                    _
#  |_   _|                  | |
#    | | __ _ _ __ __ _  ___| |_ ___
#    | |/ _` | '__/ _` |/ _ \ __/ __|
#    | | (_| | | | (_| |  __/ |_\__ \
#    \_/\__,_|_|  \__, |\___|\__|___/
#                  __/ |
#                 |___/

.PHONY: test report mocks clean-mocks report-watch $(APP_NAME)
test:
	ginkgo --race -p --cover -covermode atomic -coverprofile=coverage.out --output-dir $(ROOT_DIR)/ $(ROOT_DIR)/pkg/...

test-watch:
	ginkgo watch --race -p --cover -covermode atomic -output-dir=$(ROOT_DIR) $(ROOT_DIR)/...

report: test
	go tool cover -func=$(ROOT_DIR)/coverage.out
	go tool cover -html=$(ROOT_DIR)/coverage.out

report-watch:
	while inotifywait -e close_write $(ROOT_DIR)/coverage.out; do \
		go tool cover -func=$(ROOT_DIR)/coverage.out; \
		go tool cover -html=$(ROOT_DIR)/coverage.out; \
	done

dev-watch: test-watch report-watch

mocks: $(MOCK_INTERFACES)
$(MOCK_INTERFACES):
	rm -rf "$(@D)/mocks"
	mockgen \
		-source="$@.go" \
		-destination="$(@D)/mocks/$(@F).go"

build: 
	go build -ldflags="-X 'github.com/ShatteredRealms/$(APP_NAME)/pkg/config/default.Version=$(BASE_VERSION)'" -o $(ROOT_DIR)/bin/$(APP_NAME) $(ROOT_DIR)/cmd/$(APP_NAME)  

run:
	go run $(ROOT_DIR)/cmd/$(APP_NAME)

run-watch:
	gow run $(ROOT_DIR)/cmd/$(APP_NAME)

deploy: aws-docker-login push

docker:
	docker build --build-arg APP_VERSION=$(BASE_VERSION) -t sro-$(APP_NAME) -f Dockerfile .

aws-docker-login:
	aws ecr get-login-password | docker login --username AWS --password-stdin $(SRO_BASE_REGISTRY)

push:
	docker tag sro-$(APP_NAME) $(SRO_REGISTRY)/$(APP_NAME):latest
	docker tag sro-$(APP_NAME) $(SRO_REGISTRY)/$(APP_NAME):$(BASE_VERSION)
	docker tag sro-$(APP_NAME) $(SRO_REGISTRY)/$(APP_NAME):$(BASE_VERSION)-$(COMMIT_HASH)
	docker push $(SRO_REGISTRY)/$(APP_NAME):latest
	docker push $(SRO_REGISTRY)/$(APP_NAME):$(BASE_VERSION)
	docker push $(SRO_REGISTRY)/$(APP_NAME):$(BASE_VERSION)-$(COMMIT_HASH)

docker-push: docker push

build-image-push: build-image push 

.PHONY: clean-protos protos $(PROTO_FILES)

clean-protos:
	rm -rf "$(ROOT_DIR)/pkg/pb"
	git submodule update --remote --merge

protos: clean-protos $(PROTO_FILES) move-protos mocks

$(PROTO_FILES):
	protoc "$@" \
		-I "$(PROTO_DIR)" \
		--go_out="$(ROOT_DIR)" \
		--go-grpc_out="$(ROOT_DIR)" \
		--grpc-gateway_out="$(ROOT_DIR)" \
		--grpc-gateway_opt "logtostderr=true" \
		--openapi_out="$(ROOT_DIR)"

move-protos:
	mv -v "$(ROOT_DIR)/github.com/ShatteredRealms/$(APP_NAME)/pkg/pb" "$(ROOT_DIR)/pkg/"
	rm -r "$(ROOT_DIR)/github.com"

install-tools:
	  cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %@latest

git: git-patch
git-major:
	git tag v$(shell echo $(MAJOR_VERSION)+1 | bc).0.0
	git push
	git push --tags
git-minor:
	echo $(MINOR_VERSION)
	git tag v$(MAJOR_VERSION).$(shell echo $(MINOR_VERSION)+1 | bc).0 
	git push
	git push --tags
git-patch:
	git tag v$(MAJOR_VERSION).$(MINOR_VERSION).$(shell echo $(PATCH_VERSION)+1 | bc)
	git push
	git push --tags

