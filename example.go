package main

import "goctrl"

func main() {
    // Create UserInput, TimeStep, RateModel and NetChanOut
    u := NewUserInput()
    p := NewP_Controller(0.1)
    t := NewTimeStep(50)
    r := NewRateModel(0.01)
    m := NewMult()
    n := NewNetChanOut()
    
    // Connect Sink {n} to the Source {u}
    Connect(u,p)
    Connect(p,r)
    Connect(t,r.T)
    Connect(r,m)
    Connect(m,p.FdBck)
    Connect(m,n)
    
    // 'go' run sink to start listening to input channel
    go n.Run()
    go m.Run()
    go r.Run()
    go t.Run()
    go p.Run()
    
    // run source, if quit it will stop process
    u.Run()
}

