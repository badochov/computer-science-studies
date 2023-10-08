package badocha.hubert;

import java.util.ArrayList;
import java.util.List;

public class Main {

    public static void main(String[] args) {

        List<Character> lst = new ArrayList<>();
        for(char c : "Hello^".toCharArray()){
            lst.add(c);
        }
        Napis n1 = new NapisLst(lst);
        Napis n2 = new NapisTab("World?".toCharArray());

        assert n1.dlugosc() == 6;
        assert n2.dlugosc() == 6;

        assert n1.dajZnak(1) == 'e';
        assert n1.dajZnak(0) == 'H';
        assert n2.dajZnak(1) == 'o';
        assert n2.dajZnak(0) == 'W';

        assert n1.ustawZnak(5, ' ');
        assert n2.ustawZnak(5, '!');

        assert n1.dajZnak(5) == ' ';
        assert n2.dajZnak(5) == '!';

        n1.dodajNapis(n2);
        n2.dodajNapis(n1);

        n1.wypisz();
        n2.wypisz();
    }
}
