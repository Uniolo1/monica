package main

import (
	"fmt"
	"os"

	"github.com/maja42/goval"
)

func evaluate_input(input string) string {
	if os.Getenv("DISABLE_CALC") != "" && os.Getenv("DISABLE_CALC") != "false" {
		return "The calculator has been disabled!"
	} else {
		eval := goval.NewEvaluator()

		result, err := eval.Evaluate(input, nil, nil)
		if err != nil {
			return fmt.Sprintf("Error: %s", err)
		} else {
			return fmt.Sprintf("%v", result)
		}
	}
}
