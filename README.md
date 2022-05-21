## portscanner

### I built the Go-portscanner to scan the open ports without any time consumption and hassle of the network.

– Checking all the command line argument has passed accurately.
– Number of workers (functions) which will be set to work in go-routine.
– Usage of the channel, assign ports to scan it worker to sub-routine.
– When the go-routine finishes, set back to main method.
– Displaying all the sorted scanned open ports from the network.

Working of the [script](https://github.com/ajay-deshmukh/portscanner/blob/master/portscanner.go) I developed.

![Alt text](https://github.com/ajay-deshmukh/portscanner/blob/master/output/screenshot.png)

Locating the open port by passing generalize subroutine for non-preemptive
multitasking by allowing execution to be suspended and resumed state (halt).
It means to implement go-routines (Go Lang Method) that run concurrently not
as the parallel execution, but familiar with cooperative tasks, exception, 
event loops, iterators, infinite lists and pipes.
