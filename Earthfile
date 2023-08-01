VERSION 0.7

all-unit-test:
    BUILD ./libs/hello+unit-test
    BUILD ./services/one+unit-test
    BUILD ./services/two+unit-test

all-docker:
    BUILD ./services/one+release
    BUILD ./services/two+release

all-release-tag:
    FROM golang:1.17-alpine
    RUN go install github.com/maykonlf/semver-cli/cmd/semver@v1.0.2
    COPY .semver.yaml .
    RUN semver get release > version
    SAVE ARTIFACT version

all-release:
    FROM +all-docker
    COPY +all-release-tag/version .
    ARG VERSION="$(cat version)"
    SAVE IMAGE --push service:$VERSION

dev-up:
    LOCALLY
    RUN docker-compose up

dev-down:
    LOCALLY
    RUN docker-compose down