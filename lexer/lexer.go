package lexer

import (
    "go-monkey-compiler/token"
)

type Lexer struct {
	input 			string
	position 		int // 输入的字符串中的当前位置(指向当前字符)
	readPosition 	int // 输入的字符串中的当前读取位置(指向当前字符串后的那个字符)i
	ch				byte // 当前正在查看的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar(){
	if l.readPosition >= len(l.input) {
		l.ch = 0 // NUL的ASCCI码
	} else {
		l.ch = l.input[l.readPosition]
	}
	// 前移
	l.position = l.readPosition
	l.readPosition += 1
}

// 向前查看一个字符，但是不能移动指针
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:	tokenType,
		Literal: string(ch),
	}
}

// 根据当前的ch创建词法单元
func (l *Lexer) NextToken() token.Token {
    var tok token.Token

	// 跳过空格
	l.skipWhitespace()
    
	switch l.ch {
    case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
        	tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
        tok = newToken(token.PLUS, l.ch)
    case '-':
        tok = newToken(token.MINUS, l.ch)
    case '!':
		if l.peekChar() == '=' {
            // 记录当前ch (!)
            ch := l.ch
            l.readChar()
            literal := string(ch) + string(l.ch)
            tok = token.Token{Type: token.NOT_EQ, Literal: literal}
        } else {
            tok = newToken(token.BANG, l.ch)
        }
    case '/':
        tok = newToken(token.SLASH, l.ch)
    case '*':
        tok = newToken(token.ASTERISK, l.ch)
    case '<':
        tok = newToken(token.LT, l.ch)
    case '>':
        tok = newToken(token.GT, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
         if isLetter(l.ch) {
             tok.Literal = l.readIdentifier()
             tok.Type = token.LookupIdent(tok.Literal)
             return tok
         } else if isDigit(l.ch) {
             tok.Type = token.INT
             tok.Literal = l.readNumber()
             return tok
         } else {
             tok = newToken(token.ILLEGAL, l.ch)
         }
	}

    l.readChar()
    return tok
}

// 判断读取到的字符是不是字母
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 读取字母(标识符/关键字)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		// 如果接下来还有字母，一直移动指针
		l.readChar()
	}
	return l.input[position:l.position]
}


func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' { 
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

