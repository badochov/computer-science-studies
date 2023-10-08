import { Quiz, getQuizes } from './quiz';
import { QuizRun } from './QuizRun';
import { QuizResults } from './QuizResults';
import { quizesDiv, quizChoiseSection } from './HTMLElements';

/**
 * Function that logs errors.
 * May in the future be overwriten to function that display error on screen.
 */
export const error = console.error;

/**
 * Adds quiz to table of available quizes.
 * 
 * @param id quiz id
 * @param quiz quiz object
 */
const addQuizToTable = (quiz: Quiz) => {
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

	quizesDiv.appendChild(row);
};

/**
 * Displays available quizes in the table.
 */
const displayQuizes = () => {
	for (const quiz of getQuizes()) {
		console.log(quiz);
		addQuizToTable(quiz);
	}

	const buttons = document.getElementsByClassName('quiz-start-button') as HTMLCollectionOf<HTMLButtonElement>;
	Array.from(buttons).forEach((button) => {
		button.onclick = () => {
			new QuizRun(button.dataset.quizId);
		};
	});
};

displayQuizes();

QuizResults.displayPreviousResults();
