package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// lendo a porta
	var port string
	fmt.Println("Digite a porta do servidor 'Ecoardor Tabajara'")
	fmt.Scan(&port)
	port = fmt.Sprintf(":%v",port)
	//checando se foi passada uma porta
	if port == ":" || port == "" {
		log.Fatalln("O parâmetro número da porta é obrigatório.")
	}

	listener, err := net.Listen("tcp4", port)
	if err != nil {
		log.Fatalf("Erro iniciando servidor na porta %s:%v", port, err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Problema identificado ao se conectar na porta:%s:%v", port, err)
		}
		go NewConnection(conn)
	}


	fmt.Println(port)
}

func NewConnection(conn net.Conn) {

	for{
		mensagem, err := bufio.NewReader(conn).ReadString('\n')

		if err == io.EOF {
			fmt.Printf("Conexão %s encerrada\n", conn.RemoteAddr().String())
			break
		}

		if err != nil {
			log.Fatalf("Problema com a conexão %v:%v", conn, err)
		}
		requisiçao := []byte(fmt.Sprintf(mensagem))
		conn.Write(requisiçao)
		fmt.Printf(" -Body:%v -Address:%s\n", string(requisiçao),conn.RemoteAddr().String() )
	}
	conn.Close()
}