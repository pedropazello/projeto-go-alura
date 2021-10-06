package main

import "fmt"
import "os"

func main() {
  exibeIntroducao()
  exibeMenu()
  comando := leComando()

  switch comando {
    case 1:
      fmt.Println("Monitorando...")
    case 2:
      fmt.Println("Exibindo Logs...")
    case 0:
      fmt.Println("Saindo do programa")
      os.Exit(0)
    default:
      fmt.Println("Não conheço esse comando")
      os.Exit(-1)
  }
}


func exibeIntroducao() {
  nome := "Pedro"
  versao := 1.1

  fmt.Println("Olá, sr.", nome)
  fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
  fmt.Println("1- Iniciar Monitoramente")
  fmt.Println("2- Exibir Logs")
  fmt.Println("0- Sair do Programa")
}

func leComando() int {
  var comando int
  fmt.Scan(&comando)

  fmt.Println("O comando escolhido foi", comando)

  return comando
}