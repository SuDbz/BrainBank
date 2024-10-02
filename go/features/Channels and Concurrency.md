


## Channels <a name = "channels"></a>

### 1. Channels
- Channels in Go are a fundamental mechanism for communicating between goroutines. They provide a way to safely pass values between concurrent tasks, ensuring that data is shared in a controlled and synchronized manner.

- Goroutines are lightweight threads of execution managed by the Go runtime. They are much cheaper to create and manage than traditional operating system threads, making it easy to write concurrent programs in Go

#### Key characteristics of channels:

- Synchronous: When a goroutine sends a value on a channel, it blocks until another goroutine receives it. This ensures that data is not lost and prevents race conditions.
- Typed: Channels are strongly typed. Each channel can only carry values of a specific type.
- Unbuffered or Buffered: Channels can be unbuffered or buffered. Unbuffered channels require a sender and receiver to be ready simultaneously, while buffered channels can store a fixed number of values before blocking the sender.
- Closed: Channels can be closed, indicating that no more values will be sent. Receivers can detect a closed channel and terminate accordingly.


Example : 
```go
    func main() {
        //define a channel of type string 
        myChannel := make(chan string)
        //create a new thread which will input to channle
        go func(){
            myChannel <- "datat"
        }()

        //read from channel
        //this line of code is blocking code,since its wait for 
        //the goroutine to put the data to channel
        msg := <- myChannel

        fmt.Println(msg)
    }

```

```go
    func main() {
        myChannel := make(chan string)
        go func(){
            for i := 0; i < 10; i++ {
                myChannel <- "mychannle"+i
            }
            //close the channel 
            close(myChannel)
        }() 
        //For range loop iterates over the values sent on the channel until the channel is closed.
        //This line of code is blocking code,since its wait for the goroutine to put the data to channel.
        for value := range myChannel {
            fmt.Println(value)
        }
    }
```

```go
    func main() {
        //bufferd channel
        myChannel := make(chan string,5)
        go func(){
            for i := 0; i < 10; i++ {
                myChannel <- "mychannle"+i
            }
            //close the channel 
            close(myChannel)
        }() 
        //For range loop iterates over the values sent on the channel until the channel is closed.
        //This line of code is blocking code,since its wait for the goroutine to put the data to channel.
        for value := range myChannel {
            fmt.Println(value)
        }
    }
```

#### Common use cases for channels:

- Inter-goroutine communication: Channels are essential for coordinating the work of multiple goroutines.
- Data pipelines: Channels can be used to create data pipelines, where values are processed sequentially by a chain of goroutines.
- Synchronization: Channels can be used to synchronize the execution of goroutines, ensuring that certain events happen in a specific order.
- Cancelation signals: Channels can be used to send cancelation signals to goroutines, allowing them to terminate gracefully.

#### Buffered vs. Unbuffered Channels in Go

##### Unbuffered Channels:
- Require a sender and receiver to be ready simultaneously.
- If a sender attempts to send a value on an unbuffered channel and there's no receiver ready to receive it, the sender will block until a receiver becomes available.
- If a receiver tries to receive a value from an unbuffered channel and there's no sender ready to send a value, the receiver will block.

##### Buffered Channels:
- Can store a fixed number of values before blocking the sender.
- If a sender tries to send a value on a buffered channel and the buffer is full, the sender will block until a receiver takes a value from the buffer.
- If a receiver tries to receive a value from a buffered channel and the buffer is empty, the receiver will block until a sender sends a value.

##### When to use which:

###### Unbuffered channels
- When you want to ensure that data is processed immediately after it's sent.
- When you want to synchronize the execution of goroutines.
- When you want to implement a simple producer-consumer pattern.

###### Buffered channels
- When you want to decouple the sender and receiver, allowing them to operate at different rates.
- When you want to avoid blocking the sender if the receiver is not ready.
- When you want to implement a more complex producer-consumer pattern with buffering.



### 2. Select 
 select is a control flow statement in Go that allows you to wait for multiple channels to become ready for communication

```go 
select {
case clause1:
    // Code to execute if clause1 is ready
case clause2:
    // Code to execute if clause2 is ready
default:
    // Code to execute if no case is ready
}
```
(If no case is ready and there's no default clause, the select statement blocks until at least one case becomes ready.)

Example :
```go
func main() {
    c1 := make(chan int)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second)
        c1 <- 42
    }()

    go func() {
        time.Sleep(time.Millisecond * 500)
        c2 <- "Hello"
    }()

    select {
    case v1 := <-c1:
        fmt.Println("Received on c1:", v1)
    case v2 := <-c2:
        fmt.Println("Received on c2:", v2)
    default:
        fmt.Println("No channels ready")
    }
}
```


#### Channel References 
 [Channels and Concurrency](https://www.youtube.com/watch?v=qyM8Pi1KiiM)

<br><br><br><br><br>

## Concurrency <a name = "concurrency"></a>
Concurrency in Go refers to the ability of a program to perform multiple tasks simultaneously, even if they don't necessarily execute at the same time. This is achieved through the use of goroutines.

### 1. For-Select Loop
The for-select loop in Go is a combination of a for loop and a select statement. It allows you to iterate over a channel while also handling other concurrent operations. This is particularly useful when you need to process values from a channel but also want to be able to handle timeouts, cancelations, or other events.

```go
for {
    select {
    case value := <-ch:
        // Code to execute when a value is received from the channel
    case <-ctx.Done():
        // Code to execute when the context is canceled
    default:
        // Code to execute if no case is ready
    }
}
```

Example :

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		select {
		case value := <-ch:
			fmt.Println(value)
		case <-ctx.Done():
			fmt.Println("Context canceled")
			return
		}
	}
}
```

### 2. Done Channel
In Go, a "done channel" is a channel that is used to signal the completion or termination of a task. It's a common pattern in concurrent programming to coordinate the execution of multiple goroutines.


```go
func worker( done <-chan bool) {
    for {
        select {
        case <-done:
            fmt.Printf("Worker %d received done signal\n", id)
            return
        default:
            // Do some work
            fmt.Printf("Worker %d is working\n")
            time.Sleep(time.Second)
        }
    }
}

func main() {
    done := make(chan bool)
    //Call and infinite loop
    go worker(done)
    
    time.Sleep(5 * time.Second)
    //When called, the <-done channel will be activated and the goroutine will be terminated
    close(done)
    time.Sleep(time.Second)
}
```
###### Note
 - ```func worker( done <-chan bool)``` - The <-chan bool type indicates that done is a read-only channel that can only receive boolean values
 - ```func worker( done chan bool)```   - The <-chan bool type indicates that done is a read and write channel that can only receive boolean values


### 3. Pipeline

Go Pipelines are a powerful pattern for building concurrent and efficient data processing applications. They involve a series of connected goroutines that pass data between each other using channels. This creates a pipeline-like structure where data flows from one stage to the next.

#### Key components of a Go pipeline:
- Goroutines: The individual units of work that perform specific tasks within the pipeline.
- Channels: Used to communicate data between goroutines.
- Pipeline stages: Each stage in the pipeline is typically a goroutine that receives data from a channel, processes it, and sends the result to another channel.

```go 
unc main() {
    // Create channels for each stage
    sourceChan := make(chan int)
    transformChan := make(chan int)
    sinkChan := make(chan int)

    // Start the source goroutine
    go func() {
        for i := 0; i < 10; i++ {
            sourceChan <- i
        }
        close(sourceChan)
    }()

    // Start the transform goroutine
    go func() {
        for value := range sourceChan {
            transformChan <- value * 2
        }
        close(transformChan)
    }()

    // Start the sink goroutine
    go func() {
        for value := range transformChan {
            sinkChan <- value
        }
        close(sinkChan)
    }()

    // Read from the sink channel
    for value := range sinkChan {
        fmt.Println(value)
    }
}
```


 #### Concurrency References 
 - [Channels and Concurrency - 1](https://www.youtube.com/watch?v=qyM8Pi1KiiM)
 - [Channels and Concurrency - 2](https://www.youtube.com/watch?v=wELNUHb3kuA)
