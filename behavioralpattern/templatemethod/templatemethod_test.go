package templatemethod

import "testing"

func TestTemplateMethod(t *testing.T) {
	otpBase := OTPBase{}
	email := &Email{otpBase}
	sms := &SMS{otpBase}
	otpEmail := OTP{email}
	otpSMS := OTP{sms}

	otpEmail.SendVerificationMessageToUser()
	otpSMS.SendVerificationMessageToUser()
}
