## What is a Scripting Language?

* https://en.wikipedia.org/wiki/Scripting_language

> … a programming language designed for integrating and communicating with other programming languages, processes, routines.

* Examples:
  * Bash, JavaScript, VBScript, PHP, Perl, Python, Ruby

>  … programming languages that don't require an explicit compilation step. They are interpreted from source code/bytecode - one command at a time.

* Very blurry:
  * Compilation can be quite fast with modern hardware and modern compilation techniques.
    * You could design a interpreter for C language and use it as a "Scripting Language"
    * You could design a compiler for JavaScript and use it as a "Compiled Programming Language"
      * ooops, this already happened: V8, Chromes JavaScript engine compiles JavaScript code into machine code, rather than interpreting it (V8 is an optimizing two-phase compiler)
    * ...and then you have things like Python that sit in both camps: Python is widely used without a compilation step, but the main implementation (CPython) does that by compiling to bytecode on-the-fly and then running the bytecode in a VM, and it can write that bytecode out to files (.pyc, .pyo) for use without recompiling.
* We should not talk about the language, **what about the runtime?**
  * **Does the runtime environment "see" the source code?** If so, I'd call that "scripting;" if not, I wouldn't. So in that sense, browser-based JavaScript is "scripting" because even if engines like V8 compile it on-the-fly, the source is still delivered to the runtime environment. Similarly shells scripts.
    A traditional C program's source *isn't* delivered to the runtime.
* Its already hard to even define what a scripting language is. So why can't we consider go as a Scripting Language? 😋



---

## Why do we want to consider Go as a Scripting Language?

- **strongly typed language** —> you will find many errors even before you hit "save" and not at run time
- When our **project is written in go, why aren't the deployment scripts?** (`bash`, `make`, ...)
  - the whole CI/CD routines could be written in go aswell!
- Write the logic & execute it like `./helloworld.go`, without bothering about compiling
- Go is backwards compatible, so everything you write now, will be executable in the future (e.g. in contrast to pythin 2 & python 3)



---

## Can we use Go at least "like" a Scripting Language?

* ✅ **Invocation** feels very scripty
  `go run hello.go`

* ✅ **REPL** Read-Eval-Print-Loop

  * with [gore](https://github.com/motemen/gore) we have a REPL like in Python or Node.js
  * however: gore works by using `go run` for each input 🙈

* ⚠️ A proper script needs return values! How does `go run` deal with return values?
  Give it a try!

  ```go
  package main

  import (
  	"fmt"
  	"os"
  )

  func main() {
  	s := "world"

  	if len(os.Args) > 1 {
  		s = os.Args[1]
  	}

  	fmt.Printf("Hello, %v!", s)
  	fmt.Println("")

  	if s == "fail" {
  		os.Exit(30)
  	}
  }
  ```

  ```bash
  $ go run hello.go gopher
  Hello, gopher!
  $ echo $?
  0
  $ go run hello.go fail
  Hello, fail!
  exit status 30
  $ echo $?
  1 # <---- WTF?
  ```

  Okay, so `go run` is useless for a lot of script use cases 😥



* But what about invoking a file like a shell script `./hello.go`?

  * With a [Shebang Line](https://en.wikipedia.org/wiki/Shebang_(Unix))?

    * `#!interpreter [optional-arg]`

      `#!/bin/bash`
      `#!/usr/bin/env python`

    * `#!` tells the _shell_ which interpreter is necessary to run the script, so the system knows exactly how to execute the script regardless of the programming language used

    * Nice! What about adding `#!/usr/bin/env go run` to my `hello.go` file?

      * unfortunately not:

        ```
        $ chmod +x hello.go
        $ ./hello.go 
        package main: 
        hello.go:1:1: illegal character U+0023 '#'
        ```

        - so `go run` tries to deal with the file, but #!-lines are no proper go code :(
        - we still would have the problem with the error codes...

    * **`gorun` [⬆️](https://github.com/erning/gorun) to the rescue!**

      ```
      $ go get github.com/erning/gorun
      ```

      ​

```go
#!/usr/bin/env gorun

package main
[...]
```

```bash
# install gorun
$ go get github.com/erning/gorun
$ sudo ln -s ~/go/bin/gorun /usr/local/bin/ #optional
# make script executable
$ chmod +x hello.go
```

```bash
$ ./hello.go
Hello, world!
$ ./hello.go gopher
Hello, gopher!
$ ./hello.go fail
Hello, fail!
$ echo $?
30
```

- write compiled binaries under a safe directory in $TMPDIR (or /tmp), so that the actual script location isn't touched (may be read-only)
- avoid races between parallel executions of the same file
- automatically clean up old compiled files that remain unused for some time (without races)
- replace the process rather than using a child
- pass arguments to the compiled application properly
- handle well GOROOT, GOROOT_FINAL and the location of the toolchain

Drawbacks:

* file does not compile anymore with `go build` (damn you `#!` shebang)

  ```bash
  $ go run hello.go 
  package main: 
  hello.go:1:1: illegal character U+0023 '#'
  ```

* no `go fmt` [⬆️](https://github.com/erning/gorun/issues/8) (damn you `#!` shebang)

  ```bash
  $ go fmt hello.go 
  can't load package: package main: 
  hello.go:1:1: illegal character U+0023 '#'
  ```

* local package import doesn't work [⬆️](https://github.com/erning/gorun/issues/9)

#### Is this the best we can do?

Install `gorun` and add a shebang line to every "script"-file?



> In [Unix-like](https://en.wikipedia.org/wiki/Unix-like) operating systems, **when a text file with a shebang is used as if it is an executable, the [program loader](https://en.wikipedia.org/wiki/Loader_(computing)) parses the rest of the file's initial line as an [interpreter directive](https://en.wikipedia.org/wiki/Interpreter_directive)**; the specified interpreter program is executed, passing to it as an argument the path that was initially used when attempting to run the script,[[8\]](https://en.wikipedia.org/wiki/Shebang_(Unix)#cite_note-linux-8) so that the program may use the file as input data. For example, if a script is named with the path *path/to/script*, and it starts with the following line, `#!/bin/sh`, then the program loader is instructed to run the program */bin/sh*, passing *path/to/script* as the first argument.
>
> *Source: https://en.wikipedia.org/wiki/Shebang_(Unix)*





### ⚠️⚠️️⚠️⚠️ Disclaimer: superficial knowledge ⚠️⚠️⚠️⚠️







What is the program loader? —> ↗️ [What exactly interpret #!/bin/bash line?](https://superuser.com/questions/117721/what-exactly-interpret-bin-bash-line)

> At least in Linux, the kernel has this functionality: `fs/binfmt_script.c` specifically.
>
> _— a nice guy from the internet_

_Wanna know more? https://www.in-ulm.de/~mascheck/various/shebang/_

Can we tell the system on a lower level how to treat `.go` files? Yes, we can - at least on linux!





#### `binfmt_misc`:

> This Kernel feature allows you to **invoke almost (for restrictions see below) every program by simply typing its name in the shell**. This includes for example compiled Java(TM), Python or Emacs programs.
>
> To achieve this you must tell binfmt_misc which interpreter has to be invoked with which binary. `Binfmt_misc` recognises the binary-type by matching some bytes at the beginning of the file with a magic byte sequence (masking out specified bits) you have supplied. **`Binfmt_misc` can also recognise a filename extension aka `.com` or `.exe`**.

```bash
# check whether binfmt_mis is already mounted
$ mount | grep binfmt_misc
systemd-1 on /proc/sys/fs/binfmt_misc type autofs (rw,relatime,fd=27,pgrp=1,timeout=0,minproto=5,maxproto=5,direct)
# if yes, then:
$ echo ':golang:E::go::/usr/local/bin/gorun:OC' | sudo tee /proc/sys/fs/binfmt_misc/register
:golang:E::go::/usr/local/bin/gorun:OC
```

wtf just happened?

```
:golang  # name (an identifier string)
:E       # type of recognition (M = Magic, E = Extension)
:        # offset; important for Magic recognition
:go      # is the byte sequence binfmt_misc is matching for OR matching FILENAME EXTENSION
:        # mask (...)
:/usr/local/bin/gorun # interpreter - the program that should be invoked with the binary
:OC      # flags (O = open-binary; C = credentials (?))
```

>  The `OC` flags at the end of the string make sure, that the script will be executed according to the owner information and permission bits set on the script itself, and not the ones set on the interpreter binary. This makes Go script execution behaviour same as the rest of the executables and scripts in Linux.

Further Info:

https://www.kernel.org/doc/html/v4.14/admin-guide/binfmt-misc.html

https://git.kernel.org/pub/scm/linux/kernel/git/stable/linux-stable.git/tree/fs/binfmt_misc.c?h=linux-4.14.y

```
# REMOVE THE SHEBANG!
$ chmod u+x hello.go
$ ./hello.go
Hello, world!
$ ./hello.go gopher
Hello, gopher!
$ ./hello.go fail
Hello, fail!
$ echo $?
30
```

That's it! Now we can edit `hello.go` to our liking and see the changes will be immediately visible the next time the file is executed. Moreover, unlike the previous shebang approach, we can compile this file any time into a real executable with `go build`.

### But… What about MacOS?

### `xbinary` [⬆️](http://www.osxbook.com/software/xbinary/)

* make it available in your path: `ln -s "/Library/Application Support/xbinary/xbinary" /usr/local/bin/`

xbinary debug:

trying to load kext file

```bash
➜ sudo xbinary -E                                                                
/Library/Application Support/xbinary/xbinary.kext failed to load - (libkern/kext) not loadable (reason unspecified); check the system/kernel logs for errors or try kextutil(8).
Failed to load XBinary kext.
```

kextutil

```bash
➜ sudo kextutil /Library/Application\ Support/xbinary/xbinary.kext
Diagnostics for /Library/Application Support/xbinary/xbinary.kext:
Warnings: 
    Executable does not contain code for architecture: 
        x86_64

Code Signing Failure: not code signed
Warnings: 
    Executable does not contain code for architecture: 
        x86_64

Untrusted kexts are not allowed
ERROR: invalid signature for com.osxbook.driver.xbinary, will not load
```

kernel log:

```
➜ log show --predicate "processID == 0" --start 2018-04-15 --debug | grep xbinary
2018-04-15 21:44:44.867416+0200 0x25d6c    Default     0x0                  0      kernel: kmod_get_info is not supported on this kernel architecture (called from xbinary)
```

## Summary

| Type                    | Exit Code | Executable | Compilable | Go Standard | Local Packages | Linux | MacOS | Windows |
| ----------------------- | --------- | ---------- | ---------- | ----------- | -------------- | ----- | ----- | ------- |
| `go run`                | ✘         | ✘          | ✔          | ✔           | ✔              | ✔     | ✔     | ✔       |
| `gorun` (`#!`)          | ✔         | ✔          | ✘          | ✘           | ✘              | ✔     | ✔     | ???     |
| `gorun` + `binfmt_misc` | ✔         | ✔          | ✔          | ✔           | ✔              | ✔     | ✘     | ✘       |

Sources

* https://stackoverflow.com/questions/17253545/scripting-language-vs-programming-language
* https://www.geeksforgeeks.org/whats-the-difference-between-scripting-and-programming-languages/
* https://www.quora.com/What-is-scripting-language-How-is-it-different-from-programming-language
* https://github.com/motemen/gore/issues
* https://en.wikipedia.org/wiki/Shebang_(Unix)
* https://github.com/erning/gorun
* https://gist.github.com/posener/73ffd326d88483df6b1cb66e8ed1e0bd
* https://blog.cloudflare.com/using-go-as-a-scripting-language-in-linux/
* https://www.kernel.org/doc/html/v4.14/admin-guide/binfmt-misc.html