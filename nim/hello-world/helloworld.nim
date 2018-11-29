import strformat

proc helloworld*(name = "World"): string =
    return fmt"Hello, {name}!"
