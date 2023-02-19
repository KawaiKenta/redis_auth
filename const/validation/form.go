package validation

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Form interface {
	Validation() error
}

type UserRegister struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (form UserRegister) Validation() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("お名前は必須です")),
		validation.Field(&form.Email, validation.Required.Error("メールアドレスは必須です"),
			is.Email.Error("メールアドレスの形式が間違っています")),
		validation.Field(&form.Password, validation.Required.Error("パスワードは必須です"),
			validation.RuneLength(8, 20).Error("パスワードは8~20文字です"),
			validation.Match(regexp.MustCompile("^[0-9a-zA-Z!-/:-@[-`{-~]+$")).Error("パスワードは英数字記号列からなる必要があります"),
			validation.Match(regexp.MustCompile("[[:alpha:]]")).Error("パスワードは英字を少なくとも1文字含む必要があります"),
			validation.Match(regexp.MustCompile("[[:digit:]]")).Error("パスワードは数字を少なくとも1文字含む必要があります"),
			validation.Match(regexp.MustCompile("[[:punct:]]")).Error("パスワードは記号を少なくとも1文字含む必要があります"),
		),
		validation.Field(&form.PasswordConfirm, validation.In(form.Password).Error("確認用パスワードとパスワードが異なっています")))
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (form Login) Validation() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Email, validation.Required.Error("メールアドレスは必須です"),
			is.Email.Error("メールアドレスの形式が間違っています")),
		validation.Field(&form.Password, validation.Required.Error("パスワードは必須です"),
			validation.RuneLength(8, 20).Error("パスワードは8~20文字です"),
			validation.Match(regexp.MustCompile("^[0-9a-zA-Z!-/:-@[-`{-~]+$")).Error("パスワードは英数字記号列からなる必要があります"),
			validation.Match(regexp.MustCompile("[[:alpha:]]")).Error("パスワードは英字を少なくとも1文字含む必要があります"),
			validation.Match(regexp.MustCompile("[[:digit:]]")).Error("パスワードは数字を少なくとも1文字含む必要があります"),
			validation.Match(regexp.MustCompile("[[:punct:]]")).Error("パスワードは記号を少なくとも1文字含む必要があります"),
		),
	)
}

type ForgetPassword struct {
	Email string `json:"email"`
}

func (form ForgetPassword) Validation() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Email, validation.Required.Error("メールアドレスは必須です"),
			is.Email.Error("メールアドレスの形式が間違っています")),
	)
}

type ResetPassword struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (form ResetPassword) Validation() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Password, validation.Required.Error("パスワードは必須です"),
			validation.RuneLength(8, 20).Error("パスワードは8~20文字です"),
			validation.Match(regexp.MustCompile("^[0-9a-zA-Z!-/:-@[-`{-~]+$")).Error("パスワードは英数字記号列からなる必要があります"),
			validation.Match(regexp.MustCompile("[[:alpha:]]")).Error("パスワードは英字を少なくとも1文字含む必要があります"),
			validation.Match(regexp.MustCompile("[[:digit:]]")).Error("パスワードは数字を少なくとも1文字含む必要があります"),
			validation.Match(regexp.MustCompile("[[:punct:]]")).Error("パスワードは記号を少なくとも1文字含む必要があります"),
		),
		validation.Field(&form.PasswordConfirm, validation.In(form.Password).Error("確認用パスワードとパスワードが異なっています")),
	)
}
