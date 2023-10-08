import { Result } from "./QuizResults";
import { RawQuizResult } from "../../results";
import { csrfToken } from "./util";
import { User } from "../../login";

export interface QuestionNoAnswer {
  prompt: string;
  penalty: number;
}

export interface QuestionWithAnswers {
  prompt: string;
  penalty: number;
  answer: number;
}

export interface QuizNoAnswers {
  desc: string;
  questions: QuestionNoAnswer[];
  id: string;
}

export interface QuizWithAnswers {
  desc: string;
  questions: QuestionWithAnswers[];
  id: string;
}

export interface Answer {
  answer: number | null;
  time: number;
  correct: boolean;
}

const isQuestion = (question: any): question is QuestionNoAnswer => {
  return question && question.prompt && question.answer && question.penalty;
};

export const isQuiz = (quiz: any): quiz is QuizNoAnswers => {
  return quiz.desc && quiz.questions && quiz.id !== undefined;
};
const isQuizWithAnswers = (quiz: any): quiz is QuizWithAnswers => {
  return quiz.desc && quiz.questions && quiz.id !== undefined;
};

/**
 * Returns quiz with given id or null on error.
 * @param id quiz id
 */
export const getQuiz = async (id: string): Promise<QuizNoAnswers | null> => {
  console.log(id);
  try {
    const quizRaw = await fetch("/get_quiz/" + id);
    const quiz = await quizRaw.json();

    if (isQuiz(quiz)) {
      return quiz;
    }
  } catch (_) {}

  return null;
};

/**
 * Returns quiz with given id or null on error.
 * @param id quiz id
 */
export const getQuizWithAnswers = async (
  id: string
): Promise<QuizWithAnswers | null> => {
  try {
    const quizRaw = await fetch("/get_quiz_with_answers/" + id);
    const quiz = await quizRaw.json();

    if (isQuizWithAnswers(quiz)) {
      return quiz;
    }
  } catch (_) {}

  return null;
};

/**
 * Generator that returns all quizes.
 */
export async function* getQuizes(): AsyncGenerator<QuizNoAnswers> {
  const quizRaw = await fetch("/get_quizes");
  const obj = await quizRaw.json();
  if (obj instanceof Array) {
    for (const quiz of obj) {
      if (isQuiz(quiz)) {
        yield quiz;
      }
    }
  }
}

export const getPrevResults = async (): Promise<Result[]> => {
  const results = await fetch("/prev_results");
  const obj = await results.json();
  if (obj instanceof Array) {
    return <Result[]>obj;
  }
  return [];
};

const getRawQuizResults = (answers: Answer[]): RawQuizResult => {
  const totalTime = answers.reduce((sum, ans) => sum + ans.time, 0);
  const qr: RawQuizResult = { answers: [], times: [] };
  answers.forEach((ans) => {
    qr.answers.push(<number>ans.answer);
    qr.times.push(ans.time / totalTime);
  });

  return qr;
};
export const saveResults = async (answers: Answer[], quizId: string) => {
  const qr = getRawQuizResults(answers);
  console.log(qr, answers, JSON.stringify(qr));
  await fetch("/post_results/" + quizId, {
    method: "POST",
    credentials: "same-origin",
    headers: {
      "Content-Type": "application/json",
      "CSRF-Token": csrfToken,
    },
    body: JSON.stringify(qr),
  });
};

export const getAnswers = async (quizId: string): Promise<Answer[] | null> => {
  try {
    const raw = await fetch("/get_answers/" + quizId);
    const obj = await raw.json();
    if (obj instanceof Array) {
      return <Answer[]>obj;
    }
  } catch (_) {}
  return null;
};

export const averageTimes = async (
  quizId: string
): Promise<number[] | null> => {
  try {
    const raw = await fetch("/get_answers_mean_time/" + quizId);
    const obj = await raw.json();
    if (obj instanceof Array) {
      return <number[]>obj;
    }
  } catch (_) {}
  return null;
};
export const topScores = async (
  quizId: string
): Promise<[User, Result][] | null> => {
  try {
    const raw = await fetch("/top_scores/" + quizId);
    const obj = await raw.json();
    if (obj instanceof Array) {
      return <[User, Result][]>obj;
    }
  } catch (_) {}
  return null;
};
