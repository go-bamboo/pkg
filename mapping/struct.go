package mapping

import "github.com/mitchellh/mapstructure"

func DecodeStruct(s map[string]interface{}, t interface{}) error {
	var md mapstructure.Metadata
	cfg := mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   t,
	}
	decoder, err := mapstructure.NewDecoder(&cfg)
	if err = decoder.Decode(s); err != nil {
		return err
	}
	return nil
}
