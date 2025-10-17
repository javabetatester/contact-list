# API de Lista de Contatos

Esta é uma API para gerenciar uma lista de contatos e grupos de contatos, desenvolvida em Go com o framework Gin.

## Funcionalidades

- Gerenciamento completo de Contatos (CRUD)
- Gerenciamento completo de Grupos (CRUD)
- Adicionar contatos a grupos

## Pré-requisitos

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)

## Configuração do Ambiente

1.  **Clone o repositório:**

    ```bash
    git clone <URL_DO_SEU_REPOSITORIO>
    cd NEWPROJ
    ```

2.  **Instale as dependências:**

    ```bash
    go mod tidy
    ```

3.  **Configure o Banco de Dados:**

    A aplicação utiliza PostgreSQL. A string de conexão está definida no arquivo `internal/infrastructure/db.go` e configurada como:

    ```go
    dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
    ```

    Certifique-se de que você tenha um servidor PostgreSQL rodando com essas credenciais ou altere a string `dsn` para corresponder à sua configuração.

    A aplicação irá criar automaticamente as tabelas `contacts`, `groups` e `group_contacts` ao ser iniciada, graças à função `AutoMigrate`.

## Como Executar

Para iniciar o servidor, execute o seguinte comando na raiz do projeto:

```bash
go run cmd/api/main.go
```

O servidor estará disponível em `http://localhost:8080`.

## Endpoints da API

A seguir estão os endpoints disponíveis na API.

### Contatos

-   **`POST /contact`**: Cria um novo contato.

    *Corpo da Requisição:*

    ```json
    {
        "name": "Nome do Contato",
        "email": "contato@exemplo.com",
        "telefone": "(00) 99999-9999"
    }
    ```

-   **`GET /contact`**: Lista todos os contatos.

-   **`GET /contact/:id`**: Busca um contato pelo ID.

-   **`GET /contact/name`**: Busca um contato pelo nome.

    *Corpo da Requisição:*
    ```json
    {
        "name": "Nome do Contato"
    }
    ```

-   **`GET /contact/phone`**: Busca um contato pelo telefone.

    *Corpo da Requisição:*
    ```json
    {
        "telefone": "(00) 99999-9999"
    }
    ```

-   **`GET /contact/email`**: Busca um contato pelo email.

    *Corpo da Requisição:*
    ```json
    {
        "email": "contato@exemplo.com"
    }
    ```

-   **`PUT /contact/:id`**: Atualiza um contato existente.

    *Corpo da Requisição:*

    ```json
    {
        "name": "Novo Nome",
        "email": "novo.email@exemplo.com",
        "telefone": "(00) 00000-0000"
    }
    ```

-   **`DELETE /contact/:id`**: Deleta um contato.

### Grupos

-   **`POST /group`**: Cria um novo grupo.

    *Corpo da Requisição:*

    ```json
    {
        "name": "Nome do Grupo"
    }
    ```

-   **`GET /group`**: Lista todos os grupos.

-   **`GET /group/:id`**: Busca um grupo pelo ID.

-   **`GET /group/name`**: Busca um grupo pelo nome.

    *Corpo da Requisição:*
    ```json
    {
        "name": "Nome do Grupo"
    }
    ```

-   **`PUT /group/:id`**: Atualiza um grupo existente.

    *Corpo da Requisição:*

    ```json
    {
        "name": "Novo Nome do Grupo"
    }
    ```

-   **`DELETE /group/:id`**: Deleta um grupo.

-   **`PUT /group/:id/contact/:contactId`**: Adiciona um contato a um grupo.
    -   `:id` é o ID do grupo.
    -   `:contactId` é o ID do contato.
    -   Nenhum corpo de requisição é necessário.
````

Se precisar de mais alguma coisa, é só me dizer
