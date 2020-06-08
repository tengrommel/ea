# defer
- What is defer?
> Deferred functions are invoked immediately before the surrounded function returns, in the reverse order they were in the reverse order they were deferred
- How can we use it?
    - close file
    - making http requests
    - use defer when working with locks
    - defer works with conditions
    - use defer to clean things up
    - use defer to build no-brainer apis
- How can it make our code better?
- Common gotchas
- How does defer work under the hood?
> use heap and _defer struct link
- What were the recent performance optimisations?