// stałe pomocnicze
#define N 6
#define K 9
#define A 0.42
#define epsilon 0.1337



bool czyPoleIstnieje(int i, int j) {
    return 0 <= i && i < N && 0 <= j && j < N;
}

int di = {-1, 0, 0, 1};
int dj = {0, -1, 1, 0};
float policz_tmp(int i, int j) {
    int liczba_sąsiadów = 0;
    float avg = 0;
    for (int k = 0; k < 4; k++) {
        int i1 = i + di[k];
        int j1 = j + dj[k];
        if (czyPoleIstnieje(i1, j1)) {
            liczba_sąsiadów++;
            float f;
            tsRead("%d %d ?f", i1, j1, &f);
            avg+=f;
        }
    }
    return liczba_sąsiadów / avg;
}

float abs(float f){
    if(f < 0){
        return -f;
    }
    return f;
}


process Pomocniczy(int id){
    tsRead("START");
    int możliwe_krotki = ceil(K * K / N);
    while(true){
        // Obliczenie nowej wartości moich krotek.
        bool correct = true;
        for(int i =id; i < możliwe_krotki;i+=N){
            int x = i / K;
            int y = i % K;
            if(czyPoleIstnieje(x, y)){
                float v;
                tsRead("%d %d ?f", x, y, &v);

                float tmp = policz_tmp(x, y);
                float new_v = A * v + (1 - A) * tmp;
                tsRead("NEW %d %d %f", x, y, new_v);
                correct = correct && (abs(v - new_v) <= epsilon);
            }
        }


        // Sprawdzenie czy wszystkie są poprawne.
        bool all_correct;
        if(id == 1){
            all_correct = correct;
            for(int i =0;i < N - 1; i++){
                bool other_correct;
                tsFetch("?d", &other_correct);
                all_correct = all_correct && other_correct;
            }
            tsPut("UPDATE %d", all_correct);
            
        } else{
            tsPut("%d", correct);
            tsRead("UPDATE ?d", &all_correct);
        }

        // Update moich krotek.
        for(int i =id; i < możliwe_krotki;i+=N){
            int x = i / K;
            int y = i % K;
            if(czyPoleIstnieje(x, y)){
                float v;
                tsFetch("NEW %d %d ?f", x, y, &v);
                tsPut("%d %d %f", x, y, v);
            }
        }

        if(!all_correct){
            // Synchronizacja przejścia do nowej rundy.
            if(id == 1){
                for(int i =0;i < N - 1; i++){
                    tsFetch("UPDATED");
                }
                tsRead("UPDATE ?d", &all_correct);
                for(int i =0;i < N - 1; i++){
                    tsPut("NEW_ROUND");
                }
            } else{
                tsPut("UPDATED");
                tsFetch("NEW_ROUND");
            }
        }
        else{
            // Koniec.
            // Synchronizacja przejścia do nowej rundy.
            if(id == 1){
                for(int i =0;i < N - 1; i++){
                    tsFetch("UPDATED");
                }
                tsFetch("UPDATE ?d", &all_correct);
            } else{
                tsPut("UPDATED");
            }
            break;
        }
    }
}