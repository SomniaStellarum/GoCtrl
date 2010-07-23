package goctrl

import "fmt"

func apndChan(chSlice[]chan float64, ch chan float64) []chan float64 {
    l := len(chSlice)
    c := cap(chSlice)
    if (l == c) {
        fmt.Printf("Slice full\n")
        return chSlice
    }
    if (l != 0) {
        chSlice = chSlice[0:l+1]
    } else {
        chSlice = chSlice[0:1]
    }
    chSlice[l] = ch;
    return chSlice
}

func chanComm(chSlice[]chan float64, comm float64) {
    for _, v := range chSlice {
        v <- comm
    }
}

func NewMult() (m *Mult) {
    m = new(Mult)
    m.Snk = new(Snk)
    m.out = make([]chan float64, 0, 10)
    return m
}

type Mult struct {
    *Snk
    out []chan float64
}

func (m *Mult) Run() {
    for {
        chanComm(m.out, <-m.in)
        //output to all channels in slice
    }
}

func (m *Mult) Sourcer(s Sink) { //out chan float64) {
    m.out = apndChan(m.out, s.Sinker())
}

