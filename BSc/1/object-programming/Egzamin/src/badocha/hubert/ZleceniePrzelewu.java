package badocha.hubert;

public class ZleceniePrzelewu extends ZlecenieAbstrakcyjne {
    private final String numerKontaWyjsciowego;
    private final String numerKontaWejsciowego;

    public ZleceniePrzelewu(int kwota, String waluta, String numerKontaWyjsciowego,
                            String numerKontaWejsciowego) {
        super(kwota, waluta);
        this.numerKontaWyjsciowego = numerKontaWyjsciowego;
        this.numerKontaWejsciowego = numerKontaWejsciowego;
    }


    public String getNumerKontaWyjsciowego() {
        return numerKontaWyjsciowego;
    }

    public String getNumerKontaWejsciowego() {
        return numerKontaWejsciowego;
    }

    @Override public String nazwaOperacji() {
        return "ZleceniePrzelewu{" +
                "kwota=" + kwota +
                ", waluta='" + waluta + '\'' +
                ", numerKontaWyjsciowego='" + numerKontaWyjsciowego + '\'' +
                ", numerKontaWejsciowego='" + numerKontaWejsciowego + '\'' +
                '}';
    }
}

