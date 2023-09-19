# SCRAPER TABLE

##
Este projeto tem como objetivo extrair tabelas de qualquer site e retornar esses dados em um arquivo CSV

## Instalação

Para executar este projeto, siga as etapas abaixo:

### Clonar o repositório

Você pode clonar o repositório usando o seguinte comando:

```bash
git clone https://github.com/matheusferreira165/tablescraper.git
```

### Utilizando Docker Compose

Certifique-se de que o docker e docker-compose estao instalados na maquina, mais informacoes em: https://docs.docker.com/get-docker/

```bash
docker-compose up
```
Isso iniciará a api na porta 3000

### Sem Docker

1. Fazer a build do projeto
    ```bash 
    go build -o server ./cmd
    ```
2. Executar o arquivo gerado
   ```bash
   ./server
   ```

## Uso

Após a instalação e execução bem-sucedidas, você pode usar a API para realizar a autenticação com JWT. Aqui estão algumas operações de exemplo:

- **Gerar Token e URL para download do arquivo:** Envie uma solicitação POST para `/api/v1/download` com o link do site EX:
   ```json
    {
	"link":"https://www.w3schools.com/html/html_tables.asp"
    }
   ```
   Caso tudo derto certo voce tera um resultado parecido com esse
   ```json
   {
	"token": "160467c75a7081",
	"url": "/api/v1/download/160467c75a7081"
   }
   ```
   Sera gerado um token e uma url pronta com esse token, o token esta associado com o arquivo CSV que foi gerado na pasta ./data.

- **Fazer Download:** Envie uma solicitação GET ou abra o link no navegador para `/api/v1/download/{token}` e o download do arquivo sera executado:
  Caso faça a requicao http get, voce tera o preview do arquivo;
  Caso executa direto no navegador, o arquivo sera baixado;

### Pacotes utilizados

Neste projeto, foram utilizadas as seguintes tecnologias:

- Go (Golang)
- net/http - Response e Request
- GoQuery - Para extrair a tabela pelas tags html 
- Gorila Mux - Roteamento
- Cors - Autorizacao
- IO - Input e output para algumas de nossas responses
- OS  Manipulacao de arquivos do sistema operacional 
- Encoding/CSV - Ler e escrever no nosso arquivo CSV
- Crypto/Rand - Para gerar nosso token aleatorio
- 