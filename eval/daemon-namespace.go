package eval

import (
	"errors"
	"strconv"

	"github.com/elves/elvish/daemon/api"
	"github.com/elves/elvish/daemon/client"
)

var ErrDaemonOffline = errors.New("daemon is offline")

func makeDaemonNamespace(daemon *client.Client) Namespace {
	// Obtain process ID
	daemonPid := func() Value {
		req := &api.PidRequest{}
		res := &api.PidResponse{}
		err := daemon.Call("Pid", req, res)
		maybeThrow(err)
		return String(strconv.Itoa(res.Pid))
	}

	return Namespace{
		"pid": MakeRoVariableFromCallback(daemonPid),

		FnPrefix + "spawn": NewRoVariable(&BuiltinFn{"daemon:spawn", daemonSpawn}),
	}
}

func daemonSpawn(ec *EvalCtx, args []Value, opts map[string]Value) {
	TakeNoArg(args)
	TakeNoOpt(opts)
	ec.ToSpawn.Spawn()
}