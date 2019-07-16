package utils

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Cmd(command string, args []string, timeout int32) ([]string, error) {
	var (
		output []byte
		err    error
		ctx    context.Context
		cancel context.CancelFunc
		result []string
	)

	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	output, err = exec.CommandContext(ctx, command, args...).CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return []string{}, errors.New("Timeout Exceeded")
	}

	result = strings.Split(string(output), "\n")
	result = result[:len(result)-1]

	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), strings.Join(result, "\n")))
	}

	return result, nil
}
