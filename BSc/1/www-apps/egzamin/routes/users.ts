import express, { Router } from "express";
import { db } from "../app";
import createError from "http-errors";
const onPage = 10;

const router = Router();
router.use((req, res, next) => {
  if (req.session?.user) {
    return next();
  }
  next(createError(401));
});

const isTeacher = (login: string): boolean => {
  const res = db
    .prepare("SELECT nauczyciel FROM `osoba` WHERE login = ?")
    .get([login]);

  return !!res && res.nauczyciel;
};
const isTeacherMiddleware = (
  req: express.Request,
  res: express.Response,
  next: express.NextFunction
) => {
  const user = req.session?.user;
  if (user && isTeacher(user)) {
    return next();
  }
  next(createError(403));
};

router.get("/", function (req, res, next) {
  const username = req.session?.user;
  const page = parseInt(<string>req.query.page || "0");

  const postsCountStmt = db.prepare(
    "SELECT COUNT(*) as count FROM `wpis` WHERE login_osoby IN (SELECT login_sledzonego FROM sledzacy WHERE login_osoby = ?)"
  );
  const count = postsCountStmt.get([username]).count;
  if (page * onPage > count) {
    return next(createError(404));
  }
  const postsStmt = db.prepare(
    "SELECT * FROM `wpis` WHERE login_osoby IN (SELECT login_sledzonego FROM sledzacy WHERE login_osoby = ?) ORDER BY timestamp DESC LIMIT ? OFFSET ?"
  );
  res.locals.posts = postsStmt.all([username, onPage, page * onPage]);
  res.locals.page = page;
  res.locals.last_page = count <= (page + 1) * onPage;

  res.locals.teacher = isTeacher(username);
  res.render("user");
});

router.get("/my_entries", isTeacherMiddleware, function (req, res, next) {
  const username = req.session?.user;
  const page = parseInt(<string>req.query.page || "0");
  const postsStmt = db.prepare(
    "SELECT * FROM `wpis` WHERE login_osoby = ? ORDER BY timestamp DESC LIMIT ? OFFSET ?"
  );
  const postsCountStmt = db.prepare(
    "SELECT COUNT(*) as count FROM `wpis` WHERE login_osoby = ?"
  );

  const count = postsCountStmt.get([username]).count;
  if (page * onPage > count) {
    return next(createError(404));
  }
  res.locals.posts = postsStmt.all([username, onPage, page * onPage]);
  res.locals.last_page = count <= (page + 1) * onPage;
  res.locals.page = page;
  res.locals.teacher = isTeacher(username);
  res.render("my_entries");
});

router.post("/my_entries/new", isTeacherMiddleware, function (req, res) {
  const content = req.body.content;
  const user = req.session?.user;
  db.prepare("INSERT INTO `wpis` VALUES(?, DATETIME('now'), ?)").run([
    user,
    content,
  ]);
  res.redirect("/users/my_entries");
});

router.post("/my_entries/delete/:timestamp", isTeacherMiddleware, function (
  req,
  res
) {
  const timestamp = req.params.timestamp;
  const login = req.session?.user;
  db.prepare("DELETE FROM `wpis` WHERE login_osoby = ? AND timestamp = ?").run([
    login,
    timestamp,
  ]);
  return res.redirect("back");
});

router.post("/wyloguj", function (req, res, next) {
  delete req.session?.user;
  return res.redirect("/");
});

export default router;
