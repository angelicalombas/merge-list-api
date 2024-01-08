# Use uma imagem base do Golang
FROM golang:latest

# Copie o código fonte para o contêiner
WORKDIR /go/src/app
COPY . .

# Instale as dependências
RUN go get -d -v ./...
RUN go install -v ./...

# Compile o código
RUN go build -o main .

# Exponha a porta que a aplicação irá escutar
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./main"]