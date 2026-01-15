# Full Cycle 3.0 - Comunicação entre Sistemas

Este repositório contém os projetos e exercícios desenvolvidos durante o módulo de **Comunicação entre Sistemas** do curso **Full Cycle 3.0**.

Neste módulo, exploramos as principais formas de comunicação entre serviços e arquiteturas distribuídas, focando em performance, escalabilidade e desacoplamento.

## 📚 Tópicos Abordados

Vamos explorar na prática as seguintes tecnologias e conceitos:

## 📚 Tópicos Abordados

Vamos explorar na prática as seguintes tecnologias, ferramentas e conceitos:

### 1. REST (Representational State Transfer)
- Padrões de comunicação HTTP.
- Maturidade de Richardson.
- **Ferramentas**:
  - Laminas API Tools
  - Dev Container (Ambiente de desenvolvimento padronizado)

### 2. GraphQL
- Implementação de APIs flexíveis com um único endpoint.
- Schemas, Queries, Mutations e Resolvers.
- [📂 Ver Módulo GraphQL](./graphql-module)

### 3. gRPC (Google Remote Procedure Call)
- **Conceitos Fundamentais**: HTTP/2, Protocol Buffers, REST vs gRPC.
- **Setup**: Instalação do compilador `protoc` e plugins.
- **Desenvolvimento**:
  - Criação de arquivos `.proto`.
  - Geração de código automática.
  - Implementação de Serviços (Server e Client).
- **Tipos de Comunicação**:
  - Unary (Requisição simples).
  - Server-side Streaming.
  - Client-side Streaming.
  - Bidirectional Streaming.
- **Ferramentas**:
  - **Evans**: Um cliente universal para gRPC (como o Postman/Insomnia, mas para gRPC).

### 4. Service Discovery & Consul
- **Conceitos**:
  - O problema dos IPs dinâmicos em microsserviços.
  - Service Registry.
  - Health Checks.
- **Arquitetura Consul**:
  - Agent, Client e Server.
  - Formação de Cluster.
- **Prática**:
  - Iniciando agentes e subindo o cluster.
  - Registro de serviços e sincronização.
  - Implementação de Health Checks.
  - `Retry Join` para formação do cluster.
  - Criptografia (Gossip encryption).
  - Web UI e dicas de produção.


## 🛠️ Tecnologias Utilizadas

- **Go (Golang)**: Linguagem principal para os exemplos de backend.
- **Docker & Docker Compose**: Para orquestração dos serviços.
- **SQLite**: Banco de dados leve para os exemplos.

---
> ⭐️ **Dica**: Navegue pelas pastas do repositório para encontrar o código fonte de cada tecnologia específica.
