# Go TCP Implementation for Learning Purposes

- notice when you terminate either the client or the server, the whole connection terminates.
- cannot have multiple clients for one server (when in session/connection, unless...)

> [!IMPORTANT]
> [UPDATE 2025-01-04 00:55]:
>  - it can Accept multiple clients if and only if the Accept() is inside the infinite loop
>       - this means that it accepts the first message in the connection (like how http works)
>  - if the `Accept()` is outside the infinite loop, then it will only accept the first client that connects to it
>       - this means that 

## How it works?

## TCP Server

## TCP Client
