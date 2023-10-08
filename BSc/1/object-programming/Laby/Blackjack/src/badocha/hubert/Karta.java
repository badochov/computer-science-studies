package badocha.hubert;

public enum Karta {
    K2(2),
    K3(3),
    K4(4),
    K5(5),
    K6(6),
    K7(7),
    K8(8),
    K9(9),
    K10(10),
    W(10),
    D(10),
    K(10),
    AS(11);

    int punkty;

    Karta(int oczka) {
        punkty = oczka;
    }
}
