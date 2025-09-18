FROM node:20-alpine AS frontend
WORKDIR /app
COPY package*.json ./
RUN npm install --frozen-lockfile
COPY . .
RUN npm run build

FROM golang:1.22.7-alpine AS backend
WORKDIR /app

RUN if grep -i -q alpine /etc/issue; then \
    apk add --no-cache gcc g++; \
    fi

ENV CGO_ENABLED=1

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=frontend /app/dist ./dist

RUN go build -o server pkg/main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=backend /app/server .
COPY --from=backend /app/dist ./dist

EXPOSE 3838
CMD ["./server"]
