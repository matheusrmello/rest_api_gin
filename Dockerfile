# FROM golang:1.21

# WORKDIR /app

# COPY . ./

# RUN go mod tidy

# RUN go build -o /api_rest

# EXPOSE 8080

# ENV POSTGRES_HOST=${DB_HOST}
# ENV POSTGRES_USER=${DB_USER}
# ENV POSTGRES_PASSWORD=r${DB_PASSWORD}
# ENV POSTGRES_DB=${DB_NAME}
# ENV POSTGRES_PORT=${DB_PORT}

# CMD [ "/api_rest" ]

 