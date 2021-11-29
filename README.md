# ecoador-TCP

<h2>Como executar o servidor:</h2> <br />

1 - Digite o seguinte comando.
```
go run servidor.go 
```
2 - Ele irá pedir a porta para a comunicação.
(obs: escolha a porta que estiver disponivel em sua máquina)
```
9090
```
3 - Servidor está pronto para receber requisições.



<h2>Como executar o cliente</h2> <br/>

1 - Digite o seguinte comando. 

```
go run cliente.go
```

2 - Ele ira pedir o Ip
- Digite o IP:
```
127.0.0.1
```
- Digite a porta
(obs:a porta necessita ser a mesma inserida para executar o servidor)
```
9090
```
3 - Digite o texto que deseja ser enviado como requisição
```
teste
```
<h2>O que cliente irá retornar </h2> <br/>

```
Body: teste
Address: 127.0.0.1:9090
```

<h2>O que servidor irá retornar </h2> <br/>

```
Body:teste
Address:127.0.0.1:34622
```
