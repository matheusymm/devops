# Devops - Mood Tracker

- Projeto elaborado para a disciplina de DevOps

    - Consiste em uma aplicação web simples, que permite ao usuário registrar como está se sentindo no dia.

    - Permite o cadastro de novos usuários.

    - Permite o login de usuários.

    - Permite o registro de sentimentos por usuário.

    - Permite a listagem dos sentimentos por usuário.

## Arquitetura Docker

- **PostgreSQL:** é um servidor de banco de dados relacional, responsável por armazenar os dados da aplicação, usuários e sentimentos. Tem relacionamento somente com o backend, o qual realiza operações de criação, leitura, atualização e remoção dos dados.

- **Backend:** é a API da aplicação, implementando a lógica de negócios e gerenciamento de requisições e respostas aos seus endpoints. Tem relacionamentos com o PostgreSQL, para fazer operações sobre os dados, e com o Caddy, o qual encaminha requisições com URL http://localhost/api/**, originadas do Frontend.

- **Frontend:** é a interface do usuário (UI) da aplicação web, responsável por permitir a interação do usuário com a aplicação, criando uma conta e adicionando como está se sentindo no dia. Tem relacionamento com o Caddy, o qual encaminha as requisições aos endpoints da API com origem do Frontend e também torna a aplicação acessível via web browser.

- **Caddy:** é um proxy reverso, responsável por tornar a aplicação acessível via web browser e encaminha as requisições com /api/* ao serviço Backend e /* para o frontend. 

- **Obs.:** no cluster minikube, é utilizado o ingress para expor a aplicação ao acesso externo, de forma que apenas o frontend é acessível via browser e esse serviço quem trata as requisições para o backend por meio de um nginx.

### Artefatos Kubernetes 

- **PostgreSQL**:

    - `configmap.yaml`: define configurações de ambiente não sensíveis, como host, nome e script de inicialização do banco de dados.

    - `deployment.yaml`: define as configurações do pod de banco de dados.

    - `pv.yaml`: define o espaço de armazenamento disponível para o banco de dados.

    - `pvc.yaml`:  define uma solicitação de armazenamento para o banco de dados.

    - `secret.yaml`: armazena e gerencia dados sensíveis do banco de dados, como usuário e senha de acesso.

    - `service.yaml`: define o pod de banco de dados como um ClusterIP.

- **Backend**:

    - `deployment.yaml`: define as configurações do pod de backend.

    - `secret.yaml`: armazena e gerencia dados sensíveis do backend, como jwt secret e a url de acesso ao banco de dados.

    - `service.yaml`: define o pod de backend como um ClusterIP.

- **Frontend**:

    - `deployment.yaml`: define as configurações do pod de frontend.

    - `service.yaml`: define o pod de frontend como um ClusterIP.

- **Ingress**:

    - `ingress.yaml`: define as regras de roteamento para o tráfego externo, no caso, tem apenas uma regra que encaminha para o serviço de frontend.

## Tecnologias

- Go
- React + Vite
- Postgresql
- Docker e Docker Compose
- Caddy
- Kubernetes
- Helm

## Instruções para Docker
- No diretório raiz da aplicação:

``` Bash
# Criar o arquivo .env pra o frontend
cp ./frontend/.env.example ./frontend/.env

# Subir a aplicação sem imagem
docker compose up -d
```

- Também é possível subir a aplicação sem a necessidade do arquivo .env, o qual sobe os contêineres a partir das imagens do Docker Hub:

``` Bash
docker compose up -f docker-compose-image.yml -d 
```

- Em um navegador acessar `http://localhost`

## Instruções para Kubernetes
- No diretório raiz da aplicação:

``` Bash
# Caso não tenha um minikube inicializado
./scripts/minikube-init.sh 

./scripts/minikube-up.sh
```

- Em um navegador acessar `http://k8s.local`

## Instruções para K8S
- No diretório raiz da aplicação:

``` Bash
# Caso não tenha um minikube inicializado
./scripts/minikube-init.sh 

./scripts/helm-up.sh
```

- Em um navegador acessar `http://k8s.local`

## Scripts

- Existem alguns scripts para auxiliar na criação do cluster minikube:

    - `minikube-init.sh`: inicializa o cluster minikube com os addons necessários.

    - `minikube-down.sh`: para e deleta o cluster minikube.

    - `minikube-up.sh`: faz o deployment da aplicação no cluster minikube por meio dos arquivos de configuração presentes no diretório `./k8s`.

    - `minikube-down.sh`: remove a aplicação do cluster minikube.

    - `helm-up.sh`: faz o deployment da aplicação no cluster minikube utilizando o helm. 

    - `helm-down.sh`: remove a aplicação do cluster minikube