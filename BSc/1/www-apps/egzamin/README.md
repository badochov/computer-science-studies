## [Polecenie](https://github.com/kciebiera/egzamin-www-2019)

# Wyjaśnienia

`jeszcze nie ma`

# Ogólnie

Egzamin trwa 2020-07-02 od **9:00** do **13:45**. Czas od 13:45 do 14:00 należy poświęcić na wgrywanie wyników do moodle. Podczas egzaminu jestem na #www-2020 na rokecie.
Na rockecie i tutaj będę publikował wyjaśnienia.

Dostajesz niedokończoną aplikację oraz dokończoną bazę danych (`baza.db` i `data.sql` to te same dane). Po uruchomieniu aplikacja "prawie działa". Obejrzyj
kod aplikacji i wykonaj zadania. Każdy, kto chce zdać, musi wykonać zadanie na 3, a poozstałe punkty są mocno niezależne od siebie i wykonanie każdego z nich
podnosi ocenę o pół stopnia.

Wolno pisać w TypeScript, ale nie dodaje to punktów.

## Zawartość bazy danych

Tabele:

- `osoba` - loginy, hasła (niezaszyfrowane) osób mających konto w naszym serwisie oraz informacje o tym czy osoba jest nauczycielem czy nie
- `wpis` - login osoby (klucz zewnętrzny), czas dodania wpisu i treść wpisu
- `sledzacy` - informacje o tym kto śledzi wpisy kogo (dwa klucze zewnętrzne)

W ramach zadania powstanie serwer, który będzie modyfikował jedynie zawartość tabeli `wpis`, czyli nie będzie można np. dodawać użytkowników ani informacji o tym
kto kogo śledzi.

Nie należy modyfikować struktury bazy danych (np dodawać pól do tabel, ani zmieniać typów, ...)

# Zadania

## Minimum na 3

- w index.js odczytaj poprawnie z bazy danych najnowsze (według pola timestamp) pięć wpisów i wyświetl je użytkownikowi
- zaimplementuj logowanie (na stronie głównej), które po podaniu poprawnie loginu i hasła przenosi na `/users`, a po niepoprawnej próbie logowania pokazuje komunikat o błędnym loginie lub haśle (również na stronie głównej)
- zaimplementuj wyświetlanie na stronie `/users` wpisów śledzonych użytkowników, do strony `/users` powinni mieć dostęp jedynie zalogowani użytkownicy
- zaimplementuj paginację na stronie `/users`
- zaimplementuj wylogowywanie

## Pug +0.5

Zaimplementuj szablony w Pug:

- wykorzystaj dziedziczenie szablonów (Template Inheritance),
- dołącz CSS,
- na każdej ze stron powinna pojawiać się stopka (użyj odpowiedniego tagu z HTML) zawierająca napis: "Imię, Nazwisko egzamin WWW 2019 edycja 1"
- stopka powinna się nie wyświetlać, jeśli szerokość strony jest poniżej 480px

## JSON +0.5

Zaimplementuj dynamiczne ładowanie wpisów wyświetlanych na stronie głównej: zamiast być częścią szablonów powinny być one przesyłane jako JSON.
Implementacja powinna zawierać JS uruchamiany po stronie przeglądarki oraz inny JS uruchamiany po stronie serwera.

## Formularz +0.5

Zaimplementuj odporne na CSRF (użyj csurf) i SQL injection formularze: dodawania wpisu i usuwania własnych wpisów (`/my_entries` w `users.js`), dodawać i usuwać
wpisy powinni móc tylko nauczyciele, przy dodawaniu wpisu powinno się automatycznie na bieżący czas ustawić pole timestamp (+0.5)

## Selenium +0.5

Przygotuj testy w selenium (+0.5). Testy powinny być przygotowane do następujących podpunktów z zadania na 3:

- czy wyświetla się pięć najnowszych wpisów na stronie głównej
- czy po złym haśle jest komunikat o błędzie, a po poprawnym jest się zalogowanym
- czy działa wylogowanie

Do testów dołącz odpowiedni skrypt w package.json (założenie bazy, uruchomienie i zamknięcie serwera, ...)

## Podsumowanie

Opisz w pliku README.md co zrobiłeś/aś (każdy podpunkt 2-3 zdania)

# Uwagi

1. Nie wolno modyfikować struktury bazy danych
2. Nie wolno używać bibliotek CSS, cały CSS należy napisać samodzielnie
3. Podobnie nie wolno używać bibliotej JS/TS, cały JS/TS należy napisać samodzielnie
4. Imię i Nazwisko z punktu 2 to oczywiście dane osobowe zdającego/zdającej
5. Oddać należy zzipowane (plik zip) katalogi `public`, `routes`, `views` oraz pliki `README.md`, `package.json`, `app.*` i ewentualne dodatkowe pliki z danymi do selenium. Nie powinno się oddawać `node_modules`. Zip powinien mieć sporo mniej niż 1 MB.

GLHF

# Co Zrobiłem

## Minimum

- Odczytywanie jest pod adresem posts ponieważ zrobiłęm podpunkt z JSONem i tam zwracam pięć najnowszych wpisów.
- Logowanie jak w opisie
- Stworzyłęm middleware sprawdzajacy czy użytkowaniki jest zalogowany na całym routerze users oraz na trasie "users" zaimplemetowałem wyświetlanie śledzonych wpisów
- zaimplementowałęm paginacje na users i na my_entries, próba przejścia na dalsze storny jest blokowana jeżeli numer storny jest za duży/mały(user nie ma guzików a jeżeli będzie majstorwał przy url dostanie 404)
- wylogowywanie dostępne jest na stronie users

## PUG

- dodałem stopke to layoutu po którym wszystkie inne widoki dziedziczą w stopce napisałęm 2020 zamiast 2019
- dodałem css sprawdzajacy szerokośc strony który odpowiednio chowa/pokazuje stopke

## JSON

- dodałem plik main.js w public.js który fetchuje posty z wcześniej opisanej trasy i wyświetla dane

## Formularz

- dodałem CSURF na wszystkich trasach poniważ kiedy nie jest potrzebny nie przeszkadza a w przeciwnym przypadku jest wygodniej
- wszsytkie query do bazy używają bezpiecznej składni better-sqlite3 z "?", która escapuje

## Selenium

- zaimplementowałem testy z użyciem selenium, chai i mochi
- plik tests/serwer.ts odpala serwer na innym porcie niż zwykle uruchomienie tworzy nową instacje db i nastepnie wyłącza serwer i usuwa bazę
