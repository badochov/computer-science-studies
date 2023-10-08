import { QuizNoAnswers, getQuizes, getQuiz } from "./quiz";
import { QuizRun } from "./QuizRun";
import { QuizResults } from "./QuizResults";
import { quizesDiv } from "./HTMLElements";

/**
 * Adds quiz to table of available quizes.
 *
 * @param id quiz id
 * @param quiz quiz object
 */
const addQuizToTable = (quiz: QuizNoAnswers) => {
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

  quizesDiv.appendChild(row);
};

/**
 * Displays available quizes in the table.
 */
const displayQuizes = async () => {
  for await (const quiz of getQuizes()) {
    addQuizToTable(quiz);
  }

  const buttons = document.getElementsByClassName(
    "quiz-start-button"
  ) as HTMLCollectionOf<HTMLButtonElement>;
  Array.from(buttons).forEach((button) => {
    button.onclick = async () => {
      new QuizRun(await getQuiz(button.dataset.quizId || ""));
    };
  });
};

displayQuizes();

QuizResults.displayPreviousResults();
