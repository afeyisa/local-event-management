FROM golang:1.23.2

WORKDIR /backend
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend/app
EXPOSE 4000
CMD ["/backend/app"]
#1 Build the image
# docker build --tag backend .

#2 View local images
# docker image ls

# 3Tag images
# docker image tag backend:latest backend:v1.0

# 4 untag images
# docker image rm docker-gs-ping:v1.0
# Untagged: docker-gs-ping:v1.0

# 5 Run the image
# docker run --publish 4000:4000 backend
# 6 Run the image in detached mode
# docker run -d -p 8080:8080 backend