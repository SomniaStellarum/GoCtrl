package goctrl

import (
    "time"
)

const ms     = 1000000

func NewTimeStep(period int64) (t *TimeStep) {
    t = new(TimeStep)
    t.p = period
    t.out = make(chan float64)
    return t
}

type TimeStep struct {
    out chan float64
    p int64
    tick *time.Ticker
}

func (t *TimeStep) Run() {
    t.tick = time.NewTicker(t.p * ms)
    prev := <- t.tick.C
    for tck := range t.tick.C {
        t.out <- float64(tck - prev)
        prev = tck
    }
}

func (t *TimeStep) Sourcer(s Sink) {
    t.out = s.Sinker()
}

type TimeSink interface {
    TimeSinker() chan float64
}

type TmSnk struct {
    tIn chan float64
}

func (t *TmSnk) TimeSinker() (in chan float64) {
    t.tIn = make(chan float64)
    return t.tIn
}

