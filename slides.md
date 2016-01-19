# hello!

This is quick review of the goodies of golang

----

## Background

Go is one of the few that enables parallelism and concurrency in language syntax
level, instead of provide a library for it.


Others with similar concepts:
> Occam, **Erlang**, Newsqueak, Alef, Limbo

----

#### CSP vs Actor model

Commincating Sequential Processes `CSP` are a way to pass information trought 
a **existing** channel. (ie. Read a file descriptor)

Actor Model are a way to pass information by name (ie. Read a file by name)

----

##### Functional Programming Vs Object Oriented Programming

- FP excels at do operations over things (transform/operate)
- OOP excels at give things over operations (update/change)
- Pure languages don't exists (not Haskell nor ol' Java)

----

##### Service Oriented Architecture vs Object Oriented Architecture

- SOA Architects over processes (stateless, microservices, **#hype**)
- OOA Architects over resources (statefull, monolit, **#rant**)

----

# Stop lecturing me!
And go to the point

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
A.k.a a typical map/reduce problem (got hadoop?)
Thanks @honduco for the idea!!

----

### API

Fetch the timeline
```
https://userstream.twitter.com/1.1/user.json?replies=all
```

Fetch the likes of a twit
```

```


---

# Golang gomorrowland

- Fast to write (almost like python)
- Fast to run (almost like C)

> https://golang.org/dl/



---

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

---
Sorry for party rockin' @FeverEng 2016

