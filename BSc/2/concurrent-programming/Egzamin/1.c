#define K 6
#define W 9
#define N 42


porcja początkowa(int w, int l);




process Z(){
    // Który proces prosił o przesunięcie kolumny
    int kolumny[K] = {-1};
    // Który proces prosił o przesunięcie wiersza
    int wiersze[W] = {-1};

    bool można_przesuwać_wiersze = true;
    bool można_przesuwać_kolumny = true;

    int l_w = 0;
    int l_k = 0;
    int liczba_wierszy_przesuniętych = 0;
    int liczba_kolumn_przesuniętych = 0;

    queue wiersz_q;
    queue kol_q;
    while(true){
        int kol, wiersz, p, n;
        select{
            on przesuńWiersz(n, wiersz, p){
                if(można_przesuwać_wiersze && liczba_wierszy_przesuniętych < W){
                    można_przesuwać_kolumny = false;
                    send M[wiersz][0].p;
                    wiersze[wiersz] = n;
                    l_w++;
                    liczba_wierszy_przesuniętych++;
                }
                else{
                    wiersz_q.add(n, wiersz, p);
                }
            }
            on przesuńKolumne(n, kol, p){
                if(można_przesuwać_kolumny && liczba_kolumn_przesuniętych < K){
                    można_przesuwać_wiersze = false;
                    send M[0][kol].p;
                    kolumny[kol] = n;
                    l_k++;
                    liczba_kolumn_przesuniętych++;
                }
                else{
                    kol_q.add(n, kol, p);
                }
            }
            // Opis muzyczno słownmy ponieważ skończył mi się czas.
            on wierszPrzesunięty(wiersz){
                int n = wiersze[wiersz];
                send U[n].wiersz;
                wiersze[wiersz] = -1;
                l_w--;
                if(empty(kol_q)){
                    // b() - na kolejce jest prośba o przesuniecię obecnego wiersza
                    // sprawdzenie odbywa sie przez przesuwanie elemntów kolejki z poczatku na koniec póki nie przejdziemy przz wszystkie lub interesujący nas element nie będzie na przodzie
                    if( b(wiersz)){
                    n, wiersz, p = wiersz_q.pop();
                    można_przesuwać_kolumny = false;
                    send M[wiersz][0].p;
                    wiersze[wiersz] = n;
                    l_w++;
                    liczba_wierszy_przesuniętych++;
                }
                else{
                    if(l_w == 0){
                        można_przesuwać_kolumny = true;
                        // Odpal K kolumn
                    }
                }
            }
            on kolumnaPrzesunięta(kol){
                // Analogicznie do wiersz przesunięty
            }
        }
    }
}

process M(int w, int k){
    porcja moja_porcja = początkowa(w, k);
    enum kierunek{KOL, WIERSZ};

    enum kierunek kierunek_przesuniecia;
    // Czy wiersz / kolumna jest przesuwany.
    bool przesuwany = false;

    // Otrzymana przy przesunięciu porcja.
    porcja new_p;
    bool porcja_otrzymana;
    while(true){
        int p, n;
        select{
            on daj(n) {
                send U[n].p
            }
            on przesuńKolumne(p) {
                przesuwany = true;
                kierunek_przesuniecia = KOL;
                // Roześlij info o przesyłaniu.
                if(2*w < W){
                    send M[2 * w][k].przesuńKolumne(p);
                }
                if(2*w + 1 < W){
                    send M[2* w + 1][k].przesuńKolumne(p);
                }

                // Wyślij do poprawnej komórki swoją wartość; 
                send M[(w+p) % W][k].przesuwamy(p);
            }
            on przesuńWiersz(p) {
                przesuwany = true;
                kierunek_przesuniecia = Wiersz;
                // Roześlij info o przesyłaniu.
                if(2 * k < K){
                    send M[w][2 * k].przesuńKolumne(p);
                }
                if(2 * k + 1 < L){
                    send M[w][2 * k + 1].przesuńKolumne(p);
                }

                // Wyślij do poprawnej komórki swoją wartość; 
                
                send M[w][(k+p) % W].przesuwamy(p);
   
            }
            on przesuwamy(new_p) {
                porcja_otrzymana = true;
            }
        }

        // Zakończenie przesuwania
        if(porcja_otrzymana && przesuwany){
            p = new_p;
            if(kierunek == KOL){
                if(2 * w < W){
                    receive Przesuniete;
                }
                if(2 * w + 1 < W){
                    receive Przesuniete;
                }


                if(w != 0){
                    send M[w/2][k].Przesuniete;
                }
                else{
                    send Z.KolumnaPrzesunięta(k);
                }
            } else {
                if(2 * k < K){
                    receive Przesuniete;
                }
                if(2 * k + 1 < K){
                    receive Przesuniete;
                }


                if(k != 0){
                    send M[w][k / 2].Przesuniete;
                }
                else{
                    send Z.WierszPrzesunięty(w);
                }
            }
        }
    }
}
