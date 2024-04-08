package main

import (
    "fmt"
    "math/rand"
)

type Vetor struct {
    ultPos int
    dados []int
}

func main() {
    testeVetor()
}

func CriaVetor(capacidade int) *Vetor {
    return &Vetor{
        ultPos: -1,
        dados: make([]int, capacidade),
    }
}

func (v Vetor) GetUltPos() int {
    return v.ultPos
}

func (v Vetor) GetDados() []int {
    return v.dados
}

func (v Vetor) EstaCheio() bool {
    return v.ultPos == len(v.dados) -1
}

func (v Vetor) EstaVazio() bool {
    return v.ultPos == -1
}

func (v *Vetor) Redimensiona(novaCapacidade int) {
    temp := make([]int, novaCapacidade)

    for i := 0; i <= v.ultPos; i++ {
        temp[i] = v.dados[i]
    }
    v.dados = temp
}

func (v *Vetor) Adiciona(e int) {
    if v.EstaCheio() {
        v.Redimensiona(len(v.dados) * 2)
    }
    v.ultPos += 1
    v.dados[v.ultPos] = e;
}

func (v *Vetor) Remove() int{
    if v.EstaVazio() {
        return -1
    }

    aux := v.dados[v.ultPos]
    v.ultPos -= 1
    if (len(v.dados) >= 10 && v.ultPos <= len(v.dados) / 4) {
        v.Redimensiona(len(v.dados) / 2)
    }
    return aux
}

func (v *Vetor) PreencheVetor() {
    for _,_  = range v.dados {
        v.Adiciona(rand.Intn(len(v.dados) * 10) + 1)
    }
}

func (v *Vetor) EsvaziaVetor() {
    v.ultPos = -1
}

func (v *Vetor) BubbleSort() {
    for i := 1; i < len(v.dados); i++ {
        for j := 0; j < len(v.dados) - i; j++ {
            if v.dados[j] > v.dados[j+1] {
                temp := v.dados[j]
                v.dados[j] = v.dados[j+1]
                v.dados[j+1] = temp
            }
        }
    }
}

func (v *Vetor) Partition(p, r int) int {
    x := v.dados[r]
    i := p-1

    for j := p; j < r; j++ {
        if v.dados[j] <= x {
            i++
            temp := v.dados[i]
            v.dados[i] = v.dados[j]
            v.dados[j] = temp
        }
    }
    i++
    temp := v.dados[i]
    v.dados[i] = x
    v.dados[r] = temp
    return i
}

func (v *Vetor) QuickSort (p, r int) {
    if p < r {
        q := v.Partition(p, r)
        v.QuickSort(p, q-1)
        v.QuickSort(q+1, r)
    }
}

func (v Vetor) ImprimeVetor() string {
    var s string

    if (v.EstaVazio()) {
        s = "Vetor vazio"
    } else {
        s = fmt.Sprintf("%v", v.dados)
    }

    return s
}
