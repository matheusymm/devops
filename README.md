# Devops

- Projeto elaborado para a disciplina de DevOps
    - Consiste em uma aplicação web simples, com dois CRUDS, um para usuário e outro para os sentimentos.
    - Possui um gerenciamento dos sentimentos por usuário.

## Contêineres
- **PostgreSQL:** é um servidor de banco de dados relacional, responsável por armazenar os dados da aplicação, usuários e sentimentos. Tem relacionamento somente com o backend, o qual realiza operações de criação, leitura, atualização e remoção dos dados.
- **Backend:** é a API da aplicação, implementando a lógica de negócios e gerenciamento de requisições e respostas aos seus endpoints. Tem relacionamentos com o PostgreSQL, para fazer operações sobre os dados, e com o Caddy, o qual encaminha requisições com URL http://localhost/api/**, originadas do Frontend.
- **Frontend:** é a interface do usuário (UI) da aplicação web, responsável por permitir a interação do usuário com a aplicação, criando uma conta e adicionando como está se sentindo no dia. Tem relacionamento com o Caddy, o qual encaminha as requisições aos endpoints da API com origem do Frontend e também torna a aplicação acessível via web browser.
- **Caddy:** é um proxy reverso, responsável por tornar a aplicação acessível via web browser e encaminha as requisições com /api/* ao serviço Backend e /* para o frontend. 

## Tecnologias

- Backend em Go
- Frontend em React + Vite
- Banco de dados Postgresql
- Docker e Docker Compose
- Caddy

## Instruções
- No diretório raiz da aplicação:

``` Bash
# Criar o arquivo .env pra o frontend
cp ./frontend/.env.example ./frontend/.env

# Subir a aplicação sem imagem
docker compose up -d
```

- Também é possível subir a aplicação sem a necessidade do arquivo .env:

``` Bash
docker compose up -f docker-compose-image.yml -d 
```

- Em um navegador acessar `http://localhost`
