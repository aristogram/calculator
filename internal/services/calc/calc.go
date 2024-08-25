package calc

import (
	"context"
	"fmt"
	calcv1 "github.com/aristogram/protos/gen/go/calculator"
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

const (
	add  = "+"
	sub  = "-"
	mul  = "*"
	div  = "/"
	pow  = "^"
	root = "rt"
	log  = "log"
)

// Precedence of the operators (+, -, *, /, ...)
type priority int

var operators = map[string]priority{
	add:  1,
	sub:  1,
	mul:  2,
	div:  2,
	pow:  3,
	root: 3,
	log:  4,
}

type Calculator struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Calculator {
	return &Calculator{
		log: log,
	}
}

func (c *Calculator) Calculate(
	ctx context.Context,
	expr *calcv1.ExpressionRequest,
) (string, error) {
	const op = "calc.Calculate"

	// TODO: validate jwt token

	c.log.With(
		slog.String("op", op),
		slog.String("expr", expr.GetExpr()),
	).Info("calculating expression")

	// TODO: create converter to "normal" format for newbies that give wrong expression's format.
	parts := exprDivider(expr.GetExpr())
	if err := validateParts(parts); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	rpnStack := infixToRPN(parts)
	return calcRPN(rpnStack)
}

// exprDivider divides expression into parts by whitespace.
func exprDivider(expr string) []string {
	return strings.Fields(expr)
}

// validateParts validates all expression's parts.
// Returns error if expression's format is wrong.
// Check README to get normal expression's format.
func validateParts(parts []string) error {
	for _, part := range parts {
		isNumber := unicode.IsDigit(rune(part[0]))

		for _, r := range part {
			if isNumber {
				if !unicode.IsDigit(r) {
					return fmt.Errorf("invalid number: %q", part)
				}
			} else {
				if unicode.IsDigit(r) {
					return fmt.Errorf("invalid operator: %q", part)
				}
			}
		}
	}

	return nil
}

// infixToRPN converts infix notation (human-readable) to reverse polish notation (i.e. "2 3 * 5" + == "2 * 3 + 5")
func infixToRPN(parts []string) []string {
	resStack := make([]string, 0, len(parts))
	opsStack := make([]string, 0, len(parts))

	for _, part := range parts {
		rn := rune(part[0])

		if unicode.IsDigit(rn) {
			resStack = append(resStack, part)

			continue
		} else if rn == '(' {
			opsStack = append(opsStack, part)

			continue
		} else if rn == ')' {
			for i := len(opsStack) - 1; i >= 0; i-- {
				if opsStack[i] != "(" {
					resStack = append(resStack, opsStack[i])
					opsStack = opsStack[:i]
				} else {
					// Removes '('
					opsStack = opsStack[:i]
					break
				}
			}

			continue
		}

		for {
			if len(opsStack) == 0 {
				break
			}

			prevOp := opsStack[len(opsStack)-1]

			if operators[prevOp] >= operators[part] {
				resStack = append(resStack, prevOp)

				opsStack = opsStack[:len(opsStack)-1]
			} else {
				break
			}
		}

		opsStack = append(opsStack, part)
	}

	slices.Reverse(opsStack)
	resStack = append(resStack, opsStack...)

	return resStack
}

func calcRPN(queue []string) (string, error) {
	resStack := make([]string, 0, len(queue))

	for _, op := range queue {
		if unicode.IsDigit(rune(op[0])) {
			resStack = append(resStack, op)

			continue
		}

		if len(resStack) < 2 {
			return "", fmt.Errorf("invalid syntax, operator %v should have 2 operands", op)
		}

		op2, err := strconv.ParseFloat(resStack[len(resStack)-1], 64)

		if err != nil {
			return "", err
		}

		op1, err := strconv.ParseFloat(resStack[len(resStack)-2], 64)

		if err != nil {
			return "", err
		}

		resStack = resStack[:len(resStack)-2]

		switch op {
		case add:
			resStack = append(resStack, fmt.Sprintf("%f", op1+op2))
		case sub:
			resStack = append(resStack, fmt.Sprintf("%f", op1-op2))
		case mul:
			resStack = append(resStack, fmt.Sprintf("%f", op1*op2))
		case div:
			resStack = append(resStack, fmt.Sprintf("%f", op1/op2))
		case pow:
			resStack = append(resStack, fmt.Sprintf("%f", math.Pow(op1, op2)))
		case root:
			resStack = append(resStack, fmt.Sprintf("%f", math.Pow(op2, 1/op1)))
		case log:
			// log(a)b = log(2)b/log(2)a
			resStack = append(resStack, fmt.Sprintf("%f", math.Log(op2)/math.Log(op1)))
		default:
			return "", fmt.Errorf("unknown operator: %v", op)
		}
	}

	if len(resStack) > 1 || len(resStack) == 0 {
		return "", fmt.Errorf("invalid expression")
	}

	return resStack[0], nil
}
