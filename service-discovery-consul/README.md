# Service Discovery com Consul

Este projeto demonstra a implementação de um cluster Consul para Service Discovery, cobrindo conceitos de agentes, servidores, clientes, registro de serviços, health checks e segurança.

## Conteúdo Abordado
- **Visão Geral**: Service Registry, Health Checks e Multicloud.
- **Arquitetura**: Diferenças entre Agent, Client e Server.
- **Cluster**: Criação de um cluster com 3 servidores e 2 clientes.
- **Serviços**: Registro de serviços (Nginx) via arquivo de configuração.
- **DNS Interface**: Consultas de serviço utilizando o DNS do Consul.
- **Segurança**: Criptografia de comunicação (Gossip Encryption).

## Estrutura do Projeto

O ambiente é orquestrado via Docker Compose e consiste em:
- **3 Servidores Consul**: `consulserver01`, `consulserver02`, `consulserver03` (Simulando um cluster de alta disponibilidade).
- **2 Clientes Consul**: `consulclient01`, `consulclient02` (Simulando máquinas onde serviços rodariam).

O `docker-compose.yaml` foi configurado para manter os containers rodando em modo interativo (`tail -f /dev/null`), permitindo a execução manual dos comandos para fins didáticos.

## Passo a Passo Realizado

### 1. Inicialização do Cluster
Para subir o ambiente:
```bash
docker-compose up -d
```

### 2. Configuração Manual dos Servidores
Acesso a cada servidor para iniciar o agente manualmente:
```bash
docker exec -it consulserver01 sh
mkdir /var/lib/consul
mkdir /etc/consul.d
consul agent -server -bootstrap-expect=3 -node=consulserver01 -bind=172.19.0.3 -data-dir=/var/lib/consul -config-dir=/etc/consul.d
```
*Repetido para `consulserver02` e `consulserver03` com seus respectivos IPs.*

### 3. Formação do Cluster (Join)
União manual dos nós ao cluster:
```bash
consul join 172.19.0.2
```

### 4. Configuração dos Clientes
Nos containers de cliente (`consulclient01`, `consulclient02`):
```bash
consul agent -bind=172.19.0.6 -data-dir=/var/lib/consul -config-dir=/etc/consul.d -retry-join=172.19.0.2
```

### 5. Registro de Serviços e Health Check
Criação de arquivos de definição de serviço (ex: `services.json`) dentro de `/etc/consul.d/`:

```json
{
    "services": [
        {
            "id": "nginx",
            "name": "nginx",
            "tags": ["web"],
            "port": 80,
            "check": {
                "id": "nginx",
                "name": "HTTP 80",
                "http": "http://localhost",
                "interval": "10s",
                "timeout": "1s"
            }
        }
    ]
}
```
Recarga das configurações:
```bash
consul reload
```

### 6. Consultas DNS (Service Discovery)
Instalação do `bind-tools` para usar o `dig` e testar a resolução de nomes:
```bash
apk -U add bind-tools
dig @localhost -p 8600 nginx.service.consul
```

### 7. Automação e Segurança
Configuração final utilizando arquivos `server.json` mapeados via volume para automação do bootstrap e segurança:
- **Auto Discovery**: Uso de `retry_join` para evitar o `consul join` manual.
- **Encrypt**: Geração de chave com `consul keygen` e configuração da propriedade `encrypt` em todos os nós.

Exemplo de execução automatizada:
```bash
consul agent -config-dir=/etc/consul.d
```

## Comandos Úteis
- **Verificar membros do cluster**: `consul members`
- **Listar serviços no catálogo**: `consul catalog services`
- **Recarregar configurações**: `consul reload`
- **Gerar chave de criptografia**: `consul keygen`
- **Consulta DNS**: `dig @localhost -p 8600 <nome-servico>.service.consul`

## Interface Web
A UI do Consul pode ser acessada através do servidor principal (exposto na porta 8500):
- http://localhost:8500
