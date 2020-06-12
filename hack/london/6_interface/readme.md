# interface in Go

- 1、Is the language called Go, Go or Golang?
> Go
- 2、Is Go an object-oriented language?

interface 不能被继承，智能被实现

# 其它语言的继承
- java和c#都只能单一继承
- c++可以多继承

# Deconstruction
> What are ingredient inheritance have?

Inheritance

- Function & Property reuse
- Polymorphism

Embedding(嵌入)
> 鸭子类型看似loose coupling 

但Go仍有一种constrained组合方式，使某类别struct explicit受到某个界面约束，必须implement其member method 
