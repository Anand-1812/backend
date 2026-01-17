import express from "express";
import { User, UserSession } from "../models/User.js";
import { createHmac, randomBytes } from "node:crypto";
import { generateSessionToken, generateHashToken } from "../utils/session.js";

const router = express.Router();

router.post("/signup", async (req, res) => {
  try {

    let { name, email, password } = req.body;

    if (!name || !email || !password) {
      return res.status(400).json({ message: "All fields are required" });
    }

    email = email.trim().toLowerCase();

    const existingUser = await User.findOne({ email });
    if (existingUser) {
      return res.status(409).json({ message: "Email already exists" });
    }

    const salt = randomBytes(16).toString("hex");
    const hashPass = createHmac("sha256", salt).update(password).digest("hex");

    const user = await User.create({
      name,
      email,
      password: hashPass,
      salt,
    });

    return res.status(201).json({ message: "User created successfully" });
  } catch (error) {
    console.log("SIGNUP ERROR =>", error);
    return res.status(500).json({ message: error.message });
  }
});

router.post("/login", async (req, res) => {
  try {

    let { email, password } = req.body;

    if (!email || !password) {
      return res.status(400).json({ message: "All fields are required" });
    }

    email = email.trim().toLowerCase();

    const user = await User.findOne({ email }).select("+password +salt");
    if (!user) {
      return res.status(404).json({ message: "Invalid credentails" });
    }

    const newHash = createHmac("sha256", user.salt).update(password).digest("hex");

    if (newHash !== user.password) {
      return res.status(401).json({ message: "Incorrect password" });
    }

    const token = generateSessionToken();
    const hashToken = generateHashToken(token);

    const expires = new Date(Date.now() + 7 * 24 * 60 * 60 * 1000);

    await UserSession.create({
      user: user._id,
      sessionToken: hashToken,
      expiresAt: expires
    })

    // Stores the token in cookie
    res.cookie("session", token, {
      httpOnly: true,
      secure: process.env.NODE_ENV === "production",
      sameSite: "lax",
      expires,
    })

    console.log("Token: ", token)
    return res.status(200).json({ message: "Login successful" });
  } catch (error) {
    console.log("LOGIN ERROR =>", error);
    return res.status(500).json({ message: error.message });
  }
});

export default router;

