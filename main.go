package main

import (
    "log"
    "net/smtp"
)

func main() {
    // Set up authentication information.
    auth := smtp.PlainAuth(
        "",
        "meuemail221@bol.com.br",
        "minha_senha",
        "smtps.bol.com.br",
    )
    // Connect to the server, authenticate, set the sender and recipient,
    // and send the email all in one step.
    err := smtp.SendMail(
        "smtps.bol.com.br:587",
        auth,
        "meuemail221@bol.com.br",
        []string{"ma@gmail.com"},
        []byte("Nois memo"),
    )
    if err != nil {
        log.Fatal(err)
    }
}
