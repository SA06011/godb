package data

type dbValue interface{}

type encode func(string) []byte
type decode func([]byte) dbValue

Encode := []encode{
	encodeInt ,
	encodeFloat
}

func encodeInt(string) []byte{
	return nil
}

func encodeFloat(string) []byte{
	return nil
}
