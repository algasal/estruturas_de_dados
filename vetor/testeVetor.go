package main

import "fmt"

func testeVetor() {
    var tamanho int
    fmt.Print("Digite o tamanho do vetor: ")
    fmt.Scanln(&tamanho)

    v := CriaVetor(tamanho)
    v.PreencheVetor()
    
    fmt.Printf("1o vetor gerado: %q\n", v.ImprimeVetor())
    v.BubbleSort()
    fmt.Printf("Vetor ordenado pelo bubble: %q\n", v.ImprimeVetor())
    v.EsvaziaVetor()

    v.PreencheVetor()
    fmt.Printf("\n2o vetor gerado: %q\n", v.ImprimeVetor())
    v.QuickSort(0, tamanho-1)
    fmt.Printf("Vetor ordenado pelo quick: %q\n", v.ImprimeVetor())
    
}

