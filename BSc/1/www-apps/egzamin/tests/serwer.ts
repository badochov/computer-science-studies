import { unlinkSync } from "fs";
import { exec, execSync } from "child_process";
import { app } from "../app";
const DB_NAME = process.env.DB_NAME || "";

execSync(`npx ts-node -e 'require(\"./app.ts\").createDb()'`);
const serwer = app.listen(3001, () => {
  exec("npx mocha -r ts-node/register tests/tests.ts", (error, stdout) => {
    console.log(stdout);
    unlinkSync(DB_NAME);
    serwer.close();
    process.exit(error?.code);
  });
});
