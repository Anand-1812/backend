import cors from "cors"
import dotenv from "dotenv";
import express from "express";
import connectDB from "./db/index.js";
import cookieParser from "cookie-parser";

import authRouter from "./routes/user.routes.js";
import notesRouter from "./routes/notes.routes.js";

// dotenv config
dotenv.config({
  path: "./.env"
})

const app = express();
app.use(express.json());
app.use(cookieParser());
app.use(express.urlencoded({ extended: true }));

const allowedOrigins = ["http://localhost:5173"];

app.use(cors({
  origin: function(origin, callback) {
    if (!origin) return callback(null, true);

    if (allowedOrigins.includes(origin)) {
      return callback(null, true);
    }

    return callback(new Error("Not allowed by CORS"));
  },
  credentials: true,
}));



app.get("/api", (_req, res) => {
  res.send("Backend is running")
});

// api endpoints
app.use("/api/auth", authRouter)
app.use("/api/notes", notesRouter)

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
