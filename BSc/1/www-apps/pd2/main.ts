import express from "express";
import csurf from "csurf";
import cookieParser from "cookie-parser";
import {
  getQuiz,
  getQuizes,
  didQuiz,
  prevResults,
  getTop5ResultsForQuiz,
  getQuizNoAnswers,
  getAnswerMeanTime,
  getAnswers,
} from "./quizes";
import session from "express-session";
//@ts-ignore
import SQLiteStore from "connect-sqlite3";
import {
  loggedInMiddleware,
  registerUser,
  loginUser,
  changePassword,
  logoutUser,
  User,
} from "./login";
import {
  saveTime,
  isQuizResult,
  validateQuizResult,
  getTime,
  saveResults,
} from "./results";
import bodyParser from "body-parser";

const app = express();
const PORT = 3000;

app.set("view engine", "pug");

app.use(express.urlencoded({ extended: true }));
app;
app.use(cookieParser());
app.use(bodyParser.json());
app.use(csurf({ cookie: true }));
app.use(express.static("views"));
app.use(
  session({
    secret: "f6e42f67904abd4d6ad031b24d697c9e",
    resave: false,
    store: new (SQLiteStore(session))(),
    saveUninitialized: false,
    cookie: { maxAge: 60 * 1000 * 15 },
  })
);

app.use(async (req, res, next) => {
  res.locals["user"] = req.session?.user;
  res.locals["csrfToken"] = req.csrfToken();
  next();
});

export interface SessionRequest extends express.Request {
  session: any;
  sessionId: string;
}

const getUser = (req: SessionRequest): User => {
  return <User>req.session.user;
};

app.get("/", (req, res) => {
  res.render("index", {});
});

app.get("/get_quizes", (req, res) => {
  const quizes = getQuizes();
  quizes.forEach((quiz) => quiz.questions.forEach((q) => delete q.answer));
  res.json(quizes);
});

app.get("/get_quiz/:quiz_id", loggedInMiddleware, (req, res) => {
  const user = getUser(<SessionRequest>req);
  const quizId = req.params.quiz_id;
  if (didQuiz(quizId, user)) {
    return res.status(403).json();
  }
  res.json(getQuizNoAnswers(quizId));
  console.log(getQuizNoAnswers(quizId));
  saveTime(user.id, parseInt(req.params.quiz_id));
});

app.get("/get_quiz_with_answers/:quiz_id", loggedInMiddleware, (req, res) => {
  const user = getUser(<SessionRequest>req);
  const quizId = req.params.quiz_id;
  if (!didQuiz(quizId, user)) {
    return res.status(403).json();
  }
  res.json(getQuiz(quizId));
});
app.get("/get_answers/:quiz_id", loggedInMiddleware, (req, res) => {
  const user = getUser(<SessionRequest>req);
  const quizId = req.params.quiz_id;
  if (!didQuiz(quizId, user)) {
    return res.status(403).json();
  }
  res.json(getAnswers(user.id, quizId));
});

app.get("/results/:quiz_id", loggedInMiddleware, (req, res) => {
  return res.render("results");
});

app.get("/get_answers_mean_time/:quiz_id", (req, res) => {
  return res.json(getAnswerMeanTime(req.params.quiz_id));
});

app.post("/post_results/:quiz_id", loggedInMiddleware, (req, res) => {
  const quizId = parseInt(req.params.quiz_id);
  if (didQuiz(quizId, getUser(<SessionRequest>req))) {
    return res.status(403).json();
  }

  const obj = req.body;
  const quiz = getQuiz(req.params.quiz_id);
  if (
    quiz !== undefined &&
    isQuizResult(obj) &&
    validateQuizResult(obj, quiz)
  ) {
    const user = getUser(<SessionRequest>req);
    const time = getTime(user.id, quizId);
    saveResults(user.id, quizId, obj, time);

    return res.json(time);
  }
  return res.status(422).json();
});

app.get("/can_do_quiz/:quiz_id", loggedInMiddleware, (req, res) => {
  res.json(!didQuiz(req.params.quiz_id, getUser(<SessionRequest>req)));
});

app.get("/top_scores/:quiz_id", (req, res) => {
  res.json(getTop5ResultsForQuiz(req.params.quiz_id));
});

app.get("/prev_results/", loggedInMiddleware, (req, res) => {
  res.json(prevResults(getUser(<SessionRequest>req)));
});

app.post("/login", async (req, res) => {
  const username = req.body.username;
  const password = req.body.password;
  try {
    if (req.body["sign-up"] === "on") {
      if (!(await registerUser(username, password))) {
        return res.status(400).send("Podana nazwa użytkownika jest zajęta");
      }
    }
    if (await loginUser(<SessionRequest>req, username, password)) {
      return res.redirect("back");
    }
    return res
      .status(400)
      .send("Złe hasło lub podany użytkownik nie istnieje!");
  } catch (e) {
    console.error(e);
    return res.status(500).send("Internal server error");
  }
});

app.get("/change_password", async (req, res) => {
  console.log(req.query);
  return res.render("change_password", { error: req.query["error"] });
});

app.post("/change_password", async (req, res) => {
  const username = req.body.username;
  const password = req.body.password;
  const newPassword = req.body["new-password"];
  if (
    await changePassword(<SessionRequest>req, username, password, newPassword)
  ) {
    return res.redirect("/");
  }
  return res.redirect((req.header("Referer") || "/") + "?error=1");
});

app.post("/logout", async (req, res) => {
  logoutUser(<SessionRequest>req);
  res.redirect("back");
});

app.listen(PORT, () =>
  console.log(`Example app listening at http://localhost:${PORT}`)
);
