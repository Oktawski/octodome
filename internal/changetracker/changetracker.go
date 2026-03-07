package changetracker

import (
	"fmt"
	"reflect"
)

type ChangeTracker struct {
	userID     uint
	HasChanges bool
	Changes    []Change
}

type Change struct {
	ModifiedBy string
	Field      string
	Old        any
	New        any
}

func NewChange(modifiedBy string, field string, old any, new any) Change {
	return Change{
		ModifiedBy: modifiedBy,
		Field:      field,
		Old:        old,
		New:        new,
	}
}

func New(userID uint) *ChangeTracker {
	return &ChangeTracker{
		userID:     userID,
		HasChanges: false,
		Changes:    []Change{},
	}
}

func UpdateWhenNotEqual[T any](
	ct *ChangeTracker,
	getter func() T,
	setter func(T),
	value T,
	fieldName string,
) {
	current := getter()
	if reflect.DeepEqual(current, value) {
		return
	}

	setter(value)

	ct.HasChanges = true
	ct.Changes = append(ct.Changes, NewChange(fmt.Sprint(ct.userID), fieldName, current, value))
}

func UpdateWhen[T any](
	ct *ChangeTracker,
	condition bool,
	getter func() T,
	setter func(T),
	value T,
	fieldName string,
) {
	if !condition {
		return
	}

	current := getter()
	if reflect.DeepEqual(current, value) {
		return
	}

	setter(value)

	ct.HasChanges = true
	ct.Changes = append(ct.Changes, NewChange(fmt.Sprint(ct.userID), fieldName, current, value))
}

func (ct *ChangeTracker) DoWhen(condition bool, fn func(), changes ...Change) {
	if !condition {
		return
	}

	fn()

	ct.HasChanges = true
	ct.Changes = append(ct.Changes, changes...)
}
