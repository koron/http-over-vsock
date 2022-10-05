# Demo: HTTP over VSOCK (Hyper-V sockets)

## Getting Started

### Getting Started with AWS Nitro Enclaves

1. Clone this repository.

    ```console
    $ git clone https://github.com/koron/http-over-vsock.git
    ```

2. Start a server in Enclave

    ```console
    $ cd http-over-vsock
    $ make encalve-run
    ```

3. Build a client and make requests.

    Build a client

    ```console
    $ cd http-over-vsock/client
    $ go build
    ```

    Run it to make a HTTP request over vsock.

    ```console
    $ ./clieht http://16:1234/
    Hello VSOCK (/)
    ```

    The response will be changed when you change path of request URL.

    ```console
    $ ./clieht http://16:1234/foo/bar
    Hello VSOCK (/foo/bar)
    ```

4. (OPTIONAL) Show server logs

    Open another terminal and run this:

    ```console
    $ cd http-over-vsock
    $ make enclave-console
    ```

    To terminate logs, interrupt with Ctrl-C or so.

See Makefile for details.

### Getting Started with Windows and WSL2

* Host:  Windows 10
* Guest: WSL2 Ubuntu-22.04

1. Run HTTP server on WSL2 Ubuntu (guest) which listen VSOCK

    ```console
    $ cd server
    $ go build
    $ sudo ./server
    ```

2. Detemine VMID (GUID)

    ```console
    > hcsdiag list 
    C34EC814-C4A9-411C-BF5D-559529ECA7AB
        VM,                         Running, C34EC814-C4A9-411C-BF5D-559529ECA7AB, WSL
    ```

    on system administrator console.

3. GET from host Windows via Hyper-V socket

    ```console
    > cd client
    > go build
    > .\client.exe http://C34EC814-C4A9-411C-BF5D-559529ECA7AB:1234
    ```

## References

* <https://man7.org/linux/man-pages/man7/vsock.7.html>
* <https://pkg.go.dev/github.com/Microsoft/go-winio?GOOS=windows>
* <https://learn.microsoft.com/en-us/virtualization/hyper-v-on-windows/user-guide/make-integration-service>
* <https://github.com/nbdd0121/wsld/blob/master/docs/impl.md>
