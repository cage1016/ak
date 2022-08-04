package alfred

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"howett.net/plist"
)

// Plist is a plist data structure
type Plist map[string]interface{}

// LoadPlist loads a plist from an XML file
func LoadPlist(filename string) (p Plist) {
	var err error
	var xmlData []byte
	if xmlData, err = ioutil.ReadFile(filename); err != nil {
		logrus.Fatalf("error reading plist file: %s", err)
	}

	if _, err = plist.Unmarshal(xmlData, &p); err != nil {
		logrus.Fatalf("error deserializing plist data: %s", err)
	}

	return
}

// SavePlist saves a plist to an XML file
func SavePlist(filename string, p Plist) {
	var err error
	var xmlData []byte
	if xmlData, err = plist.MarshalIndent(p, plist.XMLFormat, "\t"); err != nil {
		logrus.Fatalf("error serializing plist data: %s", err)
	}

	if err = ioutil.WriteFile(filename, xmlData, 0644); err != nil {
		logrus.Fatalf("error writing plist file: %s", err)
	}
}
