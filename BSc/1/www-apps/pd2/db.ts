import Database from "better-sqlite3";
import {
  hashPassword,
  User,
  fillDBWithInitialUsers,
  createUsersTable,
} from "./login";
import { fillDBWithInitialQuizes, createQuizesTable } from "./quizes";
import { createSendTimesTable } from "./results";

export const db = new Database("./quizes.sqlite", { verbose: console.log });

process.on("exit", () => db.close());
process.on("SIGHUP", () => process.exit(128 + 1));
process.on("SIGINT", () => process.exit(128 + 2));
process.on("SIGTERM", () => process.exit(128 + 15));

export const createTables = () => {
  createUsersTable();
  createQuizesTable();
  createSendTimesTable();
};

const fillDBWithInitialData = async () => {
  fillDBWithInitialQuizes();
  await fillDBWithInitialUsers();
};

export const dbSetup = async () => {
  createTables();
  await fillDBWithInitialData();
};
