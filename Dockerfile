FROM golang:alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/dotaspirit/*.go ./

RUN go build -o /dotaspirit

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /dotaspirit /dotaspirit

COPY assets/ ./assets/

EXPOSE 3682

VOLUME ["/config"]

USER nonroot:nonroot

ENTRYPOINT ["/dotaspirit"]