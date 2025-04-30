# go-stress-test

Uma ferramenta de linha de comando (CLI) escrita em Go para realizar testes de carga em serviços web.
Permite que o usuário especifique a URL do serviço, o número total de requisições e o nível de concorrência desejado.

## 📋 Descrição

Este projeto foi desenvolvido como parte de um desafio técnico com o objetivo de criar uma aplicação capaz de:

- Realizar requisições HTTP para uma URL especificada.
- Distribuir as requisições de acordo com o nível de concorrência definido.
- Garantir que o número total de requisições seja cumprido.
- Gerar um relatório ao final dos testes contendo:
  - Tempo total gasto na execução.
  - Quantidade total de requisições realizadas.
  - Quantidade de requisições com status HTTP 200.
  - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## 🚀 Como Usar

### Pré-requisitos

- [Go](https://golang.org/dl/) instalado (versão 1.16 ou superior).
- [Docker](https://www.docker.com/get-started) instalado (opcional, para execução via container).

### Compilando o Projeto

1. Clone o repositório:
   ```bash
   git clone https://github.com/brunocordeiro180/go-stress-test.git
   cd go-stress-test
   ```

2. Compile o projeto:
   ```bash
   go build -o go-stress-test main.go
   ```

3. Execute a aplicação:
   ```bash
   ./go-stress-test --url=http://example.com --requests=1000 --concurrency=10
   ```

### Executando via Docker

Execute o container:
   ```bash
   docker run brunocordeiro180/go-stress-test:latest --url=http://google.com --requests=1000 --concurrency=10
   ```

## ⚙️ Parâmetros da Linha de Comando

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requisições a serem realizadas.
- `--concurrency`: Número de requisições simultâneas (nível de concorrência).

## 📝 Exemplo de Saída

```bash
==== Relatório de Teste de Carga ====
Tempo total gasto: 1.369163076s
Total de requests realizados: 100
Requests com status HTTP 200: 100
Distribuição de códigos de status:
  200: 100
```

## 🛠️ Estrutura do Projeto

- `main.go`: Arquivo principal que contém a lógica de execução do teste de carga.
- `Dockerfile`: Define a imagem Docker para a aplicação.
- `go.mod` e `go.sum`: Gerenciam as dependências do projeto.

## 📄 Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
