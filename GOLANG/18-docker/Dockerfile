FROM golang:1.23.5 AS dev
WORKDIR /work

FROM golang:1.23.5 AS build
WORKDIR /app
COPY ./app/* /app/
RUN go build -o app


FROM alpine AS runtime
COPY --from=build /app/app /
CMD ./app




