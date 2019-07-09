package main

import "fmt"

const en_hello_prefix = "Hello "
const spanish = "Spanish"
const sp_hello_prefix = "Hola "
const french = "French"
const fr_hello_prefix = "Bonjour "

func Hello(name, language string) string {
    if name == "" { name = "World" }
    return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {
    switch language {
        case french:
            prefix = fr_hello_prefix
        case spanish:
            prefix = sp_hello_prefix
        default:
            prefix = en_hello_prefix
    }
    return
}
func main() {
    fmt.Println(Hello("Chris", ""))
}
