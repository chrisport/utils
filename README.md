# utils
Collections of some util packages that are useful sometimes

## [exek](https://github.com/chrisport/utils/tree/master/exek)

Small wrapper around exec.Command to quickly execute bash commands and receive console output as string.
Intended for experimentation, not for production code (no error handling planned).

```outputAsString := exek.Call("touch asd.yxc && ls | grep asd && rm asd.yxc ")```
