import { quiz } from "./quizes";

export interface Question {
  prompt: string;
  answer: number;
  penalty: number;
}

export interface Quiz {
  desc: string;
  questions: Question[];
  id: string;
}

export interface Answer {
  answer: number | null;
  time: number;
  correct: boolean;
}

const isQuestion = (question: any): question is Question => {
  return question && question.prompt && question.answer && question.penalty;
};

const isQuiz = (quiz: any): quiz is Quiz => {
  return quiz.desc && quiz.questions && quiz.id !== undefined;
};

/**
 * Returns quiz with given id or null on error.
 * @param id quiz id
 */
export const getQuiz = (id: string): Quiz | null => {
  const arr = JSON.parse(quiz || "{}")?.quizes;
  if (arr instanceof Array) {
    const quiz = arr.find((el: any) => el.id === id);

    if (isQuiz(quiz)) {
      return quiz;
    }
  }
  return null;
};

/**
 * Generator that returns all quizes.
 */
export function* getQuizes(): Generator<Quiz> {
  const obj = JSON.parse(quiz || "{}");
  if (obj.quizes instanceof Array) {
    for (const quiz of obj.quizes) {
      if (isQuiz(quiz)) {
        yield quiz;
      }
    }
  }
}
