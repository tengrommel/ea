# section 1: an introduction to system programming and go

this section is an introduction to unix and system programming. it will help you understand what components unix operating systems have to offer, from the kernal api to the filesystem, and you will become familiar with system programming's basic concepts.

# memory management
> the operating system handles the primary and secondary memory usage of the applications. it keeps track of how much of the memory is used, by which process, and what parts are free. it also handles allocation of new memory from the processes and memory de-allocation when the processes are complete.

# techniques of management

- single allocation: all the memory, besides the part reserved for the os, is available for the application. this means that there can only be one application in execution at a time, like in microsoft disk operating system(ms-dos).

- partitioned allocation: this divides the memory into different blocks called partitions. using one of these blocks process makes it possible to execute more than compacted in order to obtain more contiguous memory space for the next processes.

- paged memory: the memory is divided into parts called frames, which have a fixed size. a process'memory is divided into parts of the same size called pages. there is a mapping between pages and frames that makes the process see its own virtual memory as contiguous. this process is also known as pagination.

# virtual memory
> unix uses the paged memory management technique, abstracting its memory for each application into contiguous virtual memory. it also uses a technique called swapping, which extends the virtual memory to the secondary memory(hard drive or solid state drives(ssd)) using a swap file.

when memory is scarce, the operating system puts pages from processes that are sleeping in the swap partition in order to make space for active processes that are requesting more memory, executing an operation called swap-out. when a page that is in the swap file is needed by a process in execution it gets loaded back into the main memory for executing it. this is called swap-in.

the main issue of swapping is the performance drop when interacting with secondary memory, but it is very useful for extending multitasking capablilities and for dealing with applications that are bigger than the physical memory, by loading just the pieces that are actually needed at a given time. creating memory-efficient applications is a way of increasing performance by avoiding or reducing swapping.

the top command shows details about available memory, swap, and memory consumption for each process:

- res is the physical primary memory used by the process.
- virt is the total memory used by the process, including the swapped memory, so it's equal to or bigger than res.
- shr is the part of virt that is actually shareable, such as loaded libraries.

# Understanding files and filesystems
> A filesystem is a method used to structure data in a disk, and a file is the abstraction used for indicating a piece of self-contained information. If the filesystem is hierarchical, it means that files are organized in a tree of directories, which are special files used for arranging stored files.

Operating systems and filesystems
> Over the last 50 years, a large number of filesystems have been invented and used, and each one has its own characteristics regarding space management, filenames and directories, metadata, and access restriction. Each modern operating system mainly uses a single type of filesystem.

# Linux
> Linux's filesystem (FS) of choice is the extended filesystem (EXT) family, but other ones are also supported, including XFS, Journaled File System(JFS), and B-tree File System(Btrfs). It is also compatible with the older File Allocation Table (FAT) family (FAT16 and FAT32) and New Technology File System(NTFS). The filesystem most commonly used remains the latest version of EXT(EXT4), which was released in 2006 and expanded its predecessor's capacities, including support for bigger disks.

# macOS
> macOS uses the Apple File System(APFS), which supports Unix permission and has journaling. It is also metadate-rich and case-preserving, while being a case-insensitive filesystem. It offers support for other filesystems, including HFS+ and FAT32, supporting NTFS for read-only operations. To write to such a filesystem, we can use an experimental feature or third-party applications.

# Windows
>The main filesystem used by Windows is NTFS. As well as being case-insensitive, the signature feature that distinguishes Windows FS from others is the use of a letter followed by a colon to represent a partition in paths, combined with the use of backslash as a folder separator, instead of a forward slash. Drive letters, and the use of C for the primary partition, comes from MS-DOS, where A and B were reserved drive letters used for floppy disk drives.

 Windows also natively supports other filesystems, such as FAT, which is a filesystem family that was very popular between the late seventies and the late nineties, and Extended File Allocation Table (exFAT), which is a format developed by Microsoft on top of FAT for removable devices.
 
 # File and hard and soft links
 > Most files are regular files, containing a certain amount of data. For instance a text file contains a sequence of human-readable characters represented by a certain encoding, while a bitmap contains some metadata about the size and the bit used for each pixel, followed by the content of each pixel.

Files are arranged inside directories that make it possible to have different namespaces to reuse filenames. These are referred to with a name, their human-readable identifier, and organized in a tree structure. The path is a unique identifier that represents a directory, and it is made by the names of all the parents of directory joined by a sequence (/ in Unix, \ in Windows), descending from the root to the desired leaf. For instance of a directory named a is located under another named b, which is under one called c, it will have a path that starts from the root and concatenates all the directories, up to the file: /c/b/a.

When more than one file points to the same content, we have a hard link, but this is not allowed in all filesystems (for example, NTFS and FAT). A soft link is a file that points to another soft link or to a hard link. Hard links can be removed or deleted without breaking the original link, but this is not true for soft links. A symbolic link is a regular file with its own data that is the path of another file. It can also link other filesystems or files and directories that do not exist (that will be a broken link).
 
 In Unix, some resources that are not actually files are represented as files, and communication with these resources is achieved by writing to or reading from their corresponding files. For instance, the /dev/sda file represents an entire disk, while /dev/stdout, dev/stdin, and /dev/stderr are standard output, input, and error. The main advantage of Everything is a file is that the same tools that can be used for files can also interact with other devices (network and pipes) or entities (processes).
 
 Unix filesystem
 > The principles contained in the section are specific to the filesystems used by Linux, such as EXT4.

# Root and inodes
> In Linux and macOS, each file and directory is represented by an inode, which is a special data structure that stores all the information about the file except its name and its actual data.

From the latest Linux kernel source, we can see how the first inodes are reserved. This is shown as follows:
    
    #define EXT4_BAD_INO 1 /* Bad blocks inode */
    #define EXT4_ROOT_INO 2 /* Root inode */
    #define EXT4_USR_QUOTA_INO 3 /* User quota inode */
    #define EXT4_GRP_QUOTA_INO 4 /* Group quota inode */
    #define EXT4_BOOT_LOADER_INO 5 /* Boot loader inode */
    #define EXT4_UNDEL_DIR_INO 6 /* Undelete directory inode */
    #define EXT4_RESIZE_INO 7 /* Reserved group descriptors inode */
    #define EXT4_JOURNAL_INO 8 /* Journal inode */

# Directory structure
> In Unix filesystems, there is a series of other directories under the root, each one used for a specific purpose, making it possible to maintain a certain interoperability between different operating systems and enabling compiled software to run on different OSes, making the binaries portable.

This is a comprehensive list of the directories with their scope:

- /bin: Executable files for all uses
- /boot: Files for booting the system
- /dev: Device drivers
- /etc: Configuration files for applications and system
- /home: Home directory for users
- /kernel: Kernel files
- /lib: Shared library files and other kernel-related files
- /mnt: Temporary filesystems, from floppy disks and CDs to flash drives
- /proc: File with process numbers for active processes
- /sbin: Executable files for administrators
- /tmp: Temporary files that should be safe to delete
- /usr: Administrative commands, shared files, library files, and others
- /var: Variable-length files(logs and print files)

# Navigation and interaction
> While using a shell, one of the directories will be working directory, when paths are relative (for example, file.sh or dir/subdir/file.txt). The working directory is used as a prefix to obtain an absolute one. This is usually shown in the prompt of the command line, but it can be printed with the pwd command(print working directory).

The cd(change directory) command can be used to change the current working directory. To create a new directory, there's the mkdir(make directory) command.

To show the list of files for a directory, there's the ls command, which accepts a series of options, including more information(-l), showing hidden files and directories(-a), and sorting by time (-t) and size (-s).

There is a series of other commands that can be used to interact with files: the touch command creates a new empty file with the given name, and to edit its content you can use a series of editors, including vi and nano, while cat, more and less are some of the commands that make it possible to read them.

# Mounting and unmounting
> The operating system splits the hard drive into logical units called partitions, and each one can be a different file system. When the operating system starts, it makes some partitions available using the mount command for each line of the /etc/fstab file, which looks more or less like this:

    # device    # mount-point  # fstype  # options      # dumpfreq # passno
    /dev/sda1   /               ext4        defaults        0           1

This configuration mounts /dev/sda1 to /disk using an ext4 filesystem and default options, no backing up(0), and root integrity check(1). The mount command can be used at any time to expose partitions in the filesystem. Its counterpart, umount, is needed to remove these partitions from the main filesystem. The empty directory used for the operation is called mount point, and it represents the root under which the filesystem is connected.

# Processes
> When an application is launched, it becomes a process: a special instance provided by the operating system that includes all the resources that are used by the running application. This program must be in Executable and Linkable Format(ELF), in order to allow the operating system to interpret its instructions.

Process properties
> Each process is a five-digit identifier process ID (PID), and it represents the process for all its life cycle. This means that cannot be two processes with the same PID at the same time. Their uniques makes it possible to access a specific process by knowing its PID. Once a process is terminated, its PID can be reused for another process, if necessary.

Similar to PID, there are other properties that characterize a process. These are as follows:

- PPID: The parent process ID of the process that started this process
- Nice number: Degree of friendliness of this process toward other processes
- Terminal or TTY: Terminal to which the process is connected
- RUID/EUID: The real/effective user ID, which belongs to the owner of the process
- RGID/EGID: The real/effective group owner, the group owner of a process

To see a list of the active processes, there's the ps (process status) command, which shows the current list of running processes for the active user:

    > ps -f
    UID PID PPID C STIME TTY TIME CMD
    user 8 4 0 Nov03 pts/0 00:00:00 bash -l -i
    user 43 8 0 08:53 pts/0 00:00:00 ps -f

# Process life cycle
> The creation of a new process can happen in two different ways:

- Using a fork: This duplicates the calling process. The child (new process) is an exact copy (memory) of the parent (calling process), except for the following:

    - PIDs are different
    - The PPID of the child equals the PID of the parent
    - The child does not inherit the following from the parent:
        - Memory locks
        - Semaphore adjustments 
        - Outstanding asynchronous I/O operations
        - Asynchronous I/O contexts
- Using an exec: This replaces the current process image with a new one, loading the program into the current process space and running it from its entry point

# Users, groups, and permissions
> Users and groups, together with permissions, are the main entities that are used in Unix operating systems to control access to resources

# Read, write, and execute
> Users and groups are used as the first two layers of protection for accessing a file. The user that owns a file has a set of permissions that differs from the  file group. Whoever is not the owner and does not belong to the group has  different permissions. These three sets of permissions are known as owner, group, and other.

# Process communications
> The operating system is responsible for communication between processes and has different mechanisms to exchange information. These processes are unidirectional, such as exit codes, signals, and pipes, or bidirectional, such as sockets.

Exit codes 
> Applications communicate their result to the operating system by returning a value called exit status. This is an integer value called exit status. This is an integer value passed to the parent process when the process ends. A list of common exit codes can be found in the /usr/include/sysexits.h file, as shown here:
    
    #define EX_OK 0 /* successful termination */
    #define EX__BASE 64 /* base value for error messages */
    #define EX_USAGE 64 /* command line usage error */
    #define EX_DATAERR 65 /* data format error */
    #define EX_NOINPUT 66 /* cannot open input */
    #define EX_NOUSER 67 /* addressee unknown */
    #define EX_NOHOST 68 /* host name unknown */
    #define EX_UNAVAILABLE 69 /* service unavailable */
    #define EX_SOFTWARE 70 /* internal software error */
    #define EX_OSERR 71 /* system error (e.g., can't fork) */
    #define EX_OSFILE 72 /* critical OS file missing */
    #define EX_CANTCREAT 73 /* can't create (user) output file */
    #define EX_IOERR 74 /* input/output error */
    #define EX_TEMPFAIL 75 /* temp failure; user is invited to retry */
    #define EX_PROTOCOL 76 /* remote error in protocol */
    #define EX_NOPERM 77 /* permission denied */
    #define EX_CONFIG 78 /* configuration error */
    #define EX__MAX 78 /* maximum listed value */

# signals
> Exit codes connect processes and their parents, but signals make it possible to interface any process with another, including itself. They are also asynchronous nd unidirectional, but they represent communication from the outside of a process.

The most common signal is SIGINT, which tells the application to terminate, and can be sent to a foreground process in a shell with the Ctrl + C key combnation.

    SIGHUP 1 Controlling terminal is closed
    SIGINT 2 Interrupt signal (Ctrl + C)
    SIGQUIT 3 Quit signal (Ctrl + D)
    SIGFPE 8 Illegal mathematical operation is attempted
    SIGKILL 9 Quits the application immediately
    SIGALRM 14 Alarm clock signal
    
The kill command allows you to send a signal to any application, and a comprehensive list of available signals can be shown with the -l flag:

# Pipes
> Pipes are the last unidirectional communcation method between processes. As the name suggests, pipes connect two ends - a process input with another process output - making it possible to process on the same host to communicate in order to exchange data.

# Socket

# Section 2: Advanced File I/O Operations

# Working with the Filesystem

## Getting and setting the working directory

*We can use the func Getwd() (dir string, err error) function of the os package to find out which path represents the current working directory*

*Changing the working diectory is done with another function of the same package, that is, func Chdir(dir string) error*

## Path manipulation
> The filepath package contains less than 20 functions, which is a small number compared to the packages of the standard library, and it's used to manipulate paths.

## Reading from files
> Getting the contents of a file can be done with an auxiliary function in the io/ioutil package, as well as with the ReadFile function, which opens, reads, and closes the file at once. *This uses a small buffer(512 bytes) and loads the whole content in memory*

This is not a good idea if the file size is very large, unknown, or if the content of the file can be processed one part at a time.

Reading a huge file from disk at once means copying all the file's content into the primary memory, which is a limited resource. This can cause memory shortages, as well as runtime errors. Reading chunks of a file at a time can help read the content of big files without causing huge memory usage. This is because the same part of the memory will be reused when reading the next chunk.

# Reader interface
> For all operations that read from a disk, there's an interface that is paramount: 

    type Reader interface {
        Read(p []byte) (n int, err error)
    }

Its job is really simple - fill the given slice of bytes with the content that's been read and return the number of bytes that's been read and an error, if one occurs.

*A reader makes it possible to process data in chunks (the size is determined by the slice), and if the same slice is reused for the operations that follow, the resulting program is consistently more memory efficient because it is using the same limited part of the memory that allocates the slice*

The file structure
> The os.File type satisfies the reader interface and is the main actor that's used to interact with file contents. The most common way to obtain an instance for reading purposes is with the os.Open function. It's very important to remember to close a file when you're done using it - this will not be obvious with short-lived programs, but if an application keeps opening files without closing the ones that it's done with, the application will reach the limit of open files imposed by the operating system and start failing the opening operations.

- One to get the limit of open files 
    
        ulimit -n
- Another to check how many files are open by a certain process 

        lsof -p PID

# Using buffers
> A data buffer, or just a buffer, is a part of memory that is used to store temporary data while it is moved. Byte buffers are implemented in the bytes package, and they are implemented by an underlying slice that is capable of growing every time the amount of data that needs to be stored will not fit.

If new buffers get allocated each time, the old ones will eventually be cleaned up by the GC itself, which is not an optimal solution. It's always better to reuse buffers instead of allocating new ones. This is because they make it possible to reset the slice while keeping the capacity as it is(the array doesn't get cleared or collected by the GC).

A buffer also offers two functions to show its underlying length and capacity. In the following example, we can see how to reuse a buffer with *Buffer.Reset*         and how to keep track of its capacity.

# Peeking content

Peeking is the ability to read content without advancing the reader cursor. Here, under the hood, the peeked data is strored in the buffer. Each reading operation checks whether there's data in this buffer and if there is any, that data is returned while removing it from the buffer. This works like a queue(first in first out).

The possibilities that this simple operation opens are endless, and they all derive from peeking until the desired sequence of data is found, and then the interested chunk is actually read. The most common uses of this operation include the following:

- The buffers keeps reading from the reader until it finds a newline character(read one line at time).
- The same operation is used until a space is found(read one word at a time).

The structure that allows an application to achieve this behavior is bufio.Scanner. This makes it possible to define what the splitting function is and has the following type: 

- The buffers keeps reading from the reader until it finds a newline character(read one line at time)
- The same operation is used until a space is found (read one word at a time)

The structure that allows an application to achieve this behavior is bufio.Scanner.

    type SplitFunc func(data []byte, atEOF bool) (advance int, toke []byte, err error)
> This function stops when an error is returned, otherwise it returns the number of bytes to advance in the content, and eventually a token.

- ScanBytes: Byte tokens
- ScanRunes: Runes tokens
- ScanWord: Words tokens
- ScanLines: Line tokens

We could implement a file reader that counts the number of lines with just a reader. The resulting program will try to emulate what the Unix wc -l command does.

# Closer and seeker
> There are two other interfaces that related to readers: io.Closer and io.Seeker:

    type Closer interface {
        Close() error
    }
    
    type Seeker interface {
        Seek(offset int64, whence int) (int64, error)
    }

These are usually combined with io.Reader, and the resulting interfaces are as follows:
    
    type ReadCloser interface {
        Reader
        Closer
    }
    type ReadSeeker interface {
        Reader
        Seeker
    }
    
The Close method ensures that the resource gets released and avoids leaks, while the Seek method makes it possible to move the cursor of the current object(for example, a Writer) to the desired offset from the start/end of the file, or from its current position.

The os.File structure implements this method so that it satisifes all the listed interfaces. It is possible to close the file when the operations are concluded, or to move the current cursor around, depending on what you are trying to achieve.

# Writing to file
> As we have seen for reading, there are different ways to write files, each one with its own flaws and strengths. In the ioutil package, for instance, we have another function called WriteFile that allows us to execute the whole operation in one line. This includes opening the file, writing its contents, and then closing it.

# Writer interface
> The same principle that is valid for reading also applies for writing - there's an interface in the io package that determines writing behaviors, as shown in the following code:

    type Writer interface {
        Write(p []byte) (n int, err error)
    }

The io.Writer interface defines one method that, given a slice of bytes, returns how many of them have been written and/or if there's been any errors. A writer makes it possible to write data one chunk at a time without there being a requirement to have it all at once. The os.File struct also happens to be a writer, and can be used in such a fashion.

We can use a slice of bytes as a buffer to write information piece by piece. In the following example, we will try to combine reading from the previous section with writing, using the io.Seeker capabilities to reverse its content before writing it.

# Buffers and format
> If you have to import a package to use one function or type, you should consider just copying the necessary code into your own package. If a package contains much more than what you need, copying allows you to reduce the final size of the binary. You can also customize the code and tailor it to your needs

# Efficient writing
> Each time the os.File method, that is, Write, is executed, this translates to a system call, which is an operation that comes with some overhead.

Generally speaking, it's a good idea, to minimize the number of operations by writing more data at once, thus reducing the time that's spent on such calls.

The bufio.Writer struct is a writer that wraps another writer, like os.File, and executes write operations only when the buffer is full. This makes it possible to execute a forced write with the Flush method, which is generally reserved until the end of the writing processes. A good pattern of using a buffer would be the following:

    var w io.WriteCloser
    // initialise write
    defer w.Close()
    b := bufio.NewWriter(w)
    defer b.Flush()
    // write operations

defer statements are executed in reverse order before returning the current function, so the first Flush ensures that whatever is still on the buffer gets written, and then Close actually closes the file. If the two operations were executed in reverse order, flush would have tried to write a closed file, returning an error, and failed to write the last chunk of information.

# File modes
> We saw that the os.OpenFile function makes it possible to choose how to open a file with the file mode, which is a uint32 where each bit has a meaning(like Unix files and folder permissions). The os package offers a series of values, each one specifying a mode, and the correct way to combine them is with | (bitwise OR).

# State
> The os package offers the FileInfo interface, which returns the metadata of a file, as shown in the following code:

    type FileInfo interface {
        Name() string
        Size() int64
        Mode() FileMode
        IsDir() bool
        Sys() interface{}
    }

The os.Stat function returns information about the file with the specified path.

# Changing properties
> In order to interact with the filesystem and change these properties, three functions are available:

- func Chmod(name string, mode FileMode) error: Changes the permissions of a file
- func Chown(name string, uid, gid int) error: Changes the owner and group of a file
- func Chtimes(name string, atime time.Time, mtime time.Time) error: Changes the access and modification time of a file

# Handling Streams
> It also focuses on the missing parts of the input and output utilities that combine them in several different ways, with the goal being to have full control of the incoming and outgoing data.

# Technical requirements

- Streams 
> Writers and readers are not just for files; they are interfaces that abstract flows of data in one direction or another. These flows, often referred to as streams, are an essential part of most applications.

- Input and readers
> Incoming streams of data are considered the io.Reader interface if the application has no control over the data flow, and will wait for an error to end the process, receiving the io.EOF value in the best case scenario, which is a special error that signals that there is no more content to read, or another error otherwise. The other option is that the reader is also capable of terminating the stream. 


Besides os.File, there are several implementations of readers spread across the standard package.

# The bytes reader
> The bytes package contains a useful structure that treats a slice of bytes as an io.Reader interface, and it implements many more I/O interfaces:

- io.Reader: This can act as a regular reader
- io.ReaderAt: This makes it possible to read from a certain position onward
- io.WriteTo: This makes it possible to write the contents with an offset
- io.Seeker: This can move the reader's cursor freely
- io.ByteScanner: This can execute a read operation for each byte separately
- io.RuneScanner: This can do the same with characters that are made of more bytes

# The strings reader
> The strings package contains another structure that is very similar to the io.Reader interface, called strings.Reader. This works exactly like the first but the underlying value is a string instead of a slice of bytes.

One of the main advantages of using a string instead of the byte reader, when dealing with strings that need to be read, is the avoidance of copying the data when initializing it. This subtle difference helps with both performance and memory usage because it does fewer allocations and requires the Garbage Collector(GC) to clean up the copy.

# Defining a reader 
> Any Go application can define a custom implementation of the *io.Reader* interface. A good general rule when implementing interface is to accept interfaces and return concrete types, avoiding unnecessary abstraction. 

Let's look at a practical example. We want to implement a custom reader that takes the content from another reader and transforms it into uppercase;

# output and writers
> The reasoning that applies to incoming streams also applies to outgoing ones. We have the io.Writer interface, in which the application can only send data, and the io.WriteCloser interface, in which it is also able to close the connection.


# The bytes writer
> We already saw that the bytes package offers Buffer, which has both reading and writing capabilities. This implements all the methods of the ByteReader interface, plus more than one Writer interface:

- io.Writer: This can act as a regular writer.
- io.WriterAt: This makes it possible to write from a certain position onward
- io.ByteWriter: This makes it possible to write single bytes

bytes.Buffer is a very flexible structure considering that it works for both, Writer and ByteWriter and works best if reused, thanks to the Reset and Truncate methods. Instead of leaving a used buffer to be recycled by the GC and make a new buffer, it is better to reset the existing one, keeping the underlying array for the buffer and setting the slice length to 0.

# The string writer
> A byte buffer executes a copy of the bytes in order to produce a string. This is why, in version 1.10, strings.Builder made its debut.

# Defining a writer
> Any custom implementation of any writer can be defined in the application. A very common case is a decorator, which is a writer that wraps another writer and alter or extends what the original writer does.

As for the reader, it is a good habit to have a constructor that accepts another writer and possibly wraps it in order to make it compatible with a lot of the strandard library structures, such as the following:

- *os.File
- *bytes.Buffer
- *strings.Builder


