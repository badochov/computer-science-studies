// const submitButton = document.querySelector("input[type=submit]");
// if (submitButton instanceof HTMLInputElement) {
//     submitButton.style.display = "none";
// }
var nameInput = document.querySelector("input[name=fullname]");
if (nameInput instanceof HTMLInputElement) {
    nameInput.onchange = function () {
        console.log(nameInput.value);
    };
}
var dateInput = document.querySelector("input[name=date]");
if (dateInput instanceof HTMLInputElement) {
    dateInput.onchange = function () {
        console.log(dateInput.value);
    };
}
var p = document.getElementsByTagName("p")[0];
if (p instanceof HTMLParagraphElement) {
    p.innerHTML = "42";
}
var newEl = document.createElement("div");
newEl.innerText = "Nowy element";
document.body.appendChild(newEl);
var timeout = setTimeout(function () {
    console.log("No już wreszcie.");
}, 2000);
var form = document.querySelector("form#booking");
var getToday = function () {
    return new Date().toISOString().split("T")[0];
};
var validationErrorEl = document.querySelector("#validation-error");
if (validationErrorEl instanceof HTMLDivElement) {
    var closeValidationErrorButton = validationErrorEl.getElementsByTagName("button")[0];
    if (closeValidationErrorButton instanceof HTMLButtonElement) {
        closeValidationErrorButton.onclick = function () {
            validationErrorEl.style.display = "none";
        };
    }
}
var displayValidationErrorMsg = function (text) {
    if (validationErrorEl instanceof HTMLDivElement) {
        var validationP = validationErrorEl.getElementsByTagName("p")[0];
        if (validationP instanceof HTMLParagraphElement) {
            validationErrorEl.style.display = "block";
            validationP.innerText = text;
            return true;
        }
    }
    return false;
};
if (form instanceof HTMLFormElement) {
    var validateForm = function () {
        if (nameInput.value === "") {
            displayValidationErrorMsg("Imię i nazwisko nie może być puste");
            return false;
        }
        if (dateInput.value < getToday()) {
            displayValidationErrorMsg("Data nie może być wcześniejsza niż dzisiaj");
            return false;
        }
        return true;
    };
    form.onsubmit = validateForm;
}
