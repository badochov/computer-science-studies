import { Answer, Quiz, getQuiz } from "./quiz";
import {
  saveWithStatsButton,
  saveButton,
  resultSection,
  initResultSpan,
  finalResultSpan,
  correctnessDiv,
  quizPrevResultsSection,
} from "./HTMLElements";

interface Result {
  finalTime: number;
  answers?: Answer[];
  quizId: string;
}

const isResult = (obj: any): obj is Result => {
  return obj.finalTime && obj.quizId;
};

export class QuizResults {
  private penalties: number[] = [];

  /**
   * @param answers array of marked answers
   * @param quiz quiz object
   * @param quizId chosen quiz id
   */
  constructor(private answers: Answer[], private quiz: Quiz | null) {
    this.bindEventHandlers();
    this.checkAnswers();
    this.givePenalties();
    this.display();
  }

  /**
   * Checks answers, sets carrect argument in each answer to either true or false.
   */
  private checkAnswers() {
    if (this.quiz !== null) {
      for (const i in this.answers) {
        this.answers[i].correct =
          this.answers[i].answer === this.quiz.questions[i].answer;
      }
    }
  }

  /**
   * Gives penalties for each wrong answer.
   */
  private givePenalties() {
    if (this.quiz !== null) {
      for (const i in this.answers) {
        this.penalties.push(
          this.answers[i].correct ? 0 : this.quiz.questions[i].penalty
        );
      }
    }
  }

  /**
   * Bind neccessary event handlers.
   */
  private bindEventHandlers() {
    saveWithStatsButton.onclick = () => {
      this.saveWithStats();
      this.goToMainScreen();
    };
    saveButton.onclick = () => {
      this.save();
      this.goToMainScreen();
    };
  }

  /**
   * Goes back to main screen.
   */
  private goToMainScreen() {
    window.location.reload();
  }

  /**
   * Saves run results with stats.
   */
  private saveWithStats() {
    this.saveResult({
      finalTime: this.finalTime,
      answers: this.answers,
      quizId: this.quiz?.id ?? "",
    });
  }

  /**
   * Saves run results without stats.
   */
  private save() {
    this.saveResult({ finalTime: this.finalTime, quizId: this.quiz?.id ?? "" });
  }

  /**
   * Saves given result.
   * @param result result
   */
  private saveResult(result: Result) {
    const prevJSON = localStorage.getItem("results") || "[]";
    const prev = JSON.parse(prevJSON) as Array<Result>;
    prev.push(result);
    localStorage.setItem("results", JSON.stringify(prev));
  }

  /**
   * Calculates time spent on quiz.
   */
  private get initTime(): number {
    return this.answers.reduce((sum, ans) => sum + ans.time / 1000, 0);
  }

  /**
   * Calculates final time, base + penalties.
   */
  public get finalTime() {
    return this.penalties.reduce(
      (sum, penalty) => sum + penalty,
      this.initTime
    );
  }

  /**
   * Displays results.
   */
  private display() {
    if (this.quiz) {
      resultSection.style.display = "block";

      initResultSpan.textContent = QuizResults.formatTime(this.initTime);
      finalResultSpan.textContent = QuizResults.formatTime(this.finalTime);

      // console.log(this.answers.entries());
      for (const [i, answer] of this.answers.entries()) {
        console.log(i, answer);
        const div = document.createElement("div");
        div.textContent = `${i + 1}. `;
        const span = document.createElement("span");
        span.className = answer.correct ? "correct" : "incorrect";
        span.textContent = answer.correct
          ? "Correct :)"
          : `Incorect: +${this.quiz.questions[i].penalty}s`;
        div.appendChild(span);
        correctnessDiv.appendChild(div);
      }
    }
  }

  /**
   * Formats time, round to 3 digits.
   * @param time time
   */
  public static formatTime(time: number): string {
    return time.toFixed(3).toString();
  }

  /**
   * Displays previous results.
   */
  public static displayPreviousResults(): void {
    const results = JSON.parse(localStorage.getItem("results") || "");
    let i = 1;
    console.log(results);
    if (results instanceof Array) {
      for (const res of results) {
        console.log(res, isResult(res));
        if (isResult(res)) {
          const quiz = getQuiz(res.quizId);
          if (quiz) {
            const row = document.createElement("div");
            const nameCol = document.createElement("div");
            const resultCol = document.createElement("div");

            resultCol.className = "col-4 d-flex justify-content-center";
            resultCol.textContent = QuizResults.formatTime(res.finalTime) + "s";

            nameCol.className = "col-8 d-flex justify-content-center";
            nameCol.textContent = quiz.desc;

            row.className = `row prev-result ${i % 2 == 0 ? "even" : ""}`;

            row.appendChild(nameCol);
            row.appendChild(resultCol);

            quizPrevResultsSection.appendChild(row);
            i++;
          }
        }
      }
    }
  }
}
