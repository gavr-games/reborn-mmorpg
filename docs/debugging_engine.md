# Debugging engine

You can debug the Go packages using [Delve](https://github.com/go-delve/delve). It provides a lot of useful commands and gives the opportunity to debug more efficiently than just logging some information from the code.

## Usage

1. Start the project in debug mode

- `make start-debug`

2. Ensure that there is the message from the engine container in the logs: `Type 'help' for list of commands`. It means that debug mode has been activated successfully and you can start debugging.

3. Attach to engine container

- `make attach-engine`

4. Use **dlv** commands to debug the engine. You can find the list of **dlv CLI** commands in the [Delve](https://github.com/go-delve/delve/tree/master/Documentation/cli) documentation.

## Most useful **dlv CLI** commands

- `sources` - prints list of source files
- `break <absolute path to file>:<line number>` - sets a breakpoint on the specified line
- `continue` - run until breakpoint or program termination
- `next` - step over to next source line
- `print <expression>` - evaluates an expression

## How the debug mode works?

1. To start the project in debug mode the separate compose file [docker-compose.debug.yml](/docker-compose.debug.yml) is used. It overrides the commands used to start the project.
2. Compose file for debug uses [separate **.sh** scripts](/engine/bin/debug/) for running Go packages in debug mode.
3. **.sh** scripts for debug use the `dlv debug <path-to-package>` command to build and run necessary packages in debug mode.

> From [issue](https://github.com/go-delve/delve/issues/2844#issuecomment-1002343103): *For **dlv** to work correctly, you must build your program with debugging symbols and with optimizations and inlining turned off. `go run` doesn't include debug info. `go build` does, but you still need `-gcflags='all=-N -l'`. If you use `dlv debug`, **dlv** will take care of all this for you.*
