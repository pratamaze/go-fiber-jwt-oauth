# Gunakan image dasar Golang
FROM golang:1.19

# Set working directory di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Unduh dependensi
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Ekspos port 8080
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
