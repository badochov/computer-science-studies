import { QuestionNoAnswer } from "./views/src/quiz";

export interface Question extends QuestionNoAnswer {
  answer: number;
}

export class QuizWithAnswers {
  constructor(
    public id: number,
    public desc: string,
    public questions: Question[] = []
  ) {}

  public addQuestion(q: Question) {
    this.questions.push(q);
  }
}
