PROJECT?=keks-events
APP?=keksevents
PORT?=8080

RELEASE?=v0.0.2
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%d_%H:%M:%S')
CONTAINER_IMAGE?=alexeav/${APP}

GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X github.com/alexeavru/${PROJECT}/internal/version.Release=${RELEASE} \
		-X github.com/alexeavru/${PROJECT}/internal/version.Commit=${COMMIT} \
		-X github.com/alexeavru/${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .
	rm -f ${APP}

run: container
	docker stop $(APP) || true && docker rm $(APP) || true
	docker run -d --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(CONTAINER_IMAGE):$(RELEASE)

compose: container
	docker compose kill keks-event
	KEKS_EVENTS_VER=$(RELEASE) docker compose up -d

test:
	go test -v -race ./...

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)
