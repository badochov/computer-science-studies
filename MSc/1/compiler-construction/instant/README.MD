# Instant

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

## Struktura plików
- src/common
  - funkcje pomocnicze używane w obu kompliatorach
- src/jvm
  - kompilator jvm
- src/llvm
  - komilator llvm
- src/parser
  - parser wygenerowany przez antlr
- vendor
  - biblioteki