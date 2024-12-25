package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pkg/errors"
)

func foo() string {
	var s *string
	return *s
}

func handlePanic(r interface{}) error {
	var errWithStack error
	if err, ok := r.(error); ok {
		errWithStack = errors.WithStack(err)
	} else {
		errWithStack = errors.Errorf("%+v", r)
	}
	return errWithStack
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	defer func() {
		if r := recover(); r != nil {
			err := handlePanic(r)
			logger.Error(
				"panic occurred",
				"msg", err.Error(),
				"stack", fmt.Sprintf("%+v", err),
			)
		}
	}()

	fmt.Println(foo())
}
