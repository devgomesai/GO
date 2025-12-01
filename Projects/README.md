# Go Learning Journey & Project Ideas

## Current Progress Summary
Based on the code in the `cmd/` directory, you have covered the following core Go concepts:

### 1. Basics (`cmd/basics`)
- **Variables & Types:** Integers, floats, strings, and type conversions.
- **Functions:** Multiple return values and named return parameters.
- **Control Flow:** `if-else` statements.
- **Error Handling:** The standard `error` pattern (`value, err := func()`).
- **String Manipulation:** Handling UTF-8 characters and runes.

### 2. Collections & Data Structures (`cmd/collections`, `cmd/structs`)
- **Structs:** Defining custom data types (`gasEngine`, `electricEngine`).
- **Methods:** Attaching functions to structs (receivers).
- **Interfaces:** Defining behavior contracts (`engine` interface) and polymorphism.
- **Slices:** Working with dynamic arrays.

### 3. Concurrency (`cmd/goroutines`, `cmd/channels`)
- **Goroutines:** Launching lightweight threads with `go`.
- **WaitGroups:** Synchronizing goroutines using `sync.WaitGroup`.
- **Mutexes:** Handling race conditions with `sync.Mutex`.
- **Channels:** Communicating between goroutines (buffered/unbuffered) and sending/receiving data.

### 4. Advanced Features (`cmd/generics`, `cmd/pointers`, `cmd/strings`)
- **Generics:** Writing reusable functions with type constraints (e.g., `[T int | float32]`).
- **Pointers:** (Assumed) Managing memory addresses and passing by reference.

---

## Recommended Projects
Here are project ideas ranging from beginner to advanced to help you solidify these concepts.

### 🟢 Beginner (Solidifying Basics)

#### 1. CLI Task Manager
**Concepts:** Structs, Slices, Functions, User Input.
- Create a CLI tool to Add, List, and Delete tasks.
- Use a `Task` struct `{ID, Title, Completed}`.
- Store tasks in a slice `[]Task`.

#### 2. Currency Converter
**Concepts:** Maps, Structs, Math.
- Store exchange rates in a `map[string]float64`.
- Create a function that takes an amount and target currency.
- Handle errors if the currency code doesn't exist.

#### 3. Shape Calculator
**Concepts:** Interfaces, Structs, Methods.
- Define a `Shape` interface with an `Area()` method.
- Create structs for `Rectangle`, `Circle`, and `Triangle`.
- Write a function `PrintArea(s Shape)` that works for any shape.

### 🟡 Intermediate (Concurrency & Logic)

#### 4. Concurrent URL Checker
**Concepts:** Goroutines, WaitGroups, HTTP (optional) or Time simulation.
- Create a list of website URLs (strings).
- Spin up a goroutine for each URL to check if it's "up" (simulate with `time.Sleep`).
- Use a `WaitGroup` to ensure the main function waits for all checks to finish.

#### 5. Safe Bank Account
**Concepts:** Structs, Mutexes, Methods.
- Create a `BankAccount` struct with a `balance` field.
- Launch multiple goroutines that simultaneously try to `Deposit` and `Withdraw` money.
- Use `sync.Mutex` to ensure the balance never becomes corrupted (race condition free).

#### 6. Generic Stack Data Structure
**Concepts:** Generics, Slices.
- Implement a `Stack[T]` struct.
- Methods: `Push(val T)`, `Pop() T`, `Peek() T`.
- Test it with both `int` and `string` types.

### 🔴 Advanced (Combining Everything)

#### 7. Worker Pool Simulation
**Concepts:** Channels, Goroutines, Select, Structs.
- Create a fixed number of "Worker" goroutines (e.g., 3 workers).
- Create a "Job" channel to send tasks.
- Workers read from the channel, process the job (random sleep), and output the result.
- Handle graceful shutdown (closing the channel).

#### 8. Real-time Chat Server (CLI simulation)
**Concepts:** Channels, Mutexes, Structs.
- Simulate a "Server" that manages a list of "Clients".
- When one client sends a message, it goes into a channel.
- The server broadcasts that message to all other connected clients.
- Use Mutexes to safely add/remove clients from the active list.

#### 9. Dining Philosophers Problem
**Concepts:** Deadlock prevention, Mutexes, Goroutines.
- A classic computer science problem involving concurrency control.
- implement 5 philosophers sitting at a round table with forks between them.
- Ensure no two philosophers hold the same fork at the same time and avoid deadlocks.
