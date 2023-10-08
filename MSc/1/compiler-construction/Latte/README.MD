# Latte

## Kompilacja
Program napisany jest w języku Go. Do kompilacji potrzebny jest toolkit do go w wersji co najmniej 1.19
https://go.dev/dl/

## Używane biblioteki
Wszystkie używane biblioteki znajdują się w katalogu vendor.

Używane explicite biblioteki to:
- github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20220911224424-aa1f1f12a846
  - Daje ona runtime antlera dla Go
- github.com/llir/llvm v0.3.3
  - Pozwala ona w łatwy sposób czytać modyfikoć i zapisywać kod llvm.
- golang.org/x/exp
  - eksperymentalna biblioteka standardowa Golanga
    - wygodne funkcje na slice'ach 

## Struktura plików
- src/ast
  - własne AST
- src/compiler
  - kompilator
- src/frontend
  - typechecker
  - analiza returnów
  - zmiana AST na własne
  - upraszczanie stałych
- vendor
    - biblioteki