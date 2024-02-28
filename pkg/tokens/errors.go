package tokens

import "fmt"

type ActualFunctionCoverageBeneathMinimumError struct {
	minimum float64
	actual  float64
}

func (err *ActualFunctionCoverageBeneathMinimumError) Error() string {
	return fmt.Sprintf("actual coverage did not meet the minimum coverage bar (expected %0.2f and got %0.2f)", err.minimum, err.actual)
}

type MissingArgument struct{}

func (err *MissingArgument) Error() string {
	return "an argument for the command was expected"
}

type InvalidMinimumArgument struct {
	argument string
}

func (err *InvalidMinimumArgument) Error() string {
	return fmt.Sprintf("received argument %q where a well-formed float between 0 and 1 was expected", err.argument)
}

type InvalidPackageCommandError struct {
	command string
}

func (err *InvalidPackageCommandError) Error() string {
	return fmt.Sprintf("the command %q is invalid or cannot be used as a package directive", err.command)
}

type UnknownCommandError struct {
	command    string
	components []string
}

func (err *UnknownCommandError) Error() string {
	return fmt.Sprintf("the command %q is unknown (part of directive %v)", err.command, err.components)
}

type InvalidDefaultCommandError struct {
	argument Level
}

func (err *InvalidDefaultCommandError) Error() string {
	return fmt.Sprintf("the argument %q is not valid for the default command", err.argument)
}

type InvalidRegexError struct {
	reg string
	err error
}

func (err *InvalidRegexError) Error() string {
	return fmt.Sprintf("regex sequence %q is not valid: %s", err.reg, err.err.Error())
}
