package config

import (
	"fmt"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	cases := map[string]struct {
		cnfg     Config
		name     string
		cnfg_out Config
	}{
		"simple": {Config{DB_URL: "postgres://example"}, "ALurkingUser", Config{DB_URL: "postgres://example", CurrentUserName: "ALurkingUser"}},
	}

	//temp file that we can test with
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Errorf("Write Failed %v", err)
		return
	}
	tmpFile, err := os.CreateTemp(homedir, ".tempconfig.json")
	if err != nil {
		panic(fmt.Errorf("failed to create temporary file: %w", err))
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte("{\"db_url\": \"postgres://example\"}"))

	testconfigFilePath = tmpFile.Name()

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {

			tc.cnfg.SetUser(tc.name)
			rslt, err := Read()
			if err != nil {
				t.Errorf("Write Failed %v", err)
				return
			}
			if rslt.CurrentUserName != tc.cnfg_out.CurrentUserName {
				t.Errorf("Name Mismatch wanted: %v , got: %v", tc.cnfg_out.CurrentUserName, rslt.CurrentUserName)
				return
			}
		})
	}

}
