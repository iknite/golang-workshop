# hello!

```
package main

import "fmt"

func main() {
	fmt.Println("hello world!")
}

```

```
$ go run src/hello/main.go
hello world!
```

---

# Background

Go is one of the few that enables parallelism and concurrency in language syntax
level, instead of provide a library for it.


Others with similar concepts:
> Occam, **Erlang**, Newsqueak, Alef, Limbo

----

#### CSP vs Actor model

Communicating Sequential Processes are a way to pass information trough
a **existing** channel. (ie. Read a file descriptor)

Actor Model are a way to pass information by name (ie. Read a file by name)

----

##### Functional Programming Vs Object Oriented Programming

- FP excels at do operations over things (digest/operate)
- OOP excels at provide things over operations (update/change)

Pure languages don't exists (not Haskell nor old Java)

----

##### Service Oriented Architecture vs Object Oriented Architecture

- SOA Architects over processes (stateless, microservices, **#hyped**)
- OOA Architects over resources (statefull, monolith, **#ranted**)

Grab the good of both worlds when needed

---

# Golang gogorama!

- Fast to write (almost like python)
- Fast to run (almost like C)

> https://golang.org/dl/

----

## Small toolset but powerful

- Dependencies are controled by an environment variable **GOPATH** that changes where **go get** installs external
requirements
```
GOPATH=$(realpath ./t dependencies) go get code.google.com/p/go-tour/gotour
```

- Test are included out of the box
```
go test src/hello/main.go
```

- And so the compiler
```
go build src/hello/main.go -o hello
```

- easy to deploy **Selfcontained executable FTW!**

----

## Great Documentation and tutorials

- [gotour](https://tour.golang.org/welcome/1)
- https://golang.org/doc/
- https://github.com/golang/go/wiki/GoTalks

----

## Object Composition

- *Structs* are data types, and it can have methods asociated. (non-virtual)
- Dynamic method calls only trought Interfaces (Structs without attributes, only virtual methods)
- Ducktyping by interface matching at variable definition, instead of class composition
- Closures! lets do more of it
- *Pointers hurray!


---

# Hands on!

```
 # TODO: deal with oauth is not a 2h thing (for me...)
```

---

# Example exercise

Give a list of the folower that gives you the most `pity likes`. 
It liked a twit with low difusion (likes, retweets)

```
# fetch the timeline
for twit in user.timeline:
    # fetch twit information
    for like in twit.likes:
        users[like.user] += 1 / twit.total_likes

print(sort(users[5:]))
```
A.k.a a typical map/reduce problem

Thanks @honduco for the idea!!

----

### API

Fetch the timeline
```
https://dev.twitter.com/rest/reference/get/statuses/user_timeline?screen_name=<screen_name>&count=200
```
#### SCAPPY

Fetch the likes of a twit
```
https://twitter.com/<screen_name>/status/<twit_id>
```

---

Sorry for party rockin' @fevereng 2016

Note:
# Resources
https://gist.github.com/kachayev/21e7fe149bc5ae0bd878
https://github.com/adityamenon/Google-IO_2012_Go-Concurrency-Patterns
http://stackoverflow.com/questions/2078978/functional-programming-vs-object-oriented-programming
http://programmers.stackexchange.com/questions/164339/relationship-between-soa-and-ooa
https://github.com/golang/go/wiki/GoTalks
https://en.wikipedia.org/wiki/Service-oriented_architecture
https://en.wikipedia.org/wiki/Actor_model_and_process_calculi_history
https://en.wikipedia.org/wiki/Actor_model
https://en.wikipedia.org/wiki/Communicating_sequential_processes
https://en.wikipedia.org/wiki/Concurrent_computing
