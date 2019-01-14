## hacklang
The Hack programming language

## run
```bash
$ dep ensure
```

Start REPL shell
```bash
$ go run cmd/main.go
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

### run file
```bash
$ go run cmd/main.go example/hello_world.hack
```

### format file
```bash
$ go run cmd/main.go --format example/hello_world.hack
```