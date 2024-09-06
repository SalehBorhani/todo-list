package phonenumber

func IsValid(phoneNumber string) bool {
	// TODO - Check the Phone Number with the regex pattern
	if len(phoneNumber) == 11 {
		return true
	}
	return false
}
