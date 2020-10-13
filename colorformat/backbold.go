package colorformat

func BackBlkBold(textobjc string) string {
	return string("\u001b[40;1m") + textobjc + string("\u001b[0m")
}

func BackRedBold(textobjc string) string {
	return string("\u001b[41;1m") + textobjc + string("\u001b[0m")
}

func BackGrnBold(textobjc string) string {
	return string("\u001b[42;1m") + textobjc + string("\u001b[0m")
}

func BackYlwBold(textobjc string) string {
	return string("\u001b[43;1m") + textobjc + string("\u001b[0m")
}

func BackBluBold(textobjc string) string {
	return string("\u001b[44;1m") + textobjc + string("\u001b[0m")
}

func BackMgtBold(textobjc string) string {
	return string("\u001b[45;1m") + textobjc + string("\u001b[0m")
}

func BackCynBold(textobjc string) string {
	return string("\u001b[46;1m") + textobjc + string("\u001b[0m")
}

func BackWitBold(textobjc string) string {
	return string("\u001b[47;1m") + textobjc + string("\u001b[0m")
}
