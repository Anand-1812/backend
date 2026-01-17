import dotenv from "dotenv";
import express from "express";
import connectDB from "./db/index.js";
import cookieParser from "cookie-parser";

import authRouter from "./routes/user.routes.js";

// dotenv config
dotenv.config({
  path: "./.env"
})

const app = express();
app.use(express.json());
app.use(cookieParser());

app.get("/api", (_req, res) => {
  res.send("Backend is running")
});

app.use("/api/auth", authRouter)

const PORT = process.env.PORT || 6969;

connectDB()
  .then(() => {
    console.log("DB connection success");
    app.listen(PORT, () => {
      console.log(`server running on port: ${PORT}`);
    })
  })
  .catch((err) => {
    console.log(`DB connection failed: ${err}`);
  });
