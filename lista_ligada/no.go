package main

import (
    "fmt"
)

type No struct {
    info int
    proximo *No
}

func CriarNo (i int) *No {
    return &No{
        info: i,
        proximo: nil,
    }
}

func (n No) Info() int {
    return n.info
}

func (n No) Proximo() *No {
    return n.proximo
}

func (n *No) SetInfo(info_ int) {
    n.info = info_
}

func (n *No) SetProximo(proximo_ *No) {
    n.proximo = proximo_
}

func (n No) Imprimir() string {
    return "|" + fmt.Sprintf("%d", n.info) + "| -> "
}

