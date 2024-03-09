# Glua
An intererpreter for a minimal subset of Lua written in go.

## Example

Create a lua file:
```lua
function fib(n)
   if n < 2 then
      return n;
   end

   local n1 = fib(n-1);
   local n2 = fib(n-2);
   return n1 + n2;
end
```

```sh
make build
./bin/glua ./fib.lua
```

For more examples take a look at the `test/` directory
