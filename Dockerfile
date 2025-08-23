FROM golang:1.25 AS builder
WORKDIR /app

ARG GIT_BUILD_COMMIT_HASH="unknown"
ARG GIT_BUILD_COMMIT_MODIFIED="false"
ARG GIT_BUILD_LATEST_VERSION_TAG="latest"

ENV GOMODCACHE=/gomod-cache
ENV GOCACHE=/go-cache

RUN go env -w GOCACHE=${GOCACHE}
RUN go env -w GOMODCACHE=${GOMODCACHE}
RUN go env -w GOFLAGS=-mod=mod

COPY --chmod=755 go.mod /app/go.mod
COPY --chmod=755 go.sum /app/go.sum

# install dependencies
RUN --mount=type=cache,target=${GOMODCACHE} --mount=type=cache,target=/root/.config/go \ 
  go mod download

# copy code and build it
COPY . .

RUN --mount=type=cache,target=${GOMODCACHE} --mount=type=cache,target=$GOCACHE \
  CGO_ENABLED=0 \
  go build -o main src/main.go

FROM clamav/clamav:latest

COPY --from=builder --chmod=755 /app/main /usr/local/bin/
COPY --from=builder --chmod=755 /app/scripts/startup.sh /usr/local/bin/startup.sh

ENTRYPOINT [ "sh", "/usr/local/bin/startup.sh" ]