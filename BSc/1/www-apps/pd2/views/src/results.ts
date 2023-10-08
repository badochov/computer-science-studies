import { getQuizWithAnswers, getAnswers } from "./quiz";
import { QuizResults } from "./QuizResults";
const main = async () => {
  const quizId = window.location.pathname
    .split("/")
    .filter((a) => a !== "")
    .slice(-1)[0];

  const quiz = await getQuizWithAnswers(quizId);
  const answers = await getAnswers(quizId);
  if (quiz === null || answers === null) {
    alert("Ten test nie istnieje");
    window.location.href = "/";
    return;
  }
  const quizResults = new QuizResults(answers, quiz);
  quizResults.mark();
};
main();
