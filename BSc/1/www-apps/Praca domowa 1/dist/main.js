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
exports.answerInput = document.getElementById('answer');
exports.previousButton = document.getElementById('previous');
exports.nextButton = document.getElementById('next');
exports.stopButton = document.getElementById('stop');
exports.promptSpan = document.getElementById('prompt');
exports.excerciseNumberSpan = document.getElementById('excercise-number');
exports.excerciseSection = document.getElementById('excercise');
exports.quizDescSection = document.getElementById('quiz-desc');
exports.quizChoiseSection = document.getElementById('quiz-choise');
exports.penaltyTimeSpan = document.getElementById('penalty-time');
exports.descParagraph = document.getElementById('desc');
exports.resultsDiv = document.getElementById('results');
exports.quizesDiv = document.getElementById('quizes');
exports.correctnessDiv = document.getElementById('correctness');
exports.resultSection = document.getElementById('result');
exports.saveButton = document.getElementById('save');
exports.saveWithStatsButton = document.getElementById('save-with-stats');
exports.answerForm = document.getElementById('answer-form');
exports.initResultSpan = document.getElementById('init-result');
exports.finalResultSpan = document.getElementById('final-result');
exports.quizPrevResultsSection = document.getElementById('quiz-prev-results');
exports.quickAccessDiv = document.getElementById('quick-access');


/***/ }),

/***/ "./src/QuizResults.ts":
/*!****************************!*\
  !*** ./src/QuizResults.ts ***!
  \****************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

Object.defineProperty(exports, "__esModule", { value: true });
const quiz_1 = __webpack_require__(/*! ./quiz */ "./src/quiz.ts");
const HTMLElements_1 = __webpack_require__(/*! ./HTMLElements */ "./src/HTMLElements.ts");
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
        this.checkAnswers();
        this.givePenalties();
        this.display();
    }
    /**
     * Checks answers, sets carrect argument in each answer to either true or false.
     */
    checkAnswers() {
        if (this.quiz !== null) {
            for (const i in this.answers) {
                this.answers[i].correct = this.answers[i].answer === this.quiz.questions[i].answer;
            }
        }
    }
    /**
     * Gives penalties for each wrong answer.
     */
    givePenalties() {
        if (this.quiz !== null) {
            for (const i in this.answers) {
                this.penalties.push(this.answers[i].correct ? 0 : this.quiz.questions[i].penalty);
            }
        }
    }
    /**
     * Bind neccessary event handlers.
     */
    bindEventHandlers() {
        HTMLElements_1.saveWithStatsButton.onclick = () => {
            this.saveWithStats();
            this.goToMainScreen();
        };
        HTMLElements_1.saveButton.onclick = () => {
            this.save();
            this.goToMainScreen();
        };
    }
    /**
     * Goes back to main screen.
     */
    goToMainScreen() {
        window.location.reload();
    }
    /**
     * Saves run results with stats.
     */
    saveWithStats() {
        var _a, _b;
        this.saveResult({ finalTime: this.finalTime, answers: this.answers, quizId: (_b = (_a = this.quiz) === null || _a === void 0 ? void 0 : _a.id) !== null && _b !== void 0 ? _b : "" });
    }
    /**
     * Saves run results without stats.
     */
    save() {
        var _a, _b;
        this.saveResult({ finalTime: this.finalTime, quizId: (_b = (_a = this.quiz) === null || _a === void 0 ? void 0 : _a.id) !== null && _b !== void 0 ? _b : "" });
    }
    /**
     * Saves given result.
     * @param result result
     */
    saveResult(result) {
        const prevJSON = localStorage.getItem('results') || '[]';
        const prev = JSON.parse(prevJSON);
        prev.push(result);
        localStorage.setItem('results', JSON.stringify(prev));
    }
    /**
     * Calculates time spent on quiz.
     */
    get initTime() {
        let initTime = 0;
        for (const ans of this.answers) {
            initTime += ans.time / 1000;
        }
        return initTime;
    }
    /**
     * Calculates final time, base + penalties.
     */
    get finalTime() {
        let finalTime = this.initTime;
        for (const p of this.penalties) {
            finalTime += p;
        }
        return finalTime;
    }
    /**
     * Displays results.
     */
    display() {
        if (this.quiz) {
            HTMLElements_1.resultSection.style.display = 'block';
            HTMLElements_1.initResultSpan.textContent = QuizResults.formatTime(this.initTime);
            HTMLElements_1.finalResultSpan.textContent = QuizResults.formatTime(this.finalTime);
            // console.log(this.answers.entries());
            for (const [i, answer] of this.answers.entries()) {
                console.log(i, answer);
                const div = document.createElement('div');
                div.textContent = `${i + 1}. `;
                const span = document.createElement('span');
                span.className = answer.correct ? 'correct' : 'incorrect';
                span.textContent = answer.correct ? 'Correct :)' : `Incorect: + ${this.quiz.questions[i].penalty}s`;
                div.appendChild(span);
                HTMLElements_1.correctnessDiv.appendChild(div);
            }
        }
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
        const results = JSON.parse(localStorage.getItem('results') || '');
        let i = 1;
        console.log(results);
        if (results instanceof Array) {
            for (const res of results) {
                console.log(res, isResult(res));
                if (isResult(res)) {
                    const quiz = quiz_1.getQuiz(res.quizId);
                    if (quiz) {
                        const row = document.createElement('div');
                        const nameCol = document.createElement('div');
                        const resultCol = document.createElement('div');
                        resultCol.className = 'col-4 d-flex justify-content-center';
                        resultCol.textContent = QuizResults.formatTime(res.finalTime) + 's';
                        nameCol.className = 'col-8 d-flex justify-content-center';
                        nameCol.textContent = quiz.desc;
                        row.className = `row prev-result ${i % 2 == 0 ? 'even' : ''}`;
                        row.appendChild(nameCol);
                        row.appendChild(resultCol);
                        HTMLElements_1.quizPrevResultsSection.appendChild(row);
                        i++;
                    }
                }
            }
        }
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

Object.defineProperty(exports, "__esModule", { value: true });
const quiz_1 = __webpack_require__(/*! ./quiz */ "./src/quiz.ts");
const QuizResults_1 = __webpack_require__(/*! ./QuizResults */ "./src/QuizResults.ts");
const HTMLElements_1 = __webpack_require__(/*! ./HTMLElements */ "./src/HTMLElements.ts");
const main_1 = __webpack_require__(/*! ./main */ "./src/main.ts");
class QuizRun {
    /**
     * @param quizId chosen quiz id
     */
    constructor(quizId = '0') {
        this.questionNumber = 0;
        this.answerStartTime = 0;
        this.quiz = null;
        this.answers = [];
        this.penalties = [];
        this.quizResults = null;
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
        this.quiz = quiz_1.getQuiz(quizId);
        if (this.quiz == null) {
            main_1.error('Podany quiz nie istnieje');
            return;
        }
        else if (this.quiz.questions.length == 0) {
            main_1.error('Podany quiz ma za mało pytań');
            return;
        }
        HTMLElements_1.quizChoiseSection.style.display = 'none';
        HTMLElements_1.quizPrevResultsSection.style.display = 'none';
        HTMLElements_1.quizDescSection.style.display = 'block';
        HTMLElements_1.descParagraph.textContent = this.quiz.desc;
        this.quiz.questions.forEach((_, i) => {
            this.answers.push({ time: 0, answer: null, correct: false });
            this.penalties.push(0);
            const access = document.createElement('span');
            access.textContent = (i + 1).toString();
            access.className = 'quick-access-node';
            access.dataset.question = i.toString();
            access.onclick = () => {
                this.saveAnswerTime();
                this.questionNumber = i;
                this.changeQuestion();
            };
            HTMLElements_1.quickAccessDiv.appendChild(access);
        });
        this.bindEventHandlers();
        HTMLElements_1.excerciseSection.style.display = 'block';
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
                quickAccessSpan.className = quickAccessSpan.className.replace(/answered/, '');
            }
            else {
                quickAccessSpan.className += ' answered';
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
        HTMLElements_1.answerInput.value = answer === null ? '' : answer.toString();
        HTMLElements_1.previousButton.disabled = this.questionNumber === 0;
        HTMLElements_1.nextButton.disabled = this.questionNumber === this.quiz.questions.length - 1;
        this.answerStartTime = performance.now();
    }
    /**
     * End quiz. Displays results.
     */
    endQuiz() {
        this.saveAnswerTime();
        HTMLElements_1.excerciseSection.style.display = 'none';
        this.quizResults = new QuizResults_1.QuizResults(this.answers, this.quiz);
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

Object.defineProperty(exports, "__esModule", { value: true });
const quiz_1 = __webpack_require__(/*! ./quiz */ "./src/quiz.ts");
const QuizRun_1 = __webpack_require__(/*! ./QuizRun */ "./src/QuizRun.ts");
const QuizResults_1 = __webpack_require__(/*! ./QuizResults */ "./src/QuizResults.ts");
const HTMLElements_1 = __webpack_require__(/*! ./HTMLElements */ "./src/HTMLElements.ts");
/**
 * Function that logs errors.
 * May in the future be overwriten to function that display error on screen.
 */
exports.error = console.error;
/**
 * Adds quiz to table of available quizes.
 *
 * @param id quiz id
 * @param quiz quiz object
 */
const addQuizToTable = (quiz) => {
    const row = document.createElement('div');
    const nameCol = document.createElement('div');
    const actionCol = document.createElement('div');
    const button = document.createElement('button');
    button.className = 'btn btn-primary quiz-start-button';
    button.dataset.quizId = quiz.id;
    button.textContent = 'Start';
    actionCol.className = 'col-4 d-flex justify-content-center align-items-center';
    actionCol.appendChild(button);
    nameCol.className = 'col-8 d-flex justify-content-center align-items-center';
    nameCol.textContent = quiz.desc;
    row.className = 'row';
    row.appendChild(nameCol);
    row.appendChild(actionCol);
    HTMLElements_1.quizesDiv.appendChild(row);
};
/**
 * Displays available quizes in the table.
 */
const displayQuizes = () => {
    for (const quiz of quiz_1.getQuizes()) {
        console.log(quiz);
        addQuizToTable(quiz);
    }
    const buttons = document.getElementsByClassName('quiz-start-button');
    Array.from(buttons).forEach((button) => {
        button.onclick = () => {
            new QuizRun_1.QuizRun(button.dataset.quizId);
        };
    });
};
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

Object.defineProperty(exports, "__esModule", { value: true });
const quizes_1 = __webpack_require__(/*! ./quizes */ "./src/quizes.ts");
const isQuestion = (question) => {
    return question && question.prompt && question.answer && question.penalty;
};
const isQuiz = (quiz) => {
    return quiz.desc && quiz.questions && quiz.id !== undefined;
};
/**
 * Returns quiz with given id or null on error.
 * @param id quiz id
 */
exports.getQuiz = (id) => {
    var _a;
    const arr = (_a = JSON.parse(quizes_1.quiz)) === null || _a === void 0 ? void 0 : _a.quizes;
    if (arr instanceof Array) {
        const quiz = arr.find((el) => el.id === id);
        if (isQuiz(quiz)) {
            return quiz;
        }
    }
    return null;
};
/**
 * Generator that returns all quizes.
 */
function* getQuizes() {
    const obj = JSON.parse(quizes_1.quiz);
    if (obj.quizes instanceof Array) {
        for (const quiz of obj.quizes) {
            if (isQuiz(quiz)) {
                yield quiz;
            }
        }
    }
}
exports.getQuizes = getQuizes;


/***/ }),

/***/ "./src/quizes.ts":
/*!***********************!*\
  !*** ./src/quizes.ts ***!
  \***********************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

Object.defineProperty(exports, "__esModule", { value: true });
/**
 * Example quizes.
 */
exports.quiz = `{
    "quizes":[{
       "id": "0",
       "desc":"Liczyć każdy może",
       "questions":[
          {
             "prompt":"2 + 3",
             "answer":5,
             "penalty":4
          },
          {
             "prompt":"2 - (-24 : 4)",
             "answer":8,
             "penalty":10
          },
          {
             "prompt":"2 - 3",
             "answer":-1,
             "penalty":4
          },
          {
             "prompt":"2 + (-24 : 4)",
             "answer":-4,
             "penalty":10
          }
       ]
    },{
       "id": "1",
       "desc":"Quick Math",
       "questions":[
          {
             "prompt":"2 + 2",
             "answer":4,
             "penalty":2
          },
          {
             "prompt":"2 * 3",
             "answer":6,
             "penalty":1
          },
          {
             "prompt":"2 ^ 3",
             "answer":8,
             "penalty":3
          },
          {
             "prompt":"(2 ^ (1/2)) ^ 2",
             "answer":2,
             "penalty":7
          }
       ]
    }]
 }`;


/***/ })

/******/ });
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vd2VicGFjay9ib290c3RyYXAiLCJ3ZWJwYWNrOi8vLy4vc3JjL0hUTUxFbGVtZW50cy50cyIsIndlYnBhY2s6Ly8vLi9zcmMvUXVpelJlc3VsdHMudHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL1F1aXpSdW4udHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL21haW4udHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL3F1aXoudHMiLCJ3ZWJwYWNrOi8vLy4vc3JjL3F1aXplcy50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiO1FBQUE7UUFDQTs7UUFFQTtRQUNBOztRQUVBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBOztRQUVBO1FBQ0E7O1FBRUE7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7OztRQUdBO1FBQ0E7O1FBRUE7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7UUFDQSwwQ0FBMEMsZ0NBQWdDO1FBQzFFO1FBQ0E7O1FBRUE7UUFDQTtRQUNBO1FBQ0Esd0RBQXdELGtCQUFrQjtRQUMxRTtRQUNBLGlEQUFpRCxjQUFjO1FBQy9EOztRQUVBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQSx5Q0FBeUMsaUNBQWlDO1FBQzFFLGdIQUFnSCxtQkFBbUIsRUFBRTtRQUNySTtRQUNBOztRQUVBO1FBQ0E7UUFDQTtRQUNBLDJCQUEyQiwwQkFBMEIsRUFBRTtRQUN2RCxpQ0FBaUMsZUFBZTtRQUNoRDtRQUNBO1FBQ0E7O1FBRUE7UUFDQSxzREFBc0QsK0RBQStEOztRQUVySDtRQUNBOzs7UUFHQTtRQUNBOzs7Ozs7Ozs7Ozs7Ozs7QUNsRmEsbUJBQVcsR0FBRyxRQUFRLENBQUMsY0FBYyxDQUFDLFFBQVEsQ0FBcUIsQ0FBQztBQUNwRSxzQkFBYyxHQUFHLFFBQVEsQ0FBQyxjQUFjLENBQUMsVUFBVSxDQUFzQixDQUFDO0FBQzFFLGtCQUFVLEdBQUcsUUFBUSxDQUFDLGNBQWMsQ0FBQyxNQUFN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QU8sSUFBSSxDQUFDO0FBQ2IsQ0FBQyxDQUFDO0FBRUY7O0dBRUc7QUFDSCxRQUFlLENBQUMsQ0FBQyxTQUFTO0lBQ3pCLE1BQU0sR0FBRyxHQUFHLElBQUksQ0FBQyxLQUFLLENBQUMsYUFBSSxDQUFDLENBQUM7SUFDN0IsSUFBSSxHQUFHLENBQUMsTUFBTSxZQUFZLEtBQUssRUFBRTtRQUNoQyxLQUFLLE1BQU0sSUFBSSxJQUFJLEdBQUcsQ0FBQyxNQUFNLEVBQUU7WUFDOUIsSUFBSSxNQUFNLENBQUMsSUFBSSxDQUFDLEVBQUU7Z0JBQ2pCLE1BQU0sSUFBSSxDQUFDO2FBQ1g7U0FDRDtLQUNEO0FBQ0YsQ0FBQztBQVRELDhCQVNDOzs7Ozs7Ozs7Ozs7Ozs7QUN4REQ7O0dBRUc7QUFDVSxZQUFJLEdBQUc7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7R0FvRGpCLENBQUMiLCJmaWxlIjoibWFpbi5qcyIsInNvdXJjZXNDb250ZW50IjpbIiBcdC8vIFRoZSBtb2R1bGUgY2FjaGVcbiBcdHZhciBpbnN0YWxsZWRNb2R1bGVzID0ge307XG5cbiBcdC8vIFRoZSByZXF1aXJlIGZ1bmN0aW9uXG4gXHRmdW5jdGlvbiBfX3dlYnBhY2tfcmVxdWlyZV9fKG1vZHVsZUlkKSB7XG5cbiBcdFx0Ly8gQ2hlY2sgaWYgbW9kdWxlIGlzIGluIGNhY2hlXG4gXHRcdGlmKGluc3RhbGxlZE1vZHVsZXNbbW9kdWxlSWRdKSB7XG4gXHRcdFx0cmV0dXJuIGluc3RhbGxlZE1vZHVsZXNbbW9kdWxlSWRdLmV4cG9ydHM7XG4gXHRcdH1cbiBcdFx0Ly8gQ3JlYXRlIGEgbmV3IG1vZHVsZSAoYW5kIHB1dCBpdCBpbnRvIHRoZSBjYWNoZSlcbiBcdFx0dmFyIG1vZHVsZSA9IGluc3RhbGxlZE1vZHVsZXNbbW9kdWxlSWRdID0ge1xuIFx0XHRcdGk6IG1vZHVsZUlkLFxuIFx0XHRcdGw6IGZhbHNlLFxuIFx0XHRcdGV4cG9ydHM6IHt9XG4gXHRcdH07XG5cbiBcdFx0Ly8gRXhlY3V0ZSB0aGUgbW9kdWxlIGZ1bmN0aW9uXG4gXHRcdG1vZHVsZXNbbW9kdWxlSWRdLmNhbGwobW9kdWxlLmV4cG9ydHMsIG1vZHVsZSwgbW9kdWxlLmV4cG9ydHMsIF9fd2VicGFja19yZXF1aXJlX18pO1xuXG4gXHRcdC8vIEZsYWcgdGhlIG1vZHVsZSBhcyBsb2FkZWRcbiBcdFx0bW9kdWxlLmwgPSB0cnVlO1xuXG4gXHRcdC8vIFJldHVybiB0aGUgZXhwb3J0cyBvZiB0aGUgbW9kdWxlXG4gXHRcdHJldHVybiBtb2R1bGUuZXhwb3J0cztcbiBcdH1cblxuXG4gXHQvLyBleHBvc2UgdGhlIG1vZHVsZXMgb2JqZWN0IChfX3dlYnBhY2tfbW9kdWxlc19fKVxuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5tID0gbW9kdWxlcztcblxuIFx0Ly8gZXhwb3NlIHRoZSBtb2R1bGUgY2FjaGVcbiBcdF9fd2VicGFja19yZXF1aXJlX18uYyA9IGluc3RhbGxlZE1vZHVsZXM7XG5cbiBcdC8vIGRlZmluZSBnZXR0ZXIgZnVuY3Rpb24gZm9yIGhhcm1vbnkgZXhwb3J0c1xuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5kID0gZnVuY3Rpb24oZXhwb3J0cywgbmFtZSwgZ2V0dGVyKSB7XG4gXHRcdGlmKCFfX3dlYnBhY2tfcmVxdWlyZV9fLm8oZXhwb3J0cywgbmFtZSkpIHtcbiBcdFx0XHRPYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgbmFtZSwgeyBlbnVtZXJhYmxlOiB0cnVlLCBnZXQ6IGdldHRlciB9KTtcbiBcdFx0fVxuIFx0fTtcblxuIFx0Ly8gZGVmaW5lIF9fZXNNb2R1bGUgb24gZXhwb3J0c1xuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5yID0gZnVuY3Rpb24oZXhwb3J0cykge1xuIFx0XHRpZih0eXBlb2YgU3ltYm9sICE9PSAndW5kZWZpbmVkJyAmJiBTeW1ib2wudG9TdHJpbmdUYWcpIHtcbiBcdFx0XHRPYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgU3ltYm9sLnRvU3RyaW5nVGFnLCB7IHZhbHVlOiAnTW9kdWxlJyB9KTtcbiBcdFx0fVxuIFx0XHRPYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgJ19fZXNNb2R1bGUnLCB7IHZhbHVlOiB0cnVlIH0pO1xuIFx0fTtcblxuIFx0Ly8gY3JlYXRlIGEgZmFrZSBuYW1lc3BhY2Ugb2JqZWN0XG4gXHQvLyBtb2RlICYgMTogdmFsdWUgaXMgYSBtb2R1bGUgaWQsIHJlcXVpcmUgaXRcbiBcdC8vIG1vZGUgJiAyOiBtZXJnZSBhbGwgcHJvcGVydGllcyBvZiB2YWx1ZSBpbnRvIHRoZSBuc1xuIFx0Ly8gbW9kZSAmIDQ6IHJldHVybiB2YWx1ZSB3aGVuIGFscmVhZHkgbnMgb2JqZWN0XG4gXHQvLyBtb2RlICYgOHwxOiBiZWhhdmUgbGlrZSByZXF1aXJlXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLnQgPSBmdW5jdGlvbih2YWx1ZSwgbW9kZSkge1xuIFx0XHRpZihtb2RlICYgMSkgdmFsdWUgPSBfX3dlYnBhY2tfcmVxdWlyZV9fKHZhbHVlKTtcbiBcdFx0aWYobW9kZSAmIDgpIHJldHVybiB2YWx1ZTtcbiBcdFx0aWYoKG1vZGUgJiA0KSAmJiB0eXBlb2YgdmFsdWUgPT09ICdvYmplY3QnICYmIHZhbHVlICYmIHZhbHVlLl9fZXNNb2R1bGUpIHJldHVybiB2YWx1ZTtcbiBcdFx0dmFyIG5zID0gT2JqZWN0LmNyZWF0ZShudWxsKTtcbiBcdFx0X193ZWJwYWNrX3JlcXVpcmVfXy5yKG5zKTtcbiBcdFx0T2JqZWN0LmRlZmluZVByb3BlcnR5KG5zLCAnZGVmYXVsdCcsIHsgZW51bWVyYWJsZTogdHJ1ZSwgdmFsdWU6IHZhbHVlIH0pO1xuIFx0XHRpZihtb2RlICYgMiAmJiB0eXBlb2YgdmFsdWUgIT0gJ3N0cmluZycpIGZvcih2YXIga2V5IGluIHZhbHVlKSBfX3dlYnBhY2tfcmVxdWlyZV9fLmQobnMsIGtleSwgZnVuY3Rpb24oa2V5KSB7IHJldHVybiB2YWx1ZVtrZXldOyB9LmJpbmQobnVsbCwga2V5KSk7XG4gXHRcdHJldHVybiBucztcbiBcdH07XG5cbiBcdC8vIGdldERlZmF1bHRFeHBvcnQgZnVuY3Rpb24gZm9yIGNvbXBhdGliaWxpdHkgd2l0aCBub24taGFybW9ueSBtb2R1bGVzXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLm4gPSBmdW5jdGlvbihtb2R1bGUpIHtcbiBcdFx0dmFyIGdldHRlciA9IG1vZHVsZSAmJiBtb2R1bGUuX19lc01vZHVsZSA/XG4gXHRcdFx0ZnVuY3Rpb24gZ2V0RGVmYXVsdCgpIHsgcmV0dXJuIG1vZHVsZVsnZGVmYXVsdCddOyB9IDpcbiBcdFx0XHRmdW5jdGlvbiBnZXRNb2R1bGVFeHBvcnRzKCkgeyByZXR1cm4gbW9kdWxlOyB9O1xuIFx0XHRfX3dlYnBhY2tfcmVxdWlyZV9fLmQoZ2V0dGVyLCAnYScsIGdldHRlcik7XG4gXHRcdHJldHVybiBnZXR0ZXI7XG4gXHR9O1xuXG4gXHQvLyBPYmplY3QucHJvdG90eXBlLmhhc093blByb3BlcnR5LmNhbGxcbiBcdF9fd2VicGFja19yZXF1aXJlX18ubyA9IGZ1bmN0aW9uKG9iamVjdCwgcHJvcGVydHkpIHsgcmV0dXJuIE9iamVjdC5wcm90b3R5cGUuaGFzT3duUHJvcGVydHkuY2FsbChvYmplY3QsIHByb3BlcnR5KTsgfTtcblxuIFx0Ly8gX193ZWJwYWNrX3B1YmxpY19wYXRoX19cbiBcdF9fd2VicGFja19yZXF1aXJlX18ucCA9IFwiXCI7XG5cblxuIFx0Ly8gTG9hZCBlbnRyeSBtb2R1bGUgYW5kIHJldHVybiBleHBvcnRzXG4gXHRyZXR1cm4gX193ZWJwYWNrX3JlcXVpcmVfXyhfX3dlYnBhY2tfcmVxdWlyZV9fLnMgPSBcIi4vc3JjL21haW4udHNcIik7XG4iLCJleHBvcnQgY29uc3QgYW5zd2VySW5wdXQgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgnYW5zd2VyJykgYXMgSFRNTElucHV0RWxlbWVudDtcbmV4cG9ydCBjb25zdCBwcmV2aW91c0J1dHRvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdwcmV2aW91cycpIGFzIEhUTUxCdXR0b25FbGVtZW50O1xuZXhwb3J0IGNvbnN0IG5leHRCdXR0b24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgnbmV4dCcpIGFzIEhUTUxCdXR0b25FbGVtZW50O1xuZXhwb3J0IGNvbnN0IHN0b3BCdXR0b24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgnc3RvcCcpIGFzIEhUTUxCdXR0b25FbGVtZW50O1xuZXhwb3J0IGNvbnN0IHByb21wdFNwYW4gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgncHJvbXB0JykgYXMgSFRNTFNwYW5FbGVtZW50O1xuZXhwb3J0IGNvbnN0IGV4Y2VyY2lzZU51bWJlclNwYW4gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgnZXhjZXJjaXNlLW51bWJlcicpIGFzIEhUTUxTcGFuRWxlbWVudDtcbmV4cG9ydCBjb25zdCBleGNlcmNpc2VTZWN0aW9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ2V4Y2VyY2lzZScpIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHF1aXpEZXNjU2VjdGlvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdxdWl6LWRlc2MnKSBhcyBIVE1MRGl2RWxlbWVudDtcbmV4cG9ydCBjb25zdCBxdWl6Q2hvaXNlU2VjdGlvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdxdWl6LWNob2lzZScpIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHBlbmFsdHlUaW1lU3BhbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdwZW5hbHR5LXRpbWUnKSBhcyBIVE1MU3BhbkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgZGVzY1BhcmFncmFwaCA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdkZXNjJykgYXMgSFRNTFBhcmFncmFwaEVsZW1lbnQ7XG5leHBvcnQgY29uc3QgcmVzdWx0c0RpdiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdyZXN1bHRzJykgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgcXVpemVzRGl2ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ3F1aXplcycpIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IGNvcnJlY3RuZXNzRGl2ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ2NvcnJlY3RuZXNzJykgYXMgSFRNTERpdkVsZW1lbnQ7XG5leHBvcnQgY29uc3QgcmVzdWx0U2VjdGlvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdyZXN1bHQnKSBhcyBIVE1MRGl2RWxlbWVudDtcbmV4cG9ydCBjb25zdCBzYXZlQnV0dG9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ3NhdmUnKSBhcyBIVE1MQnV0dG9uRWxlbWVudDtcbmV4cG9ydCBjb25zdCBzYXZlV2l0aFN0YXRzQnV0dG9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ3NhdmUtd2l0aC1zdGF0cycpIGFzIEhUTUxCdXR0b25FbGVtZW50O1xuZXhwb3J0IGNvbnN0IGFuc3dlckZvcm0gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgnYW5zd2VyLWZvcm0nKSBhcyBIVE1MRm9ybUVsZW1lbnQ7XG5cbmV4cG9ydCBjb25zdCBpbml0UmVzdWx0U3BhbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdpbml0LXJlc3VsdCcpIGFzIEhUTUxTcGFuRWxlbWVudDtcbmV4cG9ydCBjb25zdCBmaW5hbFJlc3VsdFNwYW4gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgnZmluYWwtcmVzdWx0JykgYXMgSFRNTFNwYW5FbGVtZW50O1xuXG5leHBvcnQgY29uc3QgcXVpelByZXZSZXN1bHRzU2VjdGlvbiA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdxdWl6LXByZXYtcmVzdWx0cycpIGFzIEhUTUxEaXZFbGVtZW50O1xuZXhwb3J0IGNvbnN0IHF1aWNrQWNjZXNzRGl2ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ3F1aWNrLWFjY2VzcycpIGFzIEhUTUxEaXZFbGVtZW50O1xuIiwiaW1wb3J0IHsgQW5zd2VyLCBRdWl6LCBnZXRRdWl6IH0gZnJvbSAnLi9xdWl6JztcbmltcG9ydCB7XG5cdHNhdmVXaXRoU3RhdHNCdXR0b24sXG5cdHNhdmVCdXR0b24sXG5cdHJlc3VsdFNlY3Rpb24sXG5cdGluaXRSZXN1bHRTcGFuLFxuXHRmaW5hbFJlc3VsdFNwYW4sXG5cdGNvcnJlY3RuZXNzRGl2LFxuXHRxdWl6UHJldlJlc3VsdHNTZWN0aW9uXG59IGZyb20gJy4vSFRNTEVsZW1lbnRzJztcblxuaW50ZXJmYWNlIFJlc3VsdCB7XG5cdGZpbmFsVGltZTogbnVtYmVyO1xuXHRhbnN3ZXJzPzogQW5zd2VyW107XG5cdHF1aXpJZDogc3RyaW5nO1xufVxuXG5jb25zdCBpc1Jlc3VsdCA9IChvYmo6IGFueSk6IG9iaiBpcyBSZXN1bHQgPT4ge1xuXHRyZXR1cm4gb2JqLmZpbmFsVGltZSAmJiBvYmoucXVpeklkO1xufTtcblxuZXhwb3J0IGNsYXNzIFF1aXpSZXN1bHRzIHtcblx0cHJpdmF0ZSBwZW5hbHRpZXM6IG51bWJlcltdID0gW107XG5cblx0LyoqXG5cdCAqIEBwYXJhbSBhbnN3ZXJzIGFycmF5IG9mIG1hcmtlZCBhbnN3ZXJzXG5cdCAqIEBwYXJhbSBxdWl6IHF1aXogb2JqZWN0XG5cdCAqIEBwYXJhbSBxdWl6SWQgY2hvc2VuIHF1aXogaWRcblx0ICovXG5cdGNvbnN0cnVjdG9yKHByaXZhdGUgYW5zd2VyczogQW5zd2VyW10sIHByaXZhdGUgcXVpejogUXVpeiB8IG51bGwpIHtcblx0XHR0aGlzLmJpbmRFdmVudEhhbmRsZXJzKCk7XG5cdFx0dGhpcy5jaGVja0Fuc3dlcnMoKTtcblx0XHR0aGlzLmdpdmVQZW5hbHRpZXMoKTtcblx0XHR0aGlzLmRpc3BsYXkoKTtcblx0fVxuXG5cdC8qKlxuXHQgKiBDaGVja3MgYW5zd2Vycywgc2V0cyBjYXJyZWN0IGFyZ3VtZW50IGluIGVhY2ggYW5zd2VyIHRvIGVpdGhlciB0cnVlIG9yIGZhbHNlLlxuXHQgKi9cblx0cHJpdmF0ZSBjaGVja0Fuc3dlcnMoKSB7XG5cdFx0aWYgKHRoaXMucXVpeiAhPT0gbnVsbCkge1xuXHRcdFx0Zm9yIChjb25zdCBpIGluIHRoaXMuYW5zd2Vycykge1xuXHRcdFx0XHR0aGlzLmFuc3dlcnNbaV0uY29ycmVjdCA9IHRoaXMuYW5zd2Vyc1tpXS5hbnN3ZXIgPT09IHRoaXMucXVpei5xdWVzdGlvbnNbaV0uYW5zd2VyO1xuXHRcdFx0fVxuXHRcdH1cblx0fVxuXG5cdC8qKlxuXHQgKiBHaXZlcyBwZW5hbHRpZXMgZm9yIGVhY2ggd3JvbmcgYW5zd2VyLlxuXHQgKi9cblx0cHJpdmF0ZSBnaXZlUGVuYWx0aWVzKCkge1xuXHRcdGlmICh0aGlzLnF1aXogIT09IG51bGwpIHtcblx0XHRcdGZvciAoY29uc3QgaSBpbiB0aGlzLmFuc3dlcnMpIHtcblx0XHRcdFx0dGhpcy5wZW5hbHRpZXMucHVzaCh0aGlzLmFuc3dlcnNbaV0uY29ycmVjdCA/IDAgOiB0aGlzLnF1aXoucXVlc3Rpb25zW2ldLnBlbmFsdHkpO1xuXHRcdFx0fVxuXHRcdH1cblx0fVxuXG5cdC8qKlxuXHQgKiBCaW5kIG5lY2Nlc3NhcnkgZXZlbnQgaGFuZGxlcnMuXG5cdCAqL1xuXHRwcml2YXRlIGJpbmRFdmVudEhhbmRsZXJzKCkge1xuXHRcdHNhdmVXaXRoU3RhdHNCdXR0b24ub25jbGljayA9ICgpID0+IHtcblx0XHRcdHRoaXMuc2F2ZVdpdGhTdGF0cygpO1xuXHRcdFx0dGhpcy5nb1RvTWFpblNjcmVlbigpO1xuXHRcdH07XG5cdFx0c2F2ZUJ1dHRvbi5vbmNsaWNrID0gKCkgPT4ge1xuXHRcdFx0dGhpcy5zYXZlKCk7XG5cdFx0XHR0aGlzLmdvVG9NYWluU2NyZWVuKCk7XG5cdFx0fTtcblx0fVxuXG5cdC8qKlxuXHQgKiBHb2VzIGJhY2sgdG8gbWFpbiBzY3JlZW4uXG5cdCAqL1xuXHRwcml2YXRlIGdvVG9NYWluU2NyZWVuKCkge1xuXHRcdHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQoKTtcblx0fVxuXG5cdC8qKlxuXHQgKiBTYXZlcyBydW4gcmVzdWx0cyB3aXRoIHN0YXRzLlxuXHQgKi9cblx0cHJpdmF0ZSBzYXZlV2l0aFN0YXRzKCkge1xuXHRcdHRoaXMuc2F2ZVJlc3VsdCh7IGZpbmFsVGltZTogdGhpcy5maW5hbFRpbWUsIGFuc3dlcnM6IHRoaXMuYW5zd2VycywgcXVpeklkOiB0aGlzLnF1aXo/LmlkID8/IFwiXCJ9KTtcblx0fVxuXG5cdC8qKlxuXHQgKiBTYXZlcyBydW4gcmVzdWx0cyB3aXRob3V0IHN0YXRzLlxuXHQgKi9cblx0cHJpdmF0ZSBzYXZlKCkge1xuXHRcdHRoaXMuc2F2ZVJlc3VsdCh7IGZpbmFsVGltZTogdGhpcy5maW5hbFRpbWUsIHF1aXpJZDogdGhpcy5xdWl6Py5pZCA/PyBcIlwiIH0pO1xuXHR9XG5cblx0LyoqXG5cdCAqIFNhdmVzIGdpdmVuIHJlc3VsdC5cblx0ICogQHBhcmFtIHJlc3VsdCByZXN1bHRcblx0ICovXG5cdHByaXZhdGUgc2F2ZVJlc3VsdChyZXN1bHQ6IFJlc3VsdCkge1xuXHRcdGNvbnN0IHByZXZKU09OID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ3Jlc3VsdHMnKSB8fCAnW10nO1xuXHRcdGNvbnN0IHByZXYgPSBKU09OLnBhcnNlKHByZXZKU09OKSBhcyBBcnJheTxSZXN1bHQ+O1xuXHRcdHByZXYucHVzaChyZXN1bHQpO1xuXHRcdGxvY2FsU3RvcmFnZS5zZXRJdGVtKCdyZXN1bHRzJywgSlNPTi5zdHJpbmdpZnkocHJldikpO1xuXHR9XG5cblx0LyoqXG5cdCAqIENhbGN1bGF0ZXMgdGltZSBzcGVudCBvbiBxdWl6LlxuXHQgKi9cblx0cHJpdmF0ZSBnZXQgaW5pdFRpbWUoKTogbnVtYmVyIHtcblx0XHRsZXQgaW5pdFRpbWUgPSAwO1xuXHRcdGZvciAoY29uc3QgYW5zIG9mIHRoaXMuYW5zd2Vycykge1xuXHRcdFx0aW5pdFRpbWUgKz0gYW5zLnRpbWUgLyAxMDAwO1xuXHRcdH1cblx0XHRyZXR1cm4gaW5pdFRpbWU7XG5cdH1cblxuXHQvKipcblx0ICogQ2FsY3VsYXRlcyBmaW5hbCB0aW1lLCBiYXNlICsgcGVuYWx0aWVzLlxuXHQgKi9cblx0cHVibGljIGdldCBmaW5hbFRpbWUoKSB7XG5cdFx0bGV0IGZpbmFsVGltZSA9IHRoaXMuaW5pdFRpbWU7XG5cdFx0Zm9yIChjb25zdCBwIG9mIHRoaXMucGVuYWx0aWVzKSB7XG5cdFx0XHRmaW5hbFRpbWUgKz0gcDtcblx0XHR9XG5cdFx0cmV0dXJuIGZpbmFsVGltZTtcblx0fVxuXG5cdC8qKlxuXHQgKiBEaXNwbGF5cyByZXN1bHRzLlxuXHQgKi9cblx0cHJpdmF0ZSBkaXNwbGF5KCkge1xuXHRcdGlmICh0aGlzLnF1aXopIHtcblx0XHRcdHJlc3VsdFNlY3Rpb24uc3R5bGUuZGlzcGxheSA9ICdibG9jayc7XG5cblx0XHRcdGluaXRSZXN1bHRTcGFuLnRleHRDb250ZW50ID0gUXVpelJlc3VsdHMuZm9ybWF0VGltZSh0aGlzLmluaXRUaW1lKTtcblx0XHRcdGZpbmFsUmVzdWx0U3Bhbi50ZXh0Q29udGVudCA9IFF1aXpSZXN1bHRzLmZvcm1hdFRpbWUodGhpcy5maW5hbFRpbWUpO1xuXG5cdFx0XHQvLyBjb25zb2xlLmxvZyh0aGlzLmFuc3dlcnMuZW50cmllcygpKTtcblx0XHRcdGZvciAoY29uc3QgWyBpLCBhbnN3ZXIgXSBvZiB0aGlzLmFuc3dlcnMuZW50cmllcygpKSB7XG5cdFx0XHRcdGNvbnNvbGUubG9nKGksIGFuc3dlcik7XG5cdFx0XHRcdGNvbnN0IGRpdiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuXHRcdFx0XHRkaXYudGV4dENvbnRlbnQgPSBgJHtpICsgMX0uIGA7XG5cdFx0XHRcdGNvbnN0IHNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG5cdFx0XHRcdHNwYW4uY2xhc3NOYW1lID0gYW5zd2VyLmNvcnJlY3QgPyAnY29ycmVjdCcgOiAnaW5jb3JyZWN0Jztcblx0XHRcdFx0c3Bhbi50ZXh0Q29udGVudCA9IGFuc3dlci5jb3JyZWN0ID8gJ0NvcnJlY3QgOiknIDogYEluY29yZWN0OiArICR7dGhpcy5xdWl6LnF1ZXN0aW9uc1tpXS5wZW5hbHR5fXNgO1xuXHRcdFx0XHRkaXYuYXBwZW5kQ2hpbGQoc3Bhbik7XG5cdFx0XHRcdGNvcnJlY3RuZXNzRGl2LmFwcGVuZENoaWxkKGRpdik7XG5cdFx0XHR9XG5cdFx0fVxuXHR9XG5cblx0LyoqXG5cdCAqIEZvcm1hdHMgdGltZSwgcm91bmQgdG8gMyBkaWdpdHMuXG5cdCAqIEBwYXJhbSB0aW1lIHRpbWVcblx0ICovXG5cdHB1YmxpYyBzdGF0aWMgZm9ybWF0VGltZSh0aW1lOiBudW1iZXIpOiBzdHJpbmcge1xuXHRcdHJldHVybiB0aW1lLnRvRml4ZWQoMykudG9TdHJpbmcoKTtcblx0fVxuXG5cdC8qKlxuXHQgKiBEaXNwbGF5cyBwcmV2aW91cyByZXN1bHRzLlxuXHQgKi9cblx0cHVibGljIHN0YXRpYyBkaXNwbGF5UHJldmlvdXNSZXN1bHRzKCk6IHZvaWQge1xuXHRcdGNvbnN0IHJlc3VsdHMgPSBKU09OLnBhcnNlKGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdyZXN1bHRzJykgfHwgJycpO1xuXHRcdGxldCBpID0gMTtcblx0XHRjb25zb2xlLmxvZyhyZXN1bHRzKVxuXHRcdGlmIChyZXN1bHRzIGluc3RhbmNlb2YgQXJyYXkpIHtcblx0XHRcdGZvciAoY29uc3QgcmVzIG9mIHJlc3VsdHMpIHtcblx0XHRcdFx0Y29uc29sZS5sb2cocmVzLCBpc1Jlc3VsdChyZXMpKTtcblx0XHRcdFx0aWYgKGlzUmVzdWx0KHJlcykpIHtcblx0XHRcdFx0XHRjb25zdCBxdWl6ID0gZ2V0UXVpeihyZXMucXVpeklkKTtcblx0XHRcdFx0XHRpZiAocXVpeikge1xuXHRcdFx0XHRcdFx0Y29uc3Qgcm93ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnZGl2Jyk7XG5cdFx0XHRcdFx0XHRjb25zdCBuYW1lQ29sID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnZGl2Jyk7XG5cdFx0XHRcdFx0XHRjb25zdCByZXN1bHRDb2wgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcblxuXHRcdFx0XHRcdFx0cmVzdWx0Q29sLmNsYXNzTmFtZSA9ICdjb2wtNCBkLWZsZXgganVzdGlmeS1jb250ZW50LWNlbnRlcic7XG5cdFx0XHRcdFx0XHRyZXN1bHRDb2wudGV4dENvbnRlbnQgPSBRdWl6UmVzdWx0cy5mb3JtYXRUaW1lKHJlcy5maW5hbFRpbWUpICsgJ3MnO1xuXG5cdFx0XHRcdFx0XHRuYW1lQ29sLmNsYXNzTmFtZSA9ICdjb2wtOCBkLWZsZXgganVzdGlmeS1jb250ZW50LWNlbnRlcic7XG5cdFx0XHRcdFx0XHRuYW1lQ29sLnRleHRDb250ZW50ID0gcXVpei5kZXNjO1xuXG5cdFx0XHRcdFx0XHRyb3cuY2xhc3NOYW1lID0gYHJvdyBwcmV2LXJlc3VsdCAke2kgJSAyID09IDAgPyAnZXZlbicgOiAnJ31gO1xuXG5cdFx0XHRcdFx0XHRyb3cuYXBwZW5kQ2hpbGQobmFtZUNvbCk7XG5cdFx0XHRcdFx0XHRyb3cuYXBwZW5kQ2hpbGQocmVzdWx0Q29sKTtcblxuXHRcdFx0XHRcdFx0cXVpelByZXZSZXN1bHRzU2VjdGlvbi5hcHBlbmRDaGlsZChyb3cpO1xuXHRcdFx0XHRcdFx0aSsrO1xuXHRcdFx0XHRcdH1cblx0XHRcdFx0fVxuXHRcdFx0fVxuXHRcdH1cblx0fVxufVxuIiwiaW1wb3J0IHsgUXVpeiwgQW5zd2VyLCBnZXRRdWl6IH0gZnJvbSAnLi9xdWl6JztcbmltcG9ydCB7IFF1aXpSZXN1bHRzIH0gZnJvbSAnLi9RdWl6UmVzdWx0cyc7XG5pbXBvcnQge1xuXHRkZXNjUGFyYWdyYXBoLFxuXHRhbnN3ZXJJbnB1dCxcblx0c3RvcEJ1dHRvbixcblx0ZXhjZXJjaXNlU2VjdGlvbixcblx0bmV4dEJ1dHRvbixcblx0cHJldmlvdXNCdXR0b24sXG5cdGV4Y2VyY2lzZU51bWJlclNwYW4sXG5cdHByb21wdFNwYW4sXG5cdHBlbmFsdHlUaW1lU3Bhbixcblx0cXVpekNob2lzZVNlY3Rpb24sXG5cdHF1aXpEZXNjU2VjdGlvbixcblx0YW5zd2VyRm9ybSxcblx0cXVpelByZXZSZXN1bHRzU2VjdGlvbixcblx0cXVpY2tBY2Nlc3NEaXZcbn0gZnJvbSAnLi9IVE1MRWxlbWVudHMnO1xuaW1wb3J0IHsgZXJyb3IgfSBmcm9tICcuL21haW4nO1xuXG5leHBvcnQgY2xhc3MgUXVpelJ1biB7XG5cdHByaXZhdGUgcXVlc3Rpb25OdW1iZXI6IG51bWJlciA9IDA7XG5cdHByaXZhdGUgYW5zd2VyU3RhcnRUaW1lOiBudW1iZXIgPSAwO1xuXHRwcml2YXRlIHF1aXo6IFF1aXogfCBudWxsID0gbnVsbDtcblx0cHJpdmF0ZSBhbnN3ZXJzOiBBbnN3ZXJbXSA9IFtdO1xuXHRwcml2YXRlIHBlbmFsdGllczogbnVtYmVyW10gPSBbXTtcblx0cHJpdmF0ZSBxdWl6UmVzdWx0czogUXVpelJlc3VsdHMgfCBudWxsID0gbnVsbDtcblxuXHQvKipcblx0ICogQHBhcmFtIHF1aXpJZCBjaG9zZW4gcXVpeiBpZFxuXHQgKi9cblx0Y29uc3RydWN0b3IocXVpeklkOiBzdHJpbmcgPSAnMCcpIHtcblx0XHR0aGlzLnF1aXogPSBnZXRRdWl6KHF1aXpJZCk7XG5cdFx0aWYgKHRoaXMucXVpeiA9PSBudWxsKSB7XG5cdFx0XHRlcnJvcignUG9kYW55IHF1aXogbmllIGlzdG5pZWplJyk7XG5cdFx0XHRyZXR1cm47XG5cdFx0fSBlbHNlIGlmICh0aGlzLnF1aXoucXVlc3Rpb25zLmxlbmd0aCA9PSAwKSB7XG5cdFx0XHRlcnJvcignUG9kYW55IHF1aXogbWEgemEgbWHFgm8gcHl0YcWEJyk7XG5cdFx0XHRyZXR1cm47XG5cdFx0fVxuXHRcdHF1aXpDaG9pc2VTZWN0aW9uLnN0eWxlLmRpc3BsYXkgPSAnbm9uZSc7XG5cdFx0cXVpelByZXZSZXN1bHRzU2VjdGlvbi5zdHlsZS5kaXNwbGF5ID0gJ25vbmUnO1xuXHRcdHF1aXpEZXNjU2VjdGlvbi5zdHlsZS5kaXNwbGF5ID0gJ2Jsb2NrJztcblx0XHRkZXNjUGFyYWdyYXBoLnRleHRDb250ZW50ID0gdGhpcy5xdWl6LmRlc2M7XG5cblx0XHR0aGlzLnF1aXoucXVlc3Rpb25zLmZvckVhY2goKF8sIGkpID0+IHtcblx0XHRcdHRoaXMuYW5zd2Vycy5wdXNoKHsgdGltZTogMCwgYW5zd2VyOiBudWxsLCBjb3JyZWN0OiBmYWxzZSB9KTtcblx0XHRcdHRoaXMucGVuYWx0aWVzLnB1c2goMCk7XG5cblx0XHRcdGNvbnN0IGFjY2VzcyA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcblx0XHRcdGFjY2Vzcy50ZXh0Q29udGVudCA9IChpICsgMSkudG9TdHJpbmcoKTtcblx0XHRcdGFjY2Vzcy5jbGFzc05hbWUgPSAncXVpY2stYWNjZXNzLW5vZGUnO1xuXHRcdFx0YWNjZXNzLmRhdGFzZXQucXVlc3Rpb24gPSBpLnRvU3RyaW5nKCk7XG5cblx0XHRcdGFjY2Vzcy5vbmNsaWNrID0gKCkgPT4ge1xuXHRcdFx0XHR0aGlzLnNhdmVBbnN3ZXJUaW1lKCk7XG5cdFx0XHRcdHRoaXMucXVlc3Rpb25OdW1iZXIgPSBpO1xuXG5cdFx0XHRcdHRoaXMuY2hhbmdlUXVlc3Rpb24oKTtcblx0XHRcdH07XG5cblx0XHRcdHF1aWNrQWNjZXNzRGl2LmFwcGVuZENoaWxkKGFjY2Vzcyk7XG5cdFx0fSk7XG5cblx0XHR0aGlzLmJpbmRFdmVudEhhbmRsZXJzKCk7XG5cblx0XHRleGNlcmNpc2VTZWN0aW9uLnN0eWxlLmRpc3BsYXkgPSAnYmxvY2snO1xuXHRcdHRoaXMuYW5zd2VyU3RhcnRUaW1lID0gcGVyZm9ybWFuY2Uubm93KCk7XG5cblx0XHR0aGlzLmNoYW5nZVF1ZXN0aW9uKCk7XG5cdH1cblxuXHQvKipcblx0ICogQmluZHMgbmVjY2Vzc2FyeSBldmVudCBoYW5kbGVycy5cblx0ICovXG5cdHByaXZhdGUgYmluZEV2ZW50SGFuZGxlcnMoKTogdm9pZCB7XG5cdFx0YW5zd2VySW5wdXQub25pbnB1dCA9ICgpID0+IHtcblx0XHRcdHRoaXMuc2F2ZUFuc3dlcigpO1xuXG5cdFx0XHRzdG9wQnV0dG9uLmRpc2FibGVkID0gdGhpcy5hbnN3ZXJzLnNvbWUoKGEpID0+IGEuYW5zd2VyID09PSBudWxsKTtcblx0XHRcdGNvbnN0IHF1aWNrQWNjZXNzU3BhbiA9IGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3IoXG5cdFx0XHRcdGBbZGF0YS1xdWVzdGlvbj1cIiR7dGhpcy5xdWVzdGlvbk51bWJlcn1cIl1gXG5cdFx0XHQpIGFzIEhUTUxTcGFuRWxlbWVudDtcblx0XHRcdGlmICh0aGlzLmFuc3dlcnNbdGhpcy5xdWVzdGlvbk51bWJlcl0uYW5zd2VyID09PSBudWxsKSB7XG5cdFx0XHRcdHF1aWNrQWNjZXNzU3Bhbi5jbGFzc05hbWUgPSBxdWlja0FjY2Vzc1NwYW4uY2xhc3NOYW1lLnJlcGxhY2UoL2Fuc3dlcmVkLywgJycpO1xuXHRcdFx0fSBlbHNlIHtcblx0XHRcdFx0cXVpY2tBY2Nlc3NTcGFuLmNsYXNzTmFtZSArPSAnIGFuc3dlcmVkJztcblx0XHRcdH1cblx0XHR9O1xuXG5cdFx0bmV4dEJ1dHRvbi5vbmNsaWNrID0gKCkgPT4ge1xuXHRcdFx0dGhpcy5zYXZlQW5zd2VyVGltZSgpO1xuXHRcdFx0dGhpcy5xdWVzdGlvbk51bWJlcisrO1xuXG5cdFx0XHR0aGlzLmNoYW5nZVF1ZXN0aW9uKCk7XG5cdFx0fTtcblxuXHRcdHByZXZpb3VzQnV0dG9uLm9uY2xpY2sgPSAoKSA9PiB7XG5cdFx0XHR0aGlzLnNhdmVBbnN3ZXJUaW1lKCk7XG5cdFx0XHR0aGlzLnF1ZXN0aW9uTnVtYmVyLS07XG5cblx0XHRcdHRoaXMuY2hhbmdlUXVlc3Rpb24oKTtcblx0XHR9O1xuXG5cdFx0c3RvcEJ1dHRvbi5vbmNsaWNrID0gKCkgPT4ge1xuXHRcdFx0dGhpcy5lbmRRdWl6KCk7XG5cdFx0fTtcblxuXHRcdGFuc3dlckZvcm0ub25zdWJtaXQgPSAoKSA9PiB7XG5cdFx0XHRpZiAoc3RvcEJ1dHRvbi5kaXNhYmxlZCkge1xuXHRcdFx0XHRuZXh0QnV0dG9uLmNsaWNrKCk7XG5cdFx0XHR9IGVsc2Uge1xuXHRcdFx0XHRzdG9wQnV0dG9uLmNsaWNrKCk7XG5cdFx0XHR9XG5cdFx0XHRyZXR1cm4gZmFsc2U7XG5cdFx0fTtcblx0fVxuXG5cdC8qKlxuICAgICogQ2hhbmdlcyBjdXJyZW50IHF1ZXN0aW9uIHRvIGdpdmVuLlxuICAgICogU2V0cyBwcmV2QW5zd2VyU3RhcnRUaW1lIHRvICBjdXJyZW50IHRpbWVzdGFtcC5cbiAgICAqL1xuXHRwcml2YXRlIGNoYW5nZVF1ZXN0aW9uKCk6IHZvaWQge1xuXHRcdGlmICh0aGlzLnF1aXogPT09IG51bGwpIHtcblx0XHRcdHJldHVybjtcblx0XHR9XG5cdFx0Y29uc3QgcXVlc3Rpb24gPSB0aGlzLnF1aXoucXVlc3Rpb25zW3RoaXMucXVlc3Rpb25OdW1iZXJdO1xuXHRcdGV4Y2VyY2lzZU51bWJlclNwYW4udGV4dENvbnRlbnQgPSAodGhpcy5xdWVzdGlvbk51bWJlciArIDEpLnRvU3RyaW5nKCk7XG5cdFx0cHJvbXB0U3Bhbi50ZXh0Q29udGVudCA9IHF1ZXN0aW9uLnByb21wdDtcblx0XHRwZW5hbHR5VGltZVNwYW4udGV4dENvbnRlbnQgPSBxdWVzdGlvbi5wZW5hbHR5LnRvU3RyaW5nKCk7XG5cblx0XHRjb25zdCBhbnN3ZXIgPSB0aGlzLmFuc3dlcnNbdGhpcy5xdWVzdGlvbk51bWJlcl0uYW5zd2VyO1xuXHRcdGFuc3dlcklucHV0LnZhbHVlID0gYW5zd2VyID09PSBudWxsID8gJycgOiBhbnN3ZXIudG9TdHJpbmcoKTtcblxuXHRcdHByZXZpb3VzQnV0dG9uLmRpc2FibGVkID0gdGhpcy5xdWVzdGlvbk51bWJlciA9PT0gMDtcblx0XHRuZXh0QnV0dG9uLmRpc2FibGVkID0gdGhpcy5xdWVzdGlvbk51bWJlciA9PT0gdGhpcy5xdWl6LnF1ZXN0aW9ucy5sZW5ndGggLSAxO1xuXG5cdFx0dGhpcy5hbnN3ZXJTdGFydFRpbWUgPSBwZXJmb3JtYW5jZS5ub3coKTtcblx0fVxuXG5cdC8qKlxuICAgICogU2F2ZXMgY29udGVzdGFudHMgYW5zd2VyIHRvIGN1cnJlbnQgcXVlc3Rpb24uIFxuICAgICovXG5cdHByaXZhdGUgc2F2ZUFuc3dlciA9ICgpOiB2b2lkID0+IHtcblx0XHRjb25zdCBwcmV2QW5zd2VyID0gdGhpcy5hbnN3ZXJzW3RoaXMucXVlc3Rpb25OdW1iZXJdO1xuXG5cdFx0aWYgKHByZXZBbnN3ZXIgPT0gbnVsbCkge1xuXHRcdFx0cmV0dXJuO1xuXHRcdH1cblxuXHRcdGNvbnN0IGFuc3dlciA9IHBhcnNlSW50KGFuc3dlcklucHV0LnZhbHVlKTtcblx0XHRpZiAoaXNOYU4oYW5zd2VyKSkge1xuXHRcdFx0cHJldkFuc3dlci5hbnN3ZXIgPSBudWxsO1xuXHRcdH0gZWxzZSB7XG5cdFx0XHRwcmV2QW5zd2VyLmFuc3dlciA9IGFuc3dlcjtcblx0XHR9XG5cdH07XG5cblx0LyoqXG4gICAgKiBTYXZlcyBjb250ZXN0YW50cyBhbnN3ZXIgdGltZSB0byBjdXJyZW50IHF1ZXN0aW9uLlxuICAgICovXG5cdHByaXZhdGUgc2F2ZUFuc3dlclRpbWUgPSAoKTogdm9pZCA9PiB7XG5cdFx0Y29uc3QgZW5kVGltZSA9IHBlcmZvcm1hbmNlLm5vdygpO1xuXG5cdFx0Y29uc3QgcHJldkFuc3dlciA9IHRoaXMuYW5zd2Vyc1t0aGlzLnF1ZXN0aW9uTnVtYmVyXTtcblxuXHRcdGlmIChwcmV2QW5zd2VyID09IG51bGwpIHtcblx0XHRcdHJldHVybjtcblx0XHR9XG5cblx0XHRwcmV2QW5zd2VyLnRpbWUgKz0gZW5kVGltZSAtIHRoaXMuYW5zd2VyU3RhcnRUaW1lO1xuXHR9O1xuXG5cdC8qKlxuXHQgKiBFbmQgcXVpei4gRGlzcGxheXMgcmVzdWx0cy5cblx0ICovXG5cdHByaXZhdGUgZW5kUXVpeigpOiB2b2lkIHtcblx0XHR0aGlzLnNhdmVBbnN3ZXJUaW1lKCk7XG5cdFx0ZXhjZXJjaXNlU2VjdGlvbi5zdHlsZS5kaXNwbGF5ID0gJ25vbmUnO1xuXG5cdFx0dGhpcy5xdWl6UmVzdWx0cyA9IG5ldyBRdWl6UmVzdWx0cyh0aGlzLmFuc3dlcnMsIHRoaXMucXVpeik7XG5cdH1cbn1cbiIsImltcG9ydCB7IFF1aXosIGdldFF1aXplcyB9IGZyb20gJy4vcXVpeic7XG5pbXBvcnQgeyBRdWl6UnVuIH0gZnJvbSAnLi9RdWl6UnVuJztcbmltcG9ydCB7IFF1aXpSZXN1bHRzIH0gZnJvbSAnLi9RdWl6UmVzdWx0cyc7XG5pbXBvcnQgeyBxdWl6ZXNEaXYsIHF1aXpDaG9pc2VTZWN0aW9uIH0gZnJvbSAnLi9IVE1MRWxlbWVudHMnO1xuXG4vKipcbiAqIEZ1bmN0aW9uIHRoYXQgbG9ncyBlcnJvcnMuXG4gKiBNYXkgaW4gdGhlIGZ1dHVyZSBiZSBvdmVyd3JpdGVuIHRvIGZ1bmN0aW9uIHRoYXQgZGlzcGxheSBlcnJvciBvbiBzY3JlZW4uXG4gKi9cbmV4cG9ydCBjb25zdCBlcnJvciA9IGNvbnNvbGUuZXJyb3I7XG5cbi8qKlxuICogQWRkcyBxdWl6IHRvIHRhYmxlIG9mIGF2YWlsYWJsZSBxdWl6ZXMuXG4gKiBcbiAqIEBwYXJhbSBpZCBxdWl6IGlkXG4gKiBAcGFyYW0gcXVpeiBxdWl6IG9iamVjdFxuICovXG5jb25zdCBhZGRRdWl6VG9UYWJsZSA9IChxdWl6OiBRdWl6KSA9PiB7XG5cdGNvbnN0IHJvdyA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuXHRjb25zdCBuYW1lQ29sID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnZGl2Jyk7XG5cdGNvbnN0IGFjdGlvbkNvbCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuXHRjb25zdCBidXR0b24gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdidXR0b24nKTtcblxuXHRidXR0b24uY2xhc3NOYW1lID0gJ2J0biBidG4tcHJpbWFyeSBxdWl6LXN0YXJ0LWJ1dHRvbic7XG5cdGJ1dHRvbi5kYXRhc2V0LnF1aXpJZCA9IHF1aXouaWQ7XG5cdGJ1dHRvbi50ZXh0Q29udGVudCA9ICdTdGFydCc7XG5cblx0YWN0aW9uQ29sLmNsYXNzTmFtZSA9ICdjb2wtNCBkLWZsZXgganVzdGlmeS1jb250ZW50LWNlbnRlciBhbGlnbi1pdGVtcy1jZW50ZXInO1xuXHRhY3Rpb25Db2wuYXBwZW5kQ2hpbGQoYnV0dG9uKTtcblxuXHRuYW1lQ29sLmNsYXNzTmFtZSA9ICdjb2wtOCBkLWZsZXgganVzdGlmeS1jb250ZW50LWNlbnRlciBhbGlnbi1pdGVtcy1jZW50ZXInO1xuXHRuYW1lQ29sLnRleHRDb250ZW50ID0gcXVpei5kZXNjO1xuXG5cdHJvdy5jbGFzc05hbWUgPSAncm93JztcblxuXHRyb3cuYXBwZW5kQ2hpbGQobmFtZUNvbCk7XG5cdHJvdy5hcHBlbmRDaGlsZChhY3Rpb25Db2wpO1xuXG5cdHF1aXplc0Rpdi5hcHBlbmRDaGlsZChyb3cpO1xufTtcblxuLyoqXG4gKiBEaXNwbGF5cyBhdmFpbGFibGUgcXVpemVzIGluIHRoZSB0YWJsZS5cbiAqL1xuY29uc3QgZGlzcGxheVF1aXplcyA9ICgpID0+IHtcblx0Zm9yIChjb25zdCBxdWl6IG9mIGdldFF1aXplcygpKSB7XG5cdFx0Y29uc29sZS5sb2cocXVpeik7XG5cdFx0YWRkUXVpelRvVGFibGUocXVpeik7XG5cdH1cblxuXHRjb25zdCBidXR0b25zID0gZG9jdW1lbnQuZ2V0RWxlbWVudHNCeUNsYXNzTmFtZSgncXVpei1zdGFydC1idXR0b24nKSBhcyBIVE1MQ29sbGVjdGlvbk9mPEhUTUxCdXR0b25FbGVtZW50Pjtcblx0QXJyYXkuZnJvbShidXR0b25zKS5mb3JFYWNoKChidXR0b24pID0+IHtcblx0XHRidXR0b24ub25jbGljayA9ICgpID0+IHtcblx0XHRcdG5ldyBRdWl6UnVuKGJ1dHRvbi5kYXRhc2V0LnF1aXpJZCk7XG5cdFx0fTtcblx0fSk7XG59O1xuXG5kaXNwbGF5UXVpemVzKCk7XG5cblF1aXpSZXN1bHRzLmRpc3BsYXlQcmV2aW91c1Jlc3VsdHMoKTtcbiIsImltcG9ydCB7IHF1aXogfSBmcm9tICcuL3F1aXplcyc7XG5cbmV4cG9ydCBpbnRlcmZhY2UgUXVlc3Rpb24ge1xuXHRwcm9tcHQ6IHN0cmluZztcblx0YW5zd2VyOiBudW1iZXI7XG5cdHBlbmFsdHk6IG51bWJlcjtcbn1cblxuZXhwb3J0IGludGVyZmFjZSBRdWl6IHtcblx0ZGVzYzogc3RyaW5nO1xuXHRxdWVzdGlvbnM6IFF1ZXN0aW9uW107XG5cdGlkOiBzdHJpbmc7XG59XG5cbmV4cG9ydCBpbnRlcmZhY2UgQW5zd2VyIHtcblx0YW5zd2VyOiBudW1iZXIgfCBudWxsO1xuXHR0aW1lOiBudW1iZXI7XG5cdGNvcnJlY3Q6IGJvb2xlYW47XG59XG5cbmNvbnN0IGlzUXVlc3Rpb24gPSAocXVlc3Rpb246IGFueSk6IHF1ZXN0aW9uIGlzIFF1ZXN0aW9uID0+IHtcblx0cmV0dXJuIHF1ZXN0aW9uICYmIHF1ZXN0aW9uLnByb21wdCAmJiBxdWVzdGlvbi5hbnN3ZXIgJiYgcXVlc3Rpb24ucGVuYWx0eTtcbn07XG5cbmNvbnN0IGlzUXVpeiA9IChxdWl6OiBhbnkpOiBxdWl6IGlzIFF1aXogPT4ge1xuXHRyZXR1cm4gcXVpei5kZXNjICYmIHF1aXoucXVlc3Rpb25zICYmIHF1aXouaWQgIT09IHVuZGVmaW5lZDtcbn07XG5cbi8qKlxuICogUmV0dXJucyBxdWl6IHdpdGggZ2l2ZW4gaWQgb3IgbnVsbCBvbiBlcnJvci5cbiAqIEBwYXJhbSBpZCBxdWl6IGlkXG4gKi9cbmV4cG9ydCBjb25zdCBnZXRRdWl6ID0gKGlkOiBzdHJpbmcpOiBRdWl6IHwgbnVsbCA9PiB7XG5cdGNvbnN0IGFyciA9IEpTT04ucGFyc2UocXVpeik/LnF1aXplcztcblx0aWYgKGFyciBpbnN0YW5jZW9mIEFycmF5KSB7XG5cdFx0Y29uc3QgcXVpeiA9IGFyci5maW5kKChlbDogYW55KSA9PiBlbC5pZCA9PT0gaWQpO1xuXG5cdFx0aWYgKGlzUXVpeihxdWl6KSkge1xuXHRcdFx0cmV0dXJuIHF1aXo7XG5cdFx0fVxuXHR9XG5cdHJldHVybiBudWxsO1xufTtcblxuLyoqXG4gKiBHZW5lcmF0b3IgdGhhdCByZXR1cm5zIGFsbCBxdWl6ZXMuXG4gKi9cbmV4cG9ydCBmdW5jdGlvbiogZ2V0UXVpemVzKCk6IEdlbmVyYXRvcjxRdWl6PiB7XG5cdGNvbnN0IG9iaiA9IEpTT04ucGFyc2UocXVpeik7XG5cdGlmIChvYmoucXVpemVzIGluc3RhbmNlb2YgQXJyYXkpIHtcblx0XHRmb3IgKGNvbnN0IHF1aXogb2Ygb2JqLnF1aXplcykge1xuXHRcdFx0aWYgKGlzUXVpeihxdWl6KSkge1xuXHRcdFx0XHR5aWVsZCBxdWl6O1xuXHRcdFx0fVxuXHRcdH1cblx0fVxufVxuIiwiLyoqXG4gKiBFeGFtcGxlIHF1aXplcy5cbiAqL1xuZXhwb3J0IGNvbnN0IHF1aXogPSBge1xuICAgIFwicXVpemVzXCI6W3tcbiAgICAgICBcImlkXCI6IFwiMFwiLFxuICAgICAgIFwiZGVzY1wiOlwiTGljennEhyBrYcW8ZHkgbW/FvGVcIixcbiAgICAgICBcInF1ZXN0aW9uc1wiOltcbiAgICAgICAgICB7XG4gICAgICAgICAgICAgXCJwcm9tcHRcIjpcIjIgKyAzXCIsXG4gICAgICAgICAgICAgXCJhbnN3ZXJcIjo1LFxuICAgICAgICAgICAgIFwicGVuYWx0eVwiOjRcbiAgICAgICAgICB9LFxuICAgICAgICAgIHtcbiAgICAgICAgICAgICBcInByb21wdFwiOlwiMiAtICgtMjQgOiA0KVwiLFxuICAgICAgICAgICAgIFwiYW5zd2VyXCI6OCxcbiAgICAgICAgICAgICBcInBlbmFsdHlcIjoxMFxuICAgICAgICAgIH0sXG4gICAgICAgICAge1xuICAgICAgICAgICAgIFwicHJvbXB0XCI6XCIyIC0gM1wiLFxuICAgICAgICAgICAgIFwiYW5zd2VyXCI6LTEsXG4gICAgICAgICAgICAgXCJwZW5hbHR5XCI6NFxuICAgICAgICAgIH0sXG4gICAgICAgICAge1xuICAgICAgICAgICAgIFwicHJvbXB0XCI6XCIyICsgKC0yNCA6IDQpXCIsXG4gICAgICAgICAgICAgXCJhbnN3ZXJcIjotNCxcbiAgICAgICAgICAgICBcInBlbmFsdHlcIjoxMFxuICAgICAgICAgIH1cbiAgICAgICBdXG4gICAgfSx7XG4gICAgICAgXCJpZFwiOiBcIjFcIixcbiAgICAgICBcImRlc2NcIjpcIlF1aWNrIE1hdGhcIixcbiAgICAgICBcInF1ZXN0aW9uc1wiOltcbiAgICAgICAgICB7XG4gICAgICAgICAgICAgXCJwcm9tcHRcIjpcIjIgKyAyXCIsXG4gICAgICAgICAgICAgXCJhbnN3ZXJcIjo0LFxuICAgICAgICAgICAgIFwicGVuYWx0eVwiOjJcbiAgICAgICAgICB9LFxuICAgICAgICAgIHtcbiAgICAgICAgICAgICBcInByb21wdFwiOlwiMiAqIDNcIixcbiAgICAgICAgICAgICBcImFuc3dlclwiOjYsXG4gICAgICAgICAgICAgXCJwZW5hbHR5XCI6MVxuICAgICAgICAgIH0sXG4gICAgICAgICAge1xuICAgICAgICAgICAgIFwicHJvbXB0XCI6XCIyIF4gM1wiLFxuICAgICAgICAgICAgIFwiYW5zd2VyXCI6OCxcbiAgICAgICAgICAgICBcInBlbmFsdHlcIjozXG4gICAgICAgICAgfSxcbiAgICAgICAgICB7XG4gICAgICAgICAgICAgXCJwcm9tcHRcIjpcIigyIF4gKDEvMikpIF4gMlwiLFxuICAgICAgICAgICAgIFwiYW5zd2VyXCI6MixcbiAgICAgICAgICAgICBcInBlbmFsdHlcIjo3XG4gICAgICAgICAgfVxuICAgICAgIF1cbiAgICB9XVxuIH1gO1xuIl0sInNvdXJjZVJvb3QiOiIifQ==