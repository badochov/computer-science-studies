import { Database } from "sqlite3";

export interface Price {
  price: number;
  date: Date;
}

export interface MemeDesc {
  id: number;
  name: string;
  price: number;
  url: string;
}

/**
 * Meme class
 */
export class Meme {
  /**
   * @param {number} id
   * @param {string} name
   * @param {string} price
   * @param {string }url
   */
  constructor(
    private id: number,
    private name: string,
    private _price: number,
    private url: string
  ) {}

  /**
   * Returns current price of meme
   */
  public get price(): number {
    return this._price;
  }

  /**
   * Set new price of meme
   * @param {number} price
   */
  private async setPrice(price: number, user: number): Promise<void> {
    this._price = price;
    const db = new Database("./memes.sqlite");
    return new Promise((res, rej) => {
      const insertStatement = db.prepare(
        `INSERT INTO \`memes_prices\` VALUES (?, ?, DATETIME('now', 'localtime'), ?)`,
        (err) => {
          if (err == null) {
            insertStatement.run([price, this.id, user], (err) => {
              if (err == null) {
                insertStatement.finalize((err) => {
                  err == null
                    ? db.close((err) => (err == null ? res() : rej(err)))
                    : rej(err);
                });
              } else {
                rej(err);
              }
            });
          } else {
            rej(err);
          }
        }
      );
    });
  }

  /**
   * Returns description of meme
   */
  public get desc(): MemeDesc {
    return {
      id: this.id,
      name: this.name,
      price: this.price,
      url: this.url,
    };
  }

  /**
   * Sets new price of meme fron request
   * @param {any} rawPrice
   */
  public async changePrice(rawPrice: any, user: number): Promise<void> {
    const price =
      typeof rawPrice === "number"
        ? rawPrice
        : typeof rawPrice === "string"
        ? parseFloat(rawPrice)
        : null;

    if (price !== null && !isNaN(price)) {
      return this.setPrice(price, user);
    }
  }

  /**
   * Returns history of meme prices
   */
  public getPriceHistory(): Promise<Price[]> {
    const db = new Database("./memes.sqlite");
    return new Promise((res, rej) => {
      db.all(
        `SELECT price, dt FROM \`memes_prices\` WHERE meme_id = ${this.id} ORDER BY dt DESC`,
        (err, rows) => {
          if (err == null) {
            const prices: Price[] = rows.map((row) => {
              return { price: parseInt(row.price), date: new Date(row.dt) };
            });

            db.close((err) => (err == null ? res(prices) : rej(err)));
          } else {
            rej(err);
          }
        }
      );
    });
  }
}
