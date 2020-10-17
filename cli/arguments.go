/**
Package that contains some utils for
ease the argument handling with the
CLI
*/
package cli

type Error struct {
	msg string
}

func (err Error) Error() string {
	return err.msg
}

/**
Represents an Argument Iterator,
iterates arguments and marks used
arguments as "consumed", they're
can be used again if the Back()
function is called specifying the
new nextIndex
*/
type ArgIterator struct {
	parentPtr *ArgIterator
	backing   []string
	nextIndex uint8
	HasNext   func() bool
	Back      func(uint8)
}

func (argItr ArgIterator) Next() (*string, error) {
	if !argItr.HasNext() {
		return nil, Error{"There're no more arguments in the iterator!"}
	}
	value := argItr.backing[argItr.nextIndex]
	argItr.nextIndex = argItr.nextIndex + 1
	return &value, nil
}
