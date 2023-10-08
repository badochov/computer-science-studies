import express from "express";
import { compare, hash } from "bcrypt";
import { SessionRequest } from "./main";
import { db } from "./db";

export const hashPassword = async (password: string): Promise<string> => {
  try {
    return await hash(password, 3);
  } catch (_) {
    return "";
  }
};

export const loggedInMiddleware = (
  req: express.Request,
  res: express.Response,
  next: Function
): void => {
  if (req.session?.user) {
    return next();
  }
  res.status(403).send("Strona dostępna tylko dla zalogowanych użytkowników!");
};

export const checkUserPassword = async (
  user: User,
  password: string
): Promise<boolean> => {
  return await compare(password, user.password);
};

export interface User extends UserNoPassword {
  password: string;
}
export interface UserNoPassword {
  id: number;
  username: string;
}

export const logoutUser = (req: SessionRequest) => {
  delete req.session.user;
};

export const getUsers = (): User[] => {
  return <User[]>db.prepare("SELECT * FROM `users`").all();
};

export class NoUserError extends Error {}

export const getUser = (username: string): User => {
  const user = db
    .prepare("SELECT * FROM `users` where username = ?")
    .get(username);
  if (user !== undefined) return <User>user;
  throw new NoUserError();
};

export const loginUser = async (
  req: SessionRequest,
  username: string,
  password: string
): Promise<boolean> => {
  try {
    const user = getUser(username);
    if (await checkUserPassword(user, password)) {
      req.session.user = user;
      return true;
    } else return false;
  } catch (_) {
    return false;
  }
};

export const changePassword = async (
  req: SessionRequest,
  username: string,
  password: string,
  newPassword: string
): Promise<boolean> => {
  try {
    const user = getUser(username);
    console.log(user);
    console.log(password, newPassword);
    if (await checkUserPassword(user, password)) {
      db.prepare("UPDATE `users` SET password = ? WHERE username = ?").run([
        await hashPassword(newPassword),
        username,
      ]);
      console.log("fadfadsfa");
      logoutUser(req);
      return true;
    }
  } catch (_) {}
  return false;
};

export const registerUser = async (
  username: string,
  password: string
): Promise<boolean> => {
  try {
    getUser(username);
    return false;
  } catch (_) {}

  const passwordHash = await hashPassword(password);

  db.prepare("INSERT INTO `users` (username, password) VALUES (?, ?)").run([
    username,
    passwordHash,
  ]);

  return true;
};

export const fillDBWithInitialUsers = async () => {
  const addUserStatement = db.prepare(
    "INSERT INTO `users` (username, password) VALUES (?, ?);"
  );

  const password1 = await hashPassword("user1");
  const password2 = await hashPassword("user2");
  const password3 = await hashPassword("test");

  addUserStatement.run(["user1", password1]);
  addUserStatement.run(["user2", password2]);
  addUserStatement.run(["test", password3]);
};

export const createUsersTable = () => {
  db.exec(
    "DROP TABLE IF EXISTS users;" +
      "CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT);"
  );
};
