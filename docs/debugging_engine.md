# Debugging engine

You can debug the Go packages using [Delve](https://github.com/go-delve/delve). It provides dozens of useful commands and gives the opportunity to debug more efficiently than just logging some information from the code.

## Build and run package in debug mode

Stop the project (if it is running).

To debug a particular package, you have to change the way it will be built and started first. To run the package in debug mode, you need to use `dlv debug <path-to-package>` command.

Let's consider running the [server](../engine/cmd/server/) package in debug mode as an example.

In the [start.sh](../engine/bin/start.sh) file:

1. Uncomment `dlv debug ./cmd/server/` command
2. Comment `go run cmd/server/main.go` command

> From [issue](https://github.com/go-delve/delve/issues/2844#issuecomment-1002343103): *For **dlv** to work correctly, you must build your program with debugging symbols and with optimizations and inlining turned off. `go run` doesn't include debug info. `go build` does, but you still need `-gcflags='all=-N -l'`. If you use `dlv debug`, **dlv** will take care of all this for you.*

Start the project with `docker compose up`. Ensure that there is the message from the engine container in the logs: `Type 'help' for list of commands`. It means that debug mode has been activated successfully and you can start debugging.

## Start debug session

To start a debug session, you need to attach to the container where the package you want to debug is running:

- `docker attach <container-name>`

**dlv** debug session will be accessed for the package, and you'll be able to run **dlv** commands in the **dlv CLI**.

- If you want to learn more about **dlv** commands, run `help` command or read the [docs](https://github.com/go-delve/delve)
- If you want to see all source files available to debug, run `sources`

> Note, if you want to references source file in the command (for example, for setting breakpoint), you have to use absolute path of it in the container.
