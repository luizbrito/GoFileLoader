# Go File Load

## Pre Requisitos

ferramentas: Docker, docker-compose, git, go

clone repository: git clone --recursive https://github.com/luizbrito/goFileLoad.git

## Instalação

Execute o comando docker-compose up para o serviço.

$ docker-compose up --build

start service:
$ docker-compose up

## Exucutando a aplicação

Pode ser acessado via tela pela URL :

http://localhost:2020/upload

ou utilizando insomnia, neste caso faça uma requisição post passando o arquivo csv/txt

### Teste

Para realizar teste pegue um dos arquivos que estão no diretório /tmp/in

### Execução Local

Acesse o Diretorio local e execute os comandos abaixo:
$ go mod tidy
$ go run .

A aplicação estará disponivel na URL já citada acima.

### Resultado

Para verificar todos os registros salvos execute no seu client de banco de dados a consulta:

SELECT id, cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, cnpj_loja_mais_frequente, cnpj_loja_ultima_compra
FROM public.people;

Para verificar os dados com erro, execute a consulta abaixo:

SELECT id, chave, message FROM public.error;
