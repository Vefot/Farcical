# The Farcical Programming Language

An interpreted programming language that isn't good for much.

## Usage

```go run main.go```

```
 ______ 
|  ____|
| |__   
|  __|  
| |     
|_|

Farcical v0.0.0
REPL Session: main
>>>
```

### Example
```example.fa```

```javascript
let tryThis = fn(foo, bar) {
    let result = foo * bar + 10;
    return result
}

let myNumber = tryThis(1, 2);

print(myNumber + 10)

let thisHash = {"apples": 10, "oranges": 5, 50:"bananas"}

print(thisHash["oranges"])
print("and we have", thisHash["apples"], "apples")

let thisList = [1,2,"three", 4]
print(thisList[2])
```

```go run main.go -file example.fa```

```
22 
5 
and we have 10 apples 
three
```