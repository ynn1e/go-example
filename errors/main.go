package main

import (
	"errors"
	"fmt"
)

type myError error

func main() {
	// go version go1.20.5 darwin/arm64

	var (
		myerr myError = errors.New("my error")
		err           = errors.New("error")
	)

	fmt.Printf("myerr: %s %T\n", myerr, myerr)
	fmt.Printf("err: %s %T\n", err, err)
	fmt.Println()

	/*
		myerr: my error *errors.errorString
		err: error *errors.errorString
	*/

	var (
		errNoWrap       = fmt.Errorf("%s, %s", myerr, err)
		errOneWrap      = fmt.Errorf("%w, %s", myerr, err)
		errTwoWraps     = fmt.Errorf("%w, %w", myerr, err)
		errHierarchyMsg = fmt.Errorf("%w: %s", myerr, fmt.Errorf("%w", err))
		errHierarchyErr = fmt.Errorf("%w: %s", fmt.Errorf("%w", myerr), err)
		errJoin         = fmt.Errorf("%w", errors.Join(myerr, err))
	)

	/*
		type error interface {
			Error() string
		}
		type errorString struct {
			s string
		}
		type wrapError struct {
			msg string
			err error
		}
		type wrapErrors struct {
			msg  string
			errs []error
		}
	*/

	// Wrap
	fmt.Printf("errNoWrap: %s %T\n", errNoWrap, errNoWrap)
	fmt.Printf("errOneWrap: %s %T\n", errOneWrap, errOneWrap)
	fmt.Printf("errTwoWraps: %s %T\n", errTwoWraps, errTwoWraps)
	fmt.Printf("errHierarchyMsg: %s %T\n", errHierarchyMsg, errHierarchyMsg)
	fmt.Printf("errHierarchyErr: %s %T\n", errHierarchyErr, errHierarchyErr)
	fmt.Printf("errJoin: %s %T\n", errJoin, errJoin)
	fmt.Println()

	/*
		errNoWrap: my error, error *errors.errorString
		errOneWrap: my error, error *fmt.wrapError
		errTwoWraps: my error, error *fmt.wrapErrors
		errHierarchyMsg: my error: error *fmt.wrapError
		errHierarchyErr: my error: error *fmt.wrapError
		errJoin: my error
		error *fmt.wrapError
	*/

	// Unwrap
	fmt.Printf("errNoWrap: %v\n", errors.Unwrap(errNoWrap))
	fmt.Printf("errOneWrap: %v\n", errors.Unwrap(errOneWrap))
	fmt.Printf("errTwoWraps: %v\n", errors.Unwrap(errTwoWraps))
	fmt.Printf("errHierarchyMsg.1: %v\n", errors.Unwrap(errHierarchyMsg))
	fmt.Printf("errHierarchyMsg.2: %v\n", errors.Unwrap(errors.Unwrap(errHierarchyMsg)))
	fmt.Printf("errHierarchyErr.1: %v\n", errors.Unwrap(errHierarchyErr))
	fmt.Printf("errHierarchyErr.2: %v\n", errors.Unwrap(errors.Unwrap(errHierarchyErr)))
	fmt.Printf("errHierarchyErr.3: %v\n", errors.Unwrap(errors.Unwrap(errors.Unwrap(errHierarchyErr))))
	fmt.Printf("errJoin: %v\n", errors.Unwrap(errJoin))
	fmt.Println()

	/*
		errNoWrap: <nil>
		errOneWrap: my error
		errTwoWraps: <nil>: because wrapErrors type does not have 'Unwrap() error' method
		errHierarchyMsg.1: my error
		errHierarchyMsg.2: <nil>
		errHierarchyErr.1: my error
		errHierarchyErr.2: my error
		errHierarchyErr.3: <nil>
		errJoin: my error
		error
	*/

	// Is
	fmt.Printf("errNoWrap: %t\n", errors.Is(myerr, errors.Unwrap(errNoWrap)))
	fmt.Printf("errOneWrap: %t\n", errors.Is(myerr, errors.Unwrap(errOneWrap)))
	fmt.Printf("errTwoWraps: %t\n", errors.Is(myerr, errors.Unwrap(errTwoWraps)))
	fmt.Printf("errHierarchyMsg.1: %t\n", errors.Is(myerr, errors.Unwrap(errHierarchyMsg)))
	fmt.Printf("errHierarchyMsg.2: %t\n", errors.Is(myerr, errors.Unwrap(errors.Unwrap(errHierarchyMsg))))
	fmt.Printf("errHierarchyErr.1: %t\n", errors.Is(myerr, errors.Unwrap(errHierarchyErr)))
	fmt.Printf("errHierarchyErr.2: %t\n", errors.Is(myerr, errors.Unwrap(errors.Unwrap(errHierarchyErr))))
	fmt.Printf("errHierarchyErr.3: %t\n", errors.Is(myerr, errors.Unwrap(errors.Unwrap(errors.Unwrap(errHierarchyErr)))))
	fmt.Printf("errJoin: %t\n", errors.Is(myerr, errors.Unwrap(errJoin)))
	fmt.Println()

	/*
		errNoWrap: false
		errOneWrap: true
		errTwoWraps: false
		errHierarchyMsg.1: true
		errHierarchyMsg.2: false
		errHierarchyErr.1: false
		errHierarchyErr.2: true
		errHierarchyErr.3: false
		errJoin: false
	*/

	// As
	var myErr myError
	fmt.Printf("errNoWrap: %t\n", errors.As(errNoWrap, &myErr))
	fmt.Printf("errOneWrap: %t\n", errors.As(errOneWrap, &myErr))
	fmt.Printf("errTwoWraps: %t\n", errors.As(errTwoWraps, &myErr))
	fmt.Printf("errHierarchyMsg: %t\n", errors.As(errHierarchyMsg, &myErr))
	fmt.Printf("errHierarchyErr: %t\n", errors.As(errHierarchyErr, &myErr))
	fmt.Printf("errJoin: %t\n", errors.As(errJoin, &myErr))
	fmt.Println()

	/*
		errNoWrap: true
		errOneWrap: true
		errTwoWraps: true
		errHierarchyMsg: true
		errHierarchyErr: true
		errJoin: true
	*/
}
