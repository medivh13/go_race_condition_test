# Go Race Condition Test
This project demonstrates scenarios of race condition and synchronization using Mutex in Go.

# Prerequisites
Make sure you have Go installed on your machine. You can download it from golang.org.

- Go installed on your system

### Running the Project

1. **Clone the repository:**

2. **Navigate to the project directory:**

    ```sh
    cd go_race_condition_test
    and then do go mod tidy

    ```

3. **Run the server:**

    ```sh
    go run main.go
    ```

The server will run on port 8080.

## Choose the operation type:

1. Without Mutex (Race Condition): Sends multiple simultaneous requests to update stock (a total of 10)without synchronization.
2. With Mutex: Sends multiple simultaneous requests to update stock (a total of 10) with synchronization using Mutex.
Follow the prompts and observe the output to see how race conditions can occur and how Mutex helps to synchronize concurrent access.

Example
To illustrate the usage:

Select option 1 to demonstrate race condition without Mutex.
Select option 2 to demonstrate synchronization using Mutex.