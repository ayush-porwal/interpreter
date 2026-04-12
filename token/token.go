package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT   = "INT"
	IDENT = "IDENT"

	// Operators
	ASSIGN       = "="
	DIVIDE       = "/"
	EQUALS       = "=="
	MINUS        = "-"
	MULTIPLY     = "*"
	NEGATE       = "!"
	NOT_EQUALS   = "!="
	PLUS         = "+"
	LESS_THAN    = "<"
	GREATER_THAN = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	LET      = "let"
	FUNCTION = "fn"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
