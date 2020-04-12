# TCP, SCANNERS, AND PROXIES
> Let's begin our practical application of Go with the Transmission Control Protocol(TCP), 
the predominant standard for connection-oriented, reliable communications and the foundation 
of modern networking.

TCP is everywhere, and it has well-documented libraries, code samples, and generally easy-to-understand packet flows.

You must understand TCP to fully evaluate, analyze, query, and manipulate network traffic.

As an attacker, you should understand how TCP works and be able to develop usable TCP constructs so that you 
can identify open/closed ports, recognize potentially errant results such as false-positives --- for example, 
synflood protections --- and bypass egress restrictions through port forwarding.

build a concurrent, properly throttled port scanner.

create a TCP proxy that can be used for port forwarding

re-create Netcat's "gaping security hole" feature.

Entire textbooks have been written to discuss every nuance of TCP, including packet structure and flow, 
reliability, communication reassembly, and more.

## understanding the tcp handshake
> For those who need a refresher, let's review the basic.

If the port is open, a three-way handshake takes place. 

First, the client sends a syn packet, which signals the beginning of a communication. 
The server then responds with a syn-ack, or acknowledgment of the server's response. 
The transfer of data can then occur.

If the port is closed, the server responds with a rst packet instead of a syn-ack.

If the traffic is being filtered by a firewall, the client will typically receive no 
response from the server.

These responses are important to understand when writing network-based tools.

Correlating the output of your tools to these low-level packet flows will help you validate 
that you've properly established a network connection and troubleshoot potential problems.

## bypassing firewalls with port forwarding
> People can configure firewalls to prevent a client from connecting to certain servers and ports, 
while allowing access to others.

In some cases, you can circumvent these restrictions by using an intermediary system to proxy the 
connection around or through a firewall, a technique known as port forwarding.

Many enterprise networks restrict internal assets from establishing HTTP connections to malicious sites.

A client connects, through a firewall, to the destination host stacktitan.com. This host is configured to 
forward connections to the host evil.com.

You can use port forwarding to exploit several restrictive network configurations. 

## WRITING A TCP SCANNER
> One effective way to conceptualize the interaction of TCP ports is by implementing a port scanner.
By writing one, you'll observe the steps that occur in a TCP handshake, along with the effects of encountered 
state changes, which allow you to determine whether a TCP port is available or filtered state.

Once you've written a basic scanner, you'll write one that's faster.

