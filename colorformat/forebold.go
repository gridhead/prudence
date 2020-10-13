package colorformat

func ForeBlkBold(textobjc string) string {
	return string("\u001b[30;1m") + textobjc + string("\u001b[0m")
}

func ForeRedBold(textobjc string) string {
	return string("\u001b[31;1m") + textobjc + string("\u001b[0m")
}

func ForeGrnBold(textobjc string) string {
	return string("\u001b[32;1m") + textobjc + string("\u001b[0m")
}

func ForeYlwBold(textobjc string) string {
	return string("\u001b[33;1m") + textobjc + string("\u001b[0m")
}

func ForeBluBold(textobjc string) string {
	return string("\u001b[34;1m") + textobjc + string("\u001b[0m")
}

func ForeMgtBold(textobjc string) string {
	return string("\u001b[35;1m") + textobjc + string("\u001b[0m")
}

func ForeCynBold(textobjc string) string {
	return string("\u001b[36;1m") + textobjc + string("\u001b[0m")
}

func ForeWitBold(textobjc string) string {
	return string("\u001b[37;1m") + textobjc + string("\u001b[0m")
}
