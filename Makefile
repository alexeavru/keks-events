PROJECT?=keks-events
APP?=keksevents
PORT?=8080

RELEASE?=v0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CONTAINER_IMAGE?=alexeav/${APP}

GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .
	rm -f ${APP}

run: container
	docker stop $(APP) || true && docker rm $(APP) || true
	docker run -d --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		-v $(PWD)/events.db:/events.db \
		$(CONTAINER_IMAGE):$(RELEASE)

test:
	go test -v -race ./...

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)
