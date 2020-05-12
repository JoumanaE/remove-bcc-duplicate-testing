package main

import (
	"fmt"
	"testing"
)

type tt struct {
	bcc      []string
	email    string
	expected int
}

type tbcc struct {
	bcc      []string
	expected int
}

func TestDupeBCC(t *testing.T) {
	tabletests := []tbcc{
		{
			bcc:      []string{"joum@testig.com", "testing@gmail.com", "testing@gmail.com"},
			expected: 2,
		},
		{bcc: []string{"boop@gmail.com", "beep@gmail.com", "boop@gmail.com"},
			expected: 2,
		},
		{bcc: []string{"foop@gmail.com", "feep@gmail.com", "fafafa@gmail.com", "frfrfr@gmail.com"},
			expected: 4,
		},
		{bcc: []string{"runningoutofideas@gmail.com"},
			expected: 1,
		},
		{bcc: []string{},
			expected: 0,
		},
		{
			bcc:      []string{"1@gmail.com", "2@gmail.com", "3@gmail.com", "2@gmail.com", "4@gmail.com"},
			expected: 4,
		},
		{
			bcc:      []string{"4@gmail.com", "4@gmail.com", "3@gmail.com", "4@gmail.com"},
			expected: 2,
		},
	}

	for _, tabletest := range tabletests {
		output := dedupeBCC(tabletest.bcc)
		if len(output) != tabletest.expected {
			t.Errorf("I got %v, and I expected %v ", len(output), tabletest.expected)

		}

	}

}

// dedupeBCC erases the duplicate from the bcc list.
func dedupeBCC(bcc []string) []string {
	found := map[string]bool{}
	newBCC := []string{}

	for email := range bcc {
		if found[bcc[email]] == true {

		} else {
			found[bcc[email]] = true
			newBCC = append(newBCC, bcc[email])
		}
	}

	fmt.Printf("\n New BCC List ---> %v", newBCC)
	return newBCC

}
func TestRecipientNotInBCC(t *testing.T) {
	tabletests := []tt{
		{bcc: []string{"test@test.com"},
			email:    "testing@boring.com",
			expected: 1,
		},
		{
			bcc:      []string{"joum@testig.com", "testing@gmail.com"},
			email:    "testertester@gmail.com",
			expected: 2,
		},
		{bcc: []string{"boop@gmail.com", "beep@gmail.com", "lalala@gmail.com"},
			email:    "boop@gmail.com",
			expected: 2,
		},
		{bcc: []string{"foop@gmail.com", "feep@gmail.com", "fafafa@gmail.com", "frfrfr@gmail.com"},
			email:    "foop@gmail.com",
			expected: 3,
		},
		{bcc: []string{"runningoutofideas@gmail.com"},
			email:    "runningoutofideas@gmail.com",
			expected: 0,
		},
		{bcc: []string{},
			email:    "lasttest@gmail.com",
			expected: 0,
		},
		{
			bcc:      []string{"1@gmail.com", "2@gmail.com", "3@gmail.com", "2@gmail.com"},
			email:    "2@gmail.com",
			expected: 2,
		},
		{
			bcc:      []string{"4@gmail.com", "4@gmail.com", "3@gmail.com", "4@gmail.com"},
			email:    "4@gmail.com",
			expected: 1,
		},
	}

	for _, tabletest := range tabletests {
		output := ensureRecipientNotInBCC(tabletest.email, tabletest.bcc)
		if len(output) != tabletest.expected {
			t.Errorf("I got %v, and I expected %v ", len(output), tabletest.expected)

		}

	}

}

//ensureRecipientNotInBCC verifies that the recipient of the email is not duplicated in the bcc, & if it is, deletes it.xs
func ensureRecipientNotInBCC(toEmail string, bcc []string) []string {
	newBcc := []string{}
	for i, email := range bcc {
		if toEmail == email {
			fmt.Printf("\nThere is a duplicate email ---> %v", email)
		} else {
			newBcc = append(newBcc, bcc[i])
		}
	}
	fmt.Printf("\n New BCC List ---> %v", newBcc)
	return newBcc

}
