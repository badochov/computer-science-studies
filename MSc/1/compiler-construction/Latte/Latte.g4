grammar Latte;

program: topDef* EOF;

topDef:
    type_ ID '(' arg? ')' block                         # DefFunc
    | 'class' ID ('extends' ID)? '{' classElement* '}'  # DefClass;

arg: type_ ref ID ( ',' type_ ref ID)*;

ref: ('&')?;

block: '{' stmt* '}';

stmt:
	';'										# Empty
	| expr '[' expr ']' '=' expr ';'	    # IndexAss
	| expr '.' ID '=' expr ';'				# AttrAss
	| block									# BlockStmt
	| type_ item ( ',' item)* ';'			# Decl
	| ID '=' expr ';'						# Ass
	| ID '++' ';'							# Incr
	| ID '--' ';'							# Decr
	| 'return' expr ';'						# Ret
	| 'return' ';'							# VRet
	| 'if' '(' expr ')' stmt				# Cond
	| 'if' '(' expr ')' stmt 'else' stmt	# CondElse
	| 'while' '(' expr ')' stmt				# While
	| 'for' '(' type_ ID ':' expr ')' stmt	# ForEach
	| expr ';'								# SExp;

type_:
	type_single '[]'	# TypeArray
    | type_single       # TypeSingle
    ;

type_single:
    'int'			# Int
    | 'string'		# Str
    | 'boolean'		# Bool
    | 'void'		# Void
    | ID			# ClassName
    ;

classElement:
    type_ ID (',' ID)* ';'        # ClassAttribute
    | type_ ID '(' arg? ')' block # ClassMethod;

item: ID | ID '=' expr;

expr:
	expr '[' expr ']'							# EIndex
	| expr '.' ID '(' ( expr ( ',' expr)*)? ')'	# EMethodCall
	| expr '.' ID								# EAttr
	| 'new' type_single '[' expr ']'			# ENewArray
	| 'new' type_single						    # ENewClass
	| unOp expr									# EUnOp
	| expr mulOp expr							# EMulOp
	| expr addOp expr							# EAddOp
	| expr relOp expr							# ERelOp
	| <assoc = right> expr '&&' expr			# EAnd
	| <assoc = right> expr '||' expr			# EOr
	| ID										# EId
	| INT										# EInt
	| 'true'									# ETrue
	| 'false'									# EFalse
	| ID '(' ( expr ( ',' expr)*)? ')'			# EFunCall
	| STR										# EStr
	| ('(' type_ ')')? 'null'				    # ENull
	| '(' expr ')'								# EParen;

unOp: '!' | '-';

addOp: '+' | '-';

mulOp: '*' | '/' | '%';

relOp: '<' | '<=' | '>' | '>=' | '==' | '!=';

COMMENT: ('#' ~[\r\n]* | '//' ~[\r\n]*) -> channel(HIDDEN);
MULTICOMMENT: '/*' .*? '*/' -> channel(HIDDEN);

fragment Letter: Capital | Small;
fragment Capital: [A-Z\u00C0-\u00D6\u00D8-\u00DE];
fragment Small: [a-z\u00DF-\u00F6\u00F8-\u00FF];
fragment Digit: [0-9];

INT: Digit+;
fragment ID_First: Letter | '_';
ID: ID_First (ID_First | Digit)*;

WS: (' ' | '\r' | '\t' | '\n')+ -> skip;

STR: '"' StringCharacters? '"';
fragment StringCharacters: StringCharacter+;
fragment StringCharacter: ~["\\] | '\\' [tnr"\\];
