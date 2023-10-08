import {
  Answer,
  QuizWithAnswers,
  getQuizWithAnswers,
  getPrevResults,
  averageTimes,
  topScores,
} from "./quiz";
import {
  saveWithStatsButton,
  initResultSpan,
  finalResultSpan,
  correctnessDiv,
  quizPrevResultsSection,
  correctAnswersSection,
  averageTimesSection,
  topScoresSection,
} from "./HTMLElements";
import { error } from "./util";

export interface Result {
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
  constructor(private answers: Answer[], private quiz: QuizWithAnswers) {
    this.bindEventHandlers();
  }

  public mark() {
    this.checkAnswers();
    this.givePenalties();
    this.display();
  }

  /**
   * Checks answers, sets carrect argument in each answer to either true or false.
   */
  private checkAnswers() {
    for (const i in this.answers) {
      this.answers[i].correct =
        this.answers[i].answer === this.quiz.questions[i].answer;
    }
  }

  /**
   * Gives penalties for each wrong answer.
   */
  private givePenalties() {
    for (const i in this.answers) {
      this.penalties.push(
        this.answers[i].correct ? 0 : this.quiz.questions[i].penalty
      );
    }
  }

  /**
   * Bind neccessary event handlers.
   */
  private bindEventHandlers() {
    saveWithStatsButton.onclick = () => {
      this.goToMainScreen();
    };
  }

  /**
   * Goes back to main screen.
   */
  private goToMainScreen() {
    window.location.href = "/";
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
  private displayUserResults() {
    initResultSpan.textContent = QuizResults.formatTime(this.initTime);
    finalResultSpan.textContent = QuizResults.formatTime(this.finalTime);

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
      correctnessDiv.appendChild(div);
    }
  }
  private displayCorrectAnswers() {
    // console.log(this.answers.entries());
    for (const [i, question] of this.quiz.questions.entries()) {
      const div = document.createElement("div");
      div.textContent = `${i + 1}. `;
      const span = document.createElement("span");
      span.textContent = `${question.prompt} = ${question.answer}`;
      div.appendChild(span);
      correctAnswersSection.appendChild(div);
    }
  }
  private async displayAverageTimes() {
    // console.log(this.answers.entries());
    const times = await averageTimes(this.quiz.id.toString());
    console.log(times);
    if (times === null) {
      error("Nie udało pobrać się średnich czasów na odpowiedź");
      return;
    }
    for (const [i, time] of times.entries()) {
      const div = document.createElement("div");
      div.textContent = `${i + 1}. `;
      const span = document.createElement("span");
      span.textContent = `${
        time ? QuizResults.formatTime(time / 1000) : "--.--"
      }s`;
      div.appendChild(span);
      averageTimesSection.appendChild(div);
    }
  }
  private async displayTopScores() {
    // console.log(this.answers.entries());
    const scores = await topScores(this.quiz.id.toString());
    console.log(scores);
    if (scores === null) {
      error("Nie udało pobrać się najlepszych wyników");
      return;
    }
    for (const [i, [user, result]] of scores.entries()) {
      const div = document.createElement("div");
      const h3 = document.createElement("h3");
      // div.textContent = `${i + 1}. `;
      const timeP = document.createElement("p");
      timeP.textContent = `Czas: ${QuizResults.formatTime(
        result.finalTime / 1000
      )}s`;
      const answersP = document.createElement("p");
      answersP.textContent = `Poprawne odpowiedzi: ${result.answers?.reduce(
        (sum, ans) => (ans.correct ? sum + 1 : sum),
        0
      )} / ${result.answers?.length}`;
      h3.textContent = user.username;
      div.appendChild(h3);
      div.appendChild(timeP);
      div.appendChild(answersP);
      topScoresSection.appendChild(div);
    }
  }

  /**
   * Displays results.
   */
  private async display() {
    this.displayUserResults();
    this.displayCorrectAnswers();
    await this.displayAverageTimes();
    await this.displayTopScores();
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
  public static async displayPreviousResults(): Promise<void> {
    let results;
    try {
      results = await getPrevResults();
    } catch (_) {}
    let i = 1;
    let any = false;
    if (results instanceof Array) {
      for (const res of results) {
        any = true;
        if (isResult(res)) {
          console.log(res);
          const quiz = await getQuizWithAnswers(res.quizId);
          if (quiz) {
            const finalTime =
              res.finalTime +
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

            quizPrevResultsSection.appendChild(row);
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
      quizPrevResultsSection.appendChild(row);
    }
  }
}
