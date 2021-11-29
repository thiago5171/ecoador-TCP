package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	var ip, porta string
	fmt.Println( "------------------Ecoardor Tabajara------------------")
	fmt.Println("Digite o IP:")
	fmt.Scan(&ip)
	fmt.Println("Digite a porta:")
	fmt.Scan(&porta)
	endereço := fmt.Sprint(ip , ":",porta)

	//checando se foi passada uma porta
	if endereço == ":" || endereço == "" {
		log.Fatalf("Problema ao se conectar ao endereço informado")
	}
	conn, err := net.Dial("tcp", endereço)
	if err != nil {
		log.Fatalf("Problema identificado ao se conectar no endereço :%s:%v", endereço, err)
	}

	for {
		fmt.Printf("Digite uma mensagem para o Tabajara:")
		var texto string
		fmt.Scanln(&texto)

		if texto == "calado" {
			break
		}

		if _, err := conn.Write([]byte(fmt.Sprintf("%s\n",texto))); err != nil {
			log.Fatalf("Erro ao enviar mensagem %v", err)
		}

		bufferReceive := make([]byte, 1024)
		_, err := conn.Read(bufferReceive)
		if  err != nil {
			log.Fatalf("Problema ao ler retorno em %s:%v", endereço, err)
		}
		//LIMPA TELA
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		//
		fmt.Println(" -Body:", string(bufferReceive), "-Address:", endereço)
	}


	conn.Close()

}