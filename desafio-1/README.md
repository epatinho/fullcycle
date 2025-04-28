Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.
 
Você precisará nos entregar dois sistemas em Go:
- client.go
- server.go
 
Os requisitos para cumprir este desafio são:
 
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
 
O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
 
Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.
 
Ao finalizar, envie o link do repositório para correção.



- **server.go**: Um servidor HTTP que consome a cotação do dólar da API pública e salva no SQLite.
- **client.go**: Um cliente HTTP que consulta o servidor e grava a cotação atual em um arquivo de texto.

---

## 📂 Estrutura de Pastas

```text
/cotacao-project
├── client/
│   ├── client.go
│   ├── go.mod
│   ├── Dockerfile
│   └── cotacao.txt         # (gerado em runtime)
├── server/
│   ├── server.go
│   ├── go.mod
│   ├── Dockerfile
│   └── cotacoes.db         # (gerado em runtime)
├── docker-compose.yml
└── README.md
```



