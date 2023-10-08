/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = "./js/main.ts");
/******/ })
/************************************************************************/
/******/ ({

/***/ "./js/fib.ts":
/*!*******************!*\
  !*** ./js/fib.ts ***!
  \*******************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
function fib(n) {
    return __awaiter(this, void 0, void 0, function* () {
        if (n < 2) {
            return Promise.resolve(n);
        }
        return (yield fib(n - 1)) + (yield fib(n - 2));
    });
}
exports.fib = fib;
function fib_sync(n) {
    if (n < 2) {
        return n;
    }
    return fib_sync(n - 1) + fib_sync(n - 2);
}
exports.fib_sync = fib_sync;


/***/ }),

/***/ "./js/main.ts":
/*!********************!*\
  !*** ./js/main.ts ***!
  \********************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
// const submitButton = document.querySelector("input[type=submit]");
// if (submitButton instanceof HTMLInputElement) {
//     submitButton.style.display = "none";
// }
const fib_1 = __webpack_require__(/*! ./fib */ "./js/fib.ts");
const nameInput = document.querySelector('input[name=fullname]');
if (nameInput instanceof HTMLInputElement) {
    nameInput.onchange = () => {
        console.log(nameInput.value);
    };
}
const dateInput = document.querySelector('input[name=date]');
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
const displayValidationErrorMsg = (text) => {
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
    const submitButton = form.querySelector('input[type=submit]');
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
        const data = {};
        for (const [key, value] of fd.entries()) {
            console.log(key, value);
            data[key] = value;
        }
        return prompt('Czy podane dane są poprawne ?\n\n' + JSON.stringify(data));
    };
    form.onsubmit = validateForm;
    form.oninput = () => {
        if (nameInput.value.split(' ').length > 1 && dateInput.value >= getToday()) {
            for (const select of Array.from(form.querySelectorAll('select'))) {
                if (select.value === '') {
                    return;
                }
            }
            submitButton.disabled = false;
        }
    };
}
const wait = (ms) => new Promise((resolve, reject) => window.setTimeout(resolve, ms));
function teczoweKolory2(el) {
    return __awaiter(this, void 0, void 0, function* () {
        const colors = ['red', 'orange', 'yellow', 'green', 'blue', 'indigo', 'purple'];
        let p = Promise.resolve();
        for (const color of colors) {
            p = p.then(() => wait(1000)).then(() => {
                el.style.backgroundColor = color;
            });
        }
    });
}
function teczoweKolory(el) {
    return __awaiter(this, void 0, void 0, function* () {
        const colors = ['red', 'orange', 'yellow', 'green', 'blue', 'indigo', 'purple'];
        for (const color of colors) {
            yield wait(1000);
            el.style.backgroundColor = color;
        }
    });
}
teczoweKolory(form);
const getLatestCommitAuthor = (url) => __awaiter(void 0, void 0, void 0, function* () {
    const raw = yield fetch(url);
    const json = yield raw.json();
    const authorURL = json[0].author.url;
    const rawAuthor = yield fetch(authorURL);
    const authorJson = yield rawAuthor.json();
    return authorJson;
});
const setLatestCommitAuthorProfilePicture = (author) => {
    const avatarURL = author.avatar_url;
    const ts = document.getElementById('ts');
    const img = document.createElement('img');
    img.setAttribute('alt', 'Latest TS Commit Author');
    img.setAttribute('src', avatarURL);
    ts.appendChild(img);
};
const getAuthorsRepos = (author) => __awaiter(void 0, void 0, void 0, function* () {
    const raw = yield fetch(author.repos_url);
    const json = yield raw.json();
    const repos = [];
    for (const [i, repo] of Object.entries(json)) {
        repos.push(repo['name']);
    }
    return repos.sort((a, b) => a.toLowerCase().localeCompare(b.toLowerCase()));
});
const latestTSCommit = (url) => __awaiter(void 0, void 0, void 0, function* () {
    const author = yield getLatestCommitAuthor(url);
    setLatestCommitAuthorProfilePicture(author);
    console.table(yield getAuthorsRepos(author));
});
latestTSCommit('https://api.github.com/repos/Microsoft/TypeScript/commits');
const main = document.getElementById('main-page');
const delayed = document.getElementById('delayed');
const booking = document.getElementById('booking');
let clickNumber = 0;
booking.onclick = (e) => {
    e.stopPropagation();
};
main.onclick = (e) => {
    const el = e.target;
    fib_1.fib(Math.pow(10, clickNumber));
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


/***/ })

/******/ });
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vd2VicGFjay9ib290c3RyYXAiLCJ3ZWJwYWNrOi8vLy4vanMvZmliLnRzIiwid2VicGFjazovLy8uL2pzL21haW4udHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IjtRQUFBO1FBQ0E7O1FBRUE7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTs7UUFFQTtRQUNBOztRQUVBO1FBQ0E7O1FBRUE7UUFDQTtRQUNBOzs7UUFHQTtRQUNBOztRQUVBO1FBQ0E7O1FBRUE7UUFDQTtRQUNBO1FBQ0EsMENBQTBDLGdDQUFnQztRQUMxRTtRQUNBOztRQUVBO1FBQ0E7UUFDQTtRQUNBLHdEQUF3RCxrQkFBa0I7UUFDMUU7UUFDQSxpREFBaUQsY0FBYztRQUMvRDs7UUFFQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0EseUNBQXlDLGlDQUFpQztRQUMxRSxnSEFBZ0gsbUJBQW1CLEVBQUU7UUFDckk7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7UUFDQSwyQkFBMkIsMEJBQTBCLEVBQUU7UUFDdkQsaUNBQWlDLGVBQWU7UUFDaEQ7UUFDQTtRQUNBOztRQUVBO1FBQ0Esc0RBQXNELCtEQUErRDs7UUFFckg7UUFDQTs7O1FBR0E7UUFDQTs7Ozs7Ozs7Ozs7OztBQ2xGYTtBQUNiO0FBQ0EsMkJBQTJCLCtEQUErRCxnQkFBZ0IsRUFBRSxFQUFFO0FBQzlHO0FBQ0EsbUNBQW1DLE1BQU0sNkJBQTZCLEVBQUUsWUFBWSxXQUFXLEVBQUU7QUFDakcsa0NBQWtDLE1BQU0saUNBQWlDLEVBQUUsWUFBWSxXQUFXLEVBQUU7QUFDcEcsK0JBQStCLHFGQUFxRjtBQUNwSDtBQUNBLEtBQUs7QUFDTDtBQUNBLDhDQUE4QyxjQUFjO0FBQzVEO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBLEtBQUs7QUFDTDtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7Ozs7Ozs7Ozs7Ozs7QUMxQmE7QUFDYjtBQUNBLDJCQUEyQiwrREFBK0QsZ0JBQWdCLEVBQUUsRUFBRTtBQUM5RztBQUNBLG1DQUFtQyxNQUFNLDZCQUE2QixFQUFFLFlBQVksV0FBVyxFQUFFO0FBQ2pHLGtDQUFrQyxNQUFNLGlDQUFpQyxFQUFFLFlBQVksV0FBVyxFQUFFO0FBQ3BHLCtCQUErQixxRkFBcUY7QUFDcEg7QUFDQSxLQUFLO0FBQ0w7QUFDQSw4Q0FBOEMsY0FBYztBQUM1RDtBQUNBO0FBQ0E7QUFDQTtBQUNBLGNBQWMsbUJBQU8sQ0FBQywwQkFBTztBQUM3QjtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQSxDQUFDO0FBQ0Q7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQSxhQUFhO0FBQ2I7QUFDQSxLQUFLO0FBQ0w7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBLEtBQUs7QUFDTDtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQSxDQUFDO0FBQ0Q7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQSxDQUFDO0FBQ0Q7QUFDQTtBQUNBO0FBQ0E7QUFDQSxDQUFDO0FBQ0Q7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQSxtQkFBbUIsT0FBTztBQUMxQjtBQUNBO0FBQ0E7QUFDQSIsImZpbGUiOiJtYWluLmpzIiwic291cmNlc0NvbnRlbnQiOlsiIFx0Ly8gVGhlIG1vZHVsZSBjYWNoZVxuIFx0dmFyIGluc3RhbGxlZE1vZHVsZXMgPSB7fTtcblxuIFx0Ly8gVGhlIHJlcXVpcmUgZnVuY3Rpb25cbiBcdGZ1bmN0aW9uIF9fd2VicGFja19yZXF1aXJlX18obW9kdWxlSWQpIHtcblxuIFx0XHQvLyBDaGVjayBpZiBtb2R1bGUgaXMgaW4gY2FjaGVcbiBcdFx0aWYoaW5zdGFsbGVkTW9kdWxlc1ttb2R1bGVJZF0pIHtcbiBcdFx0XHRyZXR1cm4gaW5zdGFsbGVkTW9kdWxlc1ttb2R1bGVJZF0uZXhwb3J0cztcbiBcdFx0fVxuIFx0XHQvLyBDcmVhdGUgYSBuZXcgbW9kdWxlIChhbmQgcHV0IGl0IGludG8gdGhlIGNhY2hlKVxuIFx0XHR2YXIgbW9kdWxlID0gaW5zdGFsbGVkTW9kdWxlc1ttb2R1bGVJZF0gPSB7XG4gXHRcdFx0aTogbW9kdWxlSWQsXG4gXHRcdFx0bDogZmFsc2UsXG4gXHRcdFx0ZXhwb3J0czoge31cbiBcdFx0fTtcblxuIFx0XHQvLyBFeGVjdXRlIHRoZSBtb2R1bGUgZnVuY3Rpb25cbiBcdFx0bW9kdWxlc1ttb2R1bGVJZF0uY2FsbChtb2R1bGUuZXhwb3J0cywgbW9kdWxlLCBtb2R1bGUuZXhwb3J0cywgX193ZWJwYWNrX3JlcXVpcmVfXyk7XG5cbiBcdFx0Ly8gRmxhZyB0aGUgbW9kdWxlIGFzIGxvYWRlZFxuIFx0XHRtb2R1bGUubCA9IHRydWU7XG5cbiBcdFx0Ly8gUmV0dXJuIHRoZSBleHBvcnRzIG9mIHRoZSBtb2R1bGVcbiBcdFx0cmV0dXJuIG1vZHVsZS5leHBvcnRzO1xuIFx0fVxuXG5cbiBcdC8vIGV4cG9zZSB0aGUgbW9kdWxlcyBvYmplY3QgKF9fd2VicGFja19tb2R1bGVzX18pXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLm0gPSBtb2R1bGVzO1xuXG4gXHQvLyBleHBvc2UgdGhlIG1vZHVsZSBjYWNoZVxuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5jID0gaW5zdGFsbGVkTW9kdWxlcztcblxuIFx0Ly8gZGVmaW5lIGdldHRlciBmdW5jdGlvbiBmb3IgaGFybW9ueSBleHBvcnRzXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLmQgPSBmdW5jdGlvbihleHBvcnRzLCBuYW1lLCBnZXR0ZXIpIHtcbiBcdFx0aWYoIV9fd2VicGFja19yZXF1aXJlX18ubyhleHBvcnRzLCBuYW1lKSkge1xuIFx0XHRcdE9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBuYW1lLCB7IGVudW1lcmFibGU6IHRydWUsIGdldDogZ2V0dGVyIH0pO1xuIFx0XHR9XG4gXHR9O1xuXG4gXHQvLyBkZWZpbmUgX19lc01vZHVsZSBvbiBleHBvcnRzXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLnIgPSBmdW5jdGlvbihleHBvcnRzKSB7XG4gXHRcdGlmKHR5cGVvZiBTeW1ib2wgIT09ICd1bmRlZmluZWQnICYmIFN5bWJvbC50b1N0cmluZ1RhZykge1xuIFx0XHRcdE9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBTeW1ib2wudG9TdHJpbmdUYWcsIHsgdmFsdWU6ICdNb2R1bGUnIH0pO1xuIFx0XHR9XG4gXHRcdE9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCAnX19lc01vZHVsZScsIHsgdmFsdWU6IHRydWUgfSk7XG4gXHR9O1xuXG4gXHQvLyBjcmVhdGUgYSBmYWtlIG5hbWVzcGFjZSBvYmplY3RcbiBcdC8vIG1vZGUgJiAxOiB2YWx1ZSBpcyBhIG1vZHVsZSBpZCwgcmVxdWlyZSBpdFxuIFx0Ly8gbW9kZSAmIDI6IG1lcmdlIGFsbCBwcm9wZXJ0aWVzIG9mIHZhbHVlIGludG8gdGhlIG5zXG4gXHQvLyBtb2RlICYgNDogcmV0dXJuIHZhbHVlIHdoZW4gYWxyZWFkeSBucyBvYmplY3RcbiBcdC8vIG1vZGUgJiA4fDE6IGJlaGF2ZSBsaWtlIHJlcXVpcmVcbiBcdF9fd2VicGFja19yZXF1aXJlX18udCA9IGZ1bmN0aW9uKHZhbHVlLCBtb2RlKSB7XG4gXHRcdGlmKG1vZGUgJiAxKSB2YWx1ZSA9IF9fd2VicGFja19yZXF1aXJlX18odmFsdWUpO1xuIFx0XHRpZihtb2RlICYgOCkgcmV0dXJuIHZhbHVlO1xuIFx0XHRpZigobW9kZSAmIDQpICYmIHR5cGVvZiB2YWx1ZSA9PT0gJ29iamVjdCcgJiYgdmFsdWUgJiYgdmFsdWUuX19lc01vZHVsZSkgcmV0dXJuIHZhbHVlO1xuIFx0XHR2YXIgbnMgPSBPYmplY3QuY3JlYXRlKG51bGwpO1xuIFx0XHRfX3dlYnBhY2tfcmVxdWlyZV9fLnIobnMpO1xuIFx0XHRPYmplY3QuZGVmaW5lUHJvcGVydHkobnMsICdkZWZhdWx0JywgeyBlbnVtZXJhYmxlOiB0cnVlLCB2YWx1ZTogdmFsdWUgfSk7XG4gXHRcdGlmKG1vZGUgJiAyICYmIHR5cGVvZiB2YWx1ZSAhPSAnc3RyaW5nJykgZm9yKHZhciBrZXkgaW4gdmFsdWUpIF9fd2VicGFja19yZXF1aXJlX18uZChucywga2V5LCBmdW5jdGlvbihrZXkpIHsgcmV0dXJuIHZhbHVlW2tleV07IH0uYmluZChudWxsLCBrZXkpKTtcbiBcdFx0cmV0dXJuIG5zO1xuIFx0fTtcblxuIFx0Ly8gZ2V0RGVmYXVsdEV4cG9ydCBmdW5jdGlvbiBmb3IgY29tcGF0aWJpbGl0eSB3aXRoIG5vbi1oYXJtb255IG1vZHVsZXNcbiBcdF9fd2VicGFja19yZXF1aXJlX18ubiA9IGZ1bmN0aW9uKG1vZHVsZSkge1xuIFx0XHR2YXIgZ2V0dGVyID0gbW9kdWxlICYmIG1vZHVsZS5fX2VzTW9kdWxlID9cbiBcdFx0XHRmdW5jdGlvbiBnZXREZWZhdWx0KCkgeyByZXR1cm4gbW9kdWxlWydkZWZhdWx0J107IH0gOlxuIFx0XHRcdGZ1bmN0aW9uIGdldE1vZHVsZUV4cG9ydHMoKSB7IHJldHVybiBtb2R1bGU7IH07XG4gXHRcdF9fd2VicGFja19yZXF1aXJlX18uZChnZXR0ZXIsICdhJywgZ2V0dGVyKTtcbiBcdFx0cmV0dXJuIGdldHRlcjtcbiBcdH07XG5cbiBcdC8vIE9iamVjdC5wcm90b3R5cGUuaGFzT3duUHJvcGVydHkuY2FsbFxuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5vID0gZnVuY3Rpb24ob2JqZWN0LCBwcm9wZXJ0eSkgeyByZXR1cm4gT2JqZWN0LnByb3RvdHlwZS5oYXNPd25Qcm9wZXJ0eS5jYWxsKG9iamVjdCwgcHJvcGVydHkpOyB9O1xuXG4gXHQvLyBfX3dlYnBhY2tfcHVibGljX3BhdGhfX1xuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5wID0gXCJcIjtcblxuXG4gXHQvLyBMb2FkIGVudHJ5IG1vZHVsZSBhbmQgcmV0dXJuIGV4cG9ydHNcbiBcdHJldHVybiBfX3dlYnBhY2tfcmVxdWlyZV9fKF9fd2VicGFja19yZXF1aXJlX18ucyA9IFwiLi9qcy9tYWluLnRzXCIpO1xuIiwiXCJ1c2Ugc3RyaWN0XCI7XG52YXIgX19hd2FpdGVyID0gKHRoaXMgJiYgdGhpcy5fX2F3YWl0ZXIpIHx8IGZ1bmN0aW9uICh0aGlzQXJnLCBfYXJndW1lbnRzLCBQLCBnZW5lcmF0b3IpIHtcbiAgICBmdW5jdGlvbiBhZG9wdCh2YWx1ZSkgeyByZXR1cm4gdmFsdWUgaW5zdGFuY2VvZiBQID8gdmFsdWUgOiBuZXcgUChmdW5jdGlvbiAocmVzb2x2ZSkgeyByZXNvbHZlKHZhbHVlKTsgfSk7IH1cbiAgICByZXR1cm4gbmV3IChQIHx8IChQID0gUHJvbWlzZSkpKGZ1bmN0aW9uIChyZXNvbHZlLCByZWplY3QpIHtcbiAgICAgICAgZnVuY3Rpb24gZnVsZmlsbGVkKHZhbHVlKSB7IHRyeSB7IHN0ZXAoZ2VuZXJhdG9yLm5leHQodmFsdWUpKTsgfSBjYXRjaCAoZSkgeyByZWplY3QoZSk7IH0gfVxuICAgICAgICBmdW5jdGlvbiByZWplY3RlZCh2YWx1ZSkgeyB0cnkgeyBzdGVwKGdlbmVyYXRvcltcInRocm93XCJdKHZhbHVlKSk7IH0gY2F0Y2ggKGUpIHsgcmVqZWN0KGUpOyB9IH1cbiAgICAgICAgZnVuY3Rpb24gc3RlcChyZXN1bHQpIHsgcmVzdWx0LmRvbmUgPyByZXNvbHZlKHJlc3VsdC52YWx1ZSkgOiBhZG9wdChyZXN1bHQudmFsdWUpLnRoZW4oZnVsZmlsbGVkLCByZWplY3RlZCk7IH1cbiAgICAgICAgc3RlcCgoZ2VuZXJhdG9yID0gZ2VuZXJhdG9yLmFwcGx5KHRoaXNBcmcsIF9hcmd1bWVudHMgfHwgW10pKS5uZXh0KCkpO1xuICAgIH0pO1xufTtcbk9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBcIl9fZXNNb2R1bGVcIiwgeyB2YWx1ZTogdHJ1ZSB9KTtcbmZ1bmN0aW9uIGZpYihuKSB7XG4gICAgcmV0dXJuIF9fYXdhaXRlcih0aGlzLCB2b2lkIDAsIHZvaWQgMCwgZnVuY3Rpb24qICgpIHtcbiAgICAgICAgaWYgKG4gPCAyKSB7XG4gICAgICAgICAgICByZXR1cm4gUHJvbWlzZS5yZXNvbHZlKG4pO1xuICAgICAgICB9XG4gICAgICAgIHJldHVybiAoeWllbGQgZmliKG4gLSAxKSkgKyAoeWllbGQgZmliKG4gLSAyKSk7XG4gICAgfSk7XG59XG5leHBvcnRzLmZpYiA9IGZpYjtcbmZ1bmN0aW9uIGZpYl9zeW5jKG4pIHtcbiAgICBpZiAobiA8IDIpIHtcbiAgICAgICAgcmV0dXJuIG47XG4gICAgfVxuICAgIHJldHVybiBmaWJfc3luYyhuIC0gMSkgKyBmaWJfc3luYyhuIC0gMik7XG59XG5leHBvcnRzLmZpYl9zeW5jID0gZmliX3N5bmM7XG4iLCJcInVzZSBzdHJpY3RcIjtcbnZhciBfX2F3YWl0ZXIgPSAodGhpcyAmJiB0aGlzLl9fYXdhaXRlcikgfHwgZnVuY3Rpb24gKHRoaXNBcmcsIF9hcmd1bWVudHMsIFAsIGdlbmVyYXRvcikge1xuICAgIGZ1bmN0aW9uIGFkb3B0KHZhbHVlKSB7IHJldHVybiB2YWx1ZSBpbnN0YW5jZW9mIFAgPyB2YWx1ZSA6IG5ldyBQKGZ1bmN0aW9uIChyZXNvbHZlKSB7IHJlc29sdmUodmFsdWUpOyB9KTsgfVxuICAgIHJldHVybiBuZXcgKFAgfHwgKFAgPSBQcm9taXNlKSkoZnVuY3Rpb24gKHJlc29sdmUsIHJlamVjdCkge1xuICAgICAgICBmdW5jdGlvbiBmdWxmaWxsZWQodmFsdWUpIHsgdHJ5IHsgc3RlcChnZW5lcmF0b3IubmV4dCh2YWx1ZSkpOyB9IGNhdGNoIChlKSB7IHJlamVjdChlKTsgfSB9XG4gICAgICAgIGZ1bmN0aW9uIHJlamVjdGVkKHZhbHVlKSB7IHRyeSB7IHN0ZXAoZ2VuZXJhdG9yW1widGhyb3dcIl0odmFsdWUpKTsgfSBjYXRjaCAoZSkgeyByZWplY3QoZSk7IH0gfVxuICAgICAgICBmdW5jdGlvbiBzdGVwKHJlc3VsdCkgeyByZXN1bHQuZG9uZSA/IHJlc29sdmUocmVzdWx0LnZhbHVlKSA6IGFkb3B0KHJlc3VsdC52YWx1ZSkudGhlbihmdWxmaWxsZWQsIHJlamVjdGVkKTsgfVxuICAgICAgICBzdGVwKChnZW5lcmF0b3IgPSBnZW5lcmF0b3IuYXBwbHkodGhpc0FyZywgX2FyZ3VtZW50cyB8fCBbXSkpLm5leHQoKSk7XG4gICAgfSk7XG59O1xuT2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIFwiX19lc01vZHVsZVwiLCB7IHZhbHVlOiB0cnVlIH0pO1xuLy8gY29uc3Qgc3VibWl0QnV0dG9uID0gZG9jdW1lbnQucXVlcnlTZWxlY3RvcihcImlucHV0W3R5cGU9c3VibWl0XVwiKTtcbi8vIGlmIChzdWJtaXRCdXR0b24gaW5zdGFuY2VvZiBIVE1MSW5wdXRFbGVtZW50KSB7XG4vLyAgICAgc3VibWl0QnV0dG9uLnN0eWxlLmRpc3BsYXkgPSBcIm5vbmVcIjtcbi8vIH1cbmNvbnN0IGZpYl8xID0gcmVxdWlyZShcIi4vZmliXCIpO1xuY29uc3QgbmFtZUlucHV0ID0gZG9jdW1lbnQucXVlcnlTZWxlY3RvcignaW5wdXRbbmFtZT1mdWxsbmFtZV0nKTtcbmlmIChuYW1lSW5wdXQgaW5zdGFuY2VvZiBIVE1MSW5wdXRFbGVtZW50KSB7XG4gICAgbmFtZUlucHV0Lm9uY2hhbmdlID0gKCkgPT4ge1xuICAgICAgICBjb25zb2xlLmxvZyhuYW1lSW5wdXQudmFsdWUpO1xuICAgIH07XG59XG5jb25zdCBkYXRlSW5wdXQgPSBkb2N1bWVudC5xdWVyeVNlbGVjdG9yKCdpbnB1dFtuYW1lPWRhdGVdJyk7XG5pZiAoZGF0ZUlucHV0IGluc3RhbmNlb2YgSFRNTElucHV0RWxlbWVudCkge1xuICAgIGRhdGVJbnB1dC5vbmNoYW5nZSA9ICgpID0+IHtcbiAgICAgICAgY29uc29sZS5sb2coZGF0ZUlucHV0LnZhbHVlKTtcbiAgICB9O1xufVxuY29uc3QgcCA9IGRvY3VtZW50LmdldEVsZW1lbnRzQnlUYWdOYW1lKCdwJylbMF07XG5pZiAocCBpbnN0YW5jZW9mIEhUTUxQYXJhZ3JhcGhFbGVtZW50KSB7XG4gICAgcC5pbm5lckhUTUwgPSAnNDInO1xufVxuY29uc3QgbmV3RWwgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbm5ld0VsLmlubmVyVGV4dCA9ICdOb3d5IGVsZW1lbnQnO1xuZG9jdW1lbnQuYm9keS5hcHBlbmRDaGlsZChuZXdFbCk7XG5jb25zdCB0aW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgY29uc29sZS5sb2coJ05vIGp1xbwgd3Jlc3pjaWUuJyk7XG59LCAyMDAwKTtcbmNvbnN0IGZvcm0gPSBkb2N1bWVudC5xdWVyeVNlbGVjdG9yKCdmb3JtI2Jvb2tpbmcnKTtcbmNvbnN0IGdldFRvZGF5ID0gKCkgPT4ge1xuICAgIHJldHVybiBuZXcgRGF0ZSgpLnRvSVNPU3RyaW5nKCkuc3BsaXQoJ1QnKVswXTtcbn07XG5jb25zdCB2YWxpZGF0aW9uRXJyb3JFbCA9IGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3IoJyN2YWxpZGF0aW9uLWVycm9yJyk7XG5pZiAodmFsaWRhdGlvbkVycm9yRWwgaW5zdGFuY2VvZiBIVE1MRGl2RWxlbWVudCkge1xuICAgIGNvbnN0IGNsb3NlVmFsaWRhdGlvbkVycm9yQnV0dG9uID0gdmFsaWRhdGlvbkVycm9yRWwuZ2V0RWxlbWVudHNCeVRhZ05hbWUoJ2J1dHRvbicpWzBdO1xuICAgIGlmIChjbG9zZVZhbGlkYXRpb25FcnJvckJ1dHRvbiBpbnN0YW5jZW9mIEhUTUxCdXR0b25FbGVtZW50KSB7XG4gICAgICAgIGNsb3NlVmFsaWRhdGlvbkVycm9yQnV0dG9uLm9uY2xpY2sgPSAoKSA9PiB7XG4gICAgICAgICAgICB2YWxpZGF0aW9uRXJyb3JFbC5zdHlsZS5kaXNwbGF5ID0gJ25vbmUnO1xuICAgICAgICB9O1xuICAgIH1cbn1cbmNvbnN0IGRpc3BsYXlWYWxpZGF0aW9uRXJyb3JNc2cgPSAodGV4dCkgPT4ge1xuICAgIGlmICh2YWxpZGF0aW9uRXJyb3JFbCBpbnN0YW5jZW9mIEhUTUxEaXZFbGVtZW50KSB7XG4gICAgICAgIGNvbnN0IHZhbGlkYXRpb25QID0gdmFsaWRhdGlvbkVycm9yRWwuZ2V0RWxlbWVudHNCeVRhZ05hbWUoJ3AnKVswXTtcbiAgICAgICAgaWYgKHZhbGlkYXRpb25QIGluc3RhbmNlb2YgSFRNTFBhcmFncmFwaEVsZW1lbnQpIHtcbiAgICAgICAgICAgIHZhbGlkYXRpb25FcnJvckVsLnN0eWxlLmRpc3BsYXkgPSAnYmxvY2snO1xuICAgICAgICAgICAgdmFsaWRhdGlvblAuaW5uZXJUZXh0ID0gdGV4dDtcbiAgICAgICAgICAgIHJldHVybiB0cnVlO1xuICAgICAgICB9XG4gICAgfVxuICAgIHJldHVybiBmYWxzZTtcbn07XG5pZiAoZm9ybSBpbnN0YW5jZW9mIEhUTUxGb3JtRWxlbWVudCkge1xuICAgIGNvbnN0IHN1Ym1pdEJ1dHRvbiA9IGZvcm0ucXVlcnlTZWxlY3RvcignaW5wdXRbdHlwZT1zdWJtaXRdJyk7XG4gICAgc3VibWl0QnV0dG9uLmRpc2FibGVkID0gdHJ1ZTtcbiAgICBjb25zdCB2YWxpZGF0ZUZvcm0gPSAoKSA9PiB7XG4gICAgICAgIGlmIChuYW1lSW5wdXQudmFsdWUgPT09ICcnKSB7XG4gICAgICAgICAgICBkaXNwbGF5VmFsaWRhdGlvbkVycm9yTXNnKCdJbWnEmSBpIG5hendpc2tvIG5pZSBtb8W8ZSBiecSHIHB1c3RlJyk7XG4gICAgICAgICAgICByZXR1cm4gZmFsc2U7XG4gICAgICAgIH1cbiAgICAgICAgaWYgKGRhdGVJbnB1dC52YWx1ZSA8IGdldFRvZGF5KCkpIHtcbiAgICAgICAgICAgIGRpc3BsYXlWYWxpZGF0aW9uRXJyb3JNc2coJ0RhdGEgbmllIG1vxbxlIGJ5xIcgd2N6ZcWbbmllanN6YSBuacW8IGR6aXNpYWonKTtcbiAgICAgICAgICAgIHJldHVybiBmYWxzZTtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCBmZCA9IG5ldyBGb3JtRGF0YShmb3JtKTtcbiAgICAgICAgY29uc3QgZGF0YSA9IHt9O1xuICAgICAgICBmb3IgKGNvbnN0IFtrZXksIHZhbHVlXSBvZiBmZC5lbnRyaWVzKCkpIHtcbiAgICAgICAgICAgIGNvbnNvbGUubG9nKGtleSwgdmFsdWUpO1xuICAgICAgICAgICAgZGF0YVtrZXldID0gdmFsdWU7XG4gICAgICAgIH1cbiAgICAgICAgcmV0dXJuIHByb21wdCgnQ3p5IHBvZGFuZSBkYW5lIHPEhSBwb3ByYXduZSA/XFxuXFxuJyArIEpTT04uc3RyaW5naWZ5KGRhdGEpKTtcbiAgICB9O1xuICAgIGZvcm0ub25zdWJtaXQgPSB2YWxpZGF0ZUZvcm07XG4gICAgZm9ybS5vbmlucHV0ID0gKCkgPT4ge1xuICAgICAgICBpZiAobmFtZUlucHV0LnZhbHVlLnNwbGl0KCcgJykubGVuZ3RoID4gMSAmJiBkYXRlSW5wdXQudmFsdWUgPj0gZ2V0VG9kYXkoKSkge1xuICAgICAgICAgICAgZm9yIChjb25zdCBzZWxlY3Qgb2YgQXJyYXkuZnJvbShmb3JtLnF1ZXJ5U2VsZWN0b3JBbGwoJ3NlbGVjdCcpKSkge1xuICAgICAgICAgICAgICAgIGlmIChzZWxlY3QudmFsdWUgPT09ICcnKSB7XG4gICAgICAgICAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgICAgICAgICB9XG4gICAgICAgICAgICB9XG4gICAgICAgICAgICBzdWJtaXRCdXR0b24uZGlzYWJsZWQgPSBmYWxzZTtcbiAgICAgICAgfVxuICAgIH07XG59XG5jb25zdCB3YWl0ID0gKG1zKSA9PiBuZXcgUHJvbWlzZSgocmVzb2x2ZSwgcmVqZWN0KSA9PiB3aW5kb3cuc2V0VGltZW91dChyZXNvbHZlLCBtcykpO1xuZnVuY3Rpb24gdGVjem93ZUtvbG9yeTIoZWwpIHtcbiAgICByZXR1cm4gX19hd2FpdGVyKHRoaXMsIHZvaWQgMCwgdm9pZCAwLCBmdW5jdGlvbiogKCkge1xuICAgICAgICBjb25zdCBjb2xvcnMgPSBbJ3JlZCcsICdvcmFuZ2UnLCAneWVsbG93JywgJ2dyZWVuJywgJ2JsdWUnLCAnaW5kaWdvJywgJ3B1cnBsZSddO1xuICAgICAgICBsZXQgcCA9IFByb21pc2UucmVzb2x2ZSgpO1xuICAgICAgICBmb3IgKGNvbnN0IGNvbG9yIG9mIGNvbG9ycykge1xuICAgICAgICAgICAgcCA9IHAudGhlbigoKSA9PiB3YWl0KDEwMDApKS50aGVuKCgpID0+IHtcbiAgICAgICAgICAgICAgICBlbC5zdHlsZS5iYWNrZ3JvdW5kQ29sb3IgPSBjb2xvcjtcbiAgICAgICAgICAgIH0pO1xuICAgICAgICB9XG4gICAgfSk7XG59XG5mdW5jdGlvbiB0ZWN6b3dlS29sb3J5KGVsKSB7XG4gICAgcmV0dXJuIF9fYXdhaXRlcih0aGlzLCB2b2lkIDAsIHZvaWQgMCwgZnVuY3Rpb24qICgpIHtcbiAgICAgICAgY29uc3QgY29sb3JzID0gWydyZWQnLCAnb3JhbmdlJywgJ3llbGxvdycsICdncmVlbicsICdibHVlJywgJ2luZGlnbycsICdwdXJwbGUnXTtcbiAgICAgICAgZm9yIChjb25zdCBjb2xvciBvZiBjb2xvcnMpIHtcbiAgICAgICAgICAgIHlpZWxkIHdhaXQoMTAwMCk7XG4gICAgICAgICAgICBlbC5zdHlsZS5iYWNrZ3JvdW5kQ29sb3IgPSBjb2xvcjtcbiAgICAgICAgfVxuICAgIH0pO1xufVxudGVjem93ZUtvbG9yeShmb3JtKTtcbmNvbnN0IGdldExhdGVzdENvbW1pdEF1dGhvciA9ICh1cmwpID0+IF9fYXdhaXRlcih2b2lkIDAsIHZvaWQgMCwgdm9pZCAwLCBmdW5jdGlvbiogKCkge1xuICAgIGNvbnN0IHJhdyA9IHlpZWxkIGZldGNoKHVybCk7XG4gICAgY29uc3QganNvbiA9IHlpZWxkIHJhdy5qc29uKCk7XG4gICAgY29uc3QgYXV0aG9yVVJMID0ganNvblswXS5hdXRob3IudXJsO1xuICAgIGNvbnN0IHJhd0F1dGhvciA9IHlpZWxkIGZldGNoKGF1dGhvclVSTCk7XG4gICAgY29uc3QgYXV0aG9ySnNvbiA9IHlpZWxkIHJhd0F1dGhvci5qc29uKCk7XG4gICAgcmV0dXJuIGF1dGhvckpzb247XG59KTtcbmNvbnN0IHNldExhdGVzdENvbW1pdEF1dGhvclByb2ZpbGVQaWN0dXJlID0gKGF1dGhvcikgPT4ge1xuICAgIGNvbnN0IGF2YXRhclVSTCA9IGF1dGhvci5hdmF0YXJfdXJsO1xuICAgIGNvbnN0IHRzID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ3RzJyk7XG4gICAgY29uc3QgaW1nID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnaW1nJyk7XG4gICAgaW1nLnNldEF0dHJpYnV0ZSgnYWx0JywgJ0xhdGVzdCBUUyBDb21taXQgQXV0aG9yJyk7XG4gICAgaW1nLnNldEF0dHJpYnV0ZSgnc3JjJywgYXZhdGFyVVJMKTtcbiAgICB0cy5hcHBlbmRDaGlsZChpbWcpO1xufTtcbmNvbnN0IGdldEF1dGhvcnNSZXBvcyA9IChhdXRob3IpID0+IF9fYXdhaXRlcih2b2lkIDAsIHZvaWQgMCwgdm9pZCAwLCBmdW5jdGlvbiogKCkge1xuICAgIGNvbnN0IHJhdyA9IHlpZWxkIGZldGNoKGF1dGhvci5yZXBvc191cmwpO1xuICAgIGNvbnN0IGpzb24gPSB5aWVsZCByYXcuanNvbigpO1xuICAgIGNvbnN0IHJlcG9zID0gW107XG4gICAgZm9yIChjb25zdCBbaSwgcmVwb10gb2YgT2JqZWN0LmVudHJpZXMoanNvbikpIHtcbiAgICAgICAgcmVwb3MucHVzaChyZXBvWyduYW1lJ10pO1xuICAgIH1cbiAgICByZXR1cm4gcmVwb3Muc29ydCgoYSwgYikgPT4gYS50b0xvd2VyQ2FzZSgpLmxvY2FsZUNvbXBhcmUoYi50b0xvd2VyQ2FzZSgpKSk7XG59KTtcbmNvbnN0IGxhdGVzdFRTQ29tbWl0ID0gKHVybCkgPT4gX19hd2FpdGVyKHZvaWQgMCwgdm9pZCAwLCB2b2lkIDAsIGZ1bmN0aW9uKiAoKSB7XG4gICAgY29uc3QgYXV0aG9yID0geWllbGQgZ2V0TGF0ZXN0Q29tbWl0QXV0aG9yKHVybCk7XG4gICAgc2V0TGF0ZXN0Q29tbWl0QXV0aG9yUHJvZmlsZVBpY3R1cmUoYXV0aG9yKTtcbiAgICBjb25zb2xlLnRhYmxlKHlpZWxkIGdldEF1dGhvcnNSZXBvcyhhdXRob3IpKTtcbn0pO1xubGF0ZXN0VFNDb21taXQoJ2h0dHBzOi8vYXBpLmdpdGh1Yi5jb20vcmVwb3MvTWljcm9zb2Z0L1R5cGVTY3JpcHQvY29tbWl0cycpO1xuY29uc3QgbWFpbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdtYWluLXBhZ2UnKTtcbmNvbnN0IGRlbGF5ZWQgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgnZGVsYXllZCcpO1xuY29uc3QgYm9va2luZyA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdib29raW5nJyk7XG5sZXQgY2xpY2tOdW1iZXIgPSAwO1xuYm9va2luZy5vbmNsaWNrID0gKGUpID0+IHtcbiAgICBlLnN0b3BQcm9wYWdhdGlvbigpO1xufTtcbm1haW4ub25jbGljayA9IChlKSA9PiB7XG4gICAgY29uc3QgZWwgPSBlLnRhcmdldDtcbiAgICBmaWJfMS5maWIoTWF0aC5wb3coMTAsIGNsaWNrTnVtYmVyKSk7XG4gICAgY2xpY2tOdW1iZXIrKztcbiAgICBpZiAoZGVsYXllZC5jb250YWlucyhlbCkpIHtcbiAgICAgICAgZWwuc3R5bGUuYmFja2dyb3VuZENvbG9yID0gZ2V0UmFuZG9tQ29sb3IoKTtcbiAgICB9XG59O1xuZnVuY3Rpb24gZ2V0UmFuZG9tQ29sb3IoKSB7XG4gICAgY29uc3QgbGV0dGVycyA9ICcwMTIzNDU2Nzg5QUJDREVGJztcbiAgICBsZXQgY29sb3IgPSAnIyc7XG4gICAgZm9yICh2YXIgaSA9IDA7IGkgPCA2OyBpKyspIHtcbiAgICAgICAgY29sb3IgKz0gbGV0dGVyc1tNYXRoLmZsb29yKE1hdGgucmFuZG9tKCkgKiAxNildO1xuICAgIH1cbiAgICByZXR1cm4gY29sb3I7XG59XG4iXSwic291cmNlUm9vdCI6IiJ9