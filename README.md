# Copilot Demo

To play demo:

- Checkout *start* tag
- Go at the end of *sync.go* file
- Type following comment: `// SyncedString is synchronized string`
- Type TAB until Copilot has written `SyncedString` structure, `Get` and `Set` functions
- Type following comment: `// Append string value`

Code for SyncedString was done in a minute. Let's add tests:

- Create file *sync_test.go*
- Write code `package main`
- Type TAB until Copilot has written `TestSyncedBool` and `TestSyncedString`

You can see expected result checking out tag *end* and discuss.

To generate code with ChatGPT:

- Launch ChatGPT
- Enter following prompt: `Write a TCP server in Go that prints requests on command line`
- You can call server with: `echo "test" | nc 127.0.0.1 8080`
- To get UDP version: `Using UDP now`
- You can call server with: `echo "test" | nc -4u -w1 127.0.0.1 8080`

Discussion: TCP server is broken, why?
