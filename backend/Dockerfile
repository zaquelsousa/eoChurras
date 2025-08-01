FROM golang:1.24.4

#dir dentro do conteiner que vai ficar o projeto
WORKDIR /app


# pra fazer os go get pai, nos vai aproveitar o cache
COPY go.mod go.sum ./
RUN go mod download


# resto do copdigo
# sa porra e inteligente para carai pq copia . do deslino meu pc pro conteiner .
COPY . .

# compila sa merda
RUN go build -o app

EXPOSE 8081

# paizao lembra e como se vc tivesse digitando os comandas, entao tem q colocar espaço entre os comandas
CMD ["./app"]
