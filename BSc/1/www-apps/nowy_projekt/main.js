"use strict";
var __spreadArrays = (this && this.__spreadArrays) || function () {
    for (var s = 0, i = 0, il = arguments.length; i < il; i++) s += arguments[i].length;
    for (var r = Array(s), k = 0, i = 0; i < il; i++)
        for (var a = arguments[i], j = 0, jl = a.length; j < jl; j++, k++)
            r[k] = a[j];
    return r;
};
var jsonString = "{\n    \"piloci\": [\n        \"Pirx\",\n        \"Exupery\",\n        \"Idzikowski\",\n        \"G\u0142\u00F3wczewski\"\n    ],\n    \"lotniska\": {\n        \"WAW\": [\"Warszawa\", [3690, 2800]],\n        \"NRT\": [\"Narita\", [4000, 2500]],\n        \"BQH\": [\"Biggin Hill\", [1802, 792]],\n        \"LBG\": [\"Paris-Le Bourget\", [2665, 3000, 1845]]\n    }\n}";
function sprawdzPilota(pilot) {
    if (Array.isArray(pilot)) {
        for (var _i = 0, pilot_1 = pilot; _i < pilot_1.length; _i++) {
            var p = pilot_1[_i];
            if (typeof p !== "string") {
                return false;
            }
        }
        return true;
    }
    return false;
}
function sprawdzTabliceIntow(arr) {
    if (Array.isArray(arr)) {
        for (var _i = 0, arr_1 = arr; _i < arr_1.length; _i++) {
            var val = arr_1[_i];
            if (typeof val !== "number") {
                return false;
            }
        }
        return true;
    }
    return false;
}
function sprawdzLotnisko(lotnisko) {
    if (typeof lotnisko === "object") {
        for (var _i = 0, _a = Object.entries(lotnisko); _i < _a.length; _i++) {
            var _b = _a[_i], key = _b[0], value = _b[1];
            if (typeof key !== "string") {
                return false;
            }
            if (Array.isArray(value)) {
                for (var _c = 0, value_1 = value; _c < value_1.length; _c++) {
                    var val = value_1[_c];
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
function sprawdzDaneLiniiLotniczej(dane) {
    if (dane.hasOwnProperty("piloci") && dane.hasOwnProperty("lotniska")) {
        return sprawdzPilota(dane.piloci) && sprawdzLotnisko(dane.lotniska);
    }
    return false;
}
var daneLiniiLotniczej = JSON.parse(jsonString);
if (sprawdzDaneLiniiLotniczej(daneLiniiLotniczej)) {
    var juzNaPewnoDaneLinii = daneLiniiLotniczej;
    console.log(juzNaPewnoDaneLinii);
    console.log(juzNaPewnoDaneLinii.piloci.length);
}
function zaloguj() {
    var komunikaty = [];
    for (var _i = 0; _i < arguments.length; _i++) {
        komunikaty[_i] = arguments[_i];
    }
    console.log.apply(console, __spreadArrays(["Aleź skomplikowany program"], komunikaty));
}
zaloguj("Ja", "cię", "nie", "mogę");
