package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	_log "log"
	"os"

	"filippo.io/age/internal/age"
	"golang.org/x/crypto/ssh/terminal"
) decrypt(keys []string, in io.Reader, out io.Writer) {
	identities := []age.Identity{
		// If there is an scrypt recipient (it will have to be the only one and)
		// this identity will be invoked.
		&LazyScryptIdentity{passphrasePrompt},
	}

	// TODO: use the default location if no arguments are provided:
	// os.UserConfigDir()/age/keys.txt, ~/.ssh/id_rsa, ~/.ssh/id_ed25519
	for _, name := range keys {
		ids, err := parseIdentitiesFile(name)
		if err != nil {
			logFatalf("Error: %v", err)
		}
		identities = append(identities, ids...)
	}

	r, err := age.Decrypt(in, identities...)
	if err != nil {
		logFatalf("Error: %v", err)
	}
	if _, err := io.Copy(out, r); err != nil {
		logFatalf("Error: %v", err)
	}
}

func logFatalf(format string, v ...interface{}) {
	_log.Printf(format, v...)
	_log.Fatalf("[ Did age not do what you expected? Could an error be more useful?" +
		" Tell us: https://filippo.io/age/report ]")
}
