# leilao-goexpert
Sistema Leilão

# Desafio GOLang Labs Leilão - FullCycle 

Aplicação em Go sendo: 
  - Servidor HTTP Rest Server

&nbsp;
- **Aplicação em Container com - Docker-compose e Dockerfile**
- **Banco de dados MongoDB**

## Funcionalidades

- **Consulta de Usuarios**
  - O servidor permite consultar de usuarios previamente cadastrados.

 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição.   
    curl --request GET \
    --url http://localhost:8080/user/8a19e094-3e10-42ee-94af-85ad31c0cb40 \
    --header 'User-Agent: insomnia/10.0.0'

  - Retorno esperado:
```
{
	"id": "8a19e094-3e10-42ee-94af-85ad31c0cb40",
	"name": "Usuario teste Leião"
}
 
``` 
## Como Utilizar localmente:

1. **Requisitos:** 
   - Certifique-se de ter o Go instalado em sua máquina.
   - Certifique-se de ter o Docker instalado em sua máquina.

&nbsp;
2. **Clonar o Repositório:**
&nbsp;

```bash
git clone https://github.com/tiago-g-sales/leilao-goexpert.git
```
&nbsp;
3. **Acesse a pasta do app:**
&nbsp;

```bash
cd leilao-goexpert
```
&nbsp;
4. **Rode o docker-compose para buildar e executar toda a stack de observabilidade:**
&nbsp;

```bash 
 docker-compose up
```

&nbsp;



## Como testar localmente:

### Portas
HTTP server on port :8080 <br />

### Efetur o cadastro do usuario 
    - Execuatar o mongo express via navegador (passo 5 desse readme) e criar a collection users e criar o New Document conforme o json abaixo:
    {
      "_id":"8a19e094-3e10-42ee-94af-85ad31c0cb40",
      "name":"Usuario teste Leião"
    }

### HTTP Consultar usuario previamente cadastrado
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição. 

    curl --request GET \
    --url http://localhost:8080/user/8a19e094-3e10-42ee-94af-85ad31c0cb40 \
    --header 'User-Agent: insomnia/10.0.0'


### HTTP Criar Leilão
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição. 

    curl --request POST \
    --url http://localhost:8080/auctions \
    --header 'Content-Type: application/json' \
    --header 'User-Agent: insomnia/10.0.0' \
    --data '{
            
        "productName":  "Product Computer",
        "category": "Informatic",
        "description": "Auction computer",
        "condition": 1
    }'


### HTTP Criar Lance
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição. 

    curl --request POST \
    --url http://localhost:8080/bid \
    --header 'Content-Type: application/json' \
    --header 'User-Agent: insomnia/10.0.0' \
    --data '{
        "userID": "8a19e094-3e10-42ee-94af-85ad31c0cb40",
        "auctionID": "aa46b62a-122a-4f8f-9a8c-03f0eeaaa40e",
        "amount": 50
    }'

### HTTP Consultar lista de Leiões com status ativo 
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição. 

    curl --request GET \
    --url 'http://localhost:8080/auctions?status=0' \
    --header 'User-Agent: insomnia/10.0.0'

### HTTP Consultar Leiões por id 
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição. 

    curl --request GET \
    --url http://localhost:8080/auctions/{auctionID} \
    --header 'User-Agent: insomnia/10.0.0'

### HTTP Consultar Lance por id 
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição. 

    curl --request GET \
    --url http://localhost:8080/bid/{bidID} \
    --header 'User-Agent: insomnia/10.0.0'

### HTTP Consultar Lance vencedor por ID leilão  
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição. 

    curl --request GET \
    --url http://localhost:8080/auctions/winner/{auctionID} \
    --header 'User-Agent: insomnia/10.0.0'


&nbsp;
5. **Acessar o MongoDB Express para consulta dos dados no banco de dados:**

  - http://localhost:8081/

