package parser

type TokenKind int

const (
	TkString TokenKind = iota
	TkIdentifier
	TkNumber

	TkLeftSquareBracket
	TkRightSquareBracket
	TkLeftCurlyBracket
	TkRightCurlyBracket

	TkColon
	TkComma

	TkEof
)

type Token struct {
	Kind  TokenKind
	Start int
	End   int
}

type Lexer struct {
	token  Token
	stream []byte
}

func (l *Lexer) Init(stream []byte) {
	l.stream = stream
}

func (l *Lexer) Parse(stream []byte) ([]Token, error) {
	tokenStream := make([]Token, 0, len(stream))

	for {
		err := l.Next()
		if err != nil {
			return nil, err
		}

		if l.token.Kind == TkEof {
			break
		}

		tokenStream = append(tokenStream, l.token)
	}

	return tokenStream, nil
}

func (l *Lexer) Next() error {
	return nil
}
