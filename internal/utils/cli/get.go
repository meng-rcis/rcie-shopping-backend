package cli

import (
	"os"
)

func GetArg(position int, defaultValue string) string {
	if len(os.Args) < position+1 {
		return defaultValue
	}

	return os.Args[position]
}
