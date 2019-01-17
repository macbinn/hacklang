## hacklang
The Hack programming language

## build
```bash
$ dep ensure
$ make test && make build
```

Run file
```bash
$ ./bin/hacklang example/hello_world.hack
```

Format file
```bash
$ ./bin/hacklang --format example/hello_world.hack
```

Start REPL shell (without any args)
```bash
$ ./bin/hacklang
```

## example
number
```javascript
-> 1
1
```
bool
```javascript
-> true
true
-> false
false
```
string
```javascript
-> `hello world`
`hello world`
```
list
```javascript
-> [1, 2, 3]
[1, 2, 3]
```
map
```javascript
-> {a: 1, b: 2}
{a: 1, b: 2}
```
function
```javascript
-> i => { print(i) }
<function>
```
See [example](./example) for more examples.
