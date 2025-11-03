package token

// 词法单元类型
type TokenType string

type Token struct {
   Type TokenType
   // 字面量
   Literal string
}

const (
    // 特殊类型
    ILLEGAL = "ILLEGAL" // 未知字符
    EOF     = "EOF"     // 文件结尾

    // 标识符+字面量
    IDENT = "IDENT" // add, foobar, x, y
    INT   = "INT"   // 1343456

    // 运算符
    ASSIGN = "="
    PLUS   = "+"
	MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"

    LT = "<"
    GT = ">"

	EQ	= "=="
	NOT_EQ = "!="

    // 分隔符
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // 关键字
    FUNCTION = "FUNCTION"
    LET      = "LET"
	TRUE     = "TRUE"
    FALSE    = "FALSE"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"
)

var keywords = map[string]TokenType {
	"fn":	FUNCTION,
	"let":	LET,
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
