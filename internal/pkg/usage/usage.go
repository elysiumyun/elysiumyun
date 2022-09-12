package usage

import "fmt"

// app usage menual
func Usage() {
	fmt.Println(`Usage: elysium [-c <file>] -[hv]

If you are using it for the first time, you can check usage or README.md.

     ------- < Commands Arguments > -------

argument:
  -c, config        Set external config file. (e.g elysium -c <external config file>)

optional:
  -h, help          Show this help message. 
  -v, version       Show the app version. 

For more help, you can use 'go-web help' for the detailed information
or you can check the docs: https://github.com/elysiumyun/elysium`)
}
