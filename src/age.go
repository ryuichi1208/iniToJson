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
) 

decrypt(keys []string, in io.Reader, out io.Writer) {
	identities := []age.Identity{

		&LazyScryptIdentity{passphrasePrompt},
	}


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

}
