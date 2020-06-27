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

# Data Structures and Algorithms

Data structures and algorithms are the basic units of building software, notably complex, performance software. Understanding them helps us think about how to impactfully organize and manipulate data in order to write effective, performant software. This chapter will include explanations of different data structures and algorithms, as well as how their Big O notation is impacted.

- Understanding benchmarking
> Metrics and measurement are at the root of optimization. The adage You can't improve what you can't measure rings true with performance. To be able to make intelligent decisions about performance optimizations, we must continuously measure the performance of the functions we are trying to optimize.

# Benchmark execution
> Benchmarks in Go use the axiom of starting with the word Benchmark(with a capital B) in the function call to denote that they are a benchmark and that they should use the benchmark functionality.

- benchtime t
> Run enough iterations of the test to take the defined t duration. Increasing this value will run more iterations of b.N.

- count n
> Run each test n times.

- benchmem
> Turn on memory profiling for your test.

- cpu x,y,z
> Specify a list of GOMAXPROCS values for which the benchmarks should be executed.

# Data structure operations and time complexity
> The following diagram contains some of the common data structure operations and their time complexities.

- O(1) - constant time
> Algorithms written in constant time have an upper bound that does not depend on the input size of algorithm.

- O(log n) - logarithmic time
> Logarithmic growth is often represented as a partial sum of the harmonic series

- O(n) - linear time
> Algorithms written in linear time scale linearly with the size of their dataset.

- O(n log n) - quasilinear time
> Algorithms written in quasilinear time are often used to order values within an array in Go

# Insertion sort
> Insertion sort is a sorting algorithm that constructs an array one item at a time until it results in a sorted array.

    for insertionSort(data interface, a, b int) {
        for i:=a+1;i<b;i++ {
            for j := i; j > a && data.Less(j, j-1); j-- {
                data.Swap(j,j-1)
            }
        }
    }

# Heap sort
> Go has a built-in heapSort in the standard library, as shown in the following code snipet helps us understand that heapSort is an O(n log n) sorting algorithm.


    func heapSort(data Interface, a, b int) {
        first := a
        lo := 0
        hi := b-a
        for i:=(hi-1)/2; i>=0; i-- {
            siftDown(data, i, hi, first)
        }
        for i:=hi-1; i>=0;i-- {
            data.Swap(first, first+i)
            siftDown(data, lo, i, first)
        }
    }
    
# Merge sort
> A stable sort ensures that two objects that share the same key in an input array appear in the resulting array in the same order. Stability is important if we want to make sure that a key-value order pair is organized within an array. An implementation of a stable sort can be found in the Go standard library


# Quick sort