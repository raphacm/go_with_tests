package greeting

const spanish = "Spanish"
const french = "French"
const italian = "Italian"

//Hello returns greetings
func Hello(name, language string) string {
	var greetPrefix = "Hello "

	switch language {
	case spanish:
		greetPrefix = "Hola "
	case french:
		greetPrefix = "Bonjour "
	case italian:
		greetPrefix = "Ciao "
	}

	if name == "" {
		name = "World"
	}

	return greetPrefix + name
}
