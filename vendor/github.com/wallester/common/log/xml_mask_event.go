package log

import (
	"strings"

	"github.com/clbanning/mxj"
	"github.com/juju/errors"
	commonstrings "github.com/wallester/common/strings"
)

func xmlMaskEvent(event map[string]interface{}) (map[string]interface{}, error) {
	for key, value := range event {
		stringValue, ok := value.(string)
		if ok {
			maskedMessage, err := maskXML(stringValue, sensitiveKeywords)
			if err != nil {
				return nil, errors.Annotate(err, "masking XML failed")
			}

			// Message is not in XML format
			if maskedMessage == nil {
				continue
			}

			event[key] = *maskedMessage
		}
	}

	return event, nil
}

// private

func maskXML(message string, keywords []string) (*string, error) {
	reader := strings.NewReader(message)
	mSeq, err := mxj.NewMapXmlSeqReader(reader)
	if err != nil {
		// Message is not in XML format
		return nil, nil
	}

	m := mxj.Map(mSeq)

	for _, keyword := range keywords {
		for _, leafPath := range m.LeafPaths() {
			if strings.Contains(strings.ToLower(leafPath), keyword+".#text") {
				if _, err = m.UpdateValuesForPath("#text:"+commonstrings.Mask, leafPath); err != nil {
					return nil, errors.Annotate(err, "updating values for path failed")
				}
			}
		}
	}

	xmlBytes, err := mSeq.Xml()
	if err != nil {
		return nil, errors.Annotate(err, "encoding map as XML failed")
	}

	xmlString := string(xmlBytes)
	return &xmlString, nil
}
