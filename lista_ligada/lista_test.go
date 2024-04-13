package main

import (
    "testing"
    "fmt"
)

func TestLista(t *testing.T) {
    var lista Lista

    t.Run("Lista Vazia", func(t *testing.T) {
        conteudo := lista.Imprimir()
        esperado := "Lista: está vazia"
        retornoEsperado(t, conteudo, esperado)
    })

    t.Run("Insere 5 no início", func(t *testing.T) {
        for i := 0; i < 5; i++ {
            lista.InserirInicio(i)
        }

        l := lista.Imprimir()
        esperado := "Lista: |4| -> |3| -> |2| -> |1| -> |0| -> \\"
        retornoEsperado(t, l, esperado)
    })

    t.Run("Insere 5 no final", func(t *testing.T) {
        for i := 5; i < 10; i++ {
            lista.InserirFim(i)
        }

        l := lista.Imprimir()
        esperado := "Lista: |4| -> |3| -> |2| -> |1| -> |0| -> |5| -> |6| -> |7| -> |8| -> |9| -> \\"
        retornoEsperado(t, l, esperado)
    })

    t.Run("Remove do início", func(t *testing.T) {
        removido := fmt.Sprintf("%d foi removido do início", lista.RemoverInicio())
        esperado := "4 foi removido do início"
        retornoEsperado(t, removido, esperado)
    })

    t.Run("Remove do fim", func(t *testing.T) {
        removido := fmt.Sprintf("%d foi removido do fim", lista.RemoverFim())
        esperado := "9 foi removido do fim"
        retornoEsperado(t, removido, esperado)
    })
}

