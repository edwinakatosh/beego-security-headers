FROM golang AS builder
ENV GO111MODULE=on \
    GOOS=linux GOARCH=amd64 \
    CGO_ENABLED=0
ARG COMMIT
WORKDIR /go/src/app
# Copy the code necessary to build the application
# You may want to change this to copy only what you actually need.
ADD . .
# Build the application
RUN env
RUN set -x \
    && go mod tidy -v \
    && go mod download \
    && go build -v -ldflags "-w -s" .
RUN set -x \
    && ls -Altr && pwd
# Let's create a /dist folder containing just the files necessary for runtime.
# Later, it will be copied as the / (root) of the output image.
WORKDIR /dist
RUN set -x \
    && cp -r /go/src/app/app /go/src/app/conf /go/src/app/views /dist/
# Create the minimal runtime image
FROM alpine:latest
ENV BEEGO_RUNMODE=dev
RUN set -x \
    && apk --no-cache --update upgrade
RUN set -x \
    && mkdir /app /scripts \
    # Create app user
    && adduser -D -s /bin/sh -u 1000 app \
    && sed -i -r 's/^app:!:/app:x:/' /etc/shadow
COPY --chown=1000:0 --from=builder /dist /app
USER app
WORKDIR /app
ENTRYPOINT ["/app/app"]
