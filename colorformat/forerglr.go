package colorformat

func ForeBlkRglr(textobjc string) string {
	return string("\u001b[30m") + textobjc + string("\u001b[0m")
}

func ForeRedRglr(textobjc string) string {
	return string("\u001b[31m") + textobjc + string("\u001b[0m")
}

func ForeGrnRglr(textobjc string) string {
	return string("\u001b[32m") + textobjc + string("\u001b[0m")
}

func ForeYlwRglr(textobjc string) string {
	return string("\u001b[33m") + textobjc + string("\u001b[0m")
}

func ForeBluRglr(textobjc string) string {
	return string("\u001b[34m") + textobjc + string("\u001b[0m")
}

func ForeMgtRglr(textobjc string) string {
	return string("\u001b[35m") + textobjc + string("\u001b[0m")
}

func ForeCynRglr(textobjc string) string {
	return string("\u001b[36m") + textobjc + string("\u001b[0m")
}

func ForeWitRglr(textobjc string) string {
	return string("\u001b[37m") + textobjc + string("\u001b[0m")
}
