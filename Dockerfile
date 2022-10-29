# syntax=docker/dockerfile:1

FROM golang:1.19 AS build

# Create appuser.
ENV USER=appuser
ENV UID=10001 

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /simple-fiber-http


FROM scratch

# Import the user and group files from the builder.
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

COPY --from=build /app/helicopter.json /
COPY --from=build /simple-fiber-http /simple-fiber-http

USER appuser:appuser

ENTRYPOINT [ "/simple-fiber-http" ]