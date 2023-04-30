# go channels

What if two goroutines want to communicate with each other? For that, Go has channels.

Go, channels are a means through which different goroutines communicate.

Think of them as pipes through which you can connect with different concurrent goroutines.
The communication is bidirectional by default, meaning that you can send and receive values
from the same channel.

Moreover, by default, channels send and receive until the other side is ready. This allows
goroutines to synchronize without explicit locks or condition variables.

Syntax:

var Channel_name chan Type

or

Make a channel: make(chan [value-type]), [value-type] is the data type of the values to send and receive, e.g., int

Send and receive values: channel <- and <- channel, where <- is the channel operator.

Close a channel: close(channel), After closing, no value will be sent to the channel.

Both sending and receiving are blocking operations by default.

its kind of asyn/wait methodology

- Send operation:

The send operation is used to send data from one goroutine to another goroutine with the help of a channel. Values like
int, float64, and bool can safe and easy to send through a channel because they are copied so there is no risk of
accidental concurrent access of the same value. Similarly, strings are also safe to transfer because they are immutable.
But for sending pointers or reference like a slice, map, etc. through a channel are not safe because the value of pointers or
reference may change by sending goroutine or by the receiving goroutine at the same time and the result is unpredicted.
So, when you use pointers or references in the channel you must make sure that they can only access by the one goroutine at a time.

Mychannel <- element

The above statement indicates that the data(element) send to the channel(Mychannel) with the help of a <- operator.

- Receive operation:

The receive operation is used to receive the data sent by the send operator.

element := <-Mychannel

The above statement indicates that the element receives data from the channel(Mychannel).

If the result of the received statement is not going to use is also a valid statement.
You can also write a receive statement as:

<-Mychannel

## Uni directional channel

As we know that a channel is a medium of communication between concurrently running goroutines so that they can send and receive data to each other. By default a channel is bidirectional but you can create a unidirectional channel also. A channel that can only receive data or a channel that can only send data is the unidirectional channel. The unidirectional channel can also create with the help of make() function as shown below:

// Only to receive data
c1:= make(<-chan bool)

// Only to send data
c2:= make(chan<- bool)

[Geek for geeks](https://www.geeksforgeeks.org/unidirectional-channel-in-golang)

## Buffered channel

Channels can be defined as a pipes using which Goroutines communicate. Similar to water flows from one end
to another in a pipe, data can be sent from one end and received from the another end using channels.

By default channels are unbuffered, which states that they will only accept sends (chan <-) if there is
a corresponding receive (<- chan) which are ready to receive the sent value.

Buffered channels allows to accept a limited number of values without a corresponding receiver for
those values. It is possible to create a channel with a buffe. Buffered channel are blocked only when the
buffer is full. Similarly receiving from a buffered channel are blocked only when the buffer will be empty.

Buffered channels can be created by passing an additional capacity parameter to the make( ) function which specifies the size of the buffer.

Syntax :

ch := make(chan type, capacity) // chan defines channel type

Here , capacity in the above syntax should be greater than 0 for a channel to have a buffer.
The capacity for an unbuffered channel is 0 by default and hence it omit the capacity parameter.

## Select

- it is kind of switch in golang but with channels
- The select statement let a goroutine wait on multiple communication operations
- In select each of the case statement wait for a send or receive operation from a channel.
- But in switch each of the case statement is an expression.
- Important
  - Select blocks until any of the case statements are ready
  - If multiple case statements are ready then it selects one at random and proceeds.
- use-cases
  - The select statements let a goroutine wait on multiple communication operations
  - Select along with channels and goroutines becomes _very powerful tool for managing synchronization and concurrency_
- Default case
  - Like switch statement,we can have a default case in select too.
  - This default case will be executed if no send or receive operations is ready on any of the case statement.
  - Default block makes the select non blocking as default case will be executed if all the other cases are blocked.
- Select vs Switch
  - Switch - Non blocking
  - select - can block since they are used with channels and they can block or receive operations.
  - switch - Deterministic and will run in sequence to select the matching case.
  - select - Non deterministic as select will chose the case randomly in non deterministic manner.

## Buffered vs Unbuffered channels

- Proper use of buffered channel means that you must handle the case where
  the channel is blocking, which might happen due to waiting on sender/receiver.
- Buffered channels are useful when you know how many go-routines you
  have launched, want to limit the number of go-routines you will launch, or
  want to limit the amount of work that is queued up.

## Reference

- [Geeks for geeks](https://www.geeksforgeeks.org/channel-in-golang/?ref=lbp)
