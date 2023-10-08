import type { MemeDesc } from "./Meme"; // eslint-disable-line
import { Meme } from "./Meme";
import { verbose, Database } from "sqlite3";
import { hash } from "bcrypt";
import { hashPassword } from "./login";
export const getMostExpensive = async (
  amount: number = 3
): Promise<MemeDesc[]> => {
  return (await getMemes())
    .sort((a, b) => a.price - b.price)
    .slice(0, amount)
    .map((meme) => meme.desc);
};
verbose();

export const getMemes = async (): Promise<Meme[]> => {
  const db = new Database("./memes.sqlite");
  return new Promise((res, rej) => {
    db.all(
      "SELECT (SELECT price \
        FROM   `memes_prices` \
        WHERE  meme_id = id \
        ORDER  BY dt DESC \
        LIMIT  1) AS price, \
        * \
      FROM   `memes`; ",
      (err, rows) => {
        if (err) {
          db.close(() => rej(err));
        } else {
          const memes: Meme[] = [];
          rows.forEach((row) => {
            memes.push(new Meme(row.id, row.name, row.price, row.url));
          });
          db.close(() => res(memes));
        }
      }
    );

    db.close();
  });
};

export const getMeme = async (rawId: string): Promise<Meme | undefined> => {
  const id = parseInt(rawId);

  const memes = await getMemes();

  return memes.find((meme) => meme.desc.id === id);
};

// może czasem nie zadziałać, bo ta biblioteka to callback hell, ale powinno działać w większości przypadków
export const dbSetup = () => {
  const db = new Database("./memes.sqlite");

  db.exec(
    "CREATE TABLE IF NOT EXISTS memes (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, url TEXT);" +
      "CREATE TABLE IF NOT EXISTS memes_prices (price INTEGER, meme_id INTEGER, dt VARCHAR(20), user_id INTEGER);" +
      "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT);",
    async (err) => {
      const addMemeStatement = db.prepare(
        "INSERT INTO `memes` (name, url) VALUES (?, ?);"
      );
      for (const memeData of [
        ["Gold", "https://i.redd.it/h7rplf9jt8y21.png"],
        [
          "Platinum",
          "http://www.quickmeme.com/img/90/90d3d6f6d527a64001b79f4e13bc61912842d4a5876d17c1f011ee519d69b469.jpg",
        ],
        ["Elite", "https://i.imgflip.com/30zz5g.jpg"],
      ]) {
        addMemeStatement.run(memeData);
      }
      addMemeStatement.finalize();

      const addMemePriceStatement = db.prepare(
        "INSERT INTO `memes_prices` VALUES (?, ?, DATETIME('now', 'localtime'), ?);"
      );
      for (const [i, price] of [1000, 1100, 1200].entries()) {
        addMemePriceStatement.run([price, i + 1, 1]);
      }
      addMemePriceStatement.finalize();

      const addUserStatement = db.prepare(
        "INSERT INTO `users` (username, password) VALUES (?, ?);"
      );
      const password = await hashPassword("password");
      addUserStatement.run(["admin", password]);
      addUserStatement.finalize();

      db.close();
    }
  );
};
