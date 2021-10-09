package main

import (
  "fmt"
  "os"
  "net/http"
  "time"
  "bufio"
  "io"
  "strings"
  "strconv"
  "io/ioutil"
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
        imprimeLogs()
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
    registraLogs(site, true)
  } else {
    fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    registraLogs(site, false)
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

func registraLogs(site string, carregado bool) {
  arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

  if err != nil {
    fmt.Println(err)
  }

  arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(carregado) + "\n")
  arquivo.Close()
}

func imprimeLogs() {
  arquivo, err := ioutil.ReadFile("log.txt")

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(arquivo))
}
