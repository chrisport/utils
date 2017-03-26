# exek

Small wrapper around exec.Command to quickly execute bash commands and receive console output as string.
Intended for experimentation, not for production code (no error handling planned).


## Example usage
```
	out := exek.Call("touch asd.yxc && ls | grep asd && rm asd.yxc ")

	fmt.Println(out)
	// Prints
	// asd.yxc

```

Formatted command with arguments
```
    f := "asd.zxc"
	out := exek.Call("touch %v && ls | grep asd && rm %v", f, f)

	fmt.Println(out)
```

Formatted command with arguments
```
    f := "asd.zxc"
	out := exek.Call("touch %v && ls | grep asd && rm %v", f, f)

	fmt.Println(out)
```