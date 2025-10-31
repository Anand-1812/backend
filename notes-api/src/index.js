import express from "express";

const PORT = 6969;

const app = express();

app.use("/api/auth", authRouter);

app.listen(PORT, () => {
  console.log(`Server running on port: ${PORT}`)
})
