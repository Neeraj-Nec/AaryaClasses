package userM

import (
	"testing"
)

func TestValidate(t *testing.T) {
	testCase := []struct {
		password string
		expected bool
	}{
		{"Password123!", true},
		{"weak", false},
		{"NoSpecialChar123", false},
		{"Alllowercasewithspecial!", true},
		{"NOSPECIALCHARS123", false},
		{"", false},
		{"1234566", false},
		{"QWERDDS@#", false},
		{"WEDDcEE!@#", true},
		{"12345!", false},
		{"aaaaWWW123!@", true},
		{"1123A3b@!",true },
	}

	for _, test := range testCase {
		t.Run(test.password, func(t *testing.T) {
			result := validatePassword(test.password)
			if result != test.expected {
				t.Errorf("Expected validation result for password '%s' to be %t, but got %t", test.password, test.expected, result)
			}
		})
	}
}



func TestEmail(t *testing.T) {
	testCase := []struct {
		email string
		expected bool
	}{
		{"john.doe@example.com", true},
		{"alice.smith123@example.co.uk", true},
		{"bob@email-with-dash.com", true},
		{"info@company.net", true},
		{"user123@localhost", true},
		{"neeraj.srivastava@india.nec.com", true},
		{"neerajsrivastava7@gmail.com", true},
		{"not_an_email", false},
		{"missing_domain@.com", false},
		{"spaces in@email.com", false},
		{"double@@at@example.com", false},
		{"too@many@ats@example.com", false},
		{"incomplete@domain.", false},
	}

	for _, test := range testCase {
		t.Run(test.email, func(t *testing.T) {
			_, result := validMailAddress(test.email)
			if result != test.expected {
				t.Errorf("Expected validation result for email '%s' to be %t, but got %t", test.email, test.expected, result)
			}
		})
	}
}


