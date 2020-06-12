# Learning about Performance in Go
> In this section, you will learn why performance in computer science is important. You will also learn why performance in the Go language. Moving on,  you will learn about data structures and algorithms, concurrency, STL algorithm 
 equivalents, and the matrix and vector computations in Go.

# Introduction to Performance in Go
>This book is written with intermediate to advanced Go developers in mind. 

In this chapter, we will cover the following topics:

- Understanding performance in computer science
- A brief history of Go
- The ideology behind Go performance

These topics are provided to guide you in beginning to understand the direction you need to take to write highly performant code in the Go language.

    https://github.com/bobstrecansky/HighPerformanceWithGo/

# Understanding performance in computer science
> Performance in computer science is a measure of work that can be accomplished by a computer system. Performant code is vital to many different groups of developers. 

Whether you're part of a large-scale software company that needs to quickly deliver masses of data to customers, an embedded computing device programmer who has limited computing resources available, or a hobbyist looking to squeeze more requests out of the Raspberry Pi that you are using for your pet project, performance should be at the forefront of your development mindset. 

**Performance matters, especially when your scale continues to grow.**

It is important to remember that we are sometimes limited by physical bounds. 

CPU, memory, disk I/O, and network connectivity all have performance ceilings based on the hardware that you either purchase or rent from a cloud provider. 

There are other systems that may run concurrently alongside our Go programs that can also consume resources, such as OS packages, logging utilities, monitoring tools, and other binaries—it is prudent to remember that our programs are very frequently not the only tenants on the physical machines they run on.

Optimized code generally helps in many ways, including the following:

- Decreased response time: The total amount of time it takes to respond to a request.
- Decreased latency: The time delay between a cause and effect within a system.
- Increased throughput: The rate at which data can be processed.
- Higher scalability: More work can be processed within a contained system.

There are many ways to service more requests within a computer system. Adding more individual computers (often referred to as horizontal scaling) or upgrading to more powerful computers (often referred to as vertical scaling) are common practices used to handle demand within a computer system. One of the fastest ways to service more requests without needing additional hardware is to increase code performance. Performance engineering acts as a way to help with both horizontal and vertical scaling. The more performant your code is, the more requests you can handle on a single machine. This pattern can potentially result in fewer or less expensive physical hosts to run your workload. This is a large value proposition for many businesses and hobbyists alike, as it helps to drive down the cost of operation and improves the end user experience.