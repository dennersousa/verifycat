<img src="https://github.com/gatinhodev/verifycat/assets/135276762/8f4b2368-3733-4864-8c28-c7719edf0ece" alt="verifycat-api" width="100" height="100">

# VerifyCat  

**VerifyCat** é uma API de validação versátil projetada para lidar com vários tipos de validações, incluindo CPF, CNPJ, URL, E-mails e números de cartão de crédito (até o momento).

## Sumário

- [Visão Geral](#visão-geral)
- [Arquitetura](#arquitetura)
- [Endpoints da API](#endpoints-da-api)
- [Uso](#uso)
- [Contribuições](#contribuições)
- [Licença](#licença)

## Visão Geral

VerifyCat é construído usando o framework [Gin](https://github.com/gin-gonic/gin) para lidar com solicitações HTTP e fornecer um servidor web rápido e eficiente. O projeto é estruturado para acomodar diferentes tipos de validação, cada um implementado em um arquivo separado dentro do pacote `validate`.

## Arquitetura

O ponto de entrada principal da aplicação é o arquivo `verifycat_api.go`. Ele configura o roteador Gin e define o endpoint para validação em `/validate`. A lógica de validação é delegada para manipuladores específicos no pacote `validate`.

O pacote `validate` contém arquivos individuais (`cpf.go`, `cnpj.go`, `url.go`, `email.go` e `creditcard.go`) para cada tipo de validação. Esses arquivos contêm a lógica de validação e manipulação de solicitações específicas para seu tipo de validação.

## Endpoints da API

### `POST /validate`

Este endpoint suporta validação para vários tipos de dados. A carga útil deve estar no formato JSON, contendo o `type` (tipo de validação) e `value` (dados a serem validados).

- **Exemplo de Carga Útil da Solicitação:**
  ```json
  {
    "type": "cpf",
    "value": "123.456.789-09"
  }
  ```

- **Exemplo de Resposta:**
  ```json
  {
    "isValid": true,
    "message": "CPF"
  }
  ```

- **Exemplo de Solicitação CURL**
```bash
curl -X POST http://localhost:8080/validate -H "Content-Type: application/json" -d '{"type": "cpf", "value": "123.456.789-09"}'
```

- **Exemplo de Resposta CURL:**
  ```json
  {
    "isValid": true,
    "message": "CPF"
  }
  ```

### Tipos de Validação Suportados

- `cpf`: Número de identificação brasileiro (CPF)
- `cnpj`: Número de entidade legal brasileira (CNPJ)
- `url`: URL
- `email`: Endereço de e-mail
- `creditcard`: Número de cartão de crédito

## Arquitetura RESTful

A API VerifyCat segue os princípios RESTful, incluindo:

- **Recursos Identificáveis:** Cada tipo de validação (CPF, CNPJ, URL, etc.) é tratado como um recurso identificável. As operações são realizadas nesses recursos por meio de URLs específicos.

- **Operações HTTP Padrão:** Operações CRUD (Criar, Ler, Atualizar, Excluir) são mapeadas para operações HTTP padrão. Neste código, a principal operação é a validação, realizada por meio de uma solicitação POST para o recurso `/validate`.

- **Estado Sem Sessão:** Cada solicitação do cliente ao servidor contém todas as informações necessárias para entender e processar a solicitação. Não há dependência de estados intermediários entre solicitações.

- **Representação de Recursos:** Os recursos são representados em JSON no corpo das respostas HTTP. A resposta é uma representação do estado atual do recurso (por exemplo, se um CPF é válido ou não).

- **HATEOAS (Hypermedia As The Engine Of Application State):** Embora o código fornecido não inclua explicitamente links para outros recursos no estilo HATEOAS, o conceito está incorporado no princípio geral de que o cliente interage com a API por meio de representações de recursos e estados fornecidos em respostas.

## Uso

1. **Clone o Repositório:**
   ```bash
   git clone https://github.com/your-username/verifycat.git
   cd verifycat
   ```

2. **Execute a Aplicação:**
   ```bash
   go run verifycat_api.go
   ```

3. **Faça Solicitações à API:**
   - Use seu cliente de API preferido (por exemplo, cURL, Postman) para enviar solicitações POST para `http://localhost:8080/validate` com a carga útil apropriada.

## Contribuições

Sinta-se à vontade para contribuir abrindo problemas, fornecendo feedback ou enviando pull requests.

## Licença

Este projeto está licenciado sob a Licença AGPL V3 - consulte o arquivo [LICENSE](LICENSE) para obter detalhes.
