# Desafio Rate Limiter - Pós GoExpert

<br>

## Sobre o projeto

Este é o repositório destinado ao desafio técnico Rate Limiter do curso Pós GoExpert da faculdade FullCycle. Projeto é um rate limiter em Go, que pode ser configurado para limitar o número máximo de requisições por segundo com base em um endereço IP específico ou em um token de acesso.

<br>

## Funcionalidades

- Restringe o número de requisições recebidas de um único endereço IP.

- Restringe o número de requisições recebidas baseadas em um token de acesso único.

<br>

## Como executar o projeto

### Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:

- [Git](https://git-scm.com)

- [VSCode](https://code.visualstudio.com/)

- [Docker](https://www.docker.com/)

- [Apache bench](https://httpd.apache.org/docs/2.4/programs/ab.html)

- Terminal de comandos (cmd, powershell, bash ...)

<br>

#### Acessando o repositório

  

```bash

# Clone este repositório

$  git  clone  https://github.com/pedrogutierresbr/lab-weather-api-pos-goexpert.git

```

  

<br>

  
  

#### Executando a aplicação em ambiente dev

  

```bash

# Executar a aplicação

$  docker-compose  up  -d

# Parar e remover recursos da aplicação

$  docker-compose  down

```

<br>

#### Variáveis de ambiente

  

```bash

# RATE_LIMIT_IP=1 --> limite por IP

# RATE_LIMIT_IP=2 --> limite por Token

# BLOCK_TIME --> tempo de bloqueio após limite ser atingido

```

<br>

  

#### Ações disponíveis

```bash

# Serviços estarão disponíveis na seguinte porta: 8080

# Você pode encontrar as ações no arquivo api / api.http


# Teste de limite por IP

$ for i in {1..5}; do curl -i http://localhost:8080/; done


# Teste de limite por Token

$ for i in {1..5}; do curl -i -H "API_KEY: teste_123" http://localhost:8080/; done


# Testes com Apache Bench (necessário instalção do Apache bench) :

$ ab -n 10 -c 1 http://localhost:8080/

```

<br>

## Licença

  

Este projeto esta sobe a licença [MIT](./LICENSE). 

Feito por Pedro Gutierres [Entre em contato!](https://www.linkedin.com/in/pedrogabrielgutierres/)