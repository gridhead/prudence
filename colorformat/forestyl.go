package colorformat

func ForeBold(textobjc string) string {
	return string("\u001b[1m") + textobjc + string("\u001b[0m")
}

func ForeUndl(textobjc string) string {
	return string("\u001b[4m") + textobjc + string("\u001b[0m")
}

func ForeInvr(textobjc string) string {
	return string("\u001b[7m") + textobjc + string("\u001b[0m")
}
