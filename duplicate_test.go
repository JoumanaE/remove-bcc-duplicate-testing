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

func TestBcc(t *testing.T) {
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
	}

	for _, tabletest := range tabletests {
		output := RemoveDuplicates(tabletest.email, tabletest.bcc)
		if len(output) != tabletest.expected {
			t.Errorf("I got %v, and I expected %v ", len(output), tabletest.expected)

		}

	}

}

// Bcc verifies that the recipient of the email is not duplicated in the bcc, & if it is, deletes it.xs
func RemoveDuplicates(toEmail string, bcc []string) []string {

	for i, email := range bcc {
		if toEmail == email {
			fmt.Printf("The recipient's email,%v, is duplicated.", toEmail)
			bcc[i] = bcc[len(bcc)-1]
			bcc[len(bcc)-1] = ""
			bcc = bcc[:len(bcc)-1]
			fmt.Printf("Here's the new bcc list: %v\n", bcc)

		}

	}

	return bcc
}
