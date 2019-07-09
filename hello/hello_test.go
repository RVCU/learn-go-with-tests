package main

import "testing"

func TestHello(t *testing.T) {
    assertCorrectMessage := func(t *testing.T, got, want string) {
            t.Helper() //this tells the testing framework this func is a helper
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }
    t.Run("In Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola Elodie"
        assertCorrectMessage(t, got, want)
    })
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "")
        want := "Hello Chris"
        assertCorrectMessage(t, got, want)
    })
    t.Run("saying hello world if no nname is provided", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello World"
        assertCorrectMessage(t, got, want)
    })
}
