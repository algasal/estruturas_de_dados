package main

import (
    "fmt"
    "math/rand"
)

type No struct{
    info int
    proximo *No
}

type Pilha struct {
    topo *No
}

func CriarNo(i int) *No{
    return &No{
        info: i,
        proximo: nil,
    }
}

func (n No) Imprimir() string{
    return fmt.Sprintf("|%d| -> ", n.info)
}



func (p Pilha) EstaVazia() bool{
    return p.topo == nil
}

func (p *Pilha) Empilhar(info int) {
    novo := CriarNo(info)

    if !p.EstaVazia() {
        novo.proximo = p.topo
    }
    p.topo = novo
}

func (p *Pilha) Desempilhar() int {
    temp := p.topo.info
    p.topo = p.topo.proximo
    return temp
}

func (p Pilha) ConsultarTopo() int {
    return p.topo.info
}

func (p Pilha) Imprimir() string {
    s := ""
    if p.EstaVazia() {
        s += "Pilha vazia"
    } else {
        aux := p.topo
        for aux != nil {
            s += aux.Imprimir()
            aux = aux.proximo
        }
    }
    return s
}

func main() {
    pilha := Pilha{}

    pilha.Empilhar(rand.Intn(10))

    for !pilha.EstaVazia() {
        if rand.Intn(2) != 0 {
            pilha.Empilhar(rand.Intn(10))
        } else {
            pilha.Desempilhar()
        }
        fmt.Printf("Minha pilha: %s\n", pilha.Imprimir())
    }
}
