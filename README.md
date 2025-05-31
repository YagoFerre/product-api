# Go API com PostgreSQL

Este projeto utiliza Docker Compose para orquestrar um ambiente de desenvolvimento contendo uma API desenvolvida em Go e um banco de dados PostgreSQL.

## Estrutura do Projeto

O arquivo `docker-compose.yml` define dois serviços:

- `go_api`: Um serviço que roda uma API em Go.
- `go_db`: Um serviço de banco de dados PostgreSQL.

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
