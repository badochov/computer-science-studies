/*
 * Parser Rules
 */

grammar Instant;

program: (stmt (';' stmt)*)? (';')?;

stmt: IDENT '=' expr # SAssign | expr # SExpr;

expr:
	IDENT							# EId
	| INT							# EInt
	| expr mulOp expr				# EMul
	| expr '-' expr					# ESub
	| <assoc = right> expr '+' expr	# EAdd
	| '(' expr ')'					# EParen;

mulOp: '*' | '/';

/*
 * Lexer Rules
 */
fragment Letter: Capital | Small;
fragment Capital: [A-Z];
fragment Small: [a-z];
fragment Digit: [0-9];

INT: Digit+;
IDENT: Letter (Letter | '_' | '\'' | Digit)*;

WS: (' ' | '\r' | '\t' | '\n')+ -> skip;
