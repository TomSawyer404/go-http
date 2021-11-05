# go-http

In the previous time, I wrote raw HTTP request in a text and
sent it to server by *nc* tool to test my API.  
This hard-core style do help me to understand HTTP proto. But
I think I could test API in a more effcient way.  
One night I found a tool called [httpie](https://httpie.io/).  
It's awesome and convenient. So I gonna make a new httpie with
Golang. **Only standard library**, it's a good way to learn Golang
for someone who new to Golang.  

@auhtod:     Mrbanana  
@date:       2021-9-4  
@licece:     The MIT License  

# Usage

```bash
$ make clean
$ make
$ ./go-http httpbin.org/status/418
```

Generally, it looks like `./go-http [host:port]`  
If `host` is empty, like `:8080`, the `host` is assumed as `localhost`  

# Feature

- In default the *go-http* use GET method;
- Use `./go-http :8080 name:banana age:12` to add your *request-headers*;
- Use `./go-http :8080 name=banana age=12` to send JSON data in POST method;
- Use `./go-http :8080 name-banana age#12` to send FORM data in POST method;

---
