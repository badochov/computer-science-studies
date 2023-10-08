const passengers = document.getElementsByClassName("passenger");

Array.from(passengers).forEach((p: Element) => {
    p.setAttribute('data-indentyfikator-pasazera', Math.random().toString())
})
let max_id: string | null = null;
let max_id_name = "";

Array.from(passengers).forEach((p: Element) => {
    const id = p.getAttribute("data-indentyfikator-pasazera") as string;
    if (!max_id || id > max_id) {
        if (p instanceof HTMLElement) {
            const span = p.getElementsByClassName("name")[0] as HTMLSpanElement;
            max_id_name = span.textContent as string;
            max_id = id;
        }
    }
});


alert(max_id_name)

