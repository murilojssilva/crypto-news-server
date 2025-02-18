# Usa uma imagem leve do Go
FROM golang:1.21-alpine

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos do projeto para dentro do container
COPY . .

# Baixa as dependências do projeto
RUN go mod tidy

# Compila o código-fonte e dá permissão de execução ao binário
RUN go build -o app cmd/main.go && chmod +x app

# Define a porta que o servidor vai rodar
ENV PORT=3000

# Expõe a porta
EXPOSE 3000

# Comando para iniciar o servidor
CMD ["./app"]
