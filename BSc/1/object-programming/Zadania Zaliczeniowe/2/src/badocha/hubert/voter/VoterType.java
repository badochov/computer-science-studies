package badocha.hubert.voter;

/**
 * Typ wyliczeniowy zwracający typ osoby głosującej.
 */
public enum VoterType {
    PartyElectorate(VoterStyle.SingleParty, VoterTraits.No),
    CandidateElectorate(VoterStyle.SingleCandidate, VoterTraits.No),
    MinOneTraitMultiParty(VoterStyle.Multi, VoterTraits.Single),
    MaxOneTraitMultiParty(VoterStyle.Multi, VoterTraits.Single),
    TraitsWeightedMultiParty(VoterStyle.Multi, VoterTraits.Multi),
    MinOneTraitSingleParty(VoterStyle.SingleParty, VoterTraits.Single),
    MaxOneTraitSingleParty(VoterStyle.SingleParty, VoterTraits.Single),
    TraitsWeightedSingleParty(VoterStyle.SingleParty, VoterTraits.Multi);

    private final VoterStyle style;
    private final VoterTraits traits;

    VoterType(VoterStyle voterStyle, VoterTraits voterTraits) {
        style = voterStyle;
        traits = voterTraits;
    }

    public static boolean isMin(VoterType voterType) {
        return voterType == MinOneTraitMultiParty || voterType == MinOneTraitSingleParty;
    }

    public static VoterType getVoterType(int type) {
        switch (type) {
            case 1:
                return PartyElectorate;
            case 2:
                return CandidateElectorate;
            case 3:
                return MinOneTraitMultiParty;
            case 4:
                return MaxOneTraitMultiParty;
            case 5:
                return TraitsWeightedMultiParty;
            case 6:
                return MinOneTraitSingleParty;
            case 7:
                return MaxOneTraitSingleParty;
            case 8:
                return TraitsWeightedSingleParty;
            default:
                return null;
        }
    }

    public boolean isSinglePartyVoter() {
        return style != VoterStyle.Multi;
    }

    public boolean isSingleCandidateVoter() {
        return style == VoterStyle.SingleCandidate;
    }

    public boolean usesTraits() {
        return traits != VoterTraits.No;
    }

    public boolean usesOneTrait() {
        return traits == VoterTraits.Single;
    }
}

/**
 * Typ wyliczeniowy mówiący o stylu głosowania danej osoby.
 */
enum VoterStyle {
    SingleParty,
    SingleCandidate,
    Multi
}

/**
 * Typ wyliczeniowy mówiący czy dana osoba patrzy na cechy podczas głosowania.
 */
enum VoterTraits {
    No,
    Single,
    Multi
}