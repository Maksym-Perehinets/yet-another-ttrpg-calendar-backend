# Base image with Go
FROM cosmtrek/air AS dev

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Default command to run the application using Air
CMD ["air"]
