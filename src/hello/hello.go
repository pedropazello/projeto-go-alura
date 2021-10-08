package main

import (
  "fmt"
  "os"
  "net/http"
  "time"
  "bufio"
  "io"
  "strings"
)

const monitoramentos = 3
const delay = 5

func main() {
  exibeIntroducao()

  for {
    exibeMenu()
    comando := leComando()

    switch comando {
      case 1:
        iniciarMonitoramento()
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
  fmt.Println("")

  return comando
}

func iniciarMonitoramento() {
  fmt.Println("Monitorando...")

  sites := leSitesArquivo()

  for i := 0; i < monitoramentos; i++ {
    for i, site := range sites {
      fmt.Println("Testando site", i, ":", site)
      testaSite(site)
    }

    time.Sleep(delay * time.Second)
    fmt.Println("")
  }
}

func testaSite(site string) {
  resp, err := http.Get(site)

  if err != nil {
    fmt.Println("ocorreu um erro:", err)
  }

  if resp.StatusCode == 200 {
    fmt.Println("Site:", site, "foi carregado com sucesso")
  } else {
    fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
  }
}

func leSitesArquivo() []string {
  var sites []string
  arquivo, err := os.Open("sites.txt")

  if err != nil {
    fmt.Println("Ocorreu um erro:", err)
  }

  leitor := bufio.NewReader(arquivo)

  for {
    linha, err := leitor.ReadString('\n')

    if err == io.EOF {
      break
    }

    if err != nil {
      fmt.Println("Ocorreu um erro:", err)
    }

    linha = strings.TrimSpace(linha)
    sites = append(sites, linha)
  }

  arquivo.Close()
  return sites
}
