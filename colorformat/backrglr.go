package colorformat

func BackBlkRglr(textobjc string) string {
	return string("\u001b[40m") + textobjc + string("\u001b[0m")
}

func BackRedRglr(textobjc string) string {
	return string("\u001b[41m") + textobjc + string("\u001b[0m")
}

func BackGrnRglr(textobjc string) string {
	return string("\u001b[42m") + textobjc + string("\u001b[0m")
}

func BackYlwRglr(textobjc string) string {
	return string("\u001b[43m") + textobjc + string("\u001b[0m")
}

func BackBluRglr(textobjc string) string {
	return string("\u001b[44m") + textobjc + string("\u001b[0m")
}

func BackMgtRglr(textobjc string) string {
	return string("\u001b[45m") + textobjc + string("\u001b[0m")
}

func BackCynRglr(textobjc string) string {
	return string("\u001b[46m") + textobjc + string("\u001b[0m")
}

func BackWitRglr(textobjc string) string {
	return string("\u001b[47m") + textobjc + string("\u001b[0m")
}
