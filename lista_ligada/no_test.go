package main

import "testing"


func TestNo(t *testing.T) {
    no1 := CriarNo(10)
    no2 := CriarNo(20)
    no1.SetProximo(no2)

    t.Run("No 1", func(t *testing.T) {
        no1_str := no1.Imprimir()
        esperado := "|10| -> "
        retornoEsperado(t, no1_str, esperado)
    })
    
    t.Run("No 2", func(t *testing.T) {
        no2_str := no2.Imprimir()
        esperado := "|20| -> "
        retornoEsperado(t, no2_str, esperado)
    })

    t.Run("Proximo no1", func(t *testing.T) {
        proximo := no1.Imprimir() + no1.Proximo().Imprimir()
        esperado := "|10| -> |20| -> "
        retornoEsperado(t, proximo, esperado)
    })
}

func retornoEsperado(t testing.TB, no1_str, esperado string) {
    t.Helper()
    if no1_str != esperado {
        t.Errorf("Esperava '%q', veio '%q'", esperado, no1_str)
    }
}
