// const submitButton = document.querySelector("input[type=submit]");
// if (submitButton instanceof HTMLInputElement) {
//     submitButton.style.display = "none";
// }

const nameInput = document.querySelector("input[name=fullname]") as HTMLInputElement;
if (nameInput instanceof HTMLInputElement) {
    nameInput.onchange = () => {
        console.log(nameInput.value);
    }
}

const dateInput = document.querySelector("input[name=date]") as HTMLInputElement;
if (dateInput instanceof HTMLInputElement) {
    dateInput.onchange = () => {
        console.log(dateInput.value);
    }
}

const p = document.getElementsByTagName("p")[0];
if (p instanceof HTMLParagraphElement) {
    p.innerHTML = "42";
}

const newEl = document.createElement("div");
newEl.innerText = "Nowy element"
document.body.appendChild(newEl);

const timeout = setTimeout(() => {
    console.log("No już wreszcie.");
}, 2000);

const form = document.querySelector("form#booking");

const getToday = () => {
    return new Date().toISOString().split("T")[0];
}

const validationErrorEl = document.querySelector("#validation-error");

if (validationErrorEl instanceof HTMLDivElement) {
    const closeValidationErrorButton = validationErrorEl.getElementsByTagName("button")[0];
    if (closeValidationErrorButton instanceof HTMLButtonElement) {
        closeValidationErrorButton.onclick = () => {
            validationErrorEl.style.display = "none";
        }
    }
}

const displayValidationErrorMsg = (text: string) => {
    if (validationErrorEl instanceof HTMLDivElement) {
        const validationP = validationErrorEl.getElementsByTagName("p")[0];
        if (validationP instanceof HTMLParagraphElement) {

            validationErrorEl.style.display = "block";
            validationP.innerText = text;
            return true;
        }
    }
    return false;
}

if (form instanceof HTMLFormElement) {
    const validateForm = () => {
        if (nameInput.value === "") {
            displayValidationErrorMsg("Imię i nazwisko nie może być puste");
            return false;
        }
        if (dateInput.value < getToday()) {
            displayValidationErrorMsg("Data nie może być wcześniejsza niż dzisiaj")
            return false;
        }
        return true;
    }

    form.onsubmit = validateForm;
}