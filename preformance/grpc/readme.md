# Protocol Buffers role in gRPC

- Protocol Buffers is used to define the:
    - Message(data, Request and Response)
    - Service(Service name and RPC endpoints)
    
- Efficiency of Protocol Buffers over JSON
    - ｇRPC uses Protocol Buffers for communications
    - Let's measure the payload size vs JSON
    - Parsing JSON is actually CPU intensive (because the format is human readable)
    - Parsing Protocol Buffers(binary format) is less CPU intensive because it's closer to how a machine represents data
    - By using gRPC, the use of Protocol Buffers means faster and more efficient communication, friendly with mobile devices that have a slower CPU
    
# What's HTTP/2?

gRPC leverages HTTP/2 as a backbone for communications 

# How HTTP/2 works
> HTTP 2 was release in 2015. It has been battled tested for many years! (SPDY)

- Http2 supports multiplexing
    - The client & server can push messages in parallel over the same TCP connection
    - This greatly reduces latency

- Http2 supports server push
    - Servers can push streams(multiple messages) for one request from the client
    
- Http2 supports header compression
    - Headers (text based) can now be compressed
    - There have much less impact on the packet size
    - (remember the average http request may have over 20 headers, due to cookies, content cache, and application headers)

- Http2 is binary
    - While HTTP1 text makes it easy for debugging, it's not efficient over the network

- Http 2 is secure