package mock

import (
	"bytes"
	"context"
	"time"

	"github.com/blink-io/x/session/store"
)

type expectedDelete struct {
	inputToken string
	returnErr  error
}

type expectedFind struct {
	inputToken  string
	returnB     []byte
	returnFound bool
	returnErr   error
}

type expectedCommit struct {
	inputToken  string
	inputB      []byte
	inputExpiry time.Time
	returnErr   error
}

type expectedAll struct {
	returnMB  map[string][]byte
	returnErr error
}

type istore struct {
	deleteExpectations []expectedDelete
	findExpectations   []expectedFind
	commitExpectations []expectedCommit
	allExpectations    []expectedAll
}

func New() interface {
	store.Store
	ExpectFind(string, []byte, bool, error)
	ExpectCommit(string, []byte, time.Time, error)
} {
	return newRaw()
}

func newRaw() *istore {
	return &istore{}
}

func (s *istore) ExpectDelete(token string, err error) {
	s.deleteExpectations = append(s.deleteExpectations, expectedDelete{
		inputToken: token,
		returnErr:  err,
	})
}

// Delete implements the Store interface
func (s *istore) Delete(ctx context.Context, token string) (err error) {
	var (
		indexToRemove    int
		expectationFound bool
	)
	for i, expectation := range s.deleteExpectations {
		if expectation.inputToken == token {
			indexToRemove = i
			expectationFound = true
			break
		}
	}
	if !expectationFound {
		panic("store.Delete called unexpectedly")
	}

	errToReturn := s.deleteExpectations[indexToRemove].returnErr
	s.deleteExpectations = s.deleteExpectations[:indexToRemove+copy(s.deleteExpectations[indexToRemove:], s.deleteExpectations[indexToRemove+1:])]

	return errToReturn
}

func (s *istore) ExpectFind(token string, b []byte, found bool, err error) {
	s.findExpectations = append(s.findExpectations, expectedFind{
		inputToken:  token,
		returnB:     b,
		returnFound: found,
		returnErr:   err,
	})
}

// Find implements the Store interface
func (s *istore) Find(ctx context.Context, token string) (b []byte, found bool, err error) {
	var (
		indexToRemove    int
		expectationFound bool
	)
	for i, expectation := range s.findExpectations {
		if expectation.inputToken == token {
			indexToRemove = i
			expectationFound = true
			break
		}
	}
	if !expectationFound {
		panic("store.Find called unexpectedly")
	}

	valueToReturn := s.findExpectations[indexToRemove]
	s.findExpectations = s.findExpectations[:indexToRemove+copy(s.findExpectations[indexToRemove:], s.findExpectations[indexToRemove+1:])]

	return valueToReturn.returnB, valueToReturn.returnFound, valueToReturn.returnErr
}

func (s *istore) ExpectCommit(token string, b []byte, expiry time.Time, err error) {
	s.commitExpectations = append(s.commitExpectations, expectedCommit{
		inputToken:  token,
		inputB:      b,
		inputExpiry: expiry,
		returnErr:   err,
	})
}

// Commit implements the Store interface
func (s *istore) Commit(ctx context.Context, token string, b []byte, expiry time.Time) (err error) {
	var (
		indexToRemove    int
		expectationFound bool
	)
	for i, expectation := range s.commitExpectations {
		if expectation.inputToken == token && bytes.Equal(expectation.inputB, b) && expectation.inputExpiry == expiry {
			indexToRemove = i
			expectationFound = true
			break
		}
	}
	if !expectationFound {
		panic("store.Commit called unexpectedly")
	}

	errToReturn := s.commitExpectations[indexToRemove].returnErr
	s.commitExpectations = s.commitExpectations[:indexToRemove+copy(s.commitExpectations[indexToRemove:], s.commitExpectations[indexToRemove+1:])]

	return errToReturn
}

func (s *istore) ExpectAll(mb map[string][]byte, err error) {
	s.allExpectations = append(s.allExpectations, expectedAll{
		returnMB:  mb,
		returnErr: err,
	})
}

func (s *istore) All(ctx context.Context) (map[string][]byte, error) {
	var (
		indexToRemove    int
		expectationFound bool
	)
	for i, expectation := range s.allExpectations {
		if len(expectation.returnMB) == 3 {
			indexToRemove = i
			expectationFound = true
			break
		}
	}
	if !expectationFound {
		panic("store.All called unexpectedly")
	}

	valueToReturn := s.allExpectations[indexToRemove]
	s.allExpectations = s.allExpectations[:indexToRemove+copy(s.allExpectations[indexToRemove:], s.allExpectations[indexToRemove+1:])]

	return valueToReturn.returnMB, valueToReturn.returnErr
}
