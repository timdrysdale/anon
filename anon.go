package anon

import (
	"errors"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type IdentityDictionary struct {
	Anonymous string `csv:"anonymous"`
	Identity  string `csv:"identity"`
}

type Anonymiser struct {
	identityToAnonymous map[string]string
	anonymousToIdentity map[string]string
}

func New(inputPath string) (*Anonymiser, error) {

	a := Anonymiser{}
	ID := []*IdentityDictionary{}

	f, err := os.Open(inputPath)
	if err != nil {
		return &a, errors.New(fmt.Sprintf("Can't open identity dictionary %s", inputPath))
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &ID); err != nil { // Load subs
		return &a, errors.New(fmt.Sprintf("Can't unmarshall from %s", inputPath))
	}

	a.anonymousToIdentity = make(map[string]string)
	a.identityToAnonymous = make(map[string]string)

	for _, id := range ID {
		a.anonymousToIdentity[id.Anonymous] = id.Identity
		a.identityToAnonymous[id.Identity] = id.Anonymous
	}

	return &a, nil
}

func (a *Anonymiser) GetLength() int {
	return len(a.identityToAnonymous)
}

func (a *Anonymiser) GetAnonymous(identityKey string) (string, error) {
	if anon, ok := a.identityToAnonymous[identityKey]; ok {
		return anon, nil
	} else {
		return "", errors.New("Unknown anonymous key")
	}
}

func (a *Anonymiser) GetIdentity(anonymousKey string) (string, error) {
	if id, ok := a.anonymousToIdentity[anonymousKey]; ok {
		return id, nil
	} else {
		return "", errors.New("Unknown anonymous key")
	}
}
