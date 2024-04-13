package main

//import "fmt"

type Lista struct {
    primeiro *No
}

func (l Lista) Primeiro() *No {
    return l.primeiro
}

func (l *Lista) SetPrimeiro(primeiro_ *No) {
    l.primeiro = primeiro_
}

func (l Lista) EstaVazia() bool{
    return l.primeiro == nil
}

func (l *Lista) InserirInicio(i int) {
    novo := CriarNo(i)
    if ! l.EstaVazia() {
        novo.SetProximo(l.primeiro)
    }
    l.primeiro = novo
}

func (l *Lista) InserirFim(i int) {
    novo := CriarNo(i)
    if l.EstaVazia() {
        l.primeiro = novo
    } else {
        aux := l.primeiro
        for aux.Proximo() != nil {
            aux = aux.Proximo()
        }
        aux.SetProximo(novo)
    }
}

func (l *Lista) RemoverInicio() int {
    info := l.primeiro.Info()
    l.primeiro = l.primeiro.Proximo()
    return info
}

func (l *Lista) RemoverFim() int {
    var temp int

    if l.primeiro.Proximo() == nil {
        temp = l.primeiro.Info()
        l.primeiro = nil
    } else {
        aux := l.primeiro
        for aux.Proximo().Proximo() != nil {
            aux = aux.Proximo()
        }
        temp = aux.Proximo().Info()
        aux.SetProximo(nil)
    }
    return temp
}

func (l Lista) Imprimir() string{
    s := "Lista: "

    if l.EstaVazia() {
        s += "está vazia"
    } else {
        aux := l.primeiro
        for aux != nil {
            s += aux.Imprimir()
            aux = aux.Proximo()
        }
        s += "\\"
    }
    return s
}

func main() {
/*    no1 := CriarNo(10)
    no2 := CriarNo(20)
    no3 := CriarNo(30)

    var l Lista
    fmt.Println(l.Imprimir())
    l.SetPrimeiro(no1)
    l.InserirInicio(no3.info)
    l.InserirFim(no2.info)

    fmt.Println(l.Imprimir())

    l.RemoverInicio()
    fmt.Println(l.Imprimir())

    l.RemoverFim()
    fmt.Println(l.Imprimir())*/

}       
