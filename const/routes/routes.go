package routes

// API
const SignUp = "/signup"
const Login = "/login"
const Logout = "/logout"
const RequestPasswordReset = "/reset/request"

// views
const VerifyEmail = "/verify/:uuid"
const PasswordResetForm = "/reset/form/:uuid"
const VerifyPasswordReset = "/reset/verify"
