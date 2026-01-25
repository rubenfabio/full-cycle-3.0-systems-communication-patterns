# Módulo gRPC - Full Cycle

Este módulo foca no desenvolvimento de aplicações de alta performance utilizando **gRPC (Google Remote Procedure Call)** e **Protocol Buffers** com a linguagem Go.

## 📚 Tópicos e Conceitos

Neste módulo, aplicamos os seguintes conceitos fundamentais:

*   **gRPC vs REST**: Diferenças entre comunicação baseada em recursos (JSON/HTTP1.1) e chamadas de procedimento remoto (Proto/HTTP2).
*   **Protocol Buffers (Protobuf)**: Linguagem de definição de interface (IDL) para serialização de dados estruturados. Mais leve e rápido que JSON.
*   **HTTP/2**: Uso de multiplexação e headers compactados para performance superior.
*   **Tipos de Comunicação**:
    *   **Unary**: Requisição e resposta simples (1:1).
    *   **Server-Side Streaming**: O cliente envia 1 requisição e recebe múltiplos dados via stream.
    *   **Client-Side Streaming**: O cliente envia múltiplos dados via stream e aguarda 1 resposta.
    *   **Bidirectional Streaming**: Cliente e Servidor enviam dados continuamente via stream de forma independente.

## 🛠️ Ferramentas Utilizadas

*   **Go (Golang)**
*   **Protoc**: Compilador de Protocol Buffers.
*   **Protoc Plugins**: `protoc-gen-go` e `protoc-gen-go-grpc`.
*   **Evans**: CLiente gRPC universal (substituto ao Postman/Insomnia para gRPC).
*   **SQLite**: Banco de dados para persistência.

## 🚀 Como Rodar o Projeto

### Pré-requisitos
1.  Ter o **Go** instalado.
2.  Ter o compilador `protoc` instalado.
3.  Instalar os plugins Go:
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```
4.  (Opcional) Ter o **Evans** instalado para testar.

### Passos

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/rubenfabio/full-cycle-3.0-systems-communication-patterns.git
    cd full-cycle-3.0-systems-communication-patterns/grpc-module
    ```

2.  **Instale as dependências:**
    ```bash
    go mod tidy
    ```

3.  **Geração de Código (Opcional):**
    Caso altere o arquivo `.proto`, gere novamente o código Go:
    ```bash
    protoc --go_out=. --go-grpc_out=. proto/course_category.proto
    ```

4.  **Execute o Servidor:**
    ```bash
    go run cmd/grpcServer/main.go
    ```
    O servidor iniciará na porta `50051`.

## 🧪 Testando com Evans

O Evans é um cliente interativo para gRPC. Com o servidor rodando:

1.  Inicie o Evans:
    ```bash
    evans -r repl
    ```
2.  Dentro do REPL do Evans:
    ```bash
    # Selecione o package e service
    package pb
    service CategoryService

    # Chamada Unary (Criar Categoria)
    call CreateCategory

    # Listar Categorias
    call ListCategories

    # Stream Bidirecional
    call CreateCategoryStreamBiDirectional
    ```

## 📂 Estrutura do Projeto

*   **`proto/`**: Contém o arquivo `course_category.proto` com a definição do serviço e mensagens.
*   **`internal/pb/`**: Código Go gerado automaticamente pelo `protoc`. **Não edite estes arquivos.**
*   **`internal/service/`**: Implementação real das regras de negócio e métodos gRPC (`CreateCategory`, `CreateCategoryStream`, etc.).
*   **`internal/database/`**: Camada de acesso ao banco de dados SQLite.
*   **`cmd/grpcServer/`**: Ponto de entrada (main) que sobe o servidor gRPC na porta 50051.

---
> ⭐️ **Dica**: Este módulo demonstra o poder do gRPC para comunicação backend-to-backend eficiente.
