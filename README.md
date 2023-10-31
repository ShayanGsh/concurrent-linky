# Thread-Safe Ordered Doubly Linked List in Go

This repository contains an implementation of a thread-safe ordered (sorted) doubly linked list in Go. The linked list is designed to be safe for concurrent use, meaning multiple goroutines can interact with the list simultaneously without causing data races or inconsistencies.

## Overview

The linked list is a fundamental data structure in computer science, often used for its simplicity and efficiency in certain operations. An ordered linked list is a variation where the elements are kept sorted at all times. This implementation uses Go's powerful concurrency primitives to ensure that the list remains safe to use in a multi-threaded context.

## Features

- **Thread-Safe**: The linked list is safe for concurrent use. It uses Go's `sync` package to ensure that multiple goroutines can safely add elements to the list simultaneously.
- **Ordered**: The linked list keeps its elements sorted at all times. This is achieved by comparing elements during insertion and placing each new element at the correct position to maintain the order.
<!-- - **Efficient**: The linked list is implemented in a way that minimizes unnecessary work. For example, it uses a worker pool pattern to distribute the work of adding elements across multiple goroutines. TO BE ADDED -->

## Usage

The main function of the program initializes a large pool of random numbers, creates a new linked list, and spawns a number of worker goroutines. Each worker receives values from a channel and inserts them into the linked list. Once all values have been inserted, the main function prints the contents of the list. It is a simple example of how the linked list can be used in a multi-threaded context.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.
