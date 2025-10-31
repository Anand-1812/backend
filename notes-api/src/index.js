import dotenv from "dotenv";
import express from "express";
import connectDB from "./db/index.js";

// dotenv config
dotenv.config({
  path: "./.env"
})

const app = express();
app.use(express.json());

app.get("/", (req, res) => {
  res.send("Backend is running")
});

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
