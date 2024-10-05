 [Home](../go.md)

# Goroutines
 

Goroutines are lightweight threads of execution managed by the Go runtime. They are a fundamental building block for concurrent programming in Go, allowing you to execute multiple tasks simultaneously without the overhead of traditional operating system threads.


```go
go func() {
    // Code to be executed concurrently
}()
```
### Key characteristics of goroutines:

- Lightweight: Goroutines are much smaller and cheaper to create and manage than operating system threads.
Stackless: Goroutines do not have a fixed stack size, and their stack size can grow and shrink dynamically as needed.
- Managed by the runtime: The Go runtime schedules and manages goroutines, optimizing their execution across multiple cores.
- Concurrent execution: Multiple goroutines can run concurrently, even on a single-core system.
- Channels: Goroutines communicate with each other using channels, which provide a safe and efficient way to pass data between concurrent tasks.

### Key points to remember:

- Channels: Goroutines communicate with each other using channels. Channels provide a safe and efficient way to pass data between goroutines.
- Synchronization: When multiple goroutines access shared data, you need to use synchronization primitives like mutexes or semaphores to avoid race conditions.
- Deadlocks: Be careful to avoid deadlocks, where two or more goroutines are waiting for each other to finish.
- Context: Use the context package to manage the lifecycle of goroutines and handle cancellations or timeouts.


Example:
```go

go func() {
    fmt.Println("I am goroutines")
}()

fmt.Println("I am main")

```

Or

```go
func main() {
    go func() {
        for i := 0; i < 3; i++ {
            fmt.Println("Goroutine 1:", i)
            time.Sleep(time.Second)
        }
    }()

    for i := 0; i < 3; i++ {
        fmt.Println("Main goroutine:", i)
        time.Sleep(time.Second)
    }
}

```

<br><br>
<br>

## Synchronization Mechanism

### 1. Wait Group 

A waitgroup is a synchronization mechanism used to block a goroutine until one or more other goroutines have completed their execution. It's particularly useful when you need to coordinate multiple concurrent tasks and ensure that they all finish before proceeding with further operations.

- *Add(delta int)* : Increments the internal counter by the specified delta. Typically, you'll use Add(1) to indicate that a new goroutine has started.
- *Done()* : Decrements the internal counter by 1. This is called by a goroutine when it has finished its work.
- *Wait()* : Blocks the calling goroutine until the internal counter reaches 0, indicating that all goroutines have completed.

```go 
func worker(id int, wg *sync.WaitGroup) {
        defer wg.Done() // Decrement counter when worker finishes
        fmt.Printf("Worker %d started\n", id)
        // Do some work here
        fmt.Printf("Worker %d finished\n", id)
}
```
```go

func main() {
        var wg sync.WaitGroup
        for i := 1; i <= 5; i++ {
                wg.Add(1) // Increment counter for each worker
                go worker(i, &wg)
        }
        wg.Wait() // Wait for all workers to finish (blocking code)
        fmt.Println("All workers have completed")
}

```


### 2. Mutex

Mutex (Mutual Exclusion) is a synchronization mechanism in Go that ensures that only one goroutine can access a shared resource at a time. It's used to prevent race conditions and data corruption when multiple goroutines are working with the same data.

- Lock(): Acquires the mutex, blocking the current goroutine if it's already locked by another goroutine.
- Unlock(): Releases the mutex, allowing another goroutine to acquire it.

```go
var counter int
var mutex sync.Mutex
```
```go
func increment(wg *sync.WaitGroup) {
        defer wg.Done()
        mutex.Lock()
        counter++
        mutex.Unlock()
        fmt.Println("Counter incremented:", counter)
}
```

```go
func main() {
        var wg sync.WaitGroup
        for i := 0; i < 100; i++ {
                wg.Add(1)
                go increment(&wg)
        }
        wg.Wait()
        fmt.Println("Final counter value:", counter)
}
```

##### Key points to remember:

- Always use a mutex to protect shared data when multiple goroutines are accessing it.
- Acquire the mutex before accessing the shared data and release it after you're done.
- Be careful to avoid deadlocks by ensuring that goroutines don't acquire mutexes in different orders.
- Consider using more advanced synchronization mechanisms like read-write locks or atomic operations when appropriate.



<br><br>
#### Goroutines References 
 - [Goroutines Crash Course (Mutex, Channels, Wait Group, & More!)](https://www.youtube.com/watch?v=5Z8skvm4g64)

