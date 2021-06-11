package gtfs

import "fmt"

type MissingStructError struct {
	Container string
	Id        string
}

func (e *MissingStructError) Error() string {
	return fmt.Sprintf("missing struct in %s with id %s", e.Container, e.Id)
}

type RequirementError struct {
	Struct string
	Field  string
}

func (e *RequirementError) Error() string {
	return fmt.Sprintf("requirement error: %s.%s", e.Struct, e.Field)
}

type ConditionnalRequirementError struct {
	Struct    string
	Field     string
	Condition string
}

func (e *ConditionnalRequirementError) Error() string {
	return fmt.Sprintf("conditional requirement error %s.%s: %s", e.Struct, e.Field, e.Condition)
}

type ReferenceError struct {
	Struct      string
	Field       string
	Destination string
	Value       string
}

func (e *ReferenceError) Error() string {
	return fmt.Sprintf("reference error in %s.%s to %s: %s", e.Struct, e.Field, e.Destination, e.Value)
}

type InvalidValueError struct {
	Struct string
	Field  string
	Value  interface{}
}

func (e *InvalidValueError) Error() string {
	return fmt.Sprintf("invalid value error in %s.%s: %v", e.Struct, e.Field, e.Value)
}
