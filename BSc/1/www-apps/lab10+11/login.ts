import express from "express";
import { Database } from "sqlite3";
import { compare, hash } from "bcrypt";
import { SessionRequest } from "./main";

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
  res.status(403).send("Requested page is only available for logged in users!");
};

export const checkUserPassword = async (
  user: User,
  password: string
): Promise<boolean> => {
  return await compare(password, user.password);
};

interface User {
  id: number;
  username: string;
  password: string;
}

export const logoutUser = (req: SessionRequest) => {
  delete req.session.user;
};

export const getUser = async (username: string): Promise<User> => {
  const db = new Database("./memes.sqlite");
  return new Promise((res, rej) => {
    const statement = db.prepare(
      "SELECT * FROM `users` where username = ?",
      async (err) => {
        if (err !== null) {
          return rej(err);
        }
        statement.get([username], async (err, row) => {
          statement.finalize(async () => {
            db.close(async (errClose) => {
              if (err === null && errClose === null) {
                if (row) {
                  res({
                    username: row.username,
                    password: row.password,
                    id: row.id,
                  });
                }
                rej(new Error("User doesn't exist"));
              } else {
                rej(err || errClose);
              }
            });
          });
        });
      }
    );
  });
};

export const loginUser = async (
  req: SessionRequest,
  username: string,
  password: string
): Promise<boolean> => {
  try {
    const user = await getUser(username);
    if (await checkUserPassword(user, password)) {
      req.session.user = user;
      return true;
    } else return false;
  } catch (_) {
    return false;
  }
};

export const registerUser = async (
  username: string,
  password: string
): Promise<boolean> => {
  try {
    await getUser(username);
    return false;
  } catch (_) {}

  const db = new Database("./memes.sqlite");
  const passwordHash = await hashPassword(password);
  return new Promise((res, rej) => {
    const statement = db.prepare(
      "INSERT INTO `users` (username, password) VALUES (?, ?)",
      async (err) => {
        if (err !== null) {
          return rej(err);
        }
        statement.run([username, passwordHash], async (err) => {
          statement.finalize(async () => {
            db.close(async (errClose) => {
              if (err === null && errClose === null) {
                res(true);
              } else {
                rej(err || errClose);
              }
            });
          });
        });
      }
    );
  });
};
