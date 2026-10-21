// SPDX-FileCopyrightText: NONE
// SPDX-License-Identifier: Unlicense

/*
 * This is free and unencumbered software released into the public domain.
 *
 * Anyone is free to copy, modify, publish, use, compile, sell, or distribute
 * this software, either in source code form or as a compiled binary, for any
 * purpose, commercial or non-commercial, and by any means.
 *
 * In jurisdictions that recognize copyright laws, the author or authors of this
 * software dedicate any and all copyright interest in the software to the
 * public domain. We make this dedication for the benefit of the public at large
 * and to the detriment of our heirs and successors. We intend this dedication
 * to be an overt act of relinquishment in perpetuity of all present and future
 * rights to this software under copyright law.
 *
 * THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package helpers

import (
	"fmt"
	"math"

	"github.com/maja42/goval"
)

func sqrt(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("sqrt expects 1 argument, got %d", len(args))
	}

	var x float64

	// apparently the only way to convert all arguments that this function might accept, is va a long ahh switch statement...
	switch v := args[0].(type) {
	case float64:
		x = v
	case float32:
		x = float64(v)
	case int:
		x = float64(v)
	case int8:
		x = float64(v)
	case int16:
		x = float64(v)
	case int32:
		x = float64(v)
	case int64:
		x = float64(v)
	case uint:
		x = float64(v)
	case uint8:
		x = float64(v)
	case uint16:
		x = float64(v)
	case uint32:
		x = float64(v)
	case uint64:
		x = float64(v)
	default:
		return nil, fmt.Errorf("sqrt expects a numeric argument, got %T", args[0])
	}

	return math.Sqrt(x), nil
}

var vars = map[string]interface{}{
	"pi":  math.Pi,
	"e":   math.E,
	"phi": math.Phi,
}

var functions = map[string]goval.ExpressionFunction{
	"sqrt": sqrt,
}

func EvaluatecCalculatorInput(input string) string {
	eval := goval.NewEvaluator()

	result, err := eval.Evaluate(input, vars, functions)
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	} else {
		return fmt.Sprintf("%v", result)
	}
}
