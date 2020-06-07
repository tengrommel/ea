# Topics

- Linters -a3 -min history
- A tour of the Go static-analysis ecosystem
- The side effects of heavy linting

# Removing all the fluff...

- lint was a Unix tool developed for C in the 70's.
- Mostly performance-focused recommendations

# As time went by 
- Compilers started to take over these responsibilities
- lint-like tools evolved to detect other unpleasantness

# Fast-forward to the present
- Linting is not reserved for C-like languages
> Java,Python,Haskell,Docker files,YAML,...?
- Static-analysis is an entire research domain by itself
> A quick search on a computer-science research library gives 20000 hits for 2019 alone

# Static-analysis in Go

- Static-analysis is pattern-matching on code's representation
- An accessible code representation means easier static-analysis
- Easier static-analysis means more will be produced and used

# Go has standard-library packages for its code representation
- go/ast
- go/parser
- go/token

# And many that help you write your own static-analysis pass:
- golang.org/x/tools/go/analysis
- golang.org/x/tools/go/ast
- golang.org/x/tools/go/packages