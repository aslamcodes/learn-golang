package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"todo"
)

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	s.Scan()

	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(), nil
}

func main() {
	add := flag.Bool("add", false, "Add tasks to list")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("delete", 0, "Item to be completed")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "built by @aslamcodes")
		flag.PrintDefaults()
	}

	flag.Parse()

	l := &todo.TodoList{}

	filename := ".todo.json"

	if envFilename := os.Getenv("TODO_FILENAME"); envFilename != "" {
		filename = envFilename
	}

	if err := l.Get(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {

	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *add:
		task, err := getTask(os.Stdin, flag.Args()...)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.Add(task)

		// repetive save code with error checking
		if err = l.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Print(l)
		os.Exit(0)
	}
}
