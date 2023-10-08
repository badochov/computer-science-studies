// const submitButton = document.querySelector("input[type=submit]");
// if (submitButton instanceof HTMLInputElement) {
//     submitButton.style.display = "none";
// }
import { fib } from './fib';

const nameInput = document.querySelector('input[name=fullname]') as HTMLInputElement;
if (nameInput instanceof HTMLInputElement) {
	nameInput.onchange = () => {
		console.log(nameInput.value);
	};
}

const dateInput = document.querySelector('input[name=date]') as HTMLInputElement;
if (dateInput instanceof HTMLInputElement) {
	dateInput.onchange = () => {
		console.log(dateInput.value);
	};
}

const p = document.getElementsByTagName('p')[0];
if (p instanceof HTMLParagraphElement) {
	p.innerHTML = '42';
}

const newEl = document.createElement('div');
newEl.innerText = 'Nowy element';
document.body.appendChild(newEl);

const timeout = setTimeout(() => {
	console.log('No już wreszcie.');
}, 2000);

const form = document.querySelector('form#booking');

const getToday = () => {
	return new Date().toISOString().split('T')[0];
};

const validationErrorEl = document.querySelector('#validation-error');

if (validationErrorEl instanceof HTMLDivElement) {
	const closeValidationErrorButton = validationErrorEl.getElementsByTagName('button')[0];
	if (closeValidationErrorButton instanceof HTMLButtonElement) {
		closeValidationErrorButton.onclick = () => {
			validationErrorEl.style.display = 'none';
		};
	}
}

const displayValidationErrorMsg = (text: string) => {
	if (validationErrorEl instanceof HTMLDivElement) {
		const validationP = validationErrorEl.getElementsByTagName('p')[0];
		if (validationP instanceof HTMLParagraphElement) {
			validationErrorEl.style.display = 'block';
			validationP.innerText = text;
			return true;
		}
	}
	return false;
};

if (form instanceof HTMLFormElement) {
	const submitButton = form.querySelector('input[type=submit]') as HTMLInputElement;
	submitButton.disabled = true;
	const validateForm = () => {
		if (nameInput.value === '') {
			displayValidationErrorMsg('Imię i nazwisko nie może być puste');
			return false;
		}
		if (dateInput.value < getToday()) {
			displayValidationErrorMsg('Data nie może być wcześniejsza niż dzisiaj');
			return false;
		}
		const fd = new FormData(form);
		const data: any = {};
		for (const [ key, value ] of fd.entries()) {
			console.log(key, value);
			data[key] = value;
		}
		return prompt('Czy podane dane są poprawne ?\n\n' + JSON.stringify(data));
	};

	form.onsubmit = validateForm;
	form.oninput = () => {
		if (nameInput.value.split(' ').length > 1 && dateInput.value >= getToday()) {
			for (const select of Array.from(form.querySelectorAll('select'))) {
				if ((<HTMLSelectElement>select).value === '') {
					return;
				}
			}
			submitButton.disabled = false;
		}
	};
}

const wait = (ms: number) => new Promise((resolve, reject) => window.setTimeout(resolve, ms));

async function teczoweKolory2(el: HTMLElement) {
	const colors = [ 'red', 'orange', 'yellow', 'green', 'blue', 'indigo', 'purple' ];
	let p = Promise.resolve();
	for (const color of colors) {
		p = p.then(() => wait(1000)).then(() => {
			el.style.backgroundColor = color;
		});
	}
}

async function teczoweKolory(el: HTMLElement) {
	const colors = [ 'red', 'orange', 'yellow', 'green', 'blue', 'indigo', 'purple' ];
	for (const color of colors) {
		await wait(1000);
		el.style.backgroundColor = color;
	}
}
teczoweKolory(form as HTMLElement);

const getLatestCommitAuthor = async (url: string) => {
	const raw = await fetch(url);
	const json = await raw.json();
	const authorURL = json[0].author.url;
	const rawAuthor = await fetch(authorURL);
	const authorJson = await rawAuthor.json();
	return authorJson;
};

const setLatestCommitAuthorProfilePicture = (author: any) => {
	const avatarURL = author.avatar_url;
	const ts = document.getElementById('ts') as HTMLDivElement;
	const img = document.createElement('img');
	img.setAttribute('alt', 'Latest TS Commit Author');
	img.setAttribute('src', avatarURL);
	ts.appendChild(img);
};

const getAuthorsRepos = async (author: any) => {
	const raw = await fetch(author.repos_url);
	const json = await raw.json();
	const repos = [];
	for (const [ i, repo ] of Object.entries(json)) {
		repos.push((<any>repo)['name']);
	}

	return repos.sort((a: string, b: string) => a.toLowerCase().localeCompare(b.toLowerCase()));
};

const latestTSCommit = async (url: string) => {
	const author = await getLatestCommitAuthor(url);
	setLatestCommitAuthorProfilePicture(author);
	console.table(await getAuthorsRepos(author));
};

latestTSCommit('https://api.github.com/repos/Microsoft/TypeScript/commits');

const main = document.getElementById('main-page') as HTMLElement;
const delayed = document.getElementById('delayed') as HTMLElement;
const booking = document.getElementById('booking') as HTMLElement;
let clickNumber = 0;
booking.onclick = (e) => {
	e.stopPropagation();
};
main.onclick = (e) => {
	const el = e.target as HTMLElement;
	fib(10 ** clickNumber);
	clickNumber++;
	if (delayed.contains(el)) {
		el.style.backgroundColor = getRandomColor();
	}
};

function getRandomColor() {
	const letters = '0123456789ABCDEF';
	let color = '#';
	for (var i = 0; i < 6; i++) {
		color += letters[Math.floor(Math.random() * 16)];
	}
	return color;
}
