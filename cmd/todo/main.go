package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/voukatas/todo-app-go"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo completed")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")
	update := flag.Int("update", 0, "update a todo by its index")

	flag.Parse()

	flagCount := 0

	if *add {
		flagCount++
	}
	if *complete > 0 {
		flagCount++
	}
	if *del > 0 {
		flagCount++
	}
	if *list {
		flagCount++
	}
	if *update > 0 {
		flagCount++
	}

	if flagCount == 0 {
		fmt.Println("Usage: todo-app-go [OPTIONS]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if flagCount > 1 {
		fmt.Println("Error: Only one flag can be provided at a time.")
		os.Exit(1)
	}

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Printf("error: %v\n", err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}

		todos.Add(task)
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}
		fmt.Println("TODO added successfully!")

	case *complete > 0:
		err := todos.Complete(*complete)

		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}

		err = todos.Store(todoFile)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("TODO #%d marked as completed!\n", *complete)

	case *del > 0:
		err := todos.Delete(*del)

		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}

		err = todos.Store(todoFile)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("TODO #%d deleted successfully!\n", *del)

	case *list:
		todos.Print()

	case *update > 0:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}

		// Assuming the todo package has an Update method
		err = todos.Update(*update, task)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("TODO #%d updated successfully!\n", *update)

	default:
		fmt.Println("Invalid command. Use -h for help.")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	fmt.Print("Enter task: ")

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil
}
