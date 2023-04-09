FROM golang:1.19-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the source code and config file into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 3000
EXPOSE 3000

# Run the command to start the app with the config file
CMD ["/app/main"]