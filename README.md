# Data Structures in Go

This project is a collection of data structures implemented in the Go programming language. Data structures are essential for efficient data manipulation and storage, and this repository serves as a resource for learning and using data structures in Go.

## Table of Contents

- [Data Structures in Go](#data-structures-in-go)
- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Data Structures](#data-structures)
- [Usage](#usage)

## Introduction

Data structures are the building blocks of any software application. They provide a way to organize and store data efficiently, which is crucial for optimizing algorithms and solving complex problems. This project aims to provide well-documented implementations of various data structures in Go, making it easier for developers to leverage these structures in their projects.

## Data Structures

The following data structures are currently implemented in this project:

- [x] Linked List
- [x] Doubly Linked List
- [x] Circular Linked List
- [x] Doubly Circular Linked List
- [x] Stack
- [x] Queue
- [x] ... more coming soon!

Each data structure is located in its respective directory within the project. You can explore each data structure's implementation and documentation to understand how to use them effectively.

## Usage

To use any of the data structures in your Go project, you can follow these general steps:

1. Clone this repository or download the specific data structure you need.
2. Include the data structure package in your Go project.
3. Import the data structure package in your Go code.
4. Use the data structure as needed by following the provided documentation and examples.

For example, to use a linked list in your project:

```go
import "github.com/yourusername/data-structures-in-go/linkedlist"

func main() {
    // Create a new linked list
    list := linkedlist.New()

    // Add elements to the list
    list.Push(1)
    list.Push(2)
    list.Push(3)

    // Perform operations on the linked list
    // ...

    // Print the list
    list.Print()
}