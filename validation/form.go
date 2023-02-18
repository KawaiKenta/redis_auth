package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Form interface {
	Validation() error
}

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (form UserRegister) Validation() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("お名前は必須です")),
		validation.Field(&form.Email, validation.Required.Error("メールアドレスは必須です"),
			is.Email.Error("メールアドレスの形式が間違っています")),
		validation.Field(&form.Password, validation.Required.Error("パスワードは必須です"),
			validation.RuneLength(8, 20).Error("パスワードは8~20文字です"),
			is.Alphanumeric.Error("パスワードは英数字で入力してください"),
			// validation.Match(regexp.MustCompile("^(?=.*[A-Z])[a-zA-Z0-9!-/:-@[-`{-~]$")).Error("パスワードは大文字を少なくとも1文字含む必要があります"),
			// validation.Match(regexp.MustCompile("^(?=.*[!-/:-@[-`{-~])[a-zA-Z0-9!-/:-@[-`{-~]$")).Error("パスワードは記号を少なくとも1文字含む必要があります"),
		))
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
			is.Alphanumeric.Error("パスワードは英数字で入力してください。"),
			// validation.Match(regexp.MustCompile("^(?=.*[A-Z])[a-zA-Z0-9!-/:-@[-`{-~]$")).Error("パスワードは大文字を少なくとも1文字含む必要があります"),
			// validation.Match(regexp.MustCompile("^(?=.*[!-/:-@[-`{-~])[a-zA-Z0-9!-/:-@[-`{-~]$")).Error("パスワードは記号を少なくとも1文字含む必要があります"),
		),
	)
}

type ForgetPassword struct {
	Email string `json:"email" binding:"required,email"`
}

func (form ForgetPassword) Validation() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Email, validation.Required.Error("メールアドレスは必須です"),
			is.Email.Error("メールアドレスの形式が間違っています")),
	)
}

type ResetPassword struct {
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required,min=8"`
}

func (form ResetPassword) Validation() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Password, validation.Required.Error("パスワードは必須です"),
			validation.RuneLength(8, 20).Error("パスワードは8~20文字です"),
			is.Alphanumeric.Error("パスワードは英数字で入力してください。"),
			// validation.Match(regexp.MustCompile("^(?=.*[A-Z])[a-zA-Z0-9!-/:-@[-`{-~]$")).Error("パスワードは大文字を少なくとも1文字含む必要があります"),
			// validation.Match(regexp.MustCompile("^(?=.*[!-/:-@[-`{-~])[a-zA-Z0-9!-/:-@[-`{-~]$")).Error("パスワードは記号を少なくとも1文字含む必要があります"),
		),
		validation.Field(&form.PasswordConfirm, validation.In(form.Password).Error("確認用パスワードとパスワードが異なっています")),
	)
}
