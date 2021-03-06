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

    hack/dial

Performing Concurrent Scanning
    
    hack/scan-too-fast

Port Scanning Using a Worker Pool
> To avoid inconsistencies, you'll use a pool of goroutines to manage the concurrent work being preformed.

    hack/workerpool

## BUILDING A TCP PROXY
> YOU can achieve all TCP-based communications by using Go's built-in net package

- Using io.Reader and io.Writer

## Creating the Echo Server
> As is customary for most languages, you'll start by building an echo server to learn 
how to read and write data to and from a socket.

This makes sense logically, as TCP connections are bidirectional and can be used to send(write) 
or receive(read) data.

After creating an instance of Conn, you'll be able to send and receive data over a TCP 
socket.

## Proxying a TCP Client
> Now that you have a solid foundation, you can take what you've learned up to this point 
and create a simple port forwarder to proxy a connection through an intermediary service 
or host.

As mentioned earlier in this chapter, this is useful for trying to circumvent restrictive 
egress controls or to leverage a system to bypass network segmentation.

## Replicating Netcat for Command Execution
> Netcat is the TCP/IP Swiss Army knife --- essentially, a more flexible, scriptable version 
of Telnet. It contains a feature that allows stdin and stdout of any arbitrary program to be 
redirected over TCP, enabling an attacker to, for example, turn a single command execution 
vulnerability into operating system shell access.

# HTTP C LIENTS AND REMOTE INTERACTION WITH TOOLS
- It will first introduce you to the basics of building and customizing HTTP requests and 
receiving their responses

- Then you'll learn how to parse structured response data so the client can interrogate the information
the information to determine actionable or relevant data

- Finally, you'll learn how to apply these fundamentals by building HTTP clients that interact with a 
variety of security tools and resources

## HTTP FUNDAMENTALS WITH GO
- First, HTTP is a stateless protocol: the server doesn't inherently maintain state and status for 
each request. Instead, state is tracked through a variety of means, which may include session identifiers, 
cookies, HTTP headers, and more.

- Second, communications between clients and servers can occur either synchronously or asynchronously, but 
they operate on a request/response cycle.

- Finally, Go contains convenience functions so you can quickly and easily build and send HTTP requests 
to a server and subsequently retrieve and process the response.

## Calling HTTP APIs
> Let's begin the HTTP discussion by examining basic requests

Each function takes---as a parameter---the URL as a string value and uses it for the 
request's destination.

## Using Structured Response Parsing
> It uses the ioutil.ReadAll() function to read data from the response body, does some error 
checking, and prints the HTTP status code and response body to stdout.

## Building a http client that interacts with shodan
> Prior to performing any authorized adversarial activities against an organization, any good 
attacker begins with reconnaissance

Typically, this organization, any good attacker begins with reconnaissance.

if the error message disclose the enterprise username format, and if the organization uses single-factor 
authentication for its VPN, those error messages could increase the likelihook of an internal network 
compromise through password-guessing attacks.

Designing the Project Structure
> When building an API client, you should structure it so that the function calls and logic stand alone