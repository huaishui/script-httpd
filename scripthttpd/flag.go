package scripthttpd

import (
	"os"

	"github.com/namsral/flag"
)

const EnvScript = "SCRIPT"

type Opts struct {
	Script []string
	Addr   string
}

func ParseConfig() Opts {
	opts := Opts{}

	flag.StringVar(&opts.Addr, "addr", ":8080", "the TCP network address to listen on, eg. ':80'")

	flag.Parse()

	opts.Script = flag.Args()
	if len(opts.Script) == 0 {
		envScript := os.Getenv(EnvScript)
		if envScript != "" {
			opts.Script = []string{os.Getenv(EnvScript)}
		} else {
			// No script was passed via args or env, print usage and exit.
			flag.Usage()
			os.Exit(2)
		}
	}

	return opts
}
