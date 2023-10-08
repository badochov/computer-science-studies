import { QuizWithAnswers, Question } from "./Quiz";

import { db } from "./db";
import { User, getUsers } from "./login";
import type { Result } from "./views/src/QuizResults";
import type { Answer } from "./views/src/quiz";
import type { QuestionNoAnswer } from "./views/src/quiz";

export interface UserNoPassword {
  username: string;
  id: number;
}
export interface QuizNoAnswers {
  id: number;
  desc: string;
  questions: QuestionNoAnswer[];
}

export interface QuizesRow {
  id: number;
  desc: string;
}

export interface QuestionsRow {
  quiz_id: number;
  questions_number: number;
  prompt: string;
  answer: number;
  penalty: number;
}

export interface AnswersRow {
  quiz_id: number;
  question_number: number;
  time: number;
  answer: number;
}

export const getAnswersRow = (
  userId: number,
  rawQuizId: string | number
): AnswersRow[] => {
  return <AnswersRow[]>(
    db
      .prepare("SELECT * FROM answers WHERE user_id = ? AND quiz_id = ?")
      .all([userId, rawQuizId])
  );
};

export const getQuizResults = (
  userId: number,
  rawQuizId: string | number
): Result | null => {
  const answers = <AnswersRow[]>(
    db
      .prepare("SELECT * FROM answers WHERE user_id = ? AND quiz_id = ?")
      .all([userId, rawQuizId])
  );
  if (answers.length === 0) {
    return null;
  }

  const quiz = getQuiz(rawQuizId);

  const numberAnswers: Answer[] = answers.map((answer) => {
    return {
      answer: answer.answer,
      time: answer.time,
      correct: quiz?.questions[answer.question_number].answer === answer.answer,
    };
  });
  return {
    quizId: rawQuizId.toString(),
    answers: numberAnswers,
    finalTime: numberAnswers.reduce(
      (sum, answer, i) =>
        sum +
        answer.time +
        (answer.correct ? 0 : quiz?.questions[i].penalty || 0),
      0
    ),
  };
};

export const getQuizes = (): QuizWithAnswers[] => {
  const quizes: QuizWithAnswers[] = [];

  const quizBaseData = <QuizesRow[]>db.prepare("SELECT * FROM quizes").all();
  for (const quizBase of quizBaseData) {
    const quiz = new QuizWithAnswers(quizBase.id, quizBase.desc);
    const questions = <Question[]>(
      db
        .prepare(
          "SELECT * FROM questions WHERE quiz_id = ? ORDER BY question_number"
        )
        .all(quizBase.id)
    );
    questions.map(quiz.addQuestion.bind(quiz));

    quizes.push(quiz);
  }
  return quizes;
};

export const didQuiz = (rawQuizId: string | number, user: User): boolean => {
  const count = db
    .prepare(
      "SELECT COUNT(*) as cnt FROM answers WHERE quiz_id = ? and user_id = ?"
    )
    .get([
      typeof rawQuizId === "number" ? rawQuizId : parseInt(rawQuizId),
      user.id,
    ]);
  return count.cnt !== 0;
};

export const getQuiz = (
  rawId: string | number
): QuizWithAnswers | undefined => {
  const id = typeof rawId === "number" ? rawId : parseInt(rawId);

  return getQuizes().find((quiz) => quiz.id === id);
};

export const prevResults = (user: User): Result[] => {
  const results: Result[] = [];
  for (const quiz of getQuizes()) {
    const result = getQuizResults(user.id, quiz.id);
    if (result !== null) {
      results.push(result);
    }
  }
  return results;
};

export const getAllQuizResults = (
  rawQuizId: number | string
): Array<[User, Result]> => {
  const results: Array<[User, Result]> = [];

  for (const user of getUsers()) {
    const result = getQuizResults(user.id, rawQuizId);
    if (result !== null) {
      results.push([user, result]);
    }
  }
  return results;
};

export const getTop5ResultsForQuiz = (
  rawId: string | number
): Array<[UserNoPassword, Result]> => {
  const results = getAllQuizResults(rawId);
  return results
    .sort((a, b) => a[1].finalTime - b[1].finalTime)
    .map((a) => {
      delete a[0].password;
      return a;
    })
    .slice(0, 5);
};

export const getQuizNoAnswers = (
  quizRawId: string | number
): QuizNoAnswers | undefined => {
  const quiz = getQuiz(quizRawId);
  if (quiz === undefined) {
    return quiz;
  }
  quiz.questions.forEach((q) => delete q.answer);
  return quiz;
};

export const getAnswerMeanTime = (
  quizId: string | number
): number[] | undefined => {
  const quiz = getQuiz(quizId);
  if (quiz === undefined) {
    return undefined;
  }
  const results = getAllQuizResults(quizId);
  const times: number[] = Array(quiz.questions.length).fill(0);
  results.forEach(([_, result]) => {
    if (result.answers) {
      result.answers.forEach((answer, i) => (times[i] += answer.time));
    }
  });
  return times.map((sum) => sum / (results.length || 1));
};

export const getAnswers = (
  userId: number,
  quizId: string | number
): Answer[] => {
  const rawAnswers = getAnswersRow(userId, quizId);
  const answers: Answer[] = [];
  rawAnswers.forEach(
    (rawAns) =>
      (answers[rawAns.question_number] = {
        answer: rawAns.answer,
        time: rawAns.time,
        correct: false,
      })
  );

  return answers;
};

const allQuestions = [
  [
    {
      prompt: "2 + 3",
      answer: 5,
      penalty: 4,
    },
    {
      prompt: "2 - (-24 : 4)",
      answer: 8,
      penalty: 10,
    },
    {
      prompt: "2 - 3",
      answer: -1,
      penalty: 4,
    },
    {
      prompt: "2 + (-24 : 4)",
      answer: -4,
      penalty: 10,
    },
  ],
  [
    {
      prompt: "2 + 2",
      answer: 4,
      penalty: 2,
    },
    {
      prompt: "2 * 3",
      answer: 6,
      penalty: 1,
    },
    {
      prompt: "2 ^ 3",
      answer: 8,
      penalty: 3,
    },
    {
      prompt: "(2 ^ (1/2)) ^ 2",
      answer: 2,
      penalty: 7,
    },
  ],
];

const fillDBWithInitialQuestions = () => {
  const addQuestionStatement = db.prepare(
    "INSERT INTO `questions` VALUES (?, ?, ?, ?, ?);"
  );
  for (const [i, questions] of allQuestions.entries()) {
    for (const [questionNumber, question] of questions.entries()) {
      addQuestionStatement.run([
        i + 1,
        questionNumber,
        question.prompt,
        question.answer,
        question.penalty,
      ]);
    }
  }
};

export const fillDBWithInitialQuizes = () => {
  const addQuizStatement = db.prepare(
    "INSERT INTO `quizes` (desc) VALUES (?);"
  );
  for (const memeData of ["Liczyć każdy może", "Quick math"]) {
    addQuizStatement.run(memeData);
  }

  fillDBWithInitialQuestions();
};

export const createQuizesTable = () => {
  db.exec(
    "DROP TABLE IF EXISTS quizes;" +
      "DROP TABLE IF EXISTS questions;" +
      "DROP TABLE IF EXISTS answers;"
  );
  db.exec(
    "CREATE TABLE IF NOT EXISTS quizes (id INTEGER PRIMARY KEY AUTOINCREMENT, desc TEXT);" +
      "CREATE TABLE IF NOT EXISTS questions (quiz_id INTEGER, question_number INTEGER, prompt TEXT, answer INTEGER, penalty INTEGER);" +
      "CREATE TABLE IF NOT EXISTS answers (quiz_id INTEGER, question_number INTEGER, time INTEGER, answer INTEGER, user_id INTEGER);"
  );
};
