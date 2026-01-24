import express from "express";
import { Note } from "../models/Notes.model.js";
import requireAuth from "../middlewares/auth.middleware.js";

const router = express.Router();

router.get("/", requireAuth, async (req, res) => {
  try {
    const notes = Note.find({user: req.user._id})
      .sort({createdAt: -1})

    return res.status(200).json(notes);
  } catch (error) {
    console.log(`Error while fetching notes: ${error}`);
    return res.status(500).json({message: "server error while fetching notes."});
  }
})

router.post("/", requireAuth, async (req, res) => {
  try {
    const {title, content, tag, color } = req.body;

    const createNote = await Note.create({
      user: req.user._id,
      title,
      content,
      tag,
      color
    });

    res.status(201).json(createNote);

  } catch (error) {
    console.log(`Error while posting notes = ${error}`)
    return res.status(500).json({message: "server error while posting notes."});
  }
});

router.delete("/:id", requireAuth, async (req, res) => {
  try {

    const note = await Note.findOneAndDelete({_id: req.params.id, user: user.req._id})

    if (!note)
      return res.status(400).json({message: "Note not found."});

    res.json({message: "Note deleted."})

  } catch (error) {
    console.log(`Error while deleting notes = ${error}`);
    return res.status(500).json({message: "Server error."})
  }
})
