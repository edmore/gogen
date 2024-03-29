package service

var ch = make(chan string)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Mul(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

type Ping struct{}

func (t *Ping) Pong(arg string, reply *string) error {
	*reply = arg + " pong\n"
	return nil
}
