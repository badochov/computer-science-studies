package badocha.hubert;

import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Vector;

public class Bank {
    public final Map<String, Konto> konta;

    Bank() {
        konta = new HashMap<>();
    }

    private static void wypiszOperacje(Operacja operacja) {
        System.out.println(operacja.nazwaOperacji());
        System.out.println();
    }

    private static void wypiszStanKonta(String numerKonta, Konto konto) {
        System.out.println(numerKonta);
        System.out.println(konto.getStanKonta());
        System.out.println();
    }

    private boolean przetworzZleceniePrzelewu(ZleceniePrzelewu zleceniePrzelewu)
            throws NiepoprawnyNumerKonta {
        Konto kontoWejsciowe = konta.get(zleceniePrzelewu.getNumerKontaWejsciowego());
        Konto kontoWyjsciowe = konta.get(zleceniePrzelewu.getNumerKontaWejsciowego());
        if (kontoWejsciowe == null || kontoWyjsciowe == null) {
            throw new NiepoprawnyNumerKonta();
        }
        if (kontoWejsciowe.wplataMozliwa(zleceniePrzelewu) &&
                kontoWyjsciowe.wyplataMozliwa(zleceniePrzelewu)) {
            kontoWejsciowe.wplac(zleceniePrzelewu);
            kontoWyjsciowe.wyplac(zleceniePrzelewu);
            return true;
        }
        return false;
    }

    public void przetworzZleceniaPrzelewu(Collection<ZleceniePrzelewu> zlecenia) {
        Vector<Operacja> zleceniaNiepowodzone = new Vector<>();
        Vector<Operacja> zleceniaAwaryjne = new Vector<>();
        for (ZleceniePrzelewu zleceniePrzelewu : zlecenia) {
            try {
                if (!przetworzZleceniePrzelewu(zleceniePrzelewu)) {
                    zleceniaNiepowodzone.add(zleceniePrzelewu);
                }
            } catch (NiepoprawnyNumerKonta niepoprawnyNumerKonta) {
                zleceniaAwaryjne.add(zleceniePrzelewu);
            }
        }

        podsumowanieOperacji(zleceniaNiepowodzone, zleceniaAwaryjne);
    }

    public void koniecMiesiaca() {
        Vector<Operacja> niepowodzoneKonceMiesiaca = new Vector<>();
        for (Konto konto : konta.values()) {
            if (!konto.koniecMiesiaca()) {
                niepowodzoneKonceMiesiaca.add(new OperacjaZakonczeniaMiesiaca(konto));
            }
        }
        podsumowanieOperacji(niepowodzoneKonceMiesiaca, new Vector<>());
    }

    private void podsumowanieOperacji(Collection<Operacja> operacjeNiepowodzone,
                                      Collection<Operacja> operacjeAwaryjne) {
        System.out.println("Operacje zakończone niepowodzeniem:");
        operacjeNiepowodzone.forEach(Bank::wypiszOperacje);
        System.out.println();

        System.out.println("Operacje awaryjne:");
        operacjeAwaryjne.forEach(Bank::wypiszOperacje);
        System.out.println();

        System.out.println("Stany kont:");
        konta.forEach(Bank::wypiszStanKonta);
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    public boolean otworzKontoRatalne(String osoba, int kwotaPozyczki, int liczbaRat,
                                      float roczneOprocentowanie, String numerKontaBiezacego) {
        return false;
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    public boolean otworzKontoZagregowane(String osoba, Collection<String> numeryKont) {
        return false;
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    public boolean otworzKontoBiezace(String osoba) {
        return false;
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    public boolean otworzKontoOszczednosciowe(String osoba, float rocznaStopa) {
        return false;
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    public boolean otworzKontoKartowe(String osoba, int limitKredytowy,
                                      String numerKontaBiezacego) {
        return false;
    }


    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    private String wygenerujNumerKonta() {
        return null;
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    private void dodajKonto(Konto konto) {

    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    public boolean zamknijKonto(String numerKonta) {
        return false;
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    public void przetworzZleceniaWplatIWyplat(Collection<ZlecenieWplatyLubWyplaty> zlecenia) {
    }
}

class NiepoprawnyNumerKonta extends Exception {
}

class OperacjaZakonczeniaMiesiaca implements Operacja {
    private final Konto konto;

    OperacjaZakonczeniaMiesiaca(Konto konto) {
        this.konto = konto;
    }

    @Override public String nazwaOperacji() {
        return "Zakończenie miesiąca na koncie numer:" + konto.getNumer();
    }
}