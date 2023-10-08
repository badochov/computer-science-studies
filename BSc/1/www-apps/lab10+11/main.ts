import express from "express";
import csurf from "csurf";
import cookieParser from "cookie-parser";
import { getMostExpensive, getMeme } from "./memes";
import session from "express-session";
//@ts-ignore
import SQLiteStore from "connect-sqlite3";
import { loggedInMiddleware, registerUser, loginUser } from "./login";

const app = express();
const PORT = 3000;

app.set("view engine", "pug");
app.use(express.urlencoded({ extended: true }));
app.use(cookieParser());
app.use(csurf({ cookie: true }));
app.use(express.static("views"));
app.use(
  session({
    secret: "f6e42f67904abd4d6ad031b24d697c9e",
    resave: false,
    saveUninitialized: true,
    store: new (SQLiteStore(session))(),
    cookie: { maxAge: 60 * 1000 * 15 },
  })
);

app.use(async (req, res, next) => {
  res.locals["visit"] = await updateAndGetVisits(<SessionRequest>req);
  res.locals["user"] = req.session?.user;
  res.locals["csrfToken"] = req.csrfToken();
  next();
});

export interface SessionRequest extends express.Request {
  session: any;
  sessionId: string;
}

const updateAndGetVisits = (req: SessionRequest) => {
  if (req.session.visits !== undefined) {
    req.session.visits++;
  } else {
    req.session.visits = 1;
  }
  return req.session.visits;
};

app.get("/", async (req, res) => {
  res.render("index", {
    title: "Meme market",
    message: "Hello there!",
    memes: await getMostExpensive(),
  });
});

app.get("/meme/:memeId", async (req, res) => {
  const meme = await getMeme(req.params.memeId);
  if (meme) {
    res.render("meme", {
      title: "Meme price history",
      meme: meme,
      priceHistory: await meme.getPriceHistory(),
    });
  } else {
    res.status(404).send("Requested page doesn't exists!");
  }
});

app.post("/login", async (req, res) => {
  const username = req.body.username;
  const password = req.body.password;
  try {
    if (req.body["sign-up"] === "on") {
      if (!(await registerUser(username, password))) {
        return res.status(400).send("Given username is taken");
      }
    }
    if (await loginUser(<SessionRequest>req, username, password)) {
      return res.redirect("back");
    }
    return res.status(400).send("Wrong password or username doesn't exist");
  } catch (e) {
    console.error(e);
    return res.status(500).send("Internal server error");
  }
});
app.post("/logout", async (req, res) => {
  delete req.session?.user;
  res.redirect("back");
});

app.post("/meme/:memeId", loggedInMiddleware, async (req, res) => {
  const meme = await getMeme(req.params.memeId);
  const price = req.body.price;
  if (meme) {
    meme.changePrice(price, <number>req.session?.user.id);
  }
  res.redirect(req.url);
});

app.listen(PORT, () =>
  console.log(`Example app listening at http://localhost:${PORT}`)
);
