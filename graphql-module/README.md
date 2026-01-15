# Projeto GraphQL - Full Cycle

Este projeto é um módulo prático do curso Full Cycle, demonstrando a implementação de uma API GraphQL utilizando Go (Golang), SQLite e a biblioteca `gqlgen`.

> ⭐️ **Curtiu o projeto?** Dê uma estrela no repositório para apoiar!

## 📦 Instalação

Antes de tudo, clone o repositório e acesse a pasta do módulo:

```bash
git clone https://github.com/rubenfabio/full-cycle-3.0-systems-communication-patterns.git
cd full-cycle-3.0-systems-communication-patterns/graphql-module
```

## 🚀 Como Rodar o Projeto

### Pré-requisitos
- **Go**: Certifique-se de ter o Go instalado (versão 1.25+ recomendada).
- **GCC**: Necessário para compilar o driver SQLite original (`go-sqlite3`), **OU** utilize o driver "pure Go" já configurado neste projeto (`glebarez/go-sqlite`) que dispensa o GCC.

### Passos para Executar

1.  **Instale as dependências:**
    ```bash
    go mod tidy
    ```

2.  **Execute o Servidor:**
    ```bash
    go run cmd/server/server.go
    ```

3.  **Acesse o Playground:**
    Abra seu navegador em [http://localhost:8080/](http://localhost:8080/).
    O Playground é uma IDE interativa para explorar a API e executar queries/mutations.

## 🗄️ Banco de Dados e Tabelas

O projeto utiliza **SQLite** como banco de dados, armazenado no arquivo `data.db` na raiz do projeto.

**Inicialização das Tabelas:**
Você **não precisa** rodar scripts SQL manuais para criar as tabelas inicialmente. A lógica de inicialização foi adicionada ao `cmd/server/server.go`. Ao iniciar o servidor, ele verifica e cria automaticamente as tabelas necessárias (`categories` e `courses`) se elas não existirem.

Caso queira inspecionar o banco manualmente via linha de comando, incluímos um utilitário:
```bash
go run cmd/sqlite/main.go "SELECT * FROM categories"
```

## 📂 Estrutura do Projeto

Aqui está uma explicação do que é cada pasta e arquivo importante:

*   **`cmd/server/server.go`**: O ponto de entrada da aplicação. Aqui inicializamos a conexão com o banco de dados, configuramos os resolvers e subimos o servidor HTTP.
*   **`graph/schema.graphqls`**: O coração do GraphQL. Define o **Schema** (tipos, queries, mutations e inputs). É o contrato da API.
*   **`graph/schema.resolvers.go`**: Onde a mágica acontece. Contém a implementação (código Go) que satisfaz as queries e mutations definidas no schema. É aqui que chamamos o banco de dados.
*   **`graph/model/`**: Contém as structs Go que representam os dados do GraphQL.
*   **`internal/database/`**: Implementação do acesso ao banco de dados (Repository Pattern). Separa a lógica de SQL da lógica do GraphQL.
    *   `category.go`: Métodos para criar e buscar categorias no SQLite.
    *   `course.go`: Métodos para criar e buscar cursos.
*   **`gqlgen.yml`**: Arquivo de configuração da biblioteca, definindo onde os arquivos gerados devem ser salvos e mapeamento de tipos entre GraphQL e Go.

---

## 🧠 Conceitos de GraphQL Aplicados

### 1. O que é GraphQL?

GraphQL é uma linguagem de consulta para APIs. Ao contrário do REST, onde você tem múltiplos endpoints (`/users`, `/posts`), no GraphQL você tem um **único endpoint** e o cliente pede exatamente os campos que deseja.

### 2. Schema Definition Language (SDL)

Arquivo: `graph/schema.graphqls`

Define os tipos de dados disponíveis.
```graphql
type Category {
  id: ID!
  name: String!
  courses: [Course!]!  # Relacionamento: Categoria tem vários cursos
}
```

### 3. Mutations (Escrita)

**Analogia REST**: Pense nas Mutations como os métodos **POST**, **PUT** ou **DELETE**.
Elas são usadas sempre que você precisa alterar o estado do servidor (criar, atualizar ou remover dados).

Exemplo de Mutation para criar categoria:
```graphql
mutation {
  createCategory(input: {name: "Go Lang", description: "Curso de Go"}) {
    id
    name
  }
}
```

### 4. Queries (Leitura)

**Analogia REST**: Pense nas Queries como o método **GET**.
Elas são utilizadas exclusivamente para **buscar** informações, sem causar efeitos colaterais.

Exemplo: Buscar todas as categorias E seus respectivos cursos:
```graphql
query {
  categories {
    name
    courses {
      name
      description
    }
  }
}
```

> **Dica**: Confira o arquivo `example.graphql` na raiz do projeto. Ele contém diversos exemplos de queries e mutations prontos para você copiar e colar no Playground para testar nossa API.

### 5. Resolvers

Arquivo: `graph/schema.resolvers.go`

São as funções que "resolvem" o pedido do usuário. Se você pede `categories`, o resolver `Categories` vai no banco e traz os dados.

Um detalhe importante implementado é o **Resolver de Campo (Field Resolver)**. Quando pedimos os `courses` de uma `Category`, o GraphQL não traz isso "de graça" do objeto categoria principal se não estiver carregado. Nós implementamos um método específico `Courses` atrelado à `Category` que busca os cursos daquela categoria específica `WHERE category_id = ?`.
