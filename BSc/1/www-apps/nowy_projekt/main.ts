const jsonString: string = `{
    "piloci": [
        "Pirx",
        "Exupery",
        "Idzikowski",
        "Główczewski"
    ],
    "lotniska": {
        "WAW": ["Warszawa", [3690, 2800]],
        "NRT": ["Narita", [4000, 2500]],
        "BQH": ["Biggin Hill", [1802, 792]],
        "LBG": ["Paris-Le Bourget", [2665, 3000, 1845]]
    }
}`;



interface ILotnisko {
    [index: string]: string | number[]
}

interface Pilot {
    [index: number]: string;
    length: number;
}

interface ILiniaLotnicza {
    piloci: Pilot;
    lotniska: ILotnisko
}
function sprawdzPilota(pilot: any): pilot is Pilot {
    if (Array.isArray(pilot)) {
        for (const p of pilot) {
            if (typeof p !== "string") {
                return false;
            }
        }
        return true;
    }
    return false;
}

function sprawdzTabliceIntow(arr: any): arr is number[] {
    if (Array.isArray(arr)) {
        for (const val of arr) {
            if (typeof val !== "number") {
                return false;
            }
        }
        return true;
    }
    return false;
}

function sprawdzLotnisko(lotnisko: any): lotnisko is ILotnisko {
    if (typeof lotnisko === "object") {
        for (const [key, value] of Object.entries(lotnisko)) {
            if (typeof key !== "string") {
                return false;
            }
            if (Array.isArray(value)) {
                for (const val of value) {
                    if (typeof val !== "string" && !sprawdzTabliceIntow(val)) {
                        return false;
                    }
                }
            }
        }
        return true;
    }
    return false;
}

function sprawdzDaneLiniiLotniczej(dane: any): dane is ILiniaLotnicza {
    if (dane.hasOwnProperty("piloci") && dane.hasOwnProperty("lotniska")) {
        return sprawdzPilota(dane.piloci) && sprawdzLotnisko(dane.lotniska);
    }
    return false;
}

const daneLiniiLotniczej = JSON.parse(jsonString);


if (sprawdzDaneLiniiLotniczej(daneLiniiLotniczej)) {
    const juzNaPewnoDaneLinii = daneLiniiLotniczej;
    console.log(juzNaPewnoDaneLinii);
    console.log(juzNaPewnoDaneLinii.piloci.length);
}

function zaloguj(...komunikaty: string[]) {
    console.log("Aleź skomplikowany program", ...komunikaty);
}

zaloguj("Ja", "cię", "nie", "mogę")