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
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/results.ts");
/******/ })
/************************************************************************/
/******/ ({

/***/ "./src/HTMLElements.ts":
/*!*****************************!*\
  !*** ./src/HTMLElements.ts ***!
  \*****************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

Object.defineProperty(exports, "__esModule", { value: true });
exports.quickAccessDiv = exports.topScoresSection = exports.averageTimesSection = exports.correctAnswersSection = exports.quizPrevResultsSection = exports.finalResultSpan = exports.initResultSpan = exports.answerForm = exports.saveWithStatsButton = exports.resultSection = exports.correctnessDiv = exports.quizesDiv = exports.resultsDiv = exports.descParagraph = exports.penaltyTimeSpan = exports.quizChoiseSection = exports.quizDescSection = exports.excerciseSection = exports.excerciseNumberSpan = exports.promptSpan = exports.stopButton = exports.nextButton = exports.previousButton = exports.answerInput = void 0;
exports.answerInput = document.getElementById("answer");
exports.previousButton = document.getElementById("previous");
exports.nextButton = document.getElementById("next");
exports.stopButton = document.getElementById("stop");
exports.promptSpan = document.getElementById("prompt");
exports.excerciseNumberSpan = document.getElementById("excercise-number");
exports.excerciseSection = document.getElementById("excercise");
exports.quizDescSection = document.getElementById("quiz-desc");
exports.quizChoiseSection = document.getElementById("quiz-choise");
exports.penaltyTimeSpan = document.getElementById("penalty-time");
exports.descParagraph = document.getElementById("desc");
exports.resultsDiv = document.getElementById("results");
exports.quizesDiv = document.getElementById("quizes");
exports.correctnessDiv = document.getElementById("correctness");
exports.resultSection = document.getElementById("result");
exports.saveWithStatsButton = document.getElementById("save-with-stats");
exports.answerForm = document.getElementById("answer-form");
exports.initResultSpan = document.getElementById("init-result");
exports.finalResultSpan = document.getElementById("final-result");
exports.quizPrevResultsSection = document.getElementById("quiz-prev-results");
exports.correctAnswersSection = document.getElementById("correct-answers");
exports.averageTimesSection = document.getElementById("average-times");
exports.topScoresSection = document.getElementById("top-scores");
exports.quickAccessDiv = document.getElementById("quick-access");


/***/ }),

/***/ "./src/QuizResults.ts":
/*!****************************!*\
  !*** ./src/QuizResults.ts ***!
  \****************************/
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
exports.QuizResults = void 0;
const quiz_1 = __webpack_require__(/*! ./quiz */ "./src/quiz.ts");
const HTMLElements_1 = __webpack_require__(/*! ./HTMLElements */ "./src/HTMLElements.ts");
const util_1 = __webpack_require__(/*! ./util */ "./src/util.ts");
const isResult = (obj) => {
    return obj.finalTime && obj.quizId;
};
class QuizResults {
    /**
     * @param answers array of marked answers
     * @param quiz quiz object
     * @param quizId chosen quiz id
     */
    constructor(answers, quiz) {
        this.answers = answers;
        this.quiz = quiz;
        this.penalties = [];
        this.bindEventHandlers();
    }
    mark() {
        this.checkAnswers();
        this.givePenalties();
        this.display();
    }
    /**
     * Checks answers, sets carrect argument in each answer to either true or false.
     */
    checkAnswers() {
        for (const i in this.answers) {
            this.answers[i].correct =
                this.answers[i].answer === this.quiz.questions[i].answer;
        }
    }
    /**
     * Gives penalties for each wrong answer.
     */
    givePenalties() {
        for (const i in this.answers) {
            this.penalties.push(this.answers[i].correct ? 0 : this.quiz.questions[i].penalty);
        }
    }
    /**
     * Bind neccessary event handlers.
     */
    bindEventHandlers() {
        HTMLElements_1.saveWithStatsButton.onclick = () => {
            this.goToMainScreen();
        };
    }
    /**
     * Goes back to main screen.
     */
    goToMainScreen() {
        window.location.href = "/";
    }
    /**
     * Calculates time spent on quiz.
     */
    get initTime() {
        return this.answers.reduce((sum, ans) => sum + ans.time / 1000, 0);
    }
    /**
     * Calculates final time, base + penalties.
     */
    get finalTime() {
        return this.penalties.reduce((sum, penalty) => sum + penalty, this.initTime);
    }
    displayUserResults() {
        HTMLElements_1.initResultSpan.textContent = QuizResults.formatTime(this.initTime);
        HTMLElements_1.finalResultSpan.textContent = QuizResults.formatTime(this.finalTime);
        // console.log(this.answers.entries());
        for (const [i, answer] of this.answers.entries()) {
            console.log(i, answer);
            const div = document.createElement("div");
            div.textContent = `${i + 1}. `;
            const span = document.createElement("span");
            const time = document.createElement("span");
            span.className = answer.correct ? "correct" : "incorrect";
            span.textContent = answer.correct
                ? "Poprawnie :)"
                : `Błąd: +${this.quiz.questions[i].penalty}s`;
            time.textContent = `  ${QuizResults.formatTime(answer.time / 1000)} s`;
            div.appendChild(span);
            div.appendChild(time);
            HTMLElements_1.correctnessDiv.appendChild(div);
        }
    }
    displayCorrectAnswers() {
        // console.log(this.answers.entries());
        for (const [i, question] of this.quiz.questions.entries()) {
            const div = document.createElement("div");
            div.textContent = `${i + 1}. `;
            const span = document.createElement("span");
            span.textContent = `${question.prompt} = ${question.answer}`;
            div.appendChild(span);
            HTMLElements_1.correctAnswersSection.appendChild(div);
        }
    }
    displayAverageTimes() {
        return __awaiter(this, void 0, void 0, function* () {
            // console.log(this.answers.entries());
            const times = yield quiz_1.averageTimes(this.quiz.id.toString());
            console.log(times);
            if (times === null) {
                util_1.error("Nie udało pobrać się średnich czasów na odpowiedź");
                return;
            }
            for (const [i, time] of times.entries()) {
                const div = document.createElement("div");
                div.textContent = `${i + 1}. `;
                const span = document.createElement("span");
                span.textContent = `${time ? QuizResults.formatTime(time / 1000) : "--.--"}s`;
                div.appendChild(span);
                HTMLElements_1.averageTimesSection.appendChild(div);
            }
        });
    }
    displayTopScores() {
        var _a, _b;
        return __awaiter(this, void 0, void 0, function* () {
            // console.log(this.answers.entries());
            const scores = yield quiz_1.topScores(this.quiz.id.toString());
            console.log(scores);
            if (scores === null) {
                util_1.error("Nie udało pobrać się najlepszych wyników");
                return;
            }
            for (const [i, [user, result]] of scores.entries()) {
                const div = document.createElement("div");
                const h3 = document.createElement("h3");
                // div.textContent = `${i + 1}. `;
                const timeP = document.createElement("p");
                timeP.textContent = `Czas: ${QuizResults.formatTime(result.finalTime / 1000)}s`;
                const answersP = document.createElement("p");
                answersP.textContent = `Poprawne odpowiedzi: ${(_a = result.answers) === null || _a === void 0 ? void 0 : _a.reduce((sum, ans) => (ans.correct ? sum + 1 : sum), 0)} / ${(_b = result.answers) === null || _b === void 0 ? void 0 : _b.length}`;
                h3.textContent = user.username;
                div.appendChild(h3);
                div.appendChild(timeP);
                div.appendChild(answersP);
                HTMLElements_1.topScoresSection.appendChild(div);
            }
        });
    }
    /**
     * Displays results.
     */
    display() {
        return __awaiter(this, void 0, void 0, function* () {
            this.displayUserResults();
            this.displayCorrectAnswers();
            yield this.displayAverageTimes();
            yield this.displayTopScores();
        });
    }
    /**
     * Formats time, round to 3 digits.
     * @param time time
     */
    static formatTime(time) {
        return time.toFixed(3).toString();
    }
    /**
     * Displays previous results.
     */
    static displayPreviousResults() {
        return __awaiter(this, void 0, void 0, function* () {
            let results;
            try {
                results = yield quiz_1.getPrevResults();
            }
            catch (_) { }
            let i = 1;
            let any = false;
            if (results instanceof Array) {
                for (const res of results) {
                    any = true;
                    if (isResult(res)) {
                        console.log(res);
                        const quiz = yield quiz_1.getQuizWithAnswers(res.quizId);
                        if (quiz) {
                            const finalTime = res.finalTime +
                                quiz.questions.reduce((sum, question, i) => {
                                    console.log(sum, res.answers, i);
                                    if (res.answers && !res.answers[i].correct) {
                                        return sum + question.penalty * 1000;
                                    }
                                    return sum;
                                }, 0);
                            const row = document.createElement("div");
                            const nameCol = document.createElement("div");
                            const resultCol = document.createElement("div");
                            const detailCol = document.createElement("div");
                            const link = document.createElement("a");
                            link.className = "btn btn-primary quiz-info-button";
                            link.dataset.quizId = quiz.id;
                            link.textContent = "Info";
                            link.href = "/results/" + quiz.id;
                            resultCol.className =
                                "col-4 d-flex justify-content-center align-items-center";
                            resultCol.textContent =
                                QuizResults.formatTime(finalTime / 1000) + "s";
                            nameCol.className =
                                "col-6 d-flex justify-content-center align-items-center";
                            nameCol.textContent = quiz.desc;
                            detailCol.className =
                                "col-2 d-flex justify-content-center align-items-center";
                            detailCol.appendChild(link);
                            row.className = `row prev-result ${i % 2 == 0 ? "even" : ""}`;
                            row.appendChild(nameCol);
                            row.appendChild(resultCol);
                            row.appendChild(detailCol);
                            HTMLElements_1.quizPrevResultsSection.appendChild(row);
                            i++;
                        }
                    }
                }
            }
            if (!any) {
                const row = document.createElement("div");
                row.className = `row prev-result d-flex justify-content-center`;
                row.textContent =
                    "Jeszcze nie rozwiązałeż żadnego quizu, na co czekasz? :)";
                HTMLElements_1.quizPrevResultsSection.appendChild(row);
            }
        });
    }
}
exports.QuizResults = QuizResults;


/***/ }),

/***/ "./src/quiz.ts":
/*!*********************!*\
  !*** ./src/quiz.ts ***!
  \*********************/
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
var __await = (this && this.__await) || function (v) { return this instanceof __await ? (this.v = v, this) : new __await(v); }
var __asyncGenerator = (this && this.__asyncGenerator) || function (thisArg, _arguments, generator) {
    if (!Symbol.asyncIterator) throw new TypeError("Symbol.asyncIterator is not defined.");
    var g = generator.apply(thisArg, _arguments || []), i, q = [];
    return i = {}, verb("next"), verb("throw"), verb("return"), i[Symbol.asyncIterator] = function () { return this; }, i;
    function verb(n) { if (g[n]) i[n] = function (v) { return new Promise(function (a, b) { q.push([n, v, a, b]) > 1 || resume(n, v); }); }; }
    function resume(n, v) { try { step(g[n](v)); } catch (e) { settle(q[0][3], e); } }
    function step(r) { r.value instanceof __await ? Promise.resolve(r.value.v).then(fulfill, reject) : settle(q[0][2], r); }
    function fulfill(value) { resume("next", value); }
    function reject(value) { resume("throw", value); }
    function settle(f, v) { if (f(v), q.shift(), q.length) resume(q[0][0], q[0][1]); }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.topScores = exports.averageTimes = exports.getAnswers = exports.saveResults = exports.getPrevResults = exports.getQuizes = exports.getQuizWithAnswers = exports.getQuiz = exports.isQuiz = void 0;
const util_1 = __webpack_require__(/*! ./util */ "./src/util.ts");
const isQuestion = (question) => {
    return question && question.prompt && question.answer && question.penalty;
};
exports.isQuiz = (quiz) => {
    return quiz.desc && quiz.questions && quiz.id !== undefined;
};
const isQuizWithAnswers = (quiz) => {
    return quiz.desc && quiz.questions && quiz.id !== undefined;
};
/**
 * Returns quiz with given id or null on error.
 * @param id quiz id
 */
exports.getQuiz = (id) => __awaiter(void 0, void 0, void 0, function* () {
    console.log(id);
    try {
        const quizRaw = yield fetch("/get_quiz/" + id);
        const quiz = yield quizRaw.json();
        if (exports.isQuiz(quiz)) {
            return quiz;
        }
    }
    catch (_) { }
    return null;
});
/**
 * Returns quiz with given id or null on error.
 * @param id quiz id
 */
exports.getQuizWithAnswers = (id) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const quizRaw = yield fetch("/get_quiz_with_answers/" + id);
        const quiz = yield quizRaw.json();
        if (isQuizWithAnswers(quiz)) {
            return quiz;
        }
    }
    catch (_) { }
    return null;
});
/**
 * Generator that returns all quizes.
 */
function getQuizes() {
    return __asyncGenerator(this, arguments, function* getQuizes_1() {
        const quizRaw = yield __await(fetch("/get_quizes"));
        const obj = yield __await(quizRaw.json());
        if (obj instanceof Array) {
            for (const quiz of obj) {
                if (exports.isQuiz(quiz)) {
                    yield yield __await(quiz);
                }
            }
        }
    });
}
exports.getQuizes = getQuizes;
exports.getPrevResults = () => __awaiter(void 0, void 0, void 0, function* () {
    const results = yield fetch("/prev_results");
    const obj = yield results.json();
    if (obj instanceof Array) {
        return obj;
    }
    return [];
});
const getRawQuizResults = (answers) => {
    const totalTime = answers.reduce((sum, ans) => sum + ans.time, 0);
    const qr = { answers: [], times: [] };
    answers.forEach((ans) => {
        qr.answers.push(ans.answer);
        qr.times.push(ans.time / totalTime);
    });
    return qr;
};
exports.saveResults = (answers, quizId) => __awaiter(void 0, void 0, void 0, function* () {
    const qr = getRawQuizResults(answers);
    console.log(qr, answers, JSON.stringify(qr));
    yield fetch("/post_results/" + quizId, {
        method: "POST",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json",
            "CSRF-Token": util_1.csrfToken,
        },
        body: JSON.stringify(qr),
    });
});
exports.getAnswers = (quizId) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const raw = yield fetch("/get_answers/" + quizId);
        const obj = yield raw.json();
        if (obj instanceof Array) {
            return obj;
        }
    }
    catch (_) { }
    return null;
});
exports.averageTimes = (quizId) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const raw = yield fetch("/get_answers_mean_time/" + quizId);
        const obj = yield raw.json();
        if (obj instanceof Array) {
            return obj;
        }
    }
    catch (_) { }
    return null;
});
exports.topScores = (quizId) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const raw = yield fetch("/top_scores/" + quizId);
        const obj = yield raw.json();
        if (obj instanceof Array) {
            return obj;
        }
    }
    catch (_) { }
    return null;
});


/***/ }),

/***/ "./src/results.ts":
/*!************************!*\
  !*** ./src/results.ts ***!
  \************************/
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
const quiz_1 = __webpack_require__(/*! ./quiz */ "./src/quiz.ts");
const QuizResults_1 = __webpack_require__(/*! ./QuizResults */ "./src/QuizResults.ts");
const main = () => __awaiter(void 0, void 0, void 0, function* () {
    const quizId = window.location.pathname
        .split("/")
        .filter((a) => a !== "")
        .slice(-1)[0];
    const quiz = yield quiz_1.getQuizWithAnswers(quizId);
    const answers = yield quiz_1.getAnswers(quizId);
    if (quiz === null || answers === null) {
        alert("Ten test nie istnieje");
        window.location.href = "/";
        return;
    }
    const quizResults = new QuizResults_1.QuizResults(answers, quiz);
    quizResults.mark();
});
main();


/***/ }),

/***/ "./src/util.ts":
/*!*********************!*\
  !*** ./src/util.ts ***!
  \*********************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.csrfToken = exports.error = void 0;
/**
 * Function that logs errors.
 * May in the future be overwriten to function that display error on screen.
 */
exports.error = (s) => alert(s);
exports.csrfToken = ((_a = document === null || document === void 0 ? void 0 : document.querySelector('meta[name="csrf-token"]')) === null || _a === void 0 ? void 0 : _a.getAttribute("content")) ||
    "";


/***/ })

/******/ });
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vd2VicGFjay9ib290c3RyYXAiLCJ3ZWJwYWNrOi8vLy4vc3JjL0hUTUxFbGVtZW50cy50cyIsIndlYnBhY2s6Ly8vLi9zcmMvUXVpelJlc3VsdHMudHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL3F1aXoudHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL3Jlc3VsdHMudHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL3V0aWwudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IjtRQUFBO1FBQ0E7O1FBRUE7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTs7UUFFQTtRQUNBOztRQUVBO1FBQ0E7O1FBRUE7UUFDQTtRQUNBOzs7UUFHQTtRQUNBOztRQUVBO1FBQ0E7O1FBRUE7UUFDQTtRQUNBO1FBQ0EsMENBQTBDLGdDQUFnQztRQUMxRTtRQUNBOztRQUVBO1FBQ0E7UUFDQTtRQUNBLHdEQUF3RCxrQkFBa0I7UUFDMUU7UUFDQSxpREFBaUQsY0FBYztRQUMvRDs7UUFFQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0EseUNBQXlDLGlDQUFpQztRQUMxRSxnSEFBZ0gsbUJBQW1CLEVBQUU7UUFDckk7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7UUFDQSwyQkFBMkIsMEJBQTBCLEVBQUU7UUFDdkQsaUNBQWlDLGVBQWU7UUFDaEQ7UUFDQTtRQUNBOztRQUVBO1FBQ0Esc0RBQXNELCtEQUErRDs7UUFFckg7UUFDQTs7O1FBR0E7UUFDQTs7Ozs7Ozs7Ozs7Ozs7OztBQ2xGYSxtQkFBVyxHQUFHLFFBQVEsQ0FBQyxjQUFjLENBQ2hELFFBQVEsQ0FDVyxDQUFDO0FBQ1Qsc0JBQWMsR0FBRyxRQUFRLENBQUMsY0FBYyxDQUNuRCxVQUFVLENBQ1UsQ0FBQztBQUNWLGtCQUFVLEdBQUcsUUFBUSxDQUFDLGNBQWMsQ0FBQyxNQUFNLENBQXNCLENBQUM7QUFDbEUsa0JBQVUsR0FBRyxR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QXdEO0FBQ3hELHVGQUE0QztBQUM1QyxNQUFNLElBQUksR0FBRyxHQUFTLEVBQUU7SUFDdEIsTUFBTSxNQUFNLEdBQUcsTUFBTSxDQUFDLFFBQVEsQ0FBQyxRQUFRO1NBQ3BDLEtBQUssQ0FBQyxHQUFHLENBQUM7U0FDVixNQUFNLENBQUMsQ0FBQyxDQUFDLEVBQUUsRUFBRSxDQUFDLENBQUMsS0FBSyxFQUFFLENBQUM7U0FDdkIsS0FBSyxDQUFDLENBQUMsQ0FBQyxDQUFDLENBQUMsQ0FBQyxDQUFDLENBQUM7SUFFaEIsTUFBTSxJQUFJLEdBQUcsTUFBTSx5QkFBa0IsQ0FBQyxNQUFNLENBQUMsQ0FBQztJQUM5QyxNQUFNLE9BQU8sR0FBRyxNQUFNLGlCQUFVLENBQUMsTUFBTSxDQUFDLENBQUM7SUFDekMsSUFBSSxJQUFJLEtBQUssSUFBSSxJQUFJLE9BQU8sS0FBSyxJQUFJLEVBQUU7UUFDckMsS0FBSyxDQUFDLHVCQUF1QixDQUFDLENBQUM7UUFDL0IsTUFBTSxDQUFDLFFBQVEsQ0FBQyxJQUFJLEdBQUcsR0FBRyxDQUFDO1FBQzNCLE9BQU87S0FDUjtJQUNELE1BQU0sV0FBVyxHQUFHLElBQUkseUJBQVcsQ0FBQyxPQUFPLEVBQUUsSUFBSSxDQUFDLENBQUM7SUFDbkQsV0FBVyxDQUFDLElBQUksRUFBRSxDQUFDO0FBQ3JCLENBQUMsRUFBQztBQUNGLElBQUksRUFBRSxDQUFDOzs7Ozs7Ozs7Ozs7Ozs7OztBQ2xCUDs7O0dBR0c7QUFDVSxhQUFLLEdBQUcsQ0FBQyxDQUFTLEVBQUUsRUFBRSxDQUFDLEtBQUssQ0FBQyxDQUFDLENBQUMsQ0FBQztBQUVoQyxpQkFBUyxHQUNwQixlQUFRLGFBQVIsUUFBUSx1QkFBUixRQUFRLENBQUUsYUFBYSxDQUFDLHlCQUF5QiwyQ0FBRyxZQUFZLENBQUMsU0FBUztJQUMxRSxFQUFFLENBQUMiLCJmaWxlIjoicmVzdWx0cy5qcyIsInNvdXJjZXNDb250ZW50IjpbIiBcdC8vIFRoZSBtb2R1bGUgY2FjaGVcbiBcdHZhciBpbnN0YWxsZWRNb2R1bGVzID0ge307XG5cbiBcdC8vIFRoZSByZXF1aXJlIGZ1bmN0aW9uXG4gXHRmdW5jdGlvbiBfX3dlYnBhY2tfcmVxdWlyZV9fKG1vZHVsZUlkKSB7XG5cbiBcdFx0Ly8gQ2hlY2sgaWYgbW9kdWxlIGlzIGluIGNhY2hlXG4gXHRcdGlmKGluc3RhbGxlZE1vZHVsZXNbbW9kdWxlSWRdKSB7XG4gXHRcdFx0cmV0dXJuIGluc3RhbGxlZE1vZHVsZXNbbW9kdWxlSWRdLmV4cG9ydHM7XG4gXHRcdH1cbiBcdFx0Ly8gQ3JlYXRlIGEgbmV3IG1vZHVsZSAoYW5kIHB1dCBpdCBpbnRvIHRoZSBjYWNoZSlcbiBcdFx0dmFyIG1vZHVsZSA9IGluc3RhbGxlZE1vZHVsZXNbbW9kdWxlSWRdID0ge1xuIFx0XHRcdGk6IG1vZHVsZUlkLFxuIFx0XHRcdGw6IGZhbHNlLFxuIFx0XHRcdGV4cG9ydHM6IHt9XG4gXHRcdH07XG5cbiBcdFx0Ly8gRXhlY3V0ZSB0aGUgbW9kdWxlIGZ1bmN0aW9uXG4gXHRcdG1vZHVsZXNbbW9kdWxlSWRdLmNhbGwobW9kdWxlLmV4cG9ydHMsIG1vZHVsZSwgbW9kdWxlLmV4cG9ydHMsIF9fd2VicGFja19yZXF1aXJlX18pO1xuXG4gXHRcdC8vIEZsYWcgdGhlIG1vZHVsZSBhcyBsb2FkZWRcbiBcdFx0bW9kdWxlLmwgPSB0cnVlO1xuXG4gXHRcdC8vIFJldHVybiB0aGUgZXhwb3J0cyBvZiB0aGUgbW9kdWxlXG4gXHRcdHJldHVybiBtb2R1bGUuZXhwb3J0cztcbiBcdH1cblxuXG4gXHQvLyBleHBvc2UgdGhlIG1vZHVsZXMgb2JqZWN0IChfX3dlYnBhY2tfbW9kdWxlc19fKVxuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5tID0gbW9kdWxlcztcblxuIFx0Ly8gZXhwb3NlIHRoZSBtb2R1bGUgY2FjaGVcbiBcdF9fd2VicGFja19yZXF1aXJlX18uYyA9IGluc3RhbGxlZE1vZHVsZXM7XG5cbiBcdC8vIGRlZmluZSBnZXR0ZXIgZnVuY3Rpb24gZm9yIGhhcm1vbnkgZXhwb3J0c1xuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5kID0gZnVuY3Rpb24oZXhwb3J0cywgbmFtZSwgZ2V0dGVyKSB7XG4gXHRcdGlmKCFfX3dlYnBhY2tfcmVxdWlyZV9fLm8oZXhwb3J0cywgbmFtZSkpIHtcbiBcdFx0XHRPYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgbmFtZSwgeyBlbnVtZXJhYmxlOiB0cnVlLCBnZXQ6IGdldHRlciB9KTtcbiBcdFx0fVxuIFx0fTtcblxuIFx0Ly8gZGVmaW5lIF9fZXNNb2R1bGUgb24gZXhwb3J0c1xuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5yID0gZnVuY3Rpb24oZXhwb3J0cykge1xuIFx0XHRpZih0eXBlb2YgU3ltYm9sICE9PSAndW5kZWZpbmVkJyAmJiBTeW1ib2wudG9TdHJpbmdUYWcpIHtcbiBcdFx0XHRPYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgU3ltYm9sLnRvU3RyaW5nVGFnLCB7IHZhbHVlOiAnTW9kdWxlJyB9KTtcbiBcdFx0fVxuIFx0XHRPYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgJ19fZXNNb2R1bGUnLCB7IHZhbHVlOiB0cnVlIH0pO1xuIFx0fTtcblxuIFx0Ly8gY3JlYXRlIGEgZmFrZSBuYW1lc3BhY2Ugb2JqZWN0XG4gXHQvLyBtb2RlICYgMTogdmFsdWUgaXMgYSBtb2R1bGUgaWQsIHJlcXVpcmUgaXRcbiBcdC8vIG1vZGUgJiAyOiBtZXJnZSBhbGwgcHJvcGVydGllcyBvZiB2YWx1ZSBpbnRvIHRoZSBuc1xuIFx0Ly8gbW9kZSAmIDQ6IHJldHVybiB2YWx1ZSB3aGVuIGFscmVhZHkgbnMgb2JqZWN0XG4gXHQvLyBtb2RlICYgOHwxOiBiZWhhdmUgbGlrZSByZXF1aXJlXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLnQgPSBmdW5jdGlvbih2YWx1ZSwgbW9kZSkge1xuIFx0XHRpZihtb2RlICYgMSkgdmFsdWUgPSBfX3dlYnBhY2tfcmVxdWlyZV9fKHZhbHVlKTtcbiBcdFx0aWYobW9kZSAmIDgpIHJldHVybiB2YWx1ZTtcbiBcdFx0aWYoKG1vZGUgJiA0KSAmJiB0eXBlb2YgdmFsdWUgPT09ICdvYmplY3QnICYmIHZhbHVlICYmIHZhbHVlLl9fZXNNb2R1bGUpIHJldHVybiB2YWx1ZTtcbiBcdFx0dmFyIG5zID0gT2JqZWN0LmNyZWF0ZShudWxsKTtcbiBcdFx0X193ZWJwYWNrX3JlcXVpcmVfXy5yKG5zKTtcbiBcdFx0T2JqZWN0LmRlZmluZVByb3BlcnR5KG5zLCAnZGVmYXVsdCcsIHsgZW51bWVyYWJsZTogdHJ1ZSwgdmFsdWU6IHZhbHVlIH0pO1xuIFx0XHRpZihtb2RlICYgMiAmJiB0eXBlb2YgdmFsdWUgIT0gJ3N0cmluZycpIGZvcih2YXIga2V5IGluIHZhbHVlKSBfX3dlYnBhY2tfcmVxdWlyZV9fLmQobnMsIGtleSwgZnVuY3Rpb24oa2V5KSB7IHJldHVybiB2YWx1ZVtrZXldOyB9LmJpbmQobnVsbCwga2V5KSk7XG4gXHRcdHJldHVybiBucztcbiBcdH07XG5cbiBcdC8vIGdldERlZmF1bHRFeHBvcnQgZnVuY3Rpb24gZm9yIGNvbXBhdGliaWxpdHkgd2l0aCBub24taGFybW9ueSBtb2R1bGVzXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLm4gPSBmdW5jdGlvbihtb2R1bGUpIHtcbiBcdFx0dmFyIGdldHRlciA9IG1vZHVsZSAmJiBtb2R1bGUuX19lc01vZHVsZSA/XG4gXHRcdFx0ZnVuY3Rpb24gZ2V0RGVmYXVsdCgpIHsgcmV0dXJuIG1vZHVsZVsnZGVmYXVsdCddOyB9IDpcbiBcdFx0XHRmdW5jdGlvbiBnZXRNb2R1bGVFeHBvcnRzKCkgeyByZXR1cm4gbW9kdWxlOyB9O1xuIFx0XHRfX3dlYnBhY2tfcmVxdWlyZV9fLmQoZ2V0dGVyLCAnYScsIGdldHRlcik7XG4gXHRcdHJldHVybiBnZXR0ZXI7XG4gXHR9O1xuXG4gXHQvLyBPYmplY3QucHJvdG90eXBlLmhhc093blByb3BlcnR5LmNhbGxcbiBcdF9fd2VicGFja19yZXF1aXJlX18ubyA9IGZ1bmN0aW9uKG9iamVjdCwgcHJvcGVydHkpIHsgcmV0dXJuIE9iamVjdC5wcm90b3R5cGUuaGFzT3duUHJvcGVydHkuY2FsbChvYmplY3QsIHByb3BlcnR5KTsgfTtcblxuIFx0Ly8gX193ZWJwYWNrX3B1YmxpY19wYXRoX19cbiBcdF9fd2VicGFja19yZXF1aXJlX18ucCA9IFwiXCI7XG5cblxuIFx0Ly8gTG9hZCBlbnRyeSBtb2R1bGUgYW5kIHJldHVybiBleHBvcnRzXG4gXHRyZXR1cm4gX193ZWJwYWNrX3JlcXVpcmVfXyhfX3dlYnBhY2tfcmVxdWlyZV9fLnMgPSBcIi4vc3JjL3Jlc3VsdHMudHNcIik7XG4iLCJleHBvcnQgY29uc3QgYW5zd2VySW5wdXQgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJhbnN3ZXJcIlxuKSBhcyBIVE1MSW5wdXRFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHByZXZpb3VzQnV0dG9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwicHJldmlvdXNcIlxuKSBhcyBIVE1MQnV0dG9uRWxlbWVudDtcbmV4cG9ydCBjb25zdCBuZXh0QnV0dG9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXCJuZXh0XCIpIGFzIEhUTUxCdXR0b25FbGVtZW50O1xuZXhwb3J0IGNvbnN0IHN0b3BCdXR0b24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcInN0b3BcIikgYXMgSFRNTEJ1dHRvbkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgcHJvbXB0U3BhbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFwicHJvbXB0XCIpIGFzIEhUTUxTcGFuRWxlbWVudDtcbmV4cG9ydCBjb25zdCBleGNlcmNpc2VOdW1iZXJTcGFuID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwiZXhjZXJjaXNlLW51bWJlclwiXG4pIGFzIEhUTUxTcGFuRWxlbWVudDtcbmV4cG9ydCBjb25zdCBleGNlcmNpc2VTZWN0aW9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwiZXhjZXJjaXNlXCJcbikgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgcXVpekRlc2NTZWN0aW9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwicXVpei1kZXNjXCJcbikgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgcXVpekNob2lzZVNlY3Rpb24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJxdWl6LWNob2lzZVwiXG4pIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHBlbmFsdHlUaW1lU3BhbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcInBlbmFsdHktdGltZVwiXG4pIGFzIEhUTUxTcGFuRWxlbWVudDtcbmV4cG9ydCBjb25zdCBkZXNjUGFyYWdyYXBoID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwiZGVzY1wiXG4pIGFzIEhUTUxQYXJhZ3JhcGhFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHJlc3VsdHNEaXYgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcInJlc3VsdHNcIikgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgcXVpemVzRGl2ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXCJxdWl6ZXNcIikgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgY29ycmVjdG5lc3NEaXYgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJjb3JyZWN0bmVzc1wiXG4pIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHJlc3VsdFNlY3Rpb24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJyZXN1bHRcIlxuKSBhcyBIVE1MRGl2RWxlbWVudDtcbmV4cG9ydCBjb25zdCBzYXZlV2l0aFN0YXRzQnV0dG9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwic2F2ZS13aXRoLXN0YXRzXCJcbikgYXMgSFRNTEJ1dHRvbkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgYW5zd2VyRm9ybSA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcImFuc3dlci1mb3JtXCJcbikgYXMgSFRNTEZvcm1FbGVtZW50O1xuXG5leHBvcnQgY29uc3QgaW5pdFJlc3VsdFNwYW4gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJpbml0LXJlc3VsdFwiXG4pIGFzIEhUTUxTcGFuRWxlbWVudDtcbmV4cG9ydCBjb25zdCBmaW5hbFJlc3VsdFNwYW4gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJmaW5hbC1yZXN1bHRcIlxuKSBhcyBIVE1MU3BhbkVsZW1lbnQ7XG5cbmV4cG9ydCBjb25zdCBxdWl6UHJldlJlc3VsdHNTZWN0aW9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwicXVpei1wcmV2LXJlc3VsdHNcIlxuKSBhcyBIVE1MRGl2RWxlbWVudDtcbmV4cG9ydCBjb25zdCBjb3JyZWN0QW5zd2Vyc1NlY3Rpb24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJjb3JyZWN0LWFuc3dlcnNcIlxuKSBhcyBIVE1MRGl2RWxlbWVudDtcbmV4cG9ydCBjb25zdCBhdmVyYWdlVGltZXNTZWN0aW9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwiYXZlcmFnZS10aW1lc1wiXG4pIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHRvcFNjb3Jlc1NlY3Rpb24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJ0b3Atc2NvcmVzXCJcbikgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgcXVpY2tBY2Nlc3NEaXYgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJxdWljay1hY2Nlc3NcIlxuKSBhcyBIVE1MRGl2RWxlbWVudDtcbiIsImltcG9ydCB7XG4gIEFuc3dlcixcbiAgUXVpeldpdGhBbnN3ZXJzLFxuICBnZXRRdWl6V2l0aEFuc3dlcnMsXG4gIGdldFByZXZSZXN1bHRzLFxuICBhdmVyYWdlVGltZXMsXG4gIHRvcFNjb3Jlcyxcbn0gZnJvbSBcIi4vcXVpelwiO1xuaW1wb3J0IHtcbiAgc2F2ZVdpdGhTdGF0c0J1dHRvbixcbiAgaW5pdFJlc3VsdFNwYW4sXG4gIGZpbmFsUmVzdWx0U3BhbixcbiAgY29ycmVjdG5lc3NEaXYsXG4gIHF1aXpQcmV2UmVzdWx0c1NlY3Rpb24sXG4gIGNvcnJlY3RBbnN3ZXJzU2VjdGlvbixcbiAgYXZlcmFnZVRpbWVzU2VjdGlvbixcbiAgdG9wU2NvcmVzU2VjdGlvbixcbn0gZnJvbSBcIi4vSFRNTEVsZW1lbnRzXCI7XG5pbXBvcnQgeyBlcnJvciB9IGZyb20gXCIuL3V0aWxcIjtcblxuZXhwb3J0IGludGVyZmFjZSBSZXN1bHQge1xuICBmaW5hbFRpbWU6IG51bWJlcjtcbiAgYW5zd2Vycz86IEFuc3dlcltdO1xuICBxdWl6SWQ6IHN0cmluZztcbn1cblxuY29uc3QgaXNSZXN1bHQgPSAob2JqOiBhbnkpOiBvYmogaXMgUmVzdWx0ID0+IHtcbiAgcmV0dXJuIG9iai5maW5hbFRpbWUgJiYgb2JqLnF1aXpJZDtcbn07XG5cbmV4cG9ydCBjbGFzcyBRdWl6UmVzdWx0cyB7XG4gIHByaXZhdGUgcGVuYWx0aWVzOiBudW1iZXJbXSA9IFtdO1xuXG4gIC8qKlxuICAgKiBAcGFyYW0gYW5zd2VycyBhcnJheSBvZiBtYXJrZWQgYW5zd2Vyc1xuICAgKiBAcGFyYW0gcXVpeiBxdWl6IG9iamVjdFxuICAgKiBAcGFyYW0gcXVpeklkIGNob3NlbiBxdWl6IGlkXG4gICAqL1xuICBjb25zdHJ1Y3Rvcihwcml2YXRlIGFuc3dlcnM6IEFuc3dlcltdLCBwcml2YXRlIHF1aXo6IFF1aXpXaXRoQW5zd2Vycykge1xuICAgIHRoaXMuYmluZEV2ZW50SGFuZGxlcnMoKTtcbiAgfVxuXG4gIHB1YmxpYyBtYXJrKCkge1xuICAgIHRoaXMuY2hlY2tBbnN3ZXJzKCk7XG4gICAgdGhpcy5naXZlUGVuYWx0aWVzKCk7XG4gICAgdGhpcy5kaXNwbGF5KCk7XG4gIH1cblxuICAvKipcbiAgICogQ2hlY2tzIGFuc3dlcnMsIHNldHMgY2FycmVjdCBhcmd1bWVudCBpbiBlYWNoIGFuc3dlciB0byBlaXRoZXIgdHJ1ZSBvciBmYWxzZS5cbiAgICovXG4gIHByaXZhdGUgY2hlY2tBbnN3ZXJzKCkge1xuICAgIGZvciAoY29uc3QgaSBpbiB0aGlzLmFuc3dlcnMpIHtcbiAgICAgIHRoaXMuYW5zd2Vyc1tpXS5jb3JyZWN0ID1cbiAgICAgICAgdGhpcy5hbnN3ZXJzW2ldLmFuc3dlciA9PT0gdGhpcy5xdWl6LnF1ZXN0aW9uc1tpXS5hbnN3ZXI7XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEdpdmVzIHBlbmFsdGllcyBmb3IgZWFjaCB3cm9uZyBhbnN3ZXIuXG4gICAqL1xuICBwcml2YXRlIGdpdmVQZW5hbHRpZXMoKSB7XG4gICAgZm9yIChjb25zdCBpIGluIHRoaXMuYW5zd2Vycykge1xuICAgICAgdGhpcy5wZW5hbHRpZXMucHVzaChcbiAgICAgICAgdGhpcy5hbnN3ZXJzW2ldLmNvcnJlY3QgPyAwIDogdGhpcy5xdWl6LnF1ZXN0aW9uc1tpXS5wZW5hbHR5XG4gICAgICApO1xuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBCaW5kIG5lY2Nlc3NhcnkgZXZlbnQgaGFuZGxlcnMuXG4gICAqL1xuICBwcml2YXRlIGJpbmRFdmVudEhhbmRsZXJzKCkge1xuICAgIHNhdmVXaXRoU3RhdHNCdXR0b24ub25jbGljayA9ICgpID0+IHtcbiAgICAgIHRoaXMuZ29Ub01haW5TY3JlZW4oKTtcbiAgICB9O1xuICB9XG5cbiAgLyoqXG4gICAqIEdvZXMgYmFjayB0byBtYWluIHNjcmVlbi5cbiAgICovXG4gIHByaXZhdGUgZ29Ub01haW5TY3JlZW4oKSB7XG4gICAgd2luZG93LmxvY2F0aW9uLmhyZWYgPSBcIi9cIjtcbiAgfVxuXG4gIC8qKlxuICAgKiBDYWxjdWxhdGVzIHRpbWUgc3BlbnQgb24gcXVpei5cbiAgICovXG4gIHByaXZhdGUgZ2V0IGluaXRUaW1lKCk6IG51bWJlciB7XG4gICAgcmV0dXJuIHRoaXMuYW5zd2Vycy5yZWR1Y2UoKHN1bSwgYW5zKSA9PiBzdW0gKyBhbnMudGltZSAvIDEwMDAsIDApO1xuICB9XG5cbiAgLyoqXG4gICAqIENhbGN1bGF0ZXMgZmluYWwgdGltZSwgYmFzZSArIHBlbmFsdGllcy5cbiAgICovXG4gIHB1YmxpYyBnZXQgZmluYWxUaW1lKCkge1xuICAgIHJldHVybiB0aGlzLnBlbmFsdGllcy5yZWR1Y2UoXG4gICAgICAoc3VtLCBwZW5hbHR5KSA9PiBzdW0gKyBwZW5hbHR5LFxuICAgICAgdGhpcy5pbml0VGltZVxuICAgICk7XG4gIH1cbiAgcHJpdmF0ZSBkaXNwbGF5VXNlclJlc3VsdHMoKSB7XG4gICAgaW5pdFJlc3VsdFNwYW4udGV4dENvbnRlbnQgPSBRdWl6UmVzdWx0cy5mb3JtYXRUaW1lKHRoaXMuaW5pdFRpbWUpO1xuICAgIGZpbmFsUmVzdWx0U3Bhbi50ZXh0Q29udGVudCA9IFF1aXpSZXN1bHRzLmZvcm1hdFRpbWUodGhpcy5maW5hbFRpbWUpO1xuXG4gICAgLy8gY29uc29sZS5sb2codGhpcy5hbnN3ZXJzLmVudHJpZXMoKSk7XG4gICAgZm9yIChjb25zdCBbaSwgYW5zd2VyXSBvZiB0aGlzLmFuc3dlcnMuZW50cmllcygpKSB7XG4gICAgICBjb25zb2xlLmxvZyhpLCBhbnN3ZXIpO1xuICAgICAgY29uc3QgZGl2ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImRpdlwiKTtcbiAgICAgIGRpdi50ZXh0Q29udGVudCA9IGAke2kgKyAxfS4gYDtcbiAgICAgIGNvbnN0IHNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwic3BhblwiKTtcbiAgICAgIGNvbnN0IHRpbWUgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwic3BhblwiKTtcbiAgICAgIHNwYW4uY2xhc3NOYW1lID0gYW5zd2VyLmNvcnJlY3QgPyBcImNvcnJlY3RcIiA6IFwiaW5jb3JyZWN0XCI7XG4gICAgICBzcGFuLnRleHRDb250ZW50ID0gYW5zd2VyLmNvcnJlY3RcbiAgICAgICAgPyBcIlBvcHJhd25pZSA6KVwiXG4gICAgICAgIDogYELFgsSFZDogKyR7dGhpcy5xdWl6LnF1ZXN0aW9uc1tpXS5wZW5hbHR5fXNgO1xuICAgICAgdGltZS50ZXh0Q29udGVudCA9IGAgICR7UXVpelJlc3VsdHMuZm9ybWF0VGltZShhbnN3ZXIudGltZSAvIDEwMDApfSBzYDtcbiAgICAgIGRpdi5hcHBlbmRDaGlsZChzcGFuKTtcbiAgICAgIGRpdi5hcHBlbmRDaGlsZCh0aW1lKTtcbiAgICAgIGNvcnJlY3RuZXNzRGl2LmFwcGVuZENoaWxkKGRpdik7XG4gICAgfVxuICB9XG4gIHByaXZhdGUgZGlzcGxheUNvcnJlY3RBbnN3ZXJzKCkge1xuICAgIC8vIGNvbnNvbGUubG9nKHRoaXMuYW5zd2Vycy5lbnRyaWVzKCkpO1xuICAgIGZvciAoY29uc3QgW2ksIHF1ZXN0aW9uXSBvZiB0aGlzLnF1aXoucXVlc3Rpb25zLmVudHJpZXMoKSkge1xuICAgICAgY29uc3QgZGl2ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImRpdlwiKTtcbiAgICAgIGRpdi50ZXh0Q29udGVudCA9IGAke2kgKyAxfS4gYDtcbiAgICAgIGNvbnN0IHNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwic3BhblwiKTtcbiAgICAgIHNwYW4udGV4dENvbnRlbnQgPSBgJHtxdWVzdGlvbi5wcm9tcHR9ID0gJHtxdWVzdGlvbi5hbnN3ZXJ9YDtcbiAgICAgIGRpdi5hcHBlbmRDaGlsZChzcGFuKTtcbiAgICAgIGNvcnJlY3RBbnN3ZXJzU2VjdGlvbi5hcHBlbmRDaGlsZChkaXYpO1xuICAgIH1cbiAgfVxuICBwcml2YXRlIGFzeW5jIGRpc3BsYXlBdmVyYWdlVGltZXMoKSB7XG4gICAgLy8gY29uc29sZS5sb2codGhpcy5hbnN3ZXJzLmVudHJpZXMoKSk7XG4gICAgY29uc3QgdGltZXMgPSBhd2FpdCBhdmVyYWdlVGltZXModGhpcy5xdWl6LmlkLnRvU3RyaW5nKCkpO1xuICAgIGNvbnNvbGUubG9nKHRpbWVzKTtcbiAgICBpZiAodGltZXMgPT09IG51bGwpIHtcbiAgICAgIGVycm9yKFwiTmllIHVkYcWCbyBwb2JyYcSHIHNpxJkgxZtyZWRuaWNoIGN6YXPDs3cgbmEgb2Rwb3dpZWTFulwiKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG4gICAgZm9yIChjb25zdCBbaSwgdGltZV0gb2YgdGltZXMuZW50cmllcygpKSB7XG4gICAgICBjb25zdCBkaXYgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwiZGl2XCIpO1xuICAgICAgZGl2LnRleHRDb250ZW50ID0gYCR7aSArIDF9LiBgO1xuICAgICAgY29uc3Qgc3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJzcGFuXCIpO1xuICAgICAgc3Bhbi50ZXh0Q29udGVudCA9IGAke1xuICAgICAgICB0aW1lID8gUXVpelJlc3VsdHMuZm9ybWF0VGltZSh0aW1lIC8gMTAwMCkgOiBcIi0tLi0tXCJcbiAgICAgIH1zYDtcbiAgICAgIGRpdi5hcHBlbmRDaGlsZChzcGFuKTtcbiAgICAgIGF2ZXJhZ2VUaW1lc1NlY3Rpb24uYXBwZW5kQ2hpbGQoZGl2KTtcbiAgICB9XG4gIH1cbiAgcHJpdmF0ZSBhc3luYyBkaXNwbGF5VG9wU2NvcmVzKCkge1xuICAgIC8vIGNvbnNvbGUubG9nKHRoaXMuYW5zd2Vycy5lbnRyaWVzKCkpO1xuICAgIGNvbnN0IHNjb3JlcyA9IGF3YWl0IHRvcFNjb3Jlcyh0aGlzLnF1aXouaWQudG9TdHJpbmcoKSk7XG4gICAgY29uc29sZS5sb2coc2NvcmVzKTtcbiAgICBpZiAoc2NvcmVzID09PSBudWxsKSB7XG4gICAgICBlcnJvcihcIk5pZSB1ZGHFgm8gcG9icmHEhyBzacSZIG5hamxlcHN6eWNoIHd5bmlrw7N3XCIpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICBmb3IgKGNvbnN0IFtpLCBbdXNlciwgcmVzdWx0XV0gb2Ygc2NvcmVzLmVudHJpZXMoKSkge1xuICAgICAgY29uc3QgZGl2ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImRpdlwiKTtcbiAgICAgIGNvbnN0IGgzID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImgzXCIpO1xuICAgICAgLy8gZGl2LnRleHRDb250ZW50ID0gYCR7aSArIDF9LiBgO1xuICAgICAgY29uc3QgdGltZVAgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwicFwiKTtcbiAgICAgIHRpbWVQLnRleHRDb250ZW50ID0gYEN6YXM6ICR7UXVpelJlc3VsdHMuZm9ybWF0VGltZShcbiAgICAgICAgcmVzdWx0LmZpbmFsVGltZSAvIDEwMDBcbiAgICAgICl9c2A7XG4gICAgICBjb25zdCBhbnN3ZXJzUCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJwXCIpO1xuICAgICAgYW5zd2Vyc1AudGV4dENvbnRlbnQgPSBgUG9wcmF3bmUgb2Rwb3dpZWR6aTogJHtyZXN1bHQuYW5zd2Vycz8ucmVkdWNlKFxuICAgICAgICAoc3VtLCBhbnMpID0+IChhbnMuY29ycmVjdCA/IHN1bSArIDEgOiBzdW0pLFxuICAgICAgICAwXG4gICAgICApfSAvICR7cmVzdWx0LmFuc3dlcnM/Lmxlbmd0aH1gO1xuICAgICAgaDMudGV4dENvbnRlbnQgPSB1c2VyLnVzZXJuYW1lO1xuICAgICAgZGl2LmFwcGVuZENoaWxkKGgzKTtcbiAgICAgIGRpdi5hcHBlbmRDaGlsZCh0aW1lUCk7XG4gICAgICBkaXYuYXBwZW5kQ2hpbGQoYW5zd2Vyc1ApO1xuICAgICAgdG9wU2NvcmVzU2VjdGlvbi5hcHBlbmRDaGlsZChkaXYpO1xuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBEaXNwbGF5cyByZXN1bHRzLlxuICAgKi9cbiAgcHJpdmF0ZSBhc3luYyBkaXNwbGF5KCkge1xuICAgIHRoaXMuZGlzcGxheVVzZXJSZXN1bHRzKCk7XG4gICAgdGhpcy5kaXNwbGF5Q29ycmVjdEFuc3dlcnMoKTtcbiAgICBhd2FpdCB0aGlzLmRpc3BsYXlBdmVyYWdlVGltZXMoKTtcbiAgICBhd2FpdCB0aGlzLmRpc3BsYXlUb3BTY29yZXMoKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBGb3JtYXRzIHRpbWUsIHJvdW5kIHRvIDMgZGlnaXRzLlxuICAgKiBAcGFyYW0gdGltZSB0aW1lXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIGZvcm1hdFRpbWUodGltZTogbnVtYmVyKTogc3RyaW5nIHtcbiAgICByZXR1cm4gdGltZS50b0ZpeGVkKDMpLnRvU3RyaW5nKCk7XG4gIH1cblxuICAvKipcbiAgICogRGlzcGxheXMgcHJldmlvdXMgcmVzdWx0cy5cbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgYXN5bmMgZGlzcGxheVByZXZpb3VzUmVzdWx0cygpOiBQcm9taXNlPHZvaWQ+IHtcbiAgICBsZXQgcmVzdWx0cztcbiAgICB0cnkge1xuICAgICAgcmVzdWx0cyA9IGF3YWl0IGdldFByZXZSZXN1bHRzKCk7XG4gICAgfSBjYXRjaCAoXykge31cbiAgICBsZXQgaSA9IDE7XG4gICAgbGV0IGFueSA9IGZhbHNlO1xuICAgIGlmIChyZXN1bHRzIGluc3RhbmNlb2YgQXJyYXkpIHtcbiAgICAgIGZvciAoY29uc3QgcmVzIG9mIHJlc3VsdHMpIHtcbiAgICAgICAgYW55ID0gdHJ1ZTtcbiAgICAgICAgaWYgKGlzUmVzdWx0KHJlcykpIHtcbiAgICAgICAgICBjb25zb2xlLmxvZyhyZXMpO1xuICAgICAgICAgIGNvbnN0IHF1aXogPSBhd2FpdCBnZXRRdWl6V2l0aEFuc3dlcnMocmVzLnF1aXpJZCk7XG4gICAgICAgICAgaWYgKHF1aXopIHtcbiAgICAgICAgICAgIGNvbnN0IGZpbmFsVGltZSA9XG4gICAgICAgICAgICAgIHJlcy5maW5hbFRpbWUgK1xuICAgICAgICAgICAgICBxdWl6LnF1ZXN0aW9ucy5yZWR1Y2UoKHN1bSwgcXVlc3Rpb24sIGkpID0+IHtcbiAgICAgICAgICAgICAgICBjb25zb2xlLmxvZyhzdW0sIHJlcy5hbnN3ZXJzLCBpKTtcbiAgICAgICAgICAgICAgICBpZiAocmVzLmFuc3dlcnMgJiYgIXJlcy5hbnN3ZXJzW2ldLmNvcnJlY3QpIHtcbiAgICAgICAgICAgICAgICAgIHJldHVybiBzdW0gKyBxdWVzdGlvbi5wZW5hbHR5ICogMTAwMDtcbiAgICAgICAgICAgICAgICB9XG4gICAgICAgICAgICAgICAgcmV0dXJuIHN1bTtcbiAgICAgICAgICAgICAgfSwgMCk7XG4gICAgICAgICAgICBjb25zdCByb3cgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwiZGl2XCIpO1xuICAgICAgICAgICAgY29uc3QgbmFtZUNvbCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJkaXZcIik7XG4gICAgICAgICAgICBjb25zdCByZXN1bHRDb2wgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwiZGl2XCIpO1xuICAgICAgICAgICAgY29uc3QgZGV0YWlsQ29sID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImRpdlwiKTtcbiAgICAgICAgICAgIGNvbnN0IGxpbmsgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwiYVwiKTtcblxuICAgICAgICAgICAgbGluay5jbGFzc05hbWUgPSBcImJ0biBidG4tcHJpbWFyeSBxdWl6LWluZm8tYnV0dG9uXCI7XG4gICAgICAgICAgICBsaW5rLmRhdGFzZXQucXVpeklkID0gcXVpei5pZDtcbiAgICAgICAgICAgIGxpbmsudGV4dENvbnRlbnQgPSBcIkluZm9cIjtcbiAgICAgICAgICAgIGxpbmsuaHJlZiA9IFwiL3Jlc3VsdHMvXCIgKyBxdWl6LmlkO1xuXG4gICAgICAgICAgICByZXN1bHRDb2wuY2xhc3NOYW1lID1cbiAgICAgICAgICAgICAgXCJjb2wtNCBkLWZsZXgganVzdGlmeS1jb250ZW50LWNlbnRlciBhbGlnbi1pdGVtcy1jZW50ZXJcIjtcbiAgICAgICAgICAgIHJlc3VsdENvbC50ZXh0Q29udGVudCA9XG4gICAgICAgICAgICAgIFF1aXpSZXN1bHRzLmZvcm1hdFRpbWUoZmluYWxUaW1lIC8gMTAwMCkgKyBcInNcIjtcblxuICAgICAgICAgICAgbmFtZUNvbC5jbGFzc05hbWUgPVxuICAgICAgICAgICAgICBcImNvbC02IGQtZmxleCBqdXN0aWZ5LWNvbnRlbnQtY2VudGVyIGFsaWduLWl0ZW1zLWNlbnRlclwiO1xuICAgICAgICAgICAgbmFtZUNvbC50ZXh0Q29udGVudCA9IHF1aXouZGVzYztcblxuICAgICAgICAgICAgZGV0YWlsQ29sLmNsYXNzTmFtZSA9XG4gICAgICAgICAgICAgIFwiY29sLTIgZC1mbGV4IGp1c3RpZnktY29udGVudC1jZW50ZXIgYWxpZ24taXRlbXMtY2VudGVyXCI7XG4gICAgICAgICAgICBkZXRhaWxDb2wuYXBwZW5kQ2hpbGQobGluayk7XG4gICAgICAgICAgICByb3cuY2xhc3NOYW1lID0gYHJvdyBwcmV2LXJlc3VsdCAke2kgJSAyID09IDAgPyBcImV2ZW5cIiA6IFwiXCJ9YDtcblxuICAgICAgICAgICAgcm93LmFwcGVuZENoaWxkKG5hbWVDb2wpO1xuICAgICAgICAgICAgcm93LmFwcGVuZENoaWxkKHJlc3VsdENvbCk7XG4gICAgICAgICAgICByb3cuYXBwZW5kQ2hpbGQoZGV0YWlsQ29sKTtcblxuICAgICAgICAgICAgcXVpelByZXZSZXN1bHRzU2VjdGlvbi5hcHBlbmRDaGlsZChyb3cpO1xuICAgICAgICAgICAgaSsrO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cbiAgICBpZiAoIWFueSkge1xuICAgICAgY29uc3Qgcm93ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImRpdlwiKTtcbiAgICAgIHJvdy5jbGFzc05hbWUgPSBgcm93IHByZXYtcmVzdWx0IGQtZmxleCBqdXN0aWZ5LWNvbnRlbnQtY2VudGVyYDtcbiAgICAgIHJvdy50ZXh0Q29udGVudCA9XG4gICAgICAgIFwiSmVzemN6ZSBuaWUgcm96d2nEhXphxYJlxbwgxbxhZG5lZ28gcXVpenUsIG5hIGNvIGN6ZWthc3o/IDopXCI7XG4gICAgICBxdWl6UHJldlJlc3VsdHNTZWN0aW9uLmFwcGVuZENoaWxkKHJvdyk7XG4gICAgfVxuICB9XG59XG4iLCJpbXBvcnQgeyBSZXN1bHQgfSBmcm9tIFwiLi9RdWl6UmVzdWx0c1wiO1xuaW1wb3J0IHsgUmF3UXVpelJlc3VsdCB9IGZyb20gXCIuLi8uLi9yZXN1bHRzXCI7XG5pbXBvcnQgeyBjc3JmVG9rZW4gfSBmcm9tIFwiLi91dGlsXCI7XG5pbXBvcnQgeyBVc2VyIH0gZnJvbSBcIi4uLy4uL2xvZ2luXCI7XG5cbmV4cG9ydCBpbnRlcmZhY2UgUXVlc3Rpb25Ob0Fuc3dlciB7XG4gIHByb21wdDogc3RyaW5nO1xuICBwZW5hbHR5OiBudW1iZXI7XG59XG5cbmV4cG9ydCBpbnRlcmZhY2UgUXVlc3Rpb25XaXRoQW5zd2VycyB7XG4gIHByb21wdDogc3RyaW5nO1xuICBwZW5hbHR5OiBudW1iZXI7XG4gIGFuc3dlcjogbnVtYmVyO1xufVxuXG5leHBvcnQgaW50ZXJmYWNlIFF1aXpOb0Fuc3dlcnMge1xuICBkZXNjOiBzdHJpbmc7XG4gIHF1ZXN0aW9uczogUXVlc3Rpb25Ob0Fuc3dlcltdO1xuICBpZDogc3RyaW5nO1xufVxuXG5leHBvcnQgaW50ZXJmYWNlIFF1aXpXaXRoQW5zd2VycyB7XG4gIGRlc2M6IHN0cmluZztcbiAgcXVlc3Rpb25zOiBRdWVzdGlvbldpdGhBbnN3ZXJzW107XG4gIGlkOiBzdHJpbmc7XG59XG5cbmV4cG9ydCBpbnRlcmZhY2UgQW5zd2VyIHtcbiAgYW5zd2VyOiBudW1iZXIgfCBudWxsO1xuICB0aW1lOiBudW1iZXI7XG4gIGNvcnJlY3Q6IGJvb2xlYW47XG59XG5cbmNvbnN0IGlzUXVlc3Rpb24gPSAocXVlc3Rpb246IGFueSk6IHF1ZXN0aW9uIGlzIFF1ZXN0aW9uTm9BbnN3ZXIgPT4ge1xuICByZXR1cm4gcXVlc3Rpb24gJiYgcXVlc3Rpb24ucHJvbXB0ICYmIHF1ZXN0aW9uLmFuc3dlciAmJiBxdWVzdGlvbi5wZW5hbHR5O1xufTtcblxuZXhwb3J0IGNvbnN0IGlzUXVpeiA9IChxdWl6OiBhbnkpOiBxdWl6IGlzIFF1aXpOb0Fuc3dlcnMgPT4ge1xuICByZXR1cm4gcXVpei5kZXNjICYmIHF1aXoucXVlc3Rpb25zICYmIHF1aXouaWQgIT09IHVuZGVmaW5lZDtcbn07XG5jb25zdCBpc1F1aXpXaXRoQW5zd2VycyA9IChxdWl6OiBhbnkpOiBxdWl6IGlzIFF1aXpXaXRoQW5zd2VycyA9PiB7XG4gIHJldHVybiBxdWl6LmRlc2MgJiYgcXVpei5xdWVzdGlvbnMgJiYgcXVpei5pZCAhPT0gdW5kZWZpbmVkO1xufTtcblxuLyoqXG4gKiBSZXR1cm5zIHF1aXogd2l0aCBnaXZlbiBpZCBvciBudWxsIG9uIGVycm9yLlxuICogQHBhcmFtIGlkIHF1aXogaWRcbiAqL1xuZXhwb3J0IGNvbnN0IGdldFF1aXogPSBhc3luYyAoaWQ6IHN0cmluZyk6IFByb21pc2U8UXVpek5vQW5zd2VycyB8IG51bGw+ID0+IHtcbiAgY29uc29sZS5sb2coaWQpO1xuICB0cnkge1xuICAgIGNvbnN0IHF1aXpSYXcgPSBhd2FpdCBmZXRjaChcIi9nZXRfcXVpei9cIiArIGlkKTtcbiAgICBjb25zdCBxdWl6ID0gYXdhaXQgcXVpelJhdy5qc29uKCk7XG5cbiAgICBpZiAoaXNRdWl6KHF1aXopKSB7XG4gICAgICByZXR1cm4gcXVpejtcbiAgICB9XG4gIH0gY2F0Y2ggKF8pIHt9XG5cbiAgcmV0dXJuIG51bGw7XG59O1xuXG4vKipcbiAqIFJldHVybnMgcXVpeiB3aXRoIGdpdmVuIGlkIG9yIG51bGwgb24gZXJyb3IuXG4gKiBAcGFyYW0gaWQgcXVpeiBpZFxuICovXG5leHBvcnQgY29uc3QgZ2V0UXVpeldpdGhBbnN3ZXJzID0gYXN5bmMgKFxuICBpZDogc3RyaW5nXG4pOiBQcm9taXNlPFF1aXpXaXRoQW5zd2VycyB8IG51bGw+ID0+IHtcbiAgdHJ5IHtcbiAgICBjb25zdCBxdWl6UmF3ID0gYXdhaXQgZmV0Y2goXCIvZ2V0X3F1aXpfd2l0aF9hbnN3ZXJzL1wiICsgaWQpO1xuICAgIGNvbnN0IHF1aXogPSBhd2FpdCBxdWl6UmF3Lmpzb24oKTtcblxuICAgIGlmIChpc1F1aXpXaXRoQW5zd2VycyhxdWl6KSkge1xuICAgICAgcmV0dXJuIHF1aXo7XG4gICAgfVxuICB9IGNhdGNoIChfKSB7fVxuXG4gIHJldHVybiBudWxsO1xufTtcblxuLyoqXG4gKiBHZW5lcmF0b3IgdGhhdCByZXR1cm5zIGFsbCBxdWl6ZXMuXG4gKi9cbmV4cG9ydCBhc3luYyBmdW5jdGlvbiogZ2V0UXVpemVzKCk6IEFzeW5jR2VuZXJhdG9yPFF1aXpOb0Fuc3dlcnM+IHtcbiAgY29uc3QgcXVpelJhdyA9IGF3YWl0IGZldGNoKFwiL2dldF9xdWl6ZXNcIik7XG4gIGNvbnN0IG9iaiA9IGF3YWl0IHF1aXpSYXcuanNvbigpO1xuICBpZiAob2JqIGluc3RhbmNlb2YgQXJyYXkpIHtcbiAgICBmb3IgKGNvbnN0IHF1aXogb2Ygb2JqKSB7XG4gICAgICBpZiAoaXNRdWl6KHF1aXopKSB7XG4gICAgICAgIHlpZWxkIHF1aXo7XG4gICAgICB9XG4gICAgfVxuICB9XG59XG5cbmV4cG9ydCBjb25zdCBnZXRQcmV2UmVzdWx0cyA9IGFzeW5jICgpOiBQcm9taXNlPFJlc3VsdFtdPiA9PiB7XG4gIGNvbnN0IHJlc3VsdHMgPSBhd2FpdCBmZXRjaChcIi9wcmV2X3Jlc3VsdHNcIik7XG4gIGNvbnN0IG9iaiA9IGF3YWl0IHJlc3VsdHMuanNvbigpO1xuICBpZiAob2JqIGluc3RhbmNlb2YgQXJyYXkpIHtcbiAgICByZXR1cm4gPFJlc3VsdFtdPm9iajtcbiAgfVxuICByZXR1cm4gW107XG59O1xuXG5jb25zdCBnZXRSYXdRdWl6UmVzdWx0cyA9IChhbnN3ZXJzOiBBbnN3ZXJbXSk6IFJhd1F1aXpSZXN1bHQgPT4ge1xuICBjb25zdCB0b3RhbFRpbWUgPSBhbnN3ZXJzLnJlZHVjZSgoc3VtLCBhbnMpID0+IHN1bSArIGFucy50aW1lLCAwKTtcbiAgY29uc3QgcXI6IFJhd1F1aXpSZXN1bHQgPSB7IGFuc3dlcnM6IFtdLCB0aW1lczogW10gfTtcbiAgYW5zd2Vycy5mb3JFYWNoKChhbnMpID0+IHtcbiAgICBxci5hbnN3ZXJzLnB1c2goPG51bWJlcj5hbnMuYW5zd2VyKTtcbiAgICBxci50aW1lcy5wdXNoKGFucy50aW1lIC8gdG90YWxUaW1lKTtcbiAgfSk7XG5cbiAgcmV0dXJuIHFyO1xufTtcbmV4cG9ydCBjb25zdCBzYXZlUmVzdWx0cyA9IGFzeW5jIChhbnN3ZXJzOiBBbnN3ZXJbXSwgcXVpeklkOiBzdHJpbmcpID0+IHtcbiAgY29uc3QgcXIgPSBnZXRSYXdRdWl6UmVzdWx0cyhhbnN3ZXJzKTtcbiAgY29uc29sZS5sb2cocXIsIGFuc3dlcnMsIEpTT04uc3RyaW5naWZ5KHFyKSk7XG4gIGF3YWl0IGZldGNoKFwiL3Bvc3RfcmVzdWx0cy9cIiArIHF1aXpJZCwge1xuICAgIG1ldGhvZDogXCJQT1NUXCIsXG4gICAgY3JlZGVudGlhbHM6IFwic2FtZS1vcmlnaW5cIixcbiAgICBoZWFkZXJzOiB7XG4gICAgICBcIkNvbnRlbnQtVHlwZVwiOiBcImFwcGxpY2F0aW9uL2pzb25cIixcbiAgICAgIFwiQ1NSRi1Ub2tlblwiOiBjc3JmVG9rZW4sXG4gICAgfSxcbiAgICBib2R5OiBKU09OLnN0cmluZ2lmeShxciksXG4gIH0pO1xufTtcblxuZXhwb3J0IGNvbnN0IGdldEFuc3dlcnMgPSBhc3luYyAocXVpeklkOiBzdHJpbmcpOiBQcm9taXNlPEFuc3dlcltdIHwgbnVsbD4gPT4ge1xuICB0cnkge1xuICAgIGNvbnN0IHJhdyA9IGF3YWl0IGZldGNoKFwiL2dldF9hbnN3ZXJzL1wiICsgcXVpeklkKTtcbiAgICBjb25zdCBvYmogPSBhd2FpdCByYXcuanNvbigpO1xuICAgIGlmIChvYmogaW5zdGFuY2VvZiBBcnJheSkge1xuICAgICAgcmV0dXJuIDxBbnN3ZXJbXT5vYmo7XG4gICAgfVxuICB9IGNhdGNoIChfKSB7fVxuICByZXR1cm4gbnVsbDtcbn07XG5cbmV4cG9ydCBjb25zdCBhdmVyYWdlVGltZXMgPSBhc3luYyAoXG4gIHF1aXpJZDogc3RyaW5nXG4pOiBQcm9taXNlPG51bWJlcltdIHwgbnVsbD4gPT4ge1xuICB0cnkge1xuICAgIGNvbnN0IHJhdyA9IGF3YWl0IGZldGNoKFwiL2dldF9hbnN3ZXJzX21lYW5fdGltZS9cIiArIHF1aXpJZCk7XG4gICAgY29uc3Qgb2JqID0gYXdhaXQgcmF3Lmpzb24oKTtcbiAgICBpZiAob2JqIGluc3RhbmNlb2YgQXJyYXkpIHtcbiAgICAgIHJldHVybiA8bnVtYmVyW10+b2JqO1xuICAgIH1cbiAgfSBjYXRjaCAoXykge31cbiAgcmV0dXJuIG51bGw7XG59O1xuZXhwb3J0IGNvbnN0IHRvcFNjb3JlcyA9IGFzeW5jIChcbiAgcXVpeklkOiBzdHJpbmdcbik6IFByb21pc2U8W1VzZXIsIFJlc3VsdF1bXSB8IG51bGw+ID0+IHtcbiAgdHJ5IHtcbiAgICBjb25zdCByYXcgPSBhd2FpdCBmZXRjaChcIi90b3Bfc2NvcmVzL1wiICsgcXVpeklkKTtcbiAgICBjb25zdCBvYmogPSBhd2FpdCByYXcuanNvbigpO1xuICAgIGlmIChvYmogaW5zdGFuY2VvZiBBcnJheSkge1xuICAgICAgcmV0dXJuIDxbVXNlciwgUmVzdWx0XVtdPm9iajtcbiAgICB9XG4gIH0gY2F0Y2ggKF8pIHt9XG4gIHJldHVybiBudWxsO1xufTtcbiIsImltcG9ydCB7IGdldFF1aXpXaXRoQW5zd2VycywgZ2V0QW5zd2VycyB9IGZyb20gXCIuL3F1aXpcIjtcbmltcG9ydCB7IFF1aXpSZXN1bHRzIH0gZnJvbSBcIi4vUXVpelJlc3VsdHNcIjtcbmNvbnN0IG1haW4gPSBhc3luYyAoKSA9PiB7XG4gIGNvbnN0IHF1aXpJZCA9IHdpbmRvdy5sb2NhdGlvbi5wYXRobmFtZVxuICAgIC5zcGxpdChcIi9cIilcbiAgICAuZmlsdGVyKChhKSA9PiBhICE9PSBcIlwiKVxuICAgIC5zbGljZSgtMSlbMF07XG5cbiAgY29uc3QgcXVpeiA9IGF3YWl0IGdldFF1aXpXaXRoQW5zd2VycyhxdWl6SWQpO1xuICBjb25zdCBhbnN3ZXJzID0gYXdhaXQgZ2V0QW5zd2VycyhxdWl6SWQpO1xuICBpZiAocXVpeiA9PT0gbnVsbCB8fCBhbnN3ZXJzID09PSBudWxsKSB7XG4gICAgYWxlcnQoXCJUZW4gdGVzdCBuaWUgaXN0bmllamVcIik7XG4gICAgd2luZG93LmxvY2F0aW9uLmhyZWYgPSBcIi9cIjtcbiAgICByZXR1cm47XG4gIH1cbiAgY29uc3QgcXVpelJlc3VsdHMgPSBuZXcgUXVpelJlc3VsdHMoYW5zd2VycywgcXVpeik7XG4gIHF1aXpSZXN1bHRzLm1hcmsoKTtcbn07XG5tYWluKCk7XG4iLCIvKipcbiAqIEZ1bmN0aW9uIHRoYXQgbG9ncyBlcnJvcnMuXG4gKiBNYXkgaW4gdGhlIGZ1dHVyZSBiZSBvdmVyd3JpdGVuIHRvIGZ1bmN0aW9uIHRoYXQgZGlzcGxheSBlcnJvciBvbiBzY3JlZW4uXG4gKi9cbmV4cG9ydCBjb25zdCBlcnJvciA9IChzOiBzdHJpbmcpID0+IGFsZXJ0KHMpO1xuXG5leHBvcnQgY29uc3QgY3NyZlRva2VuID1cbiAgZG9jdW1lbnQ/LnF1ZXJ5U2VsZWN0b3IoJ21ldGFbbmFtZT1cImNzcmYtdG9rZW5cIl0nKT8uZ2V0QXR0cmlidXRlKFwiY29udGVudFwiKSB8fFxuICBcIlwiO1xuIl0sInNvdXJjZVJvb3QiOiIifQ==