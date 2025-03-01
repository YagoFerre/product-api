# README.md para Projeto Docker Compose: Go API com PostgreSQL

Este projeto utiliza Docker Compose para orquestrar um ambiente de desenvolvimento contendo uma API desenvolvida em Go e um banco de dados PostgreSQL.

## Estrutura do Projeto

O arquivo `docker-compose.yml` define dois serviços:

- `go_api`: Um serviço que roda uma API em Go.
- `go_db`: Um serviço de banco de dados PostgreSQL.

### Detalhes dos Serviços

#### go_api

- **Container Name:** `go_api`
- **Image:** `go-api` (Esta imagem é construída a partir do Dockerfile no diretório raiz.)
- **Ports:** Mapeia a porta 8000 do host para a porta 8000 do container, permitindo o acesso à API.
- **Depends On:** Define uma dependência no serviço `go_db`, assegurando que o banco de dados esteja disponível antes de iniciar a API.

#### go_db

- **Container Name:** `go_db`
- **Image:** `postgres:14`
- **Environment:**
  - `POSTGRES_PASSWORD`: Senha para o usuário do banco de dados.
  - `POSTGRES_USER`: Nome do usuário do banco de dados.
  - `POSTGRES_DB`: Nome do banco de dados.
- **Ports:** Mapeia a porta 5432 do host para a porta 5432 do container, permitindo o acesso ao banco de dados.
- **Volumes:** Persiste os dados do banco de dados usando um volume Docker nomeado `pgdata`.

### Volumes

- `pgdata`: Armazena os dados do PostgreSQL, permitindo que os dados persistam entre as instalações do container.

## Como usar

### Pré-requisitos

- Docker
- Docker Compose

### Instruções de Configuração e Execução

1. **Clone o repositório**

   Caso o código da aplicação e o Dockerfile estejam hospedados em um repositório, clone o repositório usando:

   ```bash
   git clone [URL_DO_REPOSITORIO]
   cd [NOME_DO_DIRETORIO]
   ```

2. **Construa e execute os containers**

   No diretório raiz (onde o `docker-compose.yml` está localizado), execute:

   ```bash
   docker-compose up --build
   ```

   Este comando construirá a imagem `go-api` se necessário e iniciará os serviços definidos.

3. **Acessar a API**

   Após os containers estarem rodando, a API estará acessível via `http://localhost:8000`.

4. **Parar e remover os containers**

   Para parar e remover os containers, use:

   ```bash
   docker-compose down
   ```

## Manutenção e Suporte

Para manutenção e suporte, verifique os logs dos containers, ajuste configurações conforme necessário e reconstrua os containers para aplicar mudanças. A documentação do Docker e do Docker Compose são excelentes recursos para resolver qualquer problema.

## Licença

Especifique a licença sob a qual o projeto está distribuído, se aplicável.

---

Este README fornece uma visão geral e instruções básicas para configurar e executar o projeto usando Docker e Docker Compose.