package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.

    // Define um prefixo para todas as mensagens de log. Assim, toda mensagem registrada pelo logger começará com "greetings: ".
    log.SetPrefix("greetings: ")
    // Remove todas as informações extras (como data, hora, nome do arquivo, linha) das mensagens de log.
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(messages)
}