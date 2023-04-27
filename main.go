package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/DisgoOrg/disgohook"
	"github.com/DisgoOrg/disgohook/api"
	"gopkg.in/gomail.v2"
)

type Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Body     string `json:"body"`
}

func main() {
	intro()
	fmt.Println("[" + horaAgora() + "] -- [READER]: Lendo arquivo JSON")

	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	var dados []Data
	err2 := json.Unmarshal(data, &dados)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}

	fmt.Println("[" + horaAgora() + "] -- [READER]: Dados coletados com sucesso")

	exibeOpcoes()

	opcao := leEntradaUser()

	switch opcao {

	case 1:

		fmt.Println("[" + horaAgora() + "] -- [BOT]: Opção 1 escolhida")

		var emailMaze string
		var assuntoEmail string
		fmt.Println("[" + horaAgora() + "] -- [SEND]: Digite o email de destinatário:")
		fmt.Scan(&emailMaze)
		fmt.Println("[" + horaAgora() + "] -- [SEND]: Digite o assunto do email:")
		fmt.Scan(&assuntoEmail)
		for _, x := range dados {
			emailSender := x.Email
			passSender := x.Password
			bodySender := x.Body
			enviaEmail(emailSender, passSender, bodySender, emailMaze, assuntoEmail)

		}

		enviaWebhook()

	case 0:
		fmt.Println("[" + horaAgora() + "] -- [EXIT]: Fechando o bot")
		os.Exit(0)

	default:
		fmt.Println("[" + horaAgora() + "] -- [ERROR]: Comando não reconhecido")
		fmt.Println("[" + horaAgora() + "] -- [EXIT]: Fechando o bot")
		os.Exit(0)
	}
}

func intro() {

	versao := 3.1
	fmt.Println("")
	fmt.Println("["+horaAgora()+"] -- [BOT]: Bat.mail versão", versao, "iniciado")

}

func exibeOpcoes() {
	fmt.Println("[1] - Enviar email pra Maze")
	fmt.Println("[0] - Fechar bot")
}

func leEntradaUser() int {
	var opcaoEscolhida int
	fmt.Scan(&opcaoEscolhida)

	return opcaoEscolhida
}

func horaAgora() string {
	horaNow := time.Now()
	hora := horaNow.Format("15:04:05")

	return hora
}

func enviaEmail(email string, pass string, body string, emailMaze string, assuntoEmail string) {

	mail := gomail.NewMessage()

	mail.SetHeader("From", email)
	mail.SetHeader("To", emailMaze)

	mail.SetHeader("Subject", assuntoEmail)
	mail.SetBody("text/plain", body)

	a := gomail.NewDialer("smtp.gmail.com", 587, email, pass)

	horaEmailEnviado := horaAgora()
	if err := a.DialAndSend(mail); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("[" + horaEmailEnviado + "] -- [SUCCESS]: Email enviado com sucesso")
	}
}

func enviaWebhook() {

	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, "955553597481943040/ba4GVa_2djJWPo_e5RlygDkcYFSBGPKjyUBGM7TK54Mfm3DWAtFhH8T9uNIdFMZDjpfY")

	if err != nil {
		fmt.Println(err)
	}

	webhook.SendEmbeds(api.NewEmbedBuilder().
		SetTitle("Emails enviados com sucesso").
		SetDescription("Explode maze ativado").
		SetFooterIcon("https://www.cgcreativeshop.com/wp-content/uploads/2018/10/baticon15102018.jpg").
		SetFooterText("Bat.mail").
		Build(),
	)
	fmt.Println("[" + horaAgora() + "] -- [DISCORD]: Webhook enviado com sucesso")

}
