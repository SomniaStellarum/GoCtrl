package goctrl

import (
    "fmt"
    nt "netchan"
    "time"
)

type NetChanOut struct {
    *Snk
    out chan string
}

func NewNetChanOut() (n *NetChanOut) {
    n = new(NetChanOut)
    n.Snk = new(Snk)
    return n
}

func (n *NetChanOut) Run() {
    ticker := time.NewTicker(second)
    var val float64
    n.out = make(chan string)
    imp, _ := nt.NewImporter("tcp","localhost:9090")
    imp.Import("Output", n.out, nt.Send)
    for {
        select {
        case val = <-n.in:
        case <-ticker.C:
            n.out <- fmt.Sprintf("%f\n",val)
        }
    }
    //create string netchan and print values over that chan
}
