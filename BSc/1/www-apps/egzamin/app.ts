import createError from "http-errors";
import express from "express";
import * as path from "path";
import cookieParser from "cookie-parser";
import logger from "morgan";
import Database from "better-sqlite3";
import session from "express-session";
import indexRouter from "./routes/index";
import usersRouter from "./routes/users";
import { readFile } from "fs";
import csurf from "csurf";
export const db = new Database(process.env.DB_NAME || "./baza.db", {
  verbose: console.log,
});

export const app = express();
const PORT = process.env.PORT || 3000;
// view engine setup
app.set("views", path.join(__dirname, "views"));
app.set("view engine", "pug");

app.use(
  session({
    secret: "f6e42f67904abd4d6ad031b24d697c9e",
    resave: false,
    saveUninitialized: false,
  })
);
app.use(logger("dev"));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, "public")));
app.use(csurf({ cookie: true }));
app.use((req, res, next) => {
  res.locals["user"] = req.session?.user;
  res.locals["csrfToken"] = req.csrfToken();
  next();
});

app.use("/", indexRouter);
app.use("/users", usersRouter);

// catch 404 and forward to error handler
app.use(function (req, res, next) {
  next(createError(404));
});

interface ResponseError extends Error {
  status?: number;
}

// error handler
app.use(function (
  err: ResponseError,
  req: express.Request,
  res: express.Response,
  next: express.NextFunction
) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get("env") === "development" ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render("error");
});
if (require.main === module) {
  app.listen(PORT, () =>
    console.log(`Example app listening at http://localhost:${PORT}`)
  );
}

export default app;

export const createDb = () => {
  readFile("./data.sql", "utf-8", (err, data) => {
    if (err) {
      return console.error(err);
    }
    db.exec(data);
  });
};
