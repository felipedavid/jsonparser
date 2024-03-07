package parser

import "fmt"

func Parse(buf []byte, obj *map[string]any) error {
	var lex Lexer
	lex.Init(buf)

	tokenStream, err := lex.Parse(buf)
	if err != nil {
		return err
	}

	for token := range tokenStream {
		fmt.Printf("%+v\n", token)
	}

	return nil
}
