# Desafio: Sistema de Stress test

## Objetivo
Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas. O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

## Entrada de Parâmetros via CLI:

   - --url: URL do serviço a ser testado.
   - --requests: Número total de requests.
   - --concurrency: Número de chamadas simultâneas.

## Execução do Teste:

   - Realizar requests HTTP para a URL especificada.
   - Distribuir os requests de acordo com o nível de concorrência definido.
   - Garantir que o número total de requests seja cumprido.

## Execução do Teste:

   - Apresentar um relatório ao final dos testes contendo:
      - Tempo total gasto na execução
      - Quantidade total de requests realizados.
      - Quantidade de requests com status HTTP 200.
      - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Execução da aplicação:

   - Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:
      - docker run <sua imagem docker> —url=http://google.com —requests=1000 —concurrency=10

## Configuração do projeto

1. **Clone o Repositório:**

```bash
git clone https://github.com/vinicius-maker/stress-test.git
cd stress-test
```

2. **Configurar docker:**
    - no diretório raiz: stress-test

```bash
docker build -t stresstest .
```

3. **Execução do projeto via docker:**

```bash
docker run stresstest --url=<url> --requests=<total_requests> --concurrency=<total_concurrency>
```

## Execução/Testes do projeto

- Obs: para cada cenário abaixo, não foi adicionado no resultado esperado o Tempo Total da request por ser algo variável
- Obs2: existe um timeout configurado nas requisições HTTP de 60 segundos no worker

#### Cenário de Teste com URL válida (Status HTTP 200)
   
```bash
docker run stresstest --url=https://httpbin.org/status/200 --requests=10 --concurrency=5
```

- Resultado Esperado:
     - Requests totais: 10
     - Status HTTP 200: 10
     - Outros Status HTTP: 0

#### Cenário de Teste com URL inválida (Status HTTP 0)

```bash
docker run stresstest --url=http://invalid.url --requests=10 --concurrency=5
```

- Resultado Esperado:
     - Requests totais: 10
     - Status HTTP 200: 0
     - Outros Status HTTP 0: 10
- Obs: Status HTTP 0: 10 (indica falha na conexão).

#### Cenário de Teste com URL retornando Status HTTP 404 (Not Found)

```bash
docker run stresstest --url=https://httpbin.org/status/404 --requests=10 --concurrency=5
```

- Resultado Esperado:
     - Requests totais: 10
     - Status HTTP 200: 0
     - Outros Status HTTP 404: 10

#### Cenário de Teste com URL retornando Status HTTP 500 (Internal Server Error)

```bash
docker run stresstest --url=https://httpbin.org/status/500 --requests=10 --concurrency=5
```

- Resultado Esperado:
     - Requests totais: 10
     - Status HTTP 200: 0
     - Outros Status HTTP 500: 10

#### Cenário de Teste com URL retornando Status HTTP 502 (Bad Gateway)

```bash
docker run stresstest --url=https://httpbin.org/status/502 --requests=10 --concurrency=5
````

- Resultado Esperado:
     - Requests totais: 10
     - Status HTTP 200: 0
     - Outros Status HTTP 502: 10

#### Cenário de Teste com Alta Concurrency

```bash
docker run stresstest --url=https://httpbin.org/status/200 --requests=100 --concurrency=20
```

- Resultado Esperado:
     - Requests totais: 100
     - Status HTTP 200: 100
     - Outros Status HTTP: 0

#### Cenário de Teste com Concurrency Baixa

```bash
docker run stresstest --url=https://httpbin.org/status/200 --requests=10 --concurrency=1
```

- Resultado Esperado:
     - Requests totais: 10
     - Status HTTP 200: 10
     - Outros Status HTTP: 0

#### Cenário de Teste com Tempo de Resposta Lento

```bash
docker run stresstest --url=https://httpbin.org/delay/2 --requests=5 --concurrency=5
```

- Resultado Esperado:
     - Requests totais: 5
     - Status HTTP 200: 5
     - Outros Status HTTP: 0
