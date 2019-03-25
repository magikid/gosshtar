# gosshtar
A tar-pit for SSH written in golang it can trap bots attempting to connect over
SSH until they give up.  By default, it will respond to connection attempts with
a random number every 10 seconds.  You can also pass it a file which will be
read and each line of the file will be sent in 10 second intervals.

This takes advantage of [RFC4253](https://tools.ietf.org/html/rfc4253#section-4.2) (which covers SSH).  After establishing a TCP
connection, the server may send arbitrary data before sending the SSH version so
we just send as much data as we can at a slow rate.

```
  gosshtar [args]

  -debug
        Print lots of log messages
  -file string
        A file of text to send to the rouge connection
  -help
        Print the help and exit
  -host string
        Restrict the server to a specific IP (default "0.0.0.0")
  -port int
        The port to run on (default 2222)
```
