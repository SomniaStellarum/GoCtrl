package goctrl

import "fmt"

const second = 1000000000

type Runner interface {
    Run()
}

type Sink interface {
    Sinker() chan float64
    // Return the input channel
}

type Source interface {
    Sourcer(s Sink)
    // Connect the sink to the source
}

type Snk struct {
    in chan float64
}

func (s *Snk) Sinker() (in chan float64) {
    s.in = make(chan float64)
    in = s.in
    return in
}

type Src struct {
    out chan float64
}

func (s *Src) Sourcer(snk Sink) {
    s.out = snk.Sinker()
}

func Connect(src Source, snk Sink) {
    src.Sourcer(snk)
}

func NewP_Controller(kP float64) (p *P_Controller) {
    p = new(P_Controller)
    p.Snk = new(Snk)
    p.Src = new(Src)
    p.FdBck = new(Snk)
    p.kP = kP
    return p
}

type P_Controller struct {
    *Snk
    *Src
    FdBck *Snk
    kP float64
}

func (c *P_Controller) Run() {
    sp := <-c.in
    in := sp
    out := (sp - in) * c.kP
    c.out <- out
    for {
        select {
        case sp = <-c.in:
            out = (in - sp) * c.kP
        case c.out <- out:
        case in = <-c.FdBck.in:
            out = (in - sp) * c.kP
        }
    }
}

func NewRateModel(gn float64) (r *RateModel) {
    r = new(RateModel)
    r.Snk = new(Snk)
    r.Src = new(Src)
    r.T = new(Snk)
    r.gn = gn
    return r
}

type RateModel struct {
    *Snk
    *Src
    T *Snk
    gn float64
}

func (r *RateModel) Run() {
    in := <-r.in
    out := in
    for {
        select {
        case in = <-r.in:
        default:
        }
        dt := <-r.T.in
        out = r.gn * dt * in + out
        r.out <- out
    }
}

type UserInput struct {
    *Src
}

func NewUserInput() (u *UserInput) {
    u = new(UserInput)
    u.Src = new(Src)
    return u
}

func (u *UserInput) Run() {
    inStr := new(string)
    in := new(float64)
    for {
        fmt.Printf("Cmd: ")
        fmt.Scanln(inStr)
        switch *inStr {
            case "Quit":
                fmt.Printf("Goodbye!!\n")
                return
            default:
                fmt.Sscanf(*inStr, "%5f\n", in)
                fmt.Printf("%5f\n", *in)
                u.out <- *in
        }
    }
}

