# syntax=docker/dockerfile:1
FROM node as node-build

WORKDIR /src

COPY client/ .
RUN npm run build

##############################################
FROM golang:alpine as go-build

WORKDIR /app

COPY server/ .
RUN go mod download && go mod verify
RUN go build -o ./server .


##############################################
FROM scratch

WORKDIR /app

COPY --from=go-build app/.env app/server ./

COPY --from=node-build src/dist ./static

EXPOSE 8080

ENTRYPOINT [ "./server" ]