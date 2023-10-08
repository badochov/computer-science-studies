"use strict";
const passengers = document.getElementsByClassName("passenger");
Array.from(passengers).forEach((p) => {
    p.setAttribute('data-indentyfikator-pasazera', Math.random().toString());
});
let max_id = null;
let max_id_name = "";
Array.from(passengers).forEach((p) => {
    const id = p.getAttribute("data-indentyfikator-pasazera");
    if (!max_id || id > max_id) {
        if (p instanceof HTMLElement) {
            const span = p.getElementsByClassName("name")[0];
            max_id_name = span.textContent;
            max_id = id;
        }
    }
});
alert(max_id_name);
