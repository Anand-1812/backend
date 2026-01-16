import express from "express"
import { User } from "../models/User.js";
import { createHmac, randomBytes } from "node:crypto"

const router = express.Router();

router.post("/signup", async (req, res) => {
  try {
    const { name, email, password } = req.body;

    if (!name || !email || !password) {
      return res.status(400).json({ message: "All fields are required" });
    }

    const existingUser = await User.findOne({ email });
    if (existingUser) {
      return res.status(409).json({ message: "Email already exists" })
    }

    const salt = randomBytes(256).toString('hex')
    const hashPass = createHmac('sha256', salt).update(password).digest('hex')

    const user = await User.create({
      name,
      email,
      password: hashPass,
      salt
    })

    return res.status(200).json({ message: "User created successfully" })

  } catch (error) {
    return res.status(500).json({ message: "Server error" });
  }
})

export default router

