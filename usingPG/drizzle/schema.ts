import { integer, pgTable, varchar, text } from "drizzle-orm/pg-core";

const usersTable = pgTable("users", {
  id: integer().primaryKey().generatedAlwaysAsIdentity(),
  name: text().notNull(),
  email: text().notNull().unique()
});

export default usersTable;
