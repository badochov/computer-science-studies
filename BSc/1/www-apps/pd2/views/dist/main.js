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
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/main.ts");
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

/***/ "./src/QuizRun.ts":
/*!************************!*\
  !*** ./src/QuizRun.ts ***!
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
exports.QuizRun = void 0;
const quiz_1 = __webpack_require__(/*! ./quiz */ "./src/quiz.ts");
const HTMLElements_1 = __webpack_require__(/*! ./HTMLElements */ "./src/HTMLElements.ts");
const util_1 = __webpack_require__(/*! ./util */ "./src/util.ts");
class QuizRun {
    /**
     * @param quizId chosen quiz id
     */
    constructor(quiz = null) {
        this.quiz = quiz;
        this.questionNumber = 0;
        this.answerStartTime = 0;
        this.answers = [];
        this.penalties = [];
        /**
         * Saves contestants answer to current question.
         */
        this.saveAnswer = () => {
            const prevAnswer = this.answers[this.questionNumber];
            if (prevAnswer == null) {
                return;
            }
            const answer = parseInt(HTMLElements_1.answerInput.value);
            if (isNaN(answer)) {
                prevAnswer.answer = null;
            }
            else {
                prevAnswer.answer = answer;
            }
        };
        /**
         * Saves contestants answer time to current question.
         */
        this.saveAnswerTime = () => {
            const endTime = performance.now();
            const prevAnswer = this.answers[this.questionNumber];
            if (prevAnswer == null) {
                return;
            }
            prevAnswer.time += endTime - this.answerStartTime;
        };
        if (this.quiz == null) {
            console.log(util_1.error);
            util_1.error("Ten quiz już jest rozwiązany przez Ciebie");
            return;
        }
        else if (this.quiz.questions.length == 0) {
            util_1.error("Podany quiz ma za mało pytań");
            return;
        }
        HTMLElements_1.quizChoiseSection.style.display = "none";
        HTMLElements_1.quizPrevResultsSection.style.display = "none";
        HTMLElements_1.quizDescSection.style.display = "block";
        HTMLElements_1.descParagraph.textContent = this.quiz.desc;
        this.quiz.questions.forEach((_, i) => {
            this.answers.push({ time: 0, answer: null, correct: false });
            this.penalties.push(0);
            const access = document.createElement("span");
            access.textContent = (i + 1).toString();
            access.className = "quick-access-node";
            access.dataset.question = i.toString();
            access.onclick = () => {
                this.saveAnswerTime();
                this.questionNumber = i;
                this.changeQuestion();
            };
            HTMLElements_1.quickAccessDiv.appendChild(access);
        });
        this.bindEventHandlers();
        HTMLElements_1.excerciseSection.style.display = "block";
        this.answerStartTime = performance.now();
        this.changeQuestion();
    }
    /**
     * Binds neccessary event handlers.
     */
    bindEventHandlers() {
        HTMLElements_1.answerInput.oninput = () => {
            this.saveAnswer();
            HTMLElements_1.stopButton.disabled = this.answers.some((a) => a.answer === null);
            const quickAccessSpan = document.querySelector(`[data-question="${this.questionNumber}"]`);
            if (this.answers[this.questionNumber].answer === null) {
                quickAccessSpan.className = quickAccessSpan.className.replace(/answered/, "");
            }
            else {
                quickAccessSpan.className += " answered";
            }
        };
        HTMLElements_1.nextButton.onclick = () => {
            this.saveAnswerTime();
            this.questionNumber++;
            this.changeQuestion();
        };
        HTMLElements_1.previousButton.onclick = () => {
            this.saveAnswerTime();
            this.questionNumber--;
            this.changeQuestion();
        };
        HTMLElements_1.stopButton.onclick = () => {
            this.endQuiz();
        };
        HTMLElements_1.answerForm.onsubmit = () => {
            if (HTMLElements_1.stopButton.disabled) {
                HTMLElements_1.nextButton.click();
            }
            else {
                HTMLElements_1.stopButton.click();
            }
            return false;
        };
    }
    /**
     * Changes current question to given.
     * Sets prevAnswerStartTime to  current timestamp.
     */
    changeQuestion() {
        if (this.quiz === null) {
            return;
        }
        const question = this.quiz.questions[this.questionNumber];
        HTMLElements_1.excerciseNumberSpan.textContent = (this.questionNumber + 1).toString();
        HTMLElements_1.promptSpan.textContent = question.prompt;
        HTMLElements_1.penaltyTimeSpan.textContent = question.penalty.toString();
        const answer = this.answers[this.questionNumber].answer;
        HTMLElements_1.answerInput.value = answer === null ? "" : answer.toString();
        HTMLElements_1.previousButton.disabled = this.questionNumber === 0;
        HTMLElements_1.nextButton.disabled =
            this.questionNumber === this.quiz.questions.length - 1;
        this.answerStartTime = performance.now();
    }
    /**
     * End quiz. Displays results.
     */
    endQuiz() {
        var _a, _b;
        return __awaiter(this, void 0, void 0, function* () {
            this.saveAnswerTime();
            HTMLElements_1.excerciseSection.style.display = "none";
            yield quiz_1.saveResults(this.answers, ((_a = this.quiz) === null || _a === void 0 ? void 0 : _a.id) || "-1");
            window.location.href = "/results/" + ((_b = this.quiz) === null || _b === void 0 ? void 0 : _b.id);
        });
    }
}
exports.QuizRun = QuizRun;


/***/ }),

/***/ "./src/main.ts":
/*!*********************!*\
  !*** ./src/main.ts ***!
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
var __asyncValues = (this && this.__asyncValues) || function (o) {
    if (!Symbol.asyncIterator) throw new TypeError("Symbol.asyncIterator is not defined.");
    var m = o[Symbol.asyncIterator], i;
    return m ? m.call(o) : (o = typeof __values === "function" ? __values(o) : o[Symbol.iterator](), i = {}, verb("next"), verb("throw"), verb("return"), i[Symbol.asyncIterator] = function () { return this; }, i);
    function verb(n) { i[n] = o[n] && function (v) { return new Promise(function (resolve, reject) { v = o[n](v), settle(resolve, reject, v.done, v.value); }); }; }
    function settle(resolve, reject, d, v) { Promise.resolve(v).then(function(v) { resolve({ value: v, done: d }); }, reject); }
};
Object.defineProperty(exports, "__esModule", { value: true });
const quiz_1 = __webpack_require__(/*! ./quiz */ "./src/quiz.ts");
const QuizRun_1 = __webpack_require__(/*! ./QuizRun */ "./src/QuizRun.ts");
const QuizResults_1 = __webpack_require__(/*! ./QuizResults */ "./src/QuizResults.ts");
const HTMLElements_1 = __webpack_require__(/*! ./HTMLElements */ "./src/HTMLElements.ts");
/**
 * Adds quiz to table of available quizes.
 *
 * @param id quiz id
 * @param quiz quiz object
 */
const addQuizToTable = (quiz) => {
    const row = document.createElement("div");
    const nameCol = document.createElement("div");
    const actionCol = document.createElement("div");
    const button = document.createElement("button");
    button.className = "btn btn-primary quiz-start-button";
    button.dataset.quizId = quiz.id;
    button.textContent = "Start";
    actionCol.className =
        "col-4 d-flex justify-content-center align-items-center";
    actionCol.appendChild(button);
    nameCol.className = "col-8 d-flex justify-content-center align-items-center";
    nameCol.textContent = quiz.desc;
    row.className = "row";
    row.appendChild(nameCol);
    row.appendChild(actionCol);
    HTMLElements_1.quizesDiv.appendChild(row);
};
/**
 * Displays available quizes in the table.
 */
const displayQuizes = () => __awaiter(void 0, void 0, void 0, function* () {
    var e_1, _a;
    try {
        for (var _b = __asyncValues(quiz_1.getQuizes()), _c; _c = yield _b.next(), !_c.done;) {
            const quiz = _c.value;
            addQuizToTable(quiz);
        }
    }
    catch (e_1_1) { e_1 = { error: e_1_1 }; }
    finally {
        try {
            if (_c && !_c.done && (_a = _b.return)) yield _a.call(_b);
        }
        finally { if (e_1) throw e_1.error; }
    }
    const buttons = document.getElementsByClassName("quiz-start-button");
    Array.from(buttons).forEach((button) => {
        button.onclick = () => __awaiter(void 0, void 0, void 0, function* () {
            new QuizRun_1.QuizRun(yield quiz_1.getQuiz(button.dataset.quizId || ""));
        });
    });
});
displayQuizes();
QuizResults_1.QuizResults.displayPreviousResults();


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
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vd2VicGFjay9ib290c3RyYXAiLCJ3ZWJwYWNrOi8vLy4vc3JjL0hUTUxFbGVtZW50cy50cyIsIndlYnBhY2s6Ly8vLi9zcmMvUXVpelJlc3VsdHMudHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL1F1aXpSdW4udHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL21haW4udHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL3F1aXoudHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL3V0aWwudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IjtRQUFBO1FBQ0E7O1FBRUE7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTs7UUFFQTtRQUNBOztRQUVBO1FBQ0E7O1FBRUE7UUFDQTtRQUNBOzs7UUFHQTtRQUNBOztRQUVBO1FBQ0E7O1FBRUE7UUFDQTtRQUNBO1FBQ0EsMENBQTBDLGdDQUFnQztRQUMxRTtRQUNBOztRQUVBO1FBQ0E7UUFDQTtRQUNBLHdEQUF3RCxrQkFBa0I7UUFDMUU7UUFDQSxpREFBaUQsY0FBYztRQUMvRDs7UUFFQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0EseUNBQXlDLGlDQUFpQztRQUMxRSxnSEFBZ0gsbUJBQW1CLEVBQUU7UUFDckk7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7UUFDQSwyQkFBMkIsMEJBQTBCLEVBQUU7UUFDdkQsaUNBQWlDLGVBQWU7UUFDaEQ7UUFDQTtRQUNBOztRQUVBO1FBQ0Esc0RBQXNELCtEQUErRDs7UUFFckg7UUFDQTs7O1FBR0E7UUFDQTs7Ozs7Ozs7Ozs7Ozs7OztBQ2xGYSxtQkFBVyxHQUFHLFFBQVEsQ0FBQyxjQUFjLENBQ2hELFFBQVEsQ0FDVyxDQUFDO0FBQ1Qsc0JBQWMsR0FBRyxRQUFRLENBQUMsY0FBYyxDQUNuRCxVQUFVLENBQ1UsQ0FBQztBQUNWLGtCQUFVLEdBQUcsUUFBUSxDQUFDLGNBQWMsQ0FBQyxNQUFNLENB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QU0sQ0FBQyxDQUFDO1FBQ2pELE1BQU0sR0FBRyxHQUFHLE1BQU0sR0FBRyxDQUFDLElBQUksRUFBRSxDQUFDO1FBQzdCLElBQUksR0FBRyxZQUFZLEtBQUssRUFBRTtZQUN4QixPQUF5QixHQUFHLENBQUM7U0FDOUI7S0FDRjtJQUFDLE9BQU8sQ0FBQyxFQUFFLEdBQUU7SUFDZCxPQUFPLElBQUksQ0FBQztBQUNkLENBQUMsRUFBQzs7Ozs7Ozs7Ozs7Ozs7Ozs7QUNwS0Y7OztHQUdHO0FBQ1UsYUFBSyxHQUFHLENBQUMsQ0FBUyxFQUFFLEVBQUUsQ0FBQyxLQUFLLENBQUMsQ0FBQyxDQUFDLENBQUM7QUFFaEMsaUJBQVMsR0FDcEIsZUFBUSxhQUFSLFFBQVEsdUJBQVIsUUFBUSxDQUFFLGFBQWEsQ0FBQyx5QkFBeUIsMkNBQUcsWUFBWSxDQUFDLFNBQVM7SUFDMUUsRUFBRSxDQUFDIiwiZmlsZSI6Im1haW4uanMiLCJzb3VyY2VzQ29udGVudCI6WyIgXHQvLyBUaGUgbW9kdWxlIGNhY2hlXG4gXHR2YXIgaW5zdGFsbGVkTW9kdWxlcyA9IHt9O1xuXG4gXHQvLyBUaGUgcmVxdWlyZSBmdW5jdGlvblxuIFx0ZnVuY3Rpb24gX193ZWJwYWNrX3JlcXVpcmVfXyhtb2R1bGVJZCkge1xuXG4gXHRcdC8vIENoZWNrIGlmIG1vZHVsZSBpcyBpbiBjYWNoZVxuIFx0XHRpZihpbnN0YWxsZWRNb2R1bGVzW21vZHVsZUlkXSkge1xuIFx0XHRcdHJldHVybiBpbnN0YWxsZWRNb2R1bGVzW21vZHVsZUlkXS5leHBvcnRzO1xuIFx0XHR9XG4gXHRcdC8vIENyZWF0ZSBhIG5ldyBtb2R1bGUgKGFuZCBwdXQgaXQgaW50byB0aGUgY2FjaGUpXG4gXHRcdHZhciBtb2R1bGUgPSBpbnN0YWxsZWRNb2R1bGVzW21vZHVsZUlkXSA9IHtcbiBcdFx0XHRpOiBtb2R1bGVJZCxcbiBcdFx0XHRsOiBmYWxzZSxcbiBcdFx0XHRleHBvcnRzOiB7fVxuIFx0XHR9O1xuXG4gXHRcdC8vIEV4ZWN1dGUgdGhlIG1vZHVsZSBmdW5jdGlvblxuIFx0XHRtb2R1bGVzW21vZHVsZUlkXS5jYWxsKG1vZHVsZS5leHBvcnRzLCBtb2R1bGUsIG1vZHVsZS5leHBvcnRzLCBfX3dlYnBhY2tfcmVxdWlyZV9fKTtcblxuIFx0XHQvLyBGbGFnIHRoZSBtb2R1bGUgYXMgbG9hZGVkXG4gXHRcdG1vZHVsZS5sID0gdHJ1ZTtcblxuIFx0XHQvLyBSZXR1cm4gdGhlIGV4cG9ydHMgb2YgdGhlIG1vZHVsZVxuIFx0XHRyZXR1cm4gbW9kdWxlLmV4cG9ydHM7XG4gXHR9XG5cblxuIFx0Ly8gZXhwb3NlIHRoZSBtb2R1bGVzIG9iamVjdCAoX193ZWJwYWNrX21vZHVsZXNfXylcbiBcdF9fd2VicGFja19yZXF1aXJlX18ubSA9IG1vZHVsZXM7XG5cbiBcdC8vIGV4cG9zZSB0aGUgbW9kdWxlIGNhY2hlXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLmMgPSBpbnN0YWxsZWRNb2R1bGVzO1xuXG4gXHQvLyBkZWZpbmUgZ2V0dGVyIGZ1bmN0aW9uIGZvciBoYXJtb255IGV4cG9ydHNcbiBcdF9fd2VicGFja19yZXF1aXJlX18uZCA9IGZ1bmN0aW9uKGV4cG9ydHMsIG5hbWUsIGdldHRlcikge1xuIFx0XHRpZighX193ZWJwYWNrX3JlcXVpcmVfXy5vKGV4cG9ydHMsIG5hbWUpKSB7XG4gXHRcdFx0T2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIG5hbWUsIHsgZW51bWVyYWJsZTogdHJ1ZSwgZ2V0OiBnZXR0ZXIgfSk7XG4gXHRcdH1cbiBcdH07XG5cbiBcdC8vIGRlZmluZSBfX2VzTW9kdWxlIG9uIGV4cG9ydHNcbiBcdF9fd2VicGFja19yZXF1aXJlX18uciA9IGZ1bmN0aW9uKGV4cG9ydHMpIHtcbiBcdFx0aWYodHlwZW9mIFN5bWJvbCAhPT0gJ3VuZGVmaW5lZCcgJiYgU3ltYm9sLnRvU3RyaW5nVGFnKSB7XG4gXHRcdFx0T2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIFN5bWJvbC50b1N0cmluZ1RhZywgeyB2YWx1ZTogJ01vZHVsZScgfSk7XG4gXHRcdH1cbiBcdFx0T2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsICdfX2VzTW9kdWxlJywgeyB2YWx1ZTogdHJ1ZSB9KTtcbiBcdH07XG5cbiBcdC8vIGNyZWF0ZSBhIGZha2UgbmFtZXNwYWNlIG9iamVjdFxuIFx0Ly8gbW9kZSAmIDE6IHZhbHVlIGlzIGEgbW9kdWxlIGlkLCByZXF1aXJlIGl0XG4gXHQvLyBtb2RlICYgMjogbWVyZ2UgYWxsIHByb3BlcnRpZXMgb2YgdmFsdWUgaW50byB0aGUgbnNcbiBcdC8vIG1vZGUgJiA0OiByZXR1cm4gdmFsdWUgd2hlbiBhbHJlYWR5IG5zIG9iamVjdFxuIFx0Ly8gbW9kZSAmIDh8MTogYmVoYXZlIGxpa2UgcmVxdWlyZVxuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy50ID0gZnVuY3Rpb24odmFsdWUsIG1vZGUpIHtcbiBcdFx0aWYobW9kZSAmIDEpIHZhbHVlID0gX193ZWJwYWNrX3JlcXVpcmVfXyh2YWx1ZSk7XG4gXHRcdGlmKG1vZGUgJiA4KSByZXR1cm4gdmFsdWU7XG4gXHRcdGlmKChtb2RlICYgNCkgJiYgdHlwZW9mIHZhbHVlID09PSAnb2JqZWN0JyAmJiB2YWx1ZSAmJiB2YWx1ZS5fX2VzTW9kdWxlKSByZXR1cm4gdmFsdWU7XG4gXHRcdHZhciBucyA9IE9iamVjdC5jcmVhdGUobnVsbCk7XG4gXHRcdF9fd2VicGFja19yZXF1aXJlX18ucihucyk7XG4gXHRcdE9iamVjdC5kZWZpbmVQcm9wZXJ0eShucywgJ2RlZmF1bHQnLCB7IGVudW1lcmFibGU6IHRydWUsIHZhbHVlOiB2YWx1ZSB9KTtcbiBcdFx0aWYobW9kZSAmIDIgJiYgdHlwZW9mIHZhbHVlICE9ICdzdHJpbmcnKSBmb3IodmFyIGtleSBpbiB2YWx1ZSkgX193ZWJwYWNrX3JlcXVpcmVfXy5kKG5zLCBrZXksIGZ1bmN0aW9uKGtleSkgeyByZXR1cm4gdmFsdWVba2V5XTsgfS5iaW5kKG51bGwsIGtleSkpO1xuIFx0XHRyZXR1cm4gbnM7XG4gXHR9O1xuXG4gXHQvLyBnZXREZWZhdWx0RXhwb3J0IGZ1bmN0aW9uIGZvciBjb21wYXRpYmlsaXR5IHdpdGggbm9uLWhhcm1vbnkgbW9kdWxlc1xuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5uID0gZnVuY3Rpb24obW9kdWxlKSB7XG4gXHRcdHZhciBnZXR0ZXIgPSBtb2R1bGUgJiYgbW9kdWxlLl9fZXNNb2R1bGUgP1xuIFx0XHRcdGZ1bmN0aW9uIGdldERlZmF1bHQoKSB7IHJldHVybiBtb2R1bGVbJ2RlZmF1bHQnXTsgfSA6XG4gXHRcdFx0ZnVuY3Rpb24gZ2V0TW9kdWxlRXhwb3J0cygpIHsgcmV0dXJuIG1vZHVsZTsgfTtcbiBcdFx0X193ZWJwYWNrX3JlcXVpcmVfXy5kKGdldHRlciwgJ2EnLCBnZXR0ZXIpO1xuIFx0XHRyZXR1cm4gZ2V0dGVyO1xuIFx0fTtcblxuIFx0Ly8gT2JqZWN0LnByb3RvdHlwZS5oYXNPd25Qcm9wZXJ0eS5jYWxsXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLm8gPSBmdW5jdGlvbihvYmplY3QsIHByb3BlcnR5KSB7IHJldHVybiBPYmplY3QucHJvdG90eXBlLmhhc093blByb3BlcnR5LmNhbGwob2JqZWN0LCBwcm9wZXJ0eSk7IH07XG5cbiBcdC8vIF9fd2VicGFja19wdWJsaWNfcGF0aF9fXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLnAgPSBcIlwiO1xuXG5cbiBcdC8vIExvYWQgZW50cnkgbW9kdWxlIGFuZCByZXR1cm4gZXhwb3J0c1xuIFx0cmV0dXJuIF9fd2VicGFja19yZXF1aXJlX18oX193ZWJwYWNrX3JlcXVpcmVfXy5zID0gXCIuL3NyYy9tYWluLnRzXCIpO1xuIiwiZXhwb3J0IGNvbnN0IGFuc3dlcklucHV0ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwiYW5zd2VyXCJcbikgYXMgSFRNTElucHV0RWxlbWVudDtcbmV4cG9ydCBjb25zdCBwcmV2aW91c0J1dHRvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcInByZXZpb3VzXCJcbikgYXMgSFRNTEJ1dHRvbkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgbmV4dEJ1dHRvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFwibmV4dFwiKSBhcyBIVE1MQnV0dG9uRWxlbWVudDtcbmV4cG9ydCBjb25zdCBzdG9wQnV0dG9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXCJzdG9wXCIpIGFzIEhUTUxCdXR0b25FbGVtZW50O1xuZXhwb3J0IGNvbnN0IHByb21wdFNwYW4gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcInByb21wdFwiKSBhcyBIVE1MU3BhbkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgZXhjZXJjaXNlTnVtYmVyU3BhbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcImV4Y2VyY2lzZS1udW1iZXJcIlxuKSBhcyBIVE1MU3BhbkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgZXhjZXJjaXNlU2VjdGlvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcImV4Y2VyY2lzZVwiXG4pIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHF1aXpEZXNjU2VjdGlvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcInF1aXotZGVzY1wiXG4pIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHF1aXpDaG9pc2VTZWN0aW9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwicXVpei1jaG9pc2VcIlxuKSBhcyBIVE1MRGl2RWxlbWVudDtcbmV4cG9ydCBjb25zdCBwZW5hbHR5VGltZVNwYW4gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJwZW5hbHR5LXRpbWVcIlxuKSBhcyBIVE1MU3BhbkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgZGVzY1BhcmFncmFwaCA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcImRlc2NcIlxuKSBhcyBIVE1MUGFyYWdyYXBoRWxlbWVudDtcbmV4cG9ydCBjb25zdCByZXN1bHRzRGl2ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXCJyZXN1bHRzXCIpIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHF1aXplc0RpdiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFwicXVpemVzXCIpIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IGNvcnJlY3RuZXNzRGl2ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwiY29ycmVjdG5lc3NcIlxuKSBhcyBIVE1MRGl2RWxlbWVudDtcbmV4cG9ydCBjb25zdCByZXN1bHRTZWN0aW9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwicmVzdWx0XCJcbikgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3Qgc2F2ZVdpdGhTdGF0c0J1dHRvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcInNhdmUtd2l0aC1zdGF0c1wiXG4pIGFzIEhUTUxCdXR0b25FbGVtZW50O1xuZXhwb3J0IGNvbnN0IGFuc3dlckZvcm0gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcbiAgXCJhbnN3ZXItZm9ybVwiXG4pIGFzIEhUTUxGb3JtRWxlbWVudDtcblxuZXhwb3J0IGNvbnN0IGluaXRSZXN1bHRTcGFuID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwiaW5pdC1yZXN1bHRcIlxuKSBhcyBIVE1MU3BhbkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgZmluYWxSZXN1bHRTcGFuID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwiZmluYWwtcmVzdWx0XCJcbikgYXMgSFRNTFNwYW5FbGVtZW50O1xuXG5leHBvcnQgY29uc3QgcXVpelByZXZSZXN1bHRzU2VjdGlvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcInF1aXotcHJldi1yZXN1bHRzXCJcbikgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgY29ycmVjdEFuc3dlcnNTZWN0aW9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwiY29ycmVjdC1hbnN3ZXJzXCJcbikgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgYXZlcmFnZVRpbWVzU2VjdGlvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFxuICBcImF2ZXJhZ2UtdGltZXNcIlxuKSBhcyBIVE1MRGl2RWxlbWVudDtcbmV4cG9ydCBjb25zdCB0b3BTY29yZXNTZWN0aW9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwidG9wLXNjb3Jlc1wiXG4pIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHF1aWNrQWNjZXNzRGl2ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXG4gIFwicXVpY2stYWNjZXNzXCJcbikgYXMgSFRNTERpdkVsZW1lbnQ7XG4iLCJpbXBvcnQge1xuICBBbnN3ZXIsXG4gIFF1aXpXaXRoQW5zd2VycyxcbiAgZ2V0UXVpeldpdGhBbnN3ZXJzLFxuICBnZXRQcmV2UmVzdWx0cyxcbiAgYXZlcmFnZVRpbWVzLFxuICB0b3BTY29yZXMsXG59IGZyb20gXCIuL3F1aXpcIjtcbmltcG9ydCB7XG4gIHNhdmVXaXRoU3RhdHNCdXR0b24sXG4gIGluaXRSZXN1bHRTcGFuLFxuICBmaW5hbFJlc3VsdFNwYW4sXG4gIGNvcnJlY3RuZXNzRGl2LFxuICBxdWl6UHJldlJlc3VsdHNTZWN0aW9uLFxuICBjb3JyZWN0QW5zd2Vyc1NlY3Rpb24sXG4gIGF2ZXJhZ2VUaW1lc1NlY3Rpb24sXG4gIHRvcFNjb3Jlc1NlY3Rpb24sXG59IGZyb20gXCIuL0hUTUxFbGVtZW50c1wiO1xuaW1wb3J0IHsgZXJyb3IgfSBmcm9tIFwiLi91dGlsXCI7XG5cbmV4cG9ydCBpbnRlcmZhY2UgUmVzdWx0IHtcbiAgZmluYWxUaW1lOiBudW1iZXI7XG4gIGFuc3dlcnM/OiBBbnN3ZXJbXTtcbiAgcXVpeklkOiBzdHJpbmc7XG59XG5cbmNvbnN0IGlzUmVzdWx0ID0gKG9iajogYW55KTogb2JqIGlzIFJlc3VsdCA9PiB7XG4gIHJldHVybiBvYmouZmluYWxUaW1lICYmIG9iai5xdWl6SWQ7XG59O1xuXG5leHBvcnQgY2xhc3MgUXVpelJlc3VsdHMge1xuICBwcml2YXRlIHBlbmFsdGllczogbnVtYmVyW10gPSBbXTtcblxuICAvKipcbiAgICogQHBhcmFtIGFuc3dlcnMgYXJyYXkgb2YgbWFya2VkIGFuc3dlcnNcbiAgICogQHBhcmFtIHF1aXogcXVpeiBvYmplY3RcbiAgICogQHBhcmFtIHF1aXpJZCBjaG9zZW4gcXVpeiBpZFxuICAgKi9cbiAgY29uc3RydWN0b3IocHJpdmF0ZSBhbnN3ZXJzOiBBbnN3ZXJbXSwgcHJpdmF0ZSBxdWl6OiBRdWl6V2l0aEFuc3dlcnMpIHtcbiAgICB0aGlzLmJpbmRFdmVudEhhbmRsZXJzKCk7XG4gIH1cblxuICBwdWJsaWMgbWFyaygpIHtcbiAgICB0aGlzLmNoZWNrQW5zd2VycygpO1xuICAgIHRoaXMuZ2l2ZVBlbmFsdGllcygpO1xuICAgIHRoaXMuZGlzcGxheSgpO1xuICB9XG5cbiAgLyoqXG4gICAqIENoZWNrcyBhbnN3ZXJzLCBzZXRzIGNhcnJlY3QgYXJndW1lbnQgaW4gZWFjaCBhbnN3ZXIgdG8gZWl0aGVyIHRydWUgb3IgZmFsc2UuXG4gICAqL1xuICBwcml2YXRlIGNoZWNrQW5zd2VycygpIHtcbiAgICBmb3IgKGNvbnN0IGkgaW4gdGhpcy5hbnN3ZXJzKSB7XG4gICAgICB0aGlzLmFuc3dlcnNbaV0uY29ycmVjdCA9XG4gICAgICAgIHRoaXMuYW5zd2Vyc1tpXS5hbnN3ZXIgPT09IHRoaXMucXVpei5xdWVzdGlvbnNbaV0uYW5zd2VyO1xuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBHaXZlcyBwZW5hbHRpZXMgZm9yIGVhY2ggd3JvbmcgYW5zd2VyLlxuICAgKi9cbiAgcHJpdmF0ZSBnaXZlUGVuYWx0aWVzKCkge1xuICAgIGZvciAoY29uc3QgaSBpbiB0aGlzLmFuc3dlcnMpIHtcbiAgICAgIHRoaXMucGVuYWx0aWVzLnB1c2goXG4gICAgICAgIHRoaXMuYW5zd2Vyc1tpXS5jb3JyZWN0ID8gMCA6IHRoaXMucXVpei5xdWVzdGlvbnNbaV0ucGVuYWx0eVxuICAgICAgKTtcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQmluZCBuZWNjZXNzYXJ5IGV2ZW50IGhhbmRsZXJzLlxuICAgKi9cbiAgcHJpdmF0ZSBiaW5kRXZlbnRIYW5kbGVycygpIHtcbiAgICBzYXZlV2l0aFN0YXRzQnV0dG9uLm9uY2xpY2sgPSAoKSA9PiB7XG4gICAgICB0aGlzLmdvVG9NYWluU2NyZWVuKCk7XG4gICAgfTtcbiAgfVxuXG4gIC8qKlxuICAgKiBHb2VzIGJhY2sgdG8gbWFpbiBzY3JlZW4uXG4gICAqL1xuICBwcml2YXRlIGdvVG9NYWluU2NyZWVuKCkge1xuICAgIHdpbmRvdy5sb2NhdGlvbi5ocmVmID0gXCIvXCI7XG4gIH1cblxuICAvKipcbiAgICogQ2FsY3VsYXRlcyB0aW1lIHNwZW50IG9uIHF1aXouXG4gICAqL1xuICBwcml2YXRlIGdldCBpbml0VGltZSgpOiBudW1iZXIge1xuICAgIHJldHVybiB0aGlzLmFuc3dlcnMucmVkdWNlKChzdW0sIGFucykgPT4gc3VtICsgYW5zLnRpbWUgLyAxMDAwLCAwKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBDYWxjdWxhdGVzIGZpbmFsIHRpbWUsIGJhc2UgKyBwZW5hbHRpZXMuXG4gICAqL1xuICBwdWJsaWMgZ2V0IGZpbmFsVGltZSgpIHtcbiAgICByZXR1cm4gdGhpcy5wZW5hbHRpZXMucmVkdWNlKFxuICAgICAgKHN1bSwgcGVuYWx0eSkgPT4gc3VtICsgcGVuYWx0eSxcbiAgICAgIHRoaXMuaW5pdFRpbWVcbiAgICApO1xuICB9XG4gIHByaXZhdGUgZGlzcGxheVVzZXJSZXN1bHRzKCkge1xuICAgIGluaXRSZXN1bHRTcGFuLnRleHRDb250ZW50ID0gUXVpelJlc3VsdHMuZm9ybWF0VGltZSh0aGlzLmluaXRUaW1lKTtcbiAgICBmaW5hbFJlc3VsdFNwYW4udGV4dENvbnRlbnQgPSBRdWl6UmVzdWx0cy5mb3JtYXRUaW1lKHRoaXMuZmluYWxUaW1lKTtcblxuICAgIC8vIGNvbnNvbGUubG9nKHRoaXMuYW5zd2Vycy5lbnRyaWVzKCkpO1xuICAgIGZvciAoY29uc3QgW2ksIGFuc3dlcl0gb2YgdGhpcy5hbnN3ZXJzLmVudHJpZXMoKSkge1xuICAgICAgY29uc29sZS5sb2coaSwgYW5zd2VyKTtcbiAgICAgIGNvbnN0IGRpdiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJkaXZcIik7XG4gICAgICBkaXYudGV4dENvbnRlbnQgPSBgJHtpICsgMX0uIGA7XG4gICAgICBjb25zdCBzcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcInNwYW5cIik7XG4gICAgICBjb25zdCB0aW1lID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcInNwYW5cIik7XG4gICAgICBzcGFuLmNsYXNzTmFtZSA9IGFuc3dlci5jb3JyZWN0ID8gXCJjb3JyZWN0XCIgOiBcImluY29ycmVjdFwiO1xuICAgICAgc3Bhbi50ZXh0Q29udGVudCA9IGFuc3dlci5jb3JyZWN0XG4gICAgICAgID8gXCJQb3ByYXduaWUgOilcIlxuICAgICAgICA6IGBCxYLEhWQ6ICske3RoaXMucXVpei5xdWVzdGlvbnNbaV0ucGVuYWx0eX1zYDtcbiAgICAgIHRpbWUudGV4dENvbnRlbnQgPSBgICAke1F1aXpSZXN1bHRzLmZvcm1hdFRpbWUoYW5zd2VyLnRpbWUgLyAxMDAwKX0gc2A7XG4gICAgICBkaXYuYXBwZW5kQ2hpbGQoc3Bhbik7XG4gICAgICBkaXYuYXBwZW5kQ2hpbGQodGltZSk7XG4gICAgICBjb3JyZWN0bmVzc0Rpdi5hcHBlbmRDaGlsZChkaXYpO1xuICAgIH1cbiAgfVxuICBwcml2YXRlIGRpc3BsYXlDb3JyZWN0QW5zd2VycygpIHtcbiAgICAvLyBjb25zb2xlLmxvZyh0aGlzLmFuc3dlcnMuZW50cmllcygpKTtcbiAgICBmb3IgKGNvbnN0IFtpLCBxdWVzdGlvbl0gb2YgdGhpcy5xdWl6LnF1ZXN0aW9ucy5lbnRyaWVzKCkpIHtcbiAgICAgIGNvbnN0IGRpdiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJkaXZcIik7XG4gICAgICBkaXYudGV4dENvbnRlbnQgPSBgJHtpICsgMX0uIGA7XG4gICAgICBjb25zdCBzcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcInNwYW5cIik7XG4gICAgICBzcGFuLnRleHRDb250ZW50ID0gYCR7cXVlc3Rpb24ucHJvbXB0fSA9ICR7cXVlc3Rpb24uYW5zd2VyfWA7XG4gICAgICBkaXYuYXBwZW5kQ2hpbGQoc3Bhbik7XG4gICAgICBjb3JyZWN0QW5zd2Vyc1NlY3Rpb24uYXBwZW5kQ2hpbGQoZGl2KTtcbiAgICB9XG4gIH1cbiAgcHJpdmF0ZSBhc3luYyBkaXNwbGF5QXZlcmFnZVRpbWVzKCkge1xuICAgIC8vIGNvbnNvbGUubG9nKHRoaXMuYW5zd2Vycy5lbnRyaWVzKCkpO1xuICAgIGNvbnN0IHRpbWVzID0gYXdhaXQgYXZlcmFnZVRpbWVzKHRoaXMucXVpei5pZC50b1N0cmluZygpKTtcbiAgICBjb25zb2xlLmxvZyh0aW1lcyk7XG4gICAgaWYgKHRpbWVzID09PSBudWxsKSB7XG4gICAgICBlcnJvcihcIk5pZSB1ZGHFgm8gcG9icmHEhyBzacSZIMWbcmVkbmljaCBjemFzw7N3IG5hIG9kcG93aWVkxbpcIik7XG4gICAgICByZXR1cm47XG4gICAgfVxuICAgIGZvciAoY29uc3QgW2ksIHRpbWVdIG9mIHRpbWVzLmVudHJpZXMoKSkge1xuICAgICAgY29uc3QgZGl2ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImRpdlwiKTtcbiAgICAgIGRpdi50ZXh0Q29udGVudCA9IGAke2kgKyAxfS4gYDtcbiAgICAgIGNvbnN0IHNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwic3BhblwiKTtcbiAgICAgIHNwYW4udGV4dENvbnRlbnQgPSBgJHtcbiAgICAgICAgdGltZSA/IFF1aXpSZXN1bHRzLmZvcm1hdFRpbWUodGltZSAvIDEwMDApIDogXCItLS4tLVwiXG4gICAgICB9c2A7XG4gICAgICBkaXYuYXBwZW5kQ2hpbGQoc3Bhbik7XG4gICAgICBhdmVyYWdlVGltZXNTZWN0aW9uLmFwcGVuZENoaWxkKGRpdik7XG4gICAgfVxuICB9XG4gIHByaXZhdGUgYXN5bmMgZGlzcGxheVRvcFNjb3JlcygpIHtcbiAgICAvLyBjb25zb2xlLmxvZyh0aGlzLmFuc3dlcnMuZW50cmllcygpKTtcbiAgICBjb25zdCBzY29yZXMgPSBhd2FpdCB0b3BTY29yZXModGhpcy5xdWl6LmlkLnRvU3RyaW5nKCkpO1xuICAgIGNvbnNvbGUubG9nKHNjb3Jlcyk7XG4gICAgaWYgKHNjb3JlcyA9PT0gbnVsbCkge1xuICAgICAgZXJyb3IoXCJOaWUgdWRhxYJvIHBvYnJhxIcgc2nEmSBuYWpsZXBzenljaCB3eW5pa8Ozd1wiKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG4gICAgZm9yIChjb25zdCBbaSwgW3VzZXIsIHJlc3VsdF1dIG9mIHNjb3Jlcy5lbnRyaWVzKCkpIHtcbiAgICAgIGNvbnN0IGRpdiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJkaXZcIik7XG4gICAgICBjb25zdCBoMyA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJoM1wiKTtcbiAgICAgIC8vIGRpdi50ZXh0Q29udGVudCA9IGAke2kgKyAxfS4gYDtcbiAgICAgIGNvbnN0IHRpbWVQID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcInBcIik7XG4gICAgICB0aW1lUC50ZXh0Q29udGVudCA9IGBDemFzOiAke1F1aXpSZXN1bHRzLmZvcm1hdFRpbWUoXG4gICAgICAgIHJlc3VsdC5maW5hbFRpbWUgLyAxMDAwXG4gICAgICApfXNgO1xuICAgICAgY29uc3QgYW5zd2Vyc1AgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwicFwiKTtcbiAgICAgIGFuc3dlcnNQLnRleHRDb250ZW50ID0gYFBvcHJhd25lIG9kcG93aWVkemk6ICR7cmVzdWx0LmFuc3dlcnM/LnJlZHVjZShcbiAgICAgICAgKHN1bSwgYW5zKSA9PiAoYW5zLmNvcnJlY3QgPyBzdW0gKyAxIDogc3VtKSxcbiAgICAgICAgMFxuICAgICAgKX0gLyAke3Jlc3VsdC5hbnN3ZXJzPy5sZW5ndGh9YDtcbiAgICAgIGgzLnRleHRDb250ZW50ID0gdXNlci51c2VybmFtZTtcbiAgICAgIGRpdi5hcHBlbmRDaGlsZChoMyk7XG4gICAgICBkaXYuYXBwZW5kQ2hpbGQodGltZVApO1xuICAgICAgZGl2LmFwcGVuZENoaWxkKGFuc3dlcnNQKTtcbiAgICAgIHRvcFNjb3Jlc1NlY3Rpb24uYXBwZW5kQ2hpbGQoZGl2KTtcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogRGlzcGxheXMgcmVzdWx0cy5cbiAgICovXG4gIHByaXZhdGUgYXN5bmMgZGlzcGxheSgpIHtcbiAgICB0aGlzLmRpc3BsYXlVc2VyUmVzdWx0cygpO1xuICAgIHRoaXMuZGlzcGxheUNvcnJlY3RBbnN3ZXJzKCk7XG4gICAgYXdhaXQgdGhpcy5kaXNwbGF5QXZlcmFnZVRpbWVzKCk7XG4gICAgYXdhaXQgdGhpcy5kaXNwbGF5VG9wU2NvcmVzKCk7XG4gIH1cblxuICAvKipcbiAgICogRm9ybWF0cyB0aW1lLCByb3VuZCB0byAzIGRpZ2l0cy5cbiAgICogQHBhcmFtIHRpbWUgdGltZVxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBmb3JtYXRUaW1lKHRpbWU6IG51bWJlcik6IHN0cmluZyB7XG4gICAgcmV0dXJuIHRpbWUudG9GaXhlZCgzKS50b1N0cmluZygpO1xuICB9XG5cbiAgLyoqXG4gICAqIERpc3BsYXlzIHByZXZpb3VzIHJlc3VsdHMuXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIGFzeW5jIGRpc3BsYXlQcmV2aW91c1Jlc3VsdHMoKTogUHJvbWlzZTx2b2lkPiB7XG4gICAgbGV0IHJlc3VsdHM7XG4gICAgdHJ5IHtcbiAgICAgIHJlc3VsdHMgPSBhd2FpdCBnZXRQcmV2UmVzdWx0cygpO1xuICAgIH0gY2F0Y2ggKF8pIHt9XG4gICAgbGV0IGkgPSAxO1xuICAgIGxldCBhbnkgPSBmYWxzZTtcbiAgICBpZiAocmVzdWx0cyBpbnN0YW5jZW9mIEFycmF5KSB7XG4gICAgICBmb3IgKGNvbnN0IHJlcyBvZiByZXN1bHRzKSB7XG4gICAgICAgIGFueSA9IHRydWU7XG4gICAgICAgIGlmIChpc1Jlc3VsdChyZXMpKSB7XG4gICAgICAgICAgY29uc29sZS5sb2cocmVzKTtcbiAgICAgICAgICBjb25zdCBxdWl6ID0gYXdhaXQgZ2V0UXVpeldpdGhBbnN3ZXJzKHJlcy5xdWl6SWQpO1xuICAgICAgICAgIGlmIChxdWl6KSB7XG4gICAgICAgICAgICBjb25zdCBmaW5hbFRpbWUgPVxuICAgICAgICAgICAgICByZXMuZmluYWxUaW1lICtcbiAgICAgICAgICAgICAgcXVpei5xdWVzdGlvbnMucmVkdWNlKChzdW0sIHF1ZXN0aW9uLCBpKSA9PiB7XG4gICAgICAgICAgICAgICAgY29uc29sZS5sb2coc3VtLCByZXMuYW5zd2VycywgaSk7XG4gICAgICAgICAgICAgICAgaWYgKHJlcy5hbnN3ZXJzICYmICFyZXMuYW5zd2Vyc1tpXS5jb3JyZWN0KSB7XG4gICAgICAgICAgICAgICAgICByZXR1cm4gc3VtICsgcXVlc3Rpb24ucGVuYWx0eSAqIDEwMDA7XG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgICAgIHJldHVybiBzdW07XG4gICAgICAgICAgICAgIH0sIDApO1xuICAgICAgICAgICAgY29uc3Qgcm93ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImRpdlwiKTtcbiAgICAgICAgICAgIGNvbnN0IG5hbWVDb2wgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwiZGl2XCIpO1xuICAgICAgICAgICAgY29uc3QgcmVzdWx0Q29sID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImRpdlwiKTtcbiAgICAgICAgICAgIGNvbnN0IGRldGFpbENvbCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJkaXZcIik7XG4gICAgICAgICAgICBjb25zdCBsaW5rID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImFcIik7XG5cbiAgICAgICAgICAgIGxpbmsuY2xhc3NOYW1lID0gXCJidG4gYnRuLXByaW1hcnkgcXVpei1pbmZvLWJ1dHRvblwiO1xuICAgICAgICAgICAgbGluay5kYXRhc2V0LnF1aXpJZCA9IHF1aXouaWQ7XG4gICAgICAgICAgICBsaW5rLnRleHRDb250ZW50ID0gXCJJbmZvXCI7XG4gICAgICAgICAgICBsaW5rLmhyZWYgPSBcIi9yZXN1bHRzL1wiICsgcXVpei5pZDtcblxuICAgICAgICAgICAgcmVzdWx0Q29sLmNsYXNzTmFtZSA9XG4gICAgICAgICAgICAgIFwiY29sLTQgZC1mbGV4IGp1c3RpZnktY29udGVudC1jZW50ZXIgYWxpZ24taXRlbXMtY2VudGVyXCI7XG4gICAgICAgICAgICByZXN1bHRDb2wudGV4dENvbnRlbnQgPVxuICAgICAgICAgICAgICBRdWl6UmVzdWx0cy5mb3JtYXRUaW1lKGZpbmFsVGltZSAvIDEwMDApICsgXCJzXCI7XG5cbiAgICAgICAgICAgIG5hbWVDb2wuY2xhc3NOYW1lID1cbiAgICAgICAgICAgICAgXCJjb2wtNiBkLWZsZXgganVzdGlmeS1jb250ZW50LWNlbnRlciBhbGlnbi1pdGVtcy1jZW50ZXJcIjtcbiAgICAgICAgICAgIG5hbWVDb2wudGV4dENvbnRlbnQgPSBxdWl6LmRlc2M7XG5cbiAgICAgICAgICAgIGRldGFpbENvbC5jbGFzc05hbWUgPVxuICAgICAgICAgICAgICBcImNvbC0yIGQtZmxleCBqdXN0aWZ5LWNvbnRlbnQtY2VudGVyIGFsaWduLWl0ZW1zLWNlbnRlclwiO1xuICAgICAgICAgICAgZGV0YWlsQ29sLmFwcGVuZENoaWxkKGxpbmspO1xuICAgICAgICAgICAgcm93LmNsYXNzTmFtZSA9IGByb3cgcHJldi1yZXN1bHQgJHtpICUgMiA9PSAwID8gXCJldmVuXCIgOiBcIlwifWA7XG5cbiAgICAgICAgICAgIHJvdy5hcHBlbmRDaGlsZChuYW1lQ29sKTtcbiAgICAgICAgICAgIHJvdy5hcHBlbmRDaGlsZChyZXN1bHRDb2wpO1xuICAgICAgICAgICAgcm93LmFwcGVuZENoaWxkKGRldGFpbENvbCk7XG5cbiAgICAgICAgICAgIHF1aXpQcmV2UmVzdWx0c1NlY3Rpb24uYXBwZW5kQ2hpbGQocm93KTtcbiAgICAgICAgICAgIGkrKztcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG4gICAgaWYgKCFhbnkpIHtcbiAgICAgIGNvbnN0IHJvdyA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJkaXZcIik7XG4gICAgICByb3cuY2xhc3NOYW1lID0gYHJvdyBwcmV2LXJlc3VsdCBkLWZsZXgganVzdGlmeS1jb250ZW50LWNlbnRlcmA7XG4gICAgICByb3cudGV4dENvbnRlbnQgPVxuICAgICAgICBcIkplc3pjemUgbmllIHJvendpxIV6YcWCZcW8IMW8YWRuZWdvIHF1aXp1LCBuYSBjbyBjemVrYXN6PyA6KVwiO1xuICAgICAgcXVpelByZXZSZXN1bHRzU2VjdGlvbi5hcHBlbmRDaGlsZChyb3cpO1xuICAgIH1cbiAgfVxufVxuIiwiaW1wb3J0IHsgUXVpek5vQW5zd2VycywgQW5zd2VyLCBzYXZlUmVzdWx0cyB9IGZyb20gXCIuL3F1aXpcIjtcbmltcG9ydCB7XG4gIGRlc2NQYXJhZ3JhcGgsXG4gIGFuc3dlcklucHV0LFxuICBzdG9wQnV0dG9uLFxuICBleGNlcmNpc2VTZWN0aW9uLFxuICBuZXh0QnV0dG9uLFxuICBwcmV2aW91c0J1dHRvbixcbiAgZXhjZXJjaXNlTnVtYmVyU3BhbixcbiAgcHJvbXB0U3BhbixcbiAgcGVuYWx0eVRpbWVTcGFuLFxuICBxdWl6Q2hvaXNlU2VjdGlvbixcbiAgcXVpekRlc2NTZWN0aW9uLFxuICBhbnN3ZXJGb3JtLFxuICBxdWl6UHJldlJlc3VsdHNTZWN0aW9uLFxuICBxdWlja0FjY2Vzc0Rpdixcbn0gZnJvbSBcIi4vSFRNTEVsZW1lbnRzXCI7XG5pbXBvcnQgeyBlcnJvciB9IGZyb20gXCIuL3V0aWxcIjtcblxuZXhwb3J0IGNsYXNzIFF1aXpSdW4ge1xuICBwcml2YXRlIHF1ZXN0aW9uTnVtYmVyOiBudW1iZXIgPSAwO1xuICBwcml2YXRlIGFuc3dlclN0YXJ0VGltZTogbnVtYmVyID0gMDtcblxuICBwcml2YXRlIGFuc3dlcnM6IEFuc3dlcltdID0gW107XG4gIHByaXZhdGUgcGVuYWx0aWVzOiBudW1iZXJbXSA9IFtdO1xuXG4gIC8qKlxuICAgKiBAcGFyYW0gcXVpeklkIGNob3NlbiBxdWl6IGlkXG4gICAqL1xuICBjb25zdHJ1Y3Rvcihwcml2YXRlIHF1aXo6IFF1aXpOb0Fuc3dlcnMgfCBudWxsID0gbnVsbCkge1xuICAgIGlmICh0aGlzLnF1aXogPT0gbnVsbCkge1xuICAgICAgY29uc29sZS5sb2coZXJyb3IpO1xuICAgICAgZXJyb3IoXCJUZW4gcXVpeiBqdcW8IGplc3Qgcm96d2nEhXphbnkgcHJ6ZXogQ2llYmllXCIpO1xuICAgICAgcmV0dXJuO1xuICAgIH0gZWxzZSBpZiAodGhpcy5xdWl6LnF1ZXN0aW9ucy5sZW5ndGggPT0gMCkge1xuICAgICAgZXJyb3IoXCJQb2RhbnkgcXVpeiBtYSB6YSBtYcWCbyBweXRhxYRcIik7XG4gICAgICByZXR1cm47XG4gICAgfVxuICAgIHF1aXpDaG9pc2VTZWN0aW9uLnN0eWxlLmRpc3BsYXkgPSBcIm5vbmVcIjtcbiAgICBxdWl6UHJldlJlc3VsdHNTZWN0aW9uLnN0eWxlLmRpc3BsYXkgPSBcIm5vbmVcIjtcbiAgICBxdWl6RGVzY1NlY3Rpb24uc3R5bGUuZGlzcGxheSA9IFwiYmxvY2tcIjtcbiAgICBkZXNjUGFyYWdyYXBoLnRleHRDb250ZW50ID0gdGhpcy5xdWl6LmRlc2M7XG5cbiAgICB0aGlzLnF1aXoucXVlc3Rpb25zLmZvckVhY2goKF8sIGkpID0+IHtcbiAgICAgIHRoaXMuYW5zd2Vycy5wdXNoKHsgdGltZTogMCwgYW5zd2VyOiBudWxsLCBjb3JyZWN0OiBmYWxzZSB9KTtcbiAgICAgIHRoaXMucGVuYWx0aWVzLnB1c2goMCk7XG5cbiAgICAgIGNvbnN0IGFjY2VzcyA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJzcGFuXCIpO1xuICAgICAgYWNjZXNzLnRleHRDb250ZW50ID0gKGkgKyAxKS50b1N0cmluZygpO1xuICAgICAgYWNjZXNzLmNsYXNzTmFtZSA9IFwicXVpY2stYWNjZXNzLW5vZGVcIjtcbiAgICAgIGFjY2Vzcy5kYXRhc2V0LnF1ZXN0aW9uID0gaS50b1N0cmluZygpO1xuXG4gICAgICBhY2Nlc3Mub25jbGljayA9ICgpID0+IHtcbiAgICAgICAgdGhpcy5zYXZlQW5zd2VyVGltZSgpO1xuICAgICAgICB0aGlzLnF1ZXN0aW9uTnVtYmVyID0gaTtcblxuICAgICAgICB0aGlzLmNoYW5nZVF1ZXN0aW9uKCk7XG4gICAgICB9O1xuXG4gICAgICBxdWlja0FjY2Vzc0Rpdi5hcHBlbmRDaGlsZChhY2Nlc3MpO1xuICAgIH0pO1xuXG4gICAgdGhpcy5iaW5kRXZlbnRIYW5kbGVycygpO1xuXG4gICAgZXhjZXJjaXNlU2VjdGlvbi5zdHlsZS5kaXNwbGF5ID0gXCJibG9ja1wiO1xuICAgIHRoaXMuYW5zd2VyU3RhcnRUaW1lID0gcGVyZm9ybWFuY2Uubm93KCk7XG5cbiAgICB0aGlzLmNoYW5nZVF1ZXN0aW9uKCk7XG4gIH1cblxuICAvKipcbiAgICogQmluZHMgbmVjY2Vzc2FyeSBldmVudCBoYW5kbGVycy5cbiAgICovXG4gIHByaXZhdGUgYmluZEV2ZW50SGFuZGxlcnMoKTogdm9pZCB7XG4gICAgYW5zd2VySW5wdXQub25pbnB1dCA9ICgpID0+IHtcbiAgICAgIHRoaXMuc2F2ZUFuc3dlcigpO1xuXG4gICAgICBzdG9wQnV0dG9uLmRpc2FibGVkID0gdGhpcy5hbnN3ZXJzLnNvbWUoKGEpID0+IGEuYW5zd2VyID09PSBudWxsKTtcbiAgICAgIGNvbnN0IHF1aWNrQWNjZXNzU3BhbiA9IGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3IoXG4gICAgICAgIGBbZGF0YS1xdWVzdGlvbj1cIiR7dGhpcy5xdWVzdGlvbk51bWJlcn1cIl1gXG4gICAgICApIGFzIEhUTUxTcGFuRWxlbWVudDtcbiAgICAgIGlmICh0aGlzLmFuc3dlcnNbdGhpcy5xdWVzdGlvbk51bWJlcl0uYW5zd2VyID09PSBudWxsKSB7XG4gICAgICAgIHF1aWNrQWNjZXNzU3Bhbi5jbGFzc05hbWUgPSBxdWlja0FjY2Vzc1NwYW4uY2xhc3NOYW1lLnJlcGxhY2UoXG4gICAgICAgICAgL2Fuc3dlcmVkLyxcbiAgICAgICAgICBcIlwiXG4gICAgICAgICk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBxdWlja0FjY2Vzc1NwYW4uY2xhc3NOYW1lICs9IFwiIGFuc3dlcmVkXCI7XG4gICAgICB9XG4gICAgfTtcblxuICAgIG5leHRCdXR0b24ub25jbGljayA9ICgpID0+IHtcbiAgICAgIHRoaXMuc2F2ZUFuc3dlclRpbWUoKTtcbiAgICAgIHRoaXMucXVlc3Rpb25OdW1iZXIrKztcblxuICAgICAgdGhpcy5jaGFuZ2VRdWVzdGlvbigpO1xuICAgIH07XG5cbiAgICBwcmV2aW91c0J1dHRvbi5vbmNsaWNrID0gKCkgPT4ge1xuICAgICAgdGhpcy5zYXZlQW5zd2VyVGltZSgpO1xuICAgICAgdGhpcy5xdWVzdGlvbk51bWJlci0tO1xuXG4gICAgICB0aGlzLmNoYW5nZVF1ZXN0aW9uKCk7XG4gICAgfTtcblxuICAgIHN0b3BCdXR0b24ub25jbGljayA9ICgpID0+IHtcbiAgICAgIHRoaXMuZW5kUXVpeigpO1xuICAgIH07XG5cbiAgICBhbnN3ZXJGb3JtLm9uc3VibWl0ID0gKCkgPT4ge1xuICAgICAgaWYgKHN0b3BCdXR0b24uZGlzYWJsZWQpIHtcbiAgICAgICAgbmV4dEJ1dHRvbi5jbGljaygpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgc3RvcEJ1dHRvbi5jbGljaygpO1xuICAgICAgfVxuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH07XG4gIH1cblxuICAvKipcbiAgICogQ2hhbmdlcyBjdXJyZW50IHF1ZXN0aW9uIHRvIGdpdmVuLlxuICAgKiBTZXRzIHByZXZBbnN3ZXJTdGFydFRpbWUgdG8gIGN1cnJlbnQgdGltZXN0YW1wLlxuICAgKi9cbiAgcHJpdmF0ZSBjaGFuZ2VRdWVzdGlvbigpOiB2b2lkIHtcbiAgICBpZiAodGhpcy5xdWl6ID09PSBudWxsKSB7XG4gICAgICByZXR1cm47XG4gICAgfVxuICAgIGNvbnN0IHF1ZXN0aW9uID0gdGhpcy5xdWl6LnF1ZXN0aW9uc1t0aGlzLnF1ZXN0aW9uTnVtYmVyXTtcbiAgICBleGNlcmNpc2VOdW1iZXJTcGFuLnRleHRDb250ZW50ID0gKHRoaXMucXVlc3Rpb25OdW1iZXIgKyAxKS50b1N0cmluZygpO1xuICAgIHByb21wdFNwYW4udGV4dENvbnRlbnQgPSBxdWVzdGlvbi5wcm9tcHQ7XG4gICAgcGVuYWx0eVRpbWVTcGFuLnRleHRDb250ZW50ID0gcXVlc3Rpb24ucGVuYWx0eS50b1N0cmluZygpO1xuXG4gICAgY29uc3QgYW5zd2VyID0gdGhpcy5hbnN3ZXJzW3RoaXMucXVlc3Rpb25OdW1iZXJdLmFuc3dlcjtcbiAgICBhbnN3ZXJJbnB1dC52YWx1ZSA9IGFuc3dlciA9PT0gbnVsbCA/IFwiXCIgOiBhbnN3ZXIudG9TdHJpbmcoKTtcblxuICAgIHByZXZpb3VzQnV0dG9uLmRpc2FibGVkID0gdGhpcy5xdWVzdGlvbk51bWJlciA9PT0gMDtcbiAgICBuZXh0QnV0dG9uLmRpc2FibGVkID1cbiAgICAgIHRoaXMucXVlc3Rpb25OdW1iZXIgPT09IHRoaXMucXVpei5xdWVzdGlvbnMubGVuZ3RoIC0gMTtcblxuICAgIHRoaXMuYW5zd2VyU3RhcnRUaW1lID0gcGVyZm9ybWFuY2Uubm93KCk7XG4gIH1cblxuICAvKipcbiAgICogU2F2ZXMgY29udGVzdGFudHMgYW5zd2VyIHRvIGN1cnJlbnQgcXVlc3Rpb24uXG4gICAqL1xuICBwcml2YXRlIHNhdmVBbnN3ZXIgPSAoKTogdm9pZCA9PiB7XG4gICAgY29uc3QgcHJldkFuc3dlciA9IHRoaXMuYW5zd2Vyc1t0aGlzLnF1ZXN0aW9uTnVtYmVyXTtcblxuICAgIGlmIChwcmV2QW5zd2VyID09IG51bGwpIHtcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICBjb25zdCBhbnN3ZXIgPSBwYXJzZUludChhbnN3ZXJJbnB1dC52YWx1ZSk7XG4gICAgaWYgKGlzTmFOKGFuc3dlcikpIHtcbiAgICAgIHByZXZBbnN3ZXIuYW5zd2VyID0gbnVsbDtcbiAgICB9IGVsc2Uge1xuICAgICAgcHJldkFuc3dlci5hbnN3ZXIgPSBhbnN3ZXI7XG4gICAgfVxuICB9O1xuXG4gIC8qKlxuICAgKiBTYXZlcyBjb250ZXN0YW50cyBhbnN3ZXIgdGltZSB0byBjdXJyZW50IHF1ZXN0aW9uLlxuICAgKi9cbiAgcHJpdmF0ZSBzYXZlQW5zd2VyVGltZSA9ICgpOiB2b2lkID0+IHtcbiAgICBjb25zdCBlbmRUaW1lID0gcGVyZm9ybWFuY2Uubm93KCk7XG5cbiAgICBjb25zdCBwcmV2QW5zd2VyID0gdGhpcy5hbnN3ZXJzW3RoaXMucXVlc3Rpb25OdW1iZXJdO1xuXG4gICAgaWYgKHByZXZBbnN3ZXIgPT0gbnVsbCkge1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIHByZXZBbnN3ZXIudGltZSArPSBlbmRUaW1lIC0gdGhpcy5hbnN3ZXJTdGFydFRpbWU7XG4gIH07XG5cbiAgLyoqXG4gICAqIEVuZCBxdWl6LiBEaXNwbGF5cyByZXN1bHRzLlxuICAgKi9cbiAgcHJpdmF0ZSBhc3luYyBlbmRRdWl6KCkge1xuICAgIHRoaXMuc2F2ZUFuc3dlclRpbWUoKTtcbiAgICBleGNlcmNpc2VTZWN0aW9uLnN0eWxlLmRpc3BsYXkgPSBcIm5vbmVcIjtcbiAgICBhd2FpdCBzYXZlUmVzdWx0cyh0aGlzLmFuc3dlcnMsIHRoaXMucXVpej8uaWQgfHwgXCItMVwiKTtcbiAgICB3aW5kb3cubG9jYXRpb24uaHJlZiA9IFwiL3Jlc3VsdHMvXCIgKyB0aGlzLnF1aXo/LmlkO1xuICB9XG59XG4iLCJpbXBvcnQgeyBRdWl6Tm9BbnN3ZXJzLCBnZXRRdWl6ZXMsIGdldFF1aXogfSBmcm9tIFwiLi9xdWl6XCI7XG5pbXBvcnQgeyBRdWl6UnVuIH0gZnJvbSBcIi4vUXVpelJ1blwiO1xuaW1wb3J0IHsgUXVpelJlc3VsdHMgfSBmcm9tIFwiLi9RdWl6UmVzdWx0c1wiO1xuaW1wb3J0IHsgcXVpemVzRGl2IH0gZnJvbSBcIi4vSFRNTEVsZW1lbnRzXCI7XG5cbi8qKlxuICogQWRkcyBxdWl6IHRvIHRhYmxlIG9mIGF2YWlsYWJsZSBxdWl6ZXMuXG4gKlxuICogQHBhcmFtIGlkIHF1aXogaWRcbiAqIEBwYXJhbSBxdWl6IHF1aXogb2JqZWN0XG4gKi9cbmNvbnN0IGFkZFF1aXpUb1RhYmxlID0gKHF1aXo6IFF1aXpOb0Fuc3dlcnMpID0+IHtcbiAgY29uc3Qgcm93ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImRpdlwiKTtcbiAgY29uc3QgbmFtZUNvbCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJkaXZcIik7XG4gIGNvbnN0IGFjdGlvbkNvbCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJkaXZcIik7XG4gIGNvbnN0IGJ1dHRvbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJidXR0b25cIik7XG5cbiAgYnV0dG9uLmNsYXNzTmFtZSA9IFwiYnRuIGJ0bi1wcmltYXJ5IHF1aXotc3RhcnQtYnV0dG9uXCI7XG4gIGJ1dHRvbi5kYXRhc2V0LnF1aXpJZCA9IHF1aXouaWQ7XG4gIGJ1dHRvbi50ZXh0Q29udGVudCA9IFwiU3RhcnRcIjtcblxuICBhY3Rpb25Db2wuY2xhc3NOYW1lID1cbiAgICBcImNvbC00IGQtZmxleCBqdXN0aWZ5LWNvbnRlbnQtY2VudGVyIGFsaWduLWl0ZW1zLWNlbnRlclwiO1xuICBhY3Rpb25Db2wuYXBwZW5kQ2hpbGQoYnV0dG9uKTtcblxuICBuYW1lQ29sLmNsYXNzTmFtZSA9IFwiY29sLTggZC1mbGV4IGp1c3RpZnktY29udGVudC1jZW50ZXIgYWxpZ24taXRlbXMtY2VudGVyXCI7XG4gIG5hbWVDb2wudGV4dENvbnRlbnQgPSBxdWl6LmRlc2M7XG5cbiAgcm93LmNsYXNzTmFtZSA9IFwicm93XCI7XG5cbiAgcm93LmFwcGVuZENoaWxkKG5hbWVDb2wpO1xuICByb3cuYXBwZW5kQ2hpbGQoYWN0aW9uQ29sKTtcblxuICBxdWl6ZXNEaXYuYXBwZW5kQ2hpbGQocm93KTtcbn07XG5cbi8qKlxuICogRGlzcGxheXMgYXZhaWxhYmxlIHF1aXplcyBpbiB0aGUgdGFibGUuXG4gKi9cbmNvbnN0IGRpc3BsYXlRdWl6ZXMgPSBhc3luYyAoKSA9PiB7XG4gIGZvciBhd2FpdCAoY29uc3QgcXVpeiBvZiBnZXRRdWl6ZXMoKSkge1xuICAgIGFkZFF1aXpUb1RhYmxlKHF1aXopO1xuICB9XG5cbiAgY29uc3QgYnV0dG9ucyA9IGRvY3VtZW50LmdldEVsZW1lbnRzQnlDbGFzc05hbWUoXG4gICAgXCJxdWl6LXN0YXJ0LWJ1dHRvblwiXG4gICkgYXMgSFRNTENvbGxlY3Rpb25PZjxIVE1MQnV0dG9uRWxlbWVudD47XG4gIEFycmF5LmZyb20oYnV0dG9ucykuZm9yRWFjaCgoYnV0dG9uKSA9PiB7XG4gICAgYnV0dG9uLm9uY2xpY2sgPSBhc3luYyAoKSA9PiB7XG4gICAgICBuZXcgUXVpelJ1bihhd2FpdCBnZXRRdWl6KGJ1dHRvbi5kYXRhc2V0LnF1aXpJZCB8fCBcIlwiKSk7XG4gICAgfTtcbiAgfSk7XG59O1xuXG5kaXNwbGF5UXVpemVzKCk7XG5cblF1aXpSZXN1bHRzLmRpc3BsYXlQcmV2aW91c1Jlc3VsdHMoKTtcbiIsImltcG9ydCB7IFJlc3VsdCB9IGZyb20gXCIuL1F1aXpSZXN1bHRzXCI7XG5pbXBvcnQgeyBSYXdRdWl6UmVzdWx0IH0gZnJvbSBcIi4uLy4uL3Jlc3VsdHNcIjtcbmltcG9ydCB7IGNzcmZUb2tlbiB9IGZyb20gXCIuL3V0aWxcIjtcbmltcG9ydCB7IFVzZXIgfSBmcm9tIFwiLi4vLi4vbG9naW5cIjtcblxuZXhwb3J0IGludGVyZmFjZSBRdWVzdGlvbk5vQW5zd2VyIHtcbiAgcHJvbXB0OiBzdHJpbmc7XG4gIHBlbmFsdHk6IG51bWJlcjtcbn1cblxuZXhwb3J0IGludGVyZmFjZSBRdWVzdGlvbldpdGhBbnN3ZXJzIHtcbiAgcHJvbXB0OiBzdHJpbmc7XG4gIHBlbmFsdHk6IG51bWJlcjtcbiAgYW5zd2VyOiBudW1iZXI7XG59XG5cbmV4cG9ydCBpbnRlcmZhY2UgUXVpek5vQW5zd2VycyB7XG4gIGRlc2M6IHN0cmluZztcbiAgcXVlc3Rpb25zOiBRdWVzdGlvbk5vQW5zd2VyW107XG4gIGlkOiBzdHJpbmc7XG59XG5cbmV4cG9ydCBpbnRlcmZhY2UgUXVpeldpdGhBbnN3ZXJzIHtcbiAgZGVzYzogc3RyaW5nO1xuICBxdWVzdGlvbnM6IFF1ZXN0aW9uV2l0aEFuc3dlcnNbXTtcbiAgaWQ6IHN0cmluZztcbn1cblxuZXhwb3J0IGludGVyZmFjZSBBbnN3ZXIge1xuICBhbnN3ZXI6IG51bWJlciB8IG51bGw7XG4gIHRpbWU6IG51bWJlcjtcbiAgY29ycmVjdDogYm9vbGVhbjtcbn1cblxuY29uc3QgaXNRdWVzdGlvbiA9IChxdWVzdGlvbjogYW55KTogcXVlc3Rpb24gaXMgUXVlc3Rpb25Ob0Fuc3dlciA9PiB7XG4gIHJldHVybiBxdWVzdGlvbiAmJiBxdWVzdGlvbi5wcm9tcHQgJiYgcXVlc3Rpb24uYW5zd2VyICYmIHF1ZXN0aW9uLnBlbmFsdHk7XG59O1xuXG5leHBvcnQgY29uc3QgaXNRdWl6ID0gKHF1aXo6IGFueSk6IHF1aXogaXMgUXVpek5vQW5zd2VycyA9PiB7XG4gIHJldHVybiBxdWl6LmRlc2MgJiYgcXVpei5xdWVzdGlvbnMgJiYgcXVpei5pZCAhPT0gdW5kZWZpbmVkO1xufTtcbmNvbnN0IGlzUXVpeldpdGhBbnN3ZXJzID0gKHF1aXo6IGFueSk6IHF1aXogaXMgUXVpeldpdGhBbnN3ZXJzID0+IHtcbiAgcmV0dXJuIHF1aXouZGVzYyAmJiBxdWl6LnF1ZXN0aW9ucyAmJiBxdWl6LmlkICE9PSB1bmRlZmluZWQ7XG59O1xuXG4vKipcbiAqIFJldHVybnMgcXVpeiB3aXRoIGdpdmVuIGlkIG9yIG51bGwgb24gZXJyb3IuXG4gKiBAcGFyYW0gaWQgcXVpeiBpZFxuICovXG5leHBvcnQgY29uc3QgZ2V0UXVpeiA9IGFzeW5jIChpZDogc3RyaW5nKTogUHJvbWlzZTxRdWl6Tm9BbnN3ZXJzIHwgbnVsbD4gPT4ge1xuICBjb25zb2xlLmxvZyhpZCk7XG4gIHRyeSB7XG4gICAgY29uc3QgcXVpelJhdyA9IGF3YWl0IGZldGNoKFwiL2dldF9xdWl6L1wiICsgaWQpO1xuICAgIGNvbnN0IHF1aXogPSBhd2FpdCBxdWl6UmF3Lmpzb24oKTtcblxuICAgIGlmIChpc1F1aXoocXVpeikpIHtcbiAgICAgIHJldHVybiBxdWl6O1xuICAgIH1cbiAgfSBjYXRjaCAoXykge31cblxuICByZXR1cm4gbnVsbDtcbn07XG5cbi8qKlxuICogUmV0dXJucyBxdWl6IHdpdGggZ2l2ZW4gaWQgb3IgbnVsbCBvbiBlcnJvci5cbiAqIEBwYXJhbSBpZCBxdWl6IGlkXG4gKi9cbmV4cG9ydCBjb25zdCBnZXRRdWl6V2l0aEFuc3dlcnMgPSBhc3luYyAoXG4gIGlkOiBzdHJpbmdcbik6IFByb21pc2U8UXVpeldpdGhBbnN3ZXJzIHwgbnVsbD4gPT4ge1xuICB0cnkge1xuICAgIGNvbnN0IHF1aXpSYXcgPSBhd2FpdCBmZXRjaChcIi9nZXRfcXVpel93aXRoX2Fuc3dlcnMvXCIgKyBpZCk7XG4gICAgY29uc3QgcXVpeiA9IGF3YWl0IHF1aXpSYXcuanNvbigpO1xuXG4gICAgaWYgKGlzUXVpeldpdGhBbnN3ZXJzKHF1aXopKSB7XG4gICAgICByZXR1cm4gcXVpejtcbiAgICB9XG4gIH0gY2F0Y2ggKF8pIHt9XG5cbiAgcmV0dXJuIG51bGw7XG59O1xuXG4vKipcbiAqIEdlbmVyYXRvciB0aGF0IHJldHVybnMgYWxsIHF1aXplcy5cbiAqL1xuZXhwb3J0IGFzeW5jIGZ1bmN0aW9uKiBnZXRRdWl6ZXMoKTogQXN5bmNHZW5lcmF0b3I8UXVpek5vQW5zd2Vycz4ge1xuICBjb25zdCBxdWl6UmF3ID0gYXdhaXQgZmV0Y2goXCIvZ2V0X3F1aXplc1wiKTtcbiAgY29uc3Qgb2JqID0gYXdhaXQgcXVpelJhdy5qc29uKCk7XG4gIGlmIChvYmogaW5zdGFuY2VvZiBBcnJheSkge1xuICAgIGZvciAoY29uc3QgcXVpeiBvZiBvYmopIHtcbiAgICAgIGlmIChpc1F1aXoocXVpeikpIHtcbiAgICAgICAgeWllbGQgcXVpejtcbiAgICAgIH1cbiAgICB9XG4gIH1cbn1cblxuZXhwb3J0IGNvbnN0IGdldFByZXZSZXN1bHRzID0gYXN5bmMgKCk6IFByb21pc2U8UmVzdWx0W10+ID0+IHtcbiAgY29uc3QgcmVzdWx0cyA9IGF3YWl0IGZldGNoKFwiL3ByZXZfcmVzdWx0c1wiKTtcbiAgY29uc3Qgb2JqID0gYXdhaXQgcmVzdWx0cy5qc29uKCk7XG4gIGlmIChvYmogaW5zdGFuY2VvZiBBcnJheSkge1xuICAgIHJldHVybiA8UmVzdWx0W10+b2JqO1xuICB9XG4gIHJldHVybiBbXTtcbn07XG5cbmNvbnN0IGdldFJhd1F1aXpSZXN1bHRzID0gKGFuc3dlcnM6IEFuc3dlcltdKTogUmF3UXVpelJlc3VsdCA9PiB7XG4gIGNvbnN0IHRvdGFsVGltZSA9IGFuc3dlcnMucmVkdWNlKChzdW0sIGFucykgPT4gc3VtICsgYW5zLnRpbWUsIDApO1xuICBjb25zdCBxcjogUmF3UXVpelJlc3VsdCA9IHsgYW5zd2VyczogW10sIHRpbWVzOiBbXSB9O1xuICBhbnN3ZXJzLmZvckVhY2goKGFucykgPT4ge1xuICAgIHFyLmFuc3dlcnMucHVzaCg8bnVtYmVyPmFucy5hbnN3ZXIpO1xuICAgIHFyLnRpbWVzLnB1c2goYW5zLnRpbWUgLyB0b3RhbFRpbWUpO1xuICB9KTtcblxuICByZXR1cm4gcXI7XG59O1xuZXhwb3J0IGNvbnN0IHNhdmVSZXN1bHRzID0gYXN5bmMgKGFuc3dlcnM6IEFuc3dlcltdLCBxdWl6SWQ6IHN0cmluZykgPT4ge1xuICBjb25zdCBxciA9IGdldFJhd1F1aXpSZXN1bHRzKGFuc3dlcnMpO1xuICBjb25zb2xlLmxvZyhxciwgYW5zd2VycywgSlNPTi5zdHJpbmdpZnkocXIpKTtcbiAgYXdhaXQgZmV0Y2goXCIvcG9zdF9yZXN1bHRzL1wiICsgcXVpeklkLCB7XG4gICAgbWV0aG9kOiBcIlBPU1RcIixcbiAgICBjcmVkZW50aWFsczogXCJzYW1lLW9yaWdpblwiLFxuICAgIGhlYWRlcnM6IHtcbiAgICAgIFwiQ29udGVudC1UeXBlXCI6IFwiYXBwbGljYXRpb24vanNvblwiLFxuICAgICAgXCJDU1JGLVRva2VuXCI6IGNzcmZUb2tlbixcbiAgICB9LFxuICAgIGJvZHk6IEpTT04uc3RyaW5naWZ5KHFyKSxcbiAgfSk7XG59O1xuXG5leHBvcnQgY29uc3QgZ2V0QW5zd2VycyA9IGFzeW5jIChxdWl6SWQ6IHN0cmluZyk6IFByb21pc2U8QW5zd2VyW10gfCBudWxsPiA9PiB7XG4gIHRyeSB7XG4gICAgY29uc3QgcmF3ID0gYXdhaXQgZmV0Y2goXCIvZ2V0X2Fuc3dlcnMvXCIgKyBxdWl6SWQpO1xuICAgIGNvbnN0IG9iaiA9IGF3YWl0IHJhdy5qc29uKCk7XG4gICAgaWYgKG9iaiBpbnN0YW5jZW9mIEFycmF5KSB7XG4gICAgICByZXR1cm4gPEFuc3dlcltdPm9iajtcbiAgICB9XG4gIH0gY2F0Y2ggKF8pIHt9XG4gIHJldHVybiBudWxsO1xufTtcblxuZXhwb3J0IGNvbnN0IGF2ZXJhZ2VUaW1lcyA9IGFzeW5jIChcbiAgcXVpeklkOiBzdHJpbmdcbik6IFByb21pc2U8bnVtYmVyW10gfCBudWxsPiA9PiB7XG4gIHRyeSB7XG4gICAgY29uc3QgcmF3ID0gYXdhaXQgZmV0Y2goXCIvZ2V0X2Fuc3dlcnNfbWVhbl90aW1lL1wiICsgcXVpeklkKTtcbiAgICBjb25zdCBvYmogPSBhd2FpdCByYXcuanNvbigpO1xuICAgIGlmIChvYmogaW5zdGFuY2VvZiBBcnJheSkge1xuICAgICAgcmV0dXJuIDxudW1iZXJbXT5vYmo7XG4gICAgfVxuICB9IGNhdGNoIChfKSB7fVxuICByZXR1cm4gbnVsbDtcbn07XG5leHBvcnQgY29uc3QgdG9wU2NvcmVzID0gYXN5bmMgKFxuICBxdWl6SWQ6IHN0cmluZ1xuKTogUHJvbWlzZTxbVXNlciwgUmVzdWx0XVtdIHwgbnVsbD4gPT4ge1xuICB0cnkge1xuICAgIGNvbnN0IHJhdyA9IGF3YWl0IGZldGNoKFwiL3RvcF9zY29yZXMvXCIgKyBxdWl6SWQpO1xuICAgIGNvbnN0IG9iaiA9IGF3YWl0IHJhdy5qc29uKCk7XG4gICAgaWYgKG9iaiBpbnN0YW5jZW9mIEFycmF5KSB7XG4gICAgICByZXR1cm4gPFtVc2VyLCBSZXN1bHRdW10+b2JqO1xuICAgIH1cbiAgfSBjYXRjaCAoXykge31cbiAgcmV0dXJuIG51bGw7XG59O1xuIiwiLyoqXG4gKiBGdW5jdGlvbiB0aGF0IGxvZ3MgZXJyb3JzLlxuICogTWF5IGluIHRoZSBmdXR1cmUgYmUgb3ZlcndyaXRlbiB0byBmdW5jdGlvbiB0aGF0IGRpc3BsYXkgZXJyb3Igb24gc2NyZWVuLlxuICovXG5leHBvcnQgY29uc3QgZXJyb3IgPSAoczogc3RyaW5nKSA9PiBhbGVydChzKTtcblxuZXhwb3J0IGNvbnN0IGNzcmZUb2tlbiA9XG4gIGRvY3VtZW50Py5xdWVyeVNlbGVjdG9yKCdtZXRhW25hbWU9XCJjc3JmLXRva2VuXCJdJyk/LmdldEF0dHJpYnV0ZShcImNvbnRlbnRcIikgfHxcbiAgXCJcIjtcbiJdLCJzb3VyY2VSb290IjoiIn0=