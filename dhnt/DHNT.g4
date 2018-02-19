
/** Taken from "The Definitive ANTLR 4 Reference" by Terence Parr */

/*
 http://json.org
 https://tools.ietf.org/html/rfc8259
 */

grammar DHNT;

@header {
    import "strings"
}

@parser::members {
    // here returns `true` iff on the current index of the parser's
    // token stream a token of the given `type` exists on the
    // `HIDDEN` channel.
    //
    // Args:
    //  type (int): the type of the token on the `HIDDEN` channel
    //      to check.
    //
    //  Returns:
    //      `True` iff on the current index of the parser's
    //      token stream a token of the given `type` exists on the
    //      `HIDDEN` channel.
//    func (p *DHNTParser) here(tokenType int) bool {
//        // Get the token ahead of the current index.
//        possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
//        ahead := p.GetTokenStream().Get(possibleIndexEosToken)
//
//        // Check if the token resides on the HIDDEN channel and if it is of the
//        // provided tokenType.
//        return (ahead.GetChannel() == antlr.LexerHidden) && (ahead.GetTokenType() == tokenType)
//    }

    // lineTerminatorAhead returns `true` iff on the current index of the parser's
    // token stream a token exists on the `HIDDEN` channel which
    // either is a line terminator, or is a multi line comment that
    // contains a line terminator.
    func (p *DHNTParser) lineTerminatorAhead() bool {
        // Get the token ahead of the current index.
        possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
        ahead := p.GetTokenStream().Get(possibleIndexEosToken)

        if ahead.GetChannel() != antlr.LexerHidden {
            // We're only interested in tokens on the HIDDEN channel.
            return false
        }

        if ahead.GetTokenType() == DHNTParserTERMINATOR {
            // There is definitely a line terminator ahead.
            return true
        }

        if ahead.GetTokenType() == DHNTParserWS {
            // Get the token ahead of the current whitespaces.
            possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
            ahead = p.GetTokenStream().Get(possibleIndexEosToken)
        }

        // Get the token's text and type.
        text := ahead.GetText()
        tokenType := ahead.GetTokenType()

        // Check if the token is, or contains a line terminator.
        if tokenType == DHNTParserCOMMENT && strings.ContainsAny(text, "\r\n") {
            return true
        }

        return tokenType == DHNTParserTERMINATOR
    }
}

@lexer::members {

    // suppress complilation error: "strings" is imported but not used
    func notUsed() {
        _ = strings.Contains("", "")
    }

    // The most recently produced token.
    var lastToken antlr.Token

    // nextToken returns the next token from the character stream and records this last
    // token in case it resides on the default channel. This recorded token
    // is used to determine when the lexer could possible match a regex
    // literal.
    func (l *DHNTLexer) nextToken() antlr.Token {
        next := l.NextToken()

        if next.GetChannel() == antlr.TokenDefaultChannel {
            // Keep track of the last token on the default channel.
            lastToken = next
        }
        return next
    }
}

/*
  Dahenito Grammar
*/

script
    : value
    ;

object
    : '{' pair (',' pair)* '}'                        # ObjectMembers
    | '{' '}' class                                   # ObjectClass
    | '{' '}'                                         # ZeroObject
    ;

array
    : '[' value ( ',' value )* ']'                    # ArrayElements
    | '[' ']' kind                                    # ArrayKind
    | '[' ']'                                         # ZeroArray
    ;

function
    : '(' parameters? ')' results? block
    ;

channel
    : '<' bufsize '>' kind?              # BufferedChannel
    | '<' '>' kind?                      # UnbufferedChannel
    ;

pair
    : name ':' value
    ;

name
    : STRING
    | IDENTIFIER
    ;

value
    : literal typeAssertion?   # LiteralValue
    | '(' expression ')'       # ExpressionValue
    ;

class
    : name
    ;

kind
    : literal
    ;

literal
    : name
    | NUMBER
    | object | function | channel
    | array
    | TRUE
    | FALSE
    | NULL
    ;

parameters
    :  param ( ',' param )*
    ;

results
    :  param ( ',' param )*
    ;

param
    : name ( ':' kind )?                # ParamName
    | ':' kind                          # ParamAnonymous
    ;

arguments
     : expression (',' expression )*
     ;

block
    : '{' ( statement eos )* '}'
    ;

statement
    : label ':' statement                         # LabeledStmt
    | expression                                  # ExpressionStmt
    | ';'                                         # EmptyStmt
    | statement ( ',' statement )+                # ChoiceStmt
    | block                                       # NestedBlock
    | jump                                        # JumpStmt
    ;

label
    : name                    # NameLabel
    | '(' expression ')'      # ExpressionLabel
    ;

jump
    : '<<-' arguments?                             # ReturnOperation
    | '<-' expression?                             # BreakOperation
    | '->>' expression?                            # GotoOperation
    | '->' expression?                             # ContinueOperation
    | '<<<-' expression?                           # ExitOperation
    ;

bufsize
    : expression
    ;

expression
    : binary                                        # BinaryOperation
    | '@' expression                                # IncludeExpression
    | expression '#' kv? ( block | jump )     # RangeExpression
    ;

binary
    : unary                                         # UnaryOperation
    | binary ( '*' | '/' | '%' ) binary             # MultiplicativeExpression
    | binary ( '+' | '-' ) binary                   # AdditiveExpression
    | binary ( '<<' | '>>' ) binary                 # BitShiftExpression
    | binary ( '<' | '>' | '<=' | '>=' ) binary     # RelationalExpression
    | binary ( '==' | '!=' ) binary                 # EqualityExpression
    | binary '&' binary                             # BitAndExpression
    | binary '&^' binary                            # BinaryExpression
    | binary '^' binary                             # BitXOrExpression
    | binary '|' binary                             # BitOrExpression
    | binary '&&' binary                            # LogicalAndExpression
    | binary '||' binary                            # LogicalOrExpression
    | binary '?:' binary                            # ElvisExpression
    ;

unary
    : primary                                          # PrimaryOperation
    // | unary {!p.here(DHNTParserTERMINATOR)}? '++'   // PostIncrementExpression
    // | unary {!p.here(DHNTParserTERMINATOR)}? '--'   // PostDecreaseExpression
    //| '++' unary                                     // PreIncrementExpression
    //| '--' unary                                     // PreDecreaseExpression
    | '+' unary                                        # UnaryPlusExpression
    | '-' unary                                        # UnaryMinusExpression
    | '~' unary                                        # DeferExpression    //BitNotExpression
    | '!' unary                                        # NotExpression
    | '^' unary                                        # XORExpression
    | '*' unary                                        # DereferenceExpression
    | '&' unary                                        # AsyncOrAddressExpression
    ;

primary
    : operand                                                # OperandOperation
    | primary '[' expression? ']'                            # IndexExpression
    | primary '[' expression? ':' expression? ']'            # SliceExpression
    | primary '.' expression                                 # DotExpression
    | primary typeAssertion                                  # TypeAssertionExpression
    | primary '(' arguments? ')'                             # ArgumentsExpression
    | '?' primary                                            # TypeofExpression
    | '$' primary                                            # ValueofExpression
    | '$#' primary                                           # SizeofExpression
    | primary '?=' kind                                      # InstanceofExpression
    | primary '?<' primary                                   # MemberofExpression
    | primary (',' primary)* ASSIGN_OP arguments             # AssignmentExpression
    ;

operand
    : literal                # LiteralExpression
    | '(' expression ')'     # ParenthesizedExpression
    ;

typeAssertion
    : '(' ':' kind ')'
    ;

kv
    : '(' IDENTIFIER ')'                      # RangeValue
    | '(' IDENTIFIER ',' IDENTIFIER ')'       # RangeKeyValue
    ;

eos
    : ';'
    | EOF
    | {p.lineTerminatorAhead()}?
    | {p.GetTokenStream().LT(1).GetText() == "}" }?
    ;

//
// LEXER
//

ASSIGN_OP
    : '='
    | '+=' | '-=' | '|=' | '^=' | '*=' | '/=' | '%=' | '<<=' | '>>=' | '&=' | '&^='
    ;

//

LBRACE : '{';
RBRACE : '}';
LBRACK : '[';
RBRACK : ']';
LPAREN : '(';
RPAREN : ')';
GT     : '>';
LT     : '<';
COLON  : ':';
COMMA  : ',';
SEMI   : ';';

DOT    : '.';
//ELLIPSIS : '...';
AT : '@';
HASH : '#';
DOLLAR: '$';

//

ASSIGN : '=';

EQUAL : '==';
LE : '<=';
GE : '>=';
NOTEQUAL : '!=';
AND : '&&';
OR : '||';
INC : '++';
DEC : '--';
ADD : '+';
SUB : '-';
MUL : '*';
DIV : '/';
BITAND : '&';
BITOR : '|';
CARET : '^';
MOD : '%';
LSHIFT : '<<';
RSHIFT : '>>';

//
ADD_ASSIGN : '+=';
SUB_ASSIGN : '-=';
OR_ASSIGN : '|=';
XOR_ASSIGN : '^=';
MUL_ASSIGN : '*=';
DIV_ASSIGN : '/=';
MOD_ASSIGN : '%=';
LSHIFT_ASSIGN : '<<=';
RSHIFT_ASSIGN : '>>=';
AND_ASSIGN : '&=';
ANDNOT_ASSIGN : '&^=';

//
BANG     : '!';
TILDE    : '~';
QUESTION : '?';

SIZE_OF     : '$#';
INSTANCE_OF : '?=';
MEMBER_OF   : '?<';

ELVIS       : '?:';

RETURN   : '<<-';
BREAK    : '<-';
GOTO     : '->>';
CONTINUE : '->';
EXIT     : '<<<-';

//

TRUE
    : '1b'
    ;

FALSE
    : '0b'
    ;

NULL
    : '0a'
    ;

IDENTIFIER
    : '_'
    | LETTER ( LETTER | DIGIT )*
    ;

fragment LETTER
    : [a-z]
    ;

fragment DIGIT
    : [0-9]
    ;

//
STRING
    : '"' (ESC | NEWLINE | SAFECODEPOINT)* '"'
    ;

fragment NEWLINE
    : '\r' | '\n'
    ;

fragment ESC
    : '\\' (["\\/bfnrt] | UNICODE)
    ;

fragment UNICODE
    : 'u' HEX HEX HEX HEX
    ;

fragment HEX
    : [0-9a-fA-F]
    ;

fragment SAFECODEPOINT
    : ~ ["\\\u0000-\u001F]
    ;

NUMBER
    : '-'? INT ('.' [0-9] +)? EXP?
    ;

fragment INT
    : '0' | [1-9] [0-9]*
    ;

// no leading zeros
fragment EXP
    : [Ee] [+\-]? INT
    ;

//
// Whitespace and comments
//

WS
    : [ \t]+ -> skip
    ;

COMMENT
    : '/*' .*? '*/' -> channel(HIDDEN)
    ;

TERMINATOR
    : [\r\n]+ -> channel(HIDDEN)
    ;

LINE_COMMENT
    : '//' ~[\r\n]* ->  skip
    ;
