import { db } from "../db";

export const cleanAfterTests = () => {
  const id = <number>(
    db.prepare("SELECT id FROM USERS WHERE username = ?").get("test").id
  );
  db.prepare("DELETE FROM answers WHERE user_id = ?").run(id);
  db.prepare("DELETE FROM send_times WHERE user_id = ?").run(id);
};
