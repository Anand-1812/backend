import { get } from "http";
import db from "./db/index.ts";
import usersTable from "./drizzle/schema.ts";

async function createUsers({ name, email }: { name: string; email: string }) {
  await db.insert(usersTable).values({
    name,
    email
  });
}

async function getUsers() {
  const users = await db.select().from(usersTable);
  console.log(users);
  return users;
}

getUsers();
