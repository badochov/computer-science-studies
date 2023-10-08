import { Router } from "express";
import { db } from "../app";

const router = Router();
router.get("/", function (req, res, next) {
  const error = req.query.error;
  if (error) {
    res.locals.error = true;
  }
  const stmt = db.prepare(
    "SELECT login_osoby, tresc FROM wpis ORDER BY timestamp DESC LIMIT 5"
  );
  res.render("index");
});

router.get("/posts", function (req, res, next) {
  const stmt = db.prepare(
    "SELECT login_osoby, tresc FROM wpis ORDER BY timestamp DESC LIMIT 5"
  );
  res.json(stmt.all());
});

export default router;

router.post("/", function (req, res, next) {
  const login = req.body.login;
  const password = req.body.haslo;

  const correctPassword = db
    .prepare("SELECT haslo FROM `osoba` WHERE login = ?")
    .get([login]);
  console.log(req.body, correctPassword);
  if (correctPassword && correctPassword.haslo === password) {
    if (req.session) {
      req.session.user = login;
    }
    return res.redirect("/users");
  } else {
    return res.redirect("/?error=true");
  }
});
// TODO: zaimplementuj logowanie
// logowanie powinno albo przekierować do trasy / albo do trasy /users
