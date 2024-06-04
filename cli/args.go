package cli

import "github.com/alexflint/go-arg"

type Args struct {
	Endpoints []string `arg:"positional,required"`
	Username  string   `arg:"-u,--username"`
	Password  string   `arg:"-p,--password"`
}

func GetArgs() *Args {
	args := &Args{}
	arg.MustParse(args)
	return args
}

func (*Args) Version() string {
	return "etcd-tui 0.0.1"
}
