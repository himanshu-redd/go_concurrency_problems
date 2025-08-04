### Beginner Level (Problems 1-30)

These problems focus on the fundamentals of goroutines and basic channel usage.

1. [**Hello Goroutine:** Create a goroutine that prints "Hello from a goroutine!" and ensure the main goroutine waits for it to finish.](./beginner/ex1/solve.go)  
2. [**Multiple Goroutines:** Launch 5 goroutines, each printing its unique ID.](\./beginner/ex2/solve.go)
3. [**Basic Channel Communication:** Create two goroutines. One sends an integer to a channel, and the other receives and prints it.](\./beginner/ex3/solve.go)
4. [**Buffered Channel:** Use a buffered channel of capacity 3. Send 3 integers and then receive and print them.](\./beginner/ex4/solve.go)
5. [**Unbuffered vs. Buffered:** Explain the difference between unbuffered and buffered channels with a small code example.](\./beginner/ex5/solve.go)
6. [**WaitGroup Basics:** Use a `sync.WaitGroup` to wait for 3 goroutines to complete their work (e.g., printing a message).](./beginner/ex6/solve.go)
7. [**Simple Counter (Race Condition):** Implement a global counter incremented by multiple goroutines without any synchronization. Observe the race condition.](./beginner/ex7/solve.go)
8. [**Counter with Mutex:** Fix the simple counter problem using `sync.Mutex` to prevent race conditions.](./beginner/ex8/solve.go)
9. [**Concurrent Summation:** Calculate the sum of elements in an array concurrently by dividing the array into chunks and processing each chunk in a separate goroutine.](./beginner/ex9/solve.go)
10. [**Producer-Consumer (Unbuffered):** Implement a simple producer-consumer pattern using an unbuffered channel where one goroutine produces items and another consumes them.](./beginner/ex10/solve.go)
11. [**Producer-Consumer (Buffered):** Implement a producer-consumer pattern using a buffered channel.](./beginner/ex11/solve.go)
12. [**Fan-Out/Fan-In (Basic):** Create a fan-out pattern where a single producer sends data to multiple worker goroutines, and then collect results from these workers (fan-in) into a single channel.](./beginner/ex12/solve.go)
13. [**Select Statement (Basic):** Use a `select` statement to receive from two different channels.](./beginner/ex13/Solve.go)
14. [**Select with Default:** Add a `default` case to a `select` statement.](./beginner/ex14/Solve.go)
15. [**Goroutine Leak (Basic):** Create a goroutine that runs indefinitely without a way to stop it. Identify why this is a leak.](./beginner/ex15/solve.go)
16. [**Stopping a Goroutine (Channel Signal):** Stop a goroutine using a `done` channel.](./beginner/ex16/solve.go)
17. [**Goroutine ID (Tricky):** How would you assign unique IDs to goroutines? (Hint: think about closure variables or a shared counter with mutex).](./beginner/ex18/solve.go)
18. [**Timeouts with Select:** Implement a `select` statement that times out after a certain duration if no message is received.](./beginner/ex18/solve.go)
19. [**Worker Pool (Fixed Size):** Create a worker pool with a fixed number of worker goroutines processing jobs from a job channel.](./beginner/ex19/solve.go)
20. **Error Handling in Goroutines:** How would you handle errors that occur within a goroutine and communicate them back to the main goroutine?
21. [**Ping-Pong Game (Two Goroutines):** Implement a classic ping-pong game where two goroutines send messages back and forth using channels.](./beginner/ex21/solve.go)
22. [**Concurrent File Read:** Read multiple small text files concurrently and print their contents.](./beginner/ex22/solve.go)
23. **Simple Web Scraper (Concurrent):** Fetch the content of a few URLs concurrently. (No parsing needed, just fetching).
24. [**Atomic Operations (Basic):** Use `sync/atomic` to atomically increment a counter. Compare with mutex.](./beginner/ex24/solve.go)
25. [**Broadcast Message:** Send a single message from one goroutine to multiple listening goroutines.](./beginner/ex25/solve.go)
26. [**Context with Timeout:** Use `context.WithTimeout` to cancel an operation if it takes too long.](./beginner/ex26/solve.go)
27. [**Context with Cancel:** Use `context.WithCancel` to explicitly cancel a running goroutine.](./beginner/ex27/solve.go)
28. **Concurrent Map Read/Write (Race):** Demonstrate a race condition when reading from and writing to a `map` concurrently without protection.
29. **Concurrent Map Read/Write (Mutex):** Protect a `map` with a `sync.RWMutex` for concurrent read and write operations.
30. **Fan-Out/Fan-In (Counting):** Fan out goroutines to count words in different parts of a text, then fan in to get the total word count.

---

### Intermediate Level (Problems 31-70)

These problems introduce more complex scenarios, common concurrency patterns, and error handling.

31. **Dining Philosophers (Simplified):** Implement a simplified version of the Dining Philosophers problem (e.g., 2 philosophers, 2 chopsticks).
32. **Bounded Buffer:** Implement a bounded buffer using channels (e.g., a fixed-size queue for producers and consumers).
33. **Rate Limiter (Token Bucket):** Implement a simple rate limiter using a channel that acts as a token bucket.
34. **Pipeline (Simple):** Create a pipeline of three stages where each stage processes data from the previous stage and sends it to the next.
35. **Graceful Shutdown (Workers):** Implement a set of worker goroutines that can be gracefully shut down when a signal is received (e.g., via a context or a dedicated stop channel).
36. **Error Group:** Use `errgroup.Group` (from `golang.org/x/sync/errgroup`) to run multiple goroutines concurrently and collect the first error that occurs.
37. **Multiplexing Channels:** Combine outputs from multiple producer channels into a single consumer channel using `select`.
38. **Circuit Breaker (Simplified):** Implement a very basic circuit breaker pattern where a function call fails fast if too many consecutive errors occur.
39. **Debouncing Concurrent Events:** Implement a debouncing mechanism for concurrent events, so that rapid-fire events only trigger a single action after a short delay.
40. **Bulk Data Processing:** Process a large dataset in chunks concurrently, saving results to a shared data structure.
41. **Concurrent File Search:** Search for a specific keyword across multiple files in a directory concurrently.
42. **Distributed Counter (Mock):** Simulate a distributed counter where multiple "nodes" (goroutines) increment a shared logical counter, requiring synchronization.
43. **Producer-Consumer with Backpressure:** Implement a producer-consumer where the producer slows down if the consumer can't keep up (using buffered channels and potentially `select`).
44. **Future/Promise Equivalent:** Implement a simplified "future" or "promise" pattern using channels, where a result is computed asynchronously and can be retrieved later.
45. **Timeout on Multiple Operations:** Use `select` to time out if any of several concurrent operations don't complete within a given duration.
46. **Leader Election (Simplified):** Simulate a simplified leader election scenario among a fixed set of goroutines using channels for communication.
47. **Concurrent Cache:** Implement a simple in-memory cache that allows concurrent reads and writes using `sync.RWMutex`.
48. **Token-Based Access Control:** Control access to a limited number of resources using a channel as a semaphore.
49. **Concurrent Mergesort (Conceptual):** Explain how you would implement a concurrent mergesort algorithm using goroutines. (No full implementation needed, focus on the concurrency aspect).
50. **Streaming Data Processing:** Read a continuous stream of data (simulated with a channel) and process it concurrently using a worker pool.
51. **Concurrent Web Server (Basic):** Create a simple HTTP server that handles requests concurrently using goroutines.
52. **Idempotent Operation Execution:** Design a system where a specific operation, when called concurrently, only executes once. (Hint: `sync.Once`).
53. **Sequential Execution of Concurrent Tasks:** How would you ensure that a set of tasks that are launched concurrently actually complete in a specific sequential order?
54. **Concurrent Download Manager (Basic):** Download multiple files concurrently, limiting the number of simultaneous downloads.
55. **Pipeline with Error Handling:** Enhance the simple pipeline to propagate errors gracefully through the stages.
56. **Fan-Out/Fan-In with Aggregation:** Fan out processing to multiple goroutines, then aggregate all their results (e.g., finding the maximum value from a set of concurrently processed numbers).
57. **Custom Sync Primitive:** Implement a basic semaphore using channels.
58. **Cancellation Hierarchy:** Design a system where a parent context can cancel child contexts, allowing for hierarchical cancellation of concurrent operations.
59. **Concurrent Log Processing:** Read a large log file line by line, process each line concurrently (e.g., categorize log levels), and then output summary statistics.
60. **Concurrent Graph Traversal (BFS/DFS):** Implement a concurrent Breadth-First Search (BFS) or Depth-First Search (DFS) on a simple graph data structure.
61. **Deadlock Detection (Manual):** Create a simple scenario that leads to a deadlock and explain why it occurs.
62. **Avoiding Starvation:** Design a scenario where one goroutine might starve for resources and suggest how to avoid it.
63. **Concurrent Map with Mutex (Generic):** Create a generic concurrent map that works for any key-value pair, protected by a mutex.
64. **Dynamic Worker Pool:** Implement a worker pool where the number of active workers can be adjusted dynamically based on load.
65. **Throttling Concurrent Requests:** Implement a mechanism to throttle the number of concurrent requests to an external API.
66. **Event Bus (Simple):** Implement a basic in-memory event bus using channels for publishing and subscribing to events.
67. **Concurrent Data Validation:** Validate multiple pieces of data concurrently, collecting all validation errors.
68. **Retry Mechanism:** Implement a concurrent function that retries a failed operation a certain number of times with an exponential backoff.
69. **Concurrent Message Broadcasting (Buffered):** Broadcast messages to multiple consumers, using buffered channels to prevent producers from blocking too often.
70. **Weighted Semaphore:** Implement a weighted semaphore where different tasks require different "weights" or "tokens" to execute concurrently.

---

### Advanced Level (Problems 71-100)

These problems delve into more intricate concurrency patterns, performance considerations, and real-world scenarios.

71. **Concurrent LRU Cache:** Implement a thread-safe Least Recently Used (LRU) cache.
72. **Distributed Locking (Conceptual):** Discuss how you would implement a distributed lock using Go's concurrency primitives, considering a simplified distributed system.
73. **Concurrent Database Migrations (Mock):** Simulate running database migrations concurrently, ensuring proper ordering and atomicity.
74. **Goroutine Pool with Context:** Create a goroutine pool where each worker receives a `context.Context` to handle cancellation and timeouts for individual tasks.
75. **Pipeline with Fan-In and Fan-Out:** Design a complex pipeline that incorporates both fan-in and fan-out stages.
76. **Concurrent Task Scheduling:** Implement a basic task scheduler that executes tasks concurrently based on their priority or dependencies.
77. **Concurrent Image Processing Pipeline:** Create a pipeline for image processing (e.g., resize, grayscale, watermark) where each step is a concurrent stage.
78. **Consistent Hashing (Conceptual):** Discuss how consistent hashing could be implemented in Go for distributing work among a dynamic set of concurrent workers.
79. **Non-Blocking Channel Operations:** Use `select` with `default` to perform non-blocking sends and receives on channels.
80. **Ordered Concurrent Processing:** Process items concurrently but ensure their output is in the original input order. (Hint: use a result channel with indices).
81. **Self-Healing Worker Pool:** Implement a worker pool that automatically restarts failed worker goroutines.
82. **Shared Memory with Atomic Pointers:** Use `atomic.Value` to safely share an immutable data structure between goroutines.
83. **Concurrent File System Walker:** Walk a directory tree concurrently, processing files as they are discovered.
84. **Bi-directional Channel Communication:** Design a scenario where two goroutines communicate bidirectionally using separate channels.
85. **Concurrent Benchmarking Tool:** Create a simple tool that runs a function concurrently N times and measures its average execution time.
86. **Distributed Pub/Sub (Mock):** Simulate a simple distributed publish-subscribe system where producers publish messages to topics, and consumers subscribe to them.
87. **Concurrent DNS Resolver (Simplified):** Implement a simplified concurrent DNS resolver that takes a list of hostnames and resolves them in parallel.
88. **Error Aggregation in Pipelines:** In a multi-stage pipeline, how would you collect all errors that occurred in any stage before signaling completion?
89. **Generics with Concurrency (Go 1.18+):** Reimplement a generic concurrent data structure (e.g., a concurrent queue or stack) using Go generics.
90. **Backpressure with `context.WithDeadline`:** Use `context.WithDeadline` to implement backpressure in a producer-consumer scenario, where tasks are dropped if they can't be processed by a certain time.
91. **Concurrent Stream Processing with Checkpointing (Conceptual):** Discuss how you would handle fault tolerance and checkpointing in a long-running concurrent data stream processing application.
92. **Deadlock Prevention Strategies:** Describe and implement common deadlock prevention strategies (e.g., resource ordering, timeout, avoidance).
93. **Concurrent Load Testing Tool (Basic):** Build a basic load testing tool that can send a configurable number of concurrent requests to a target URL.
94. **Thundering Herd Problem (Conceptual):** Explain the thundering herd problem in a concurrent system and how to mitigate it using Go concurrency patterns.
95. **Concurrent Database Connection Pool (Conceptual):** Design a simple concurrent database connection pool.
96. **Quorum-Based Operations (Mock):** Simulate a quorum-based system where an operation requires a certain number of acknowledgements from concurrent participants before proceeding.
97. **Context for Request-Scoped Data:** Pass request-scoped data through a chain of concurrent operations using `context.WithValue`.
98. **Understanding the Go Scheduler:** Design a program that helps visualize or understand the Go scheduler's behavior with many goroutines.
99. **Memory Model Implications:** Create a scenario where understanding Go's memory model is crucial to avoid subtle bugs when sharing data concurrently.
100. **Design a Concurrent Key-Value Store (In-Memory):** Design and partially implement an in-memory key-value store that supports concurrent read and write operations, with considerations for consistency and performance.


