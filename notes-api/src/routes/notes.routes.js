import express from "express";
import { Note } from "../models/Notes.model.js";
import requireAuth from "../middlewares/auth.middleware.js";

const router = express.Router();

router.get("/", requireAuth, async (req, res) => {
  try {
    const notes = await Note.find({user: req.user._id})
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

// UPDATE a note (Added for the Edit functionality)
router.patch("/:id", requireAuth, async (req, res) => {
  try {
    const { title, content, tag, color } = req.body;
    const note = await Note.findOneAndUpdate(
      { _id: req.params.id, user: req.user._id },
      { title, content, tag, color },
      { new: true }
    );

    if (!note) return res.status(404).json({ message: "Note not found." });
    res.json(note);
  } catch (error) {
    return res.status(500).json({ message: "Server error while updating." });
  }
});

router.delete("/:id", requireAuth, async (req, res) => {
  try {

    const note = await Note.findOneAndDelete({_id: req.params.id, user: req.user._id})

    if (!note)
      return res.status(400).json({message: "Note not found."});

    res.json({message: "Note deleted."})

  } catch (error) {
    console.log(`Error while deleting notes = ${error}`);
    return res.status(500).json({message: "Server error."})
  }
});

// PATCH /api/notes/:id/pin - Toggle pin status
router.patch("/:id/pin", requireAuth, async (req, res) => {
  try {
    const note = await Note.findOne({ _id: req.params.id, user: req.user._id });
    if (!note) return res.status(404).json({ message: "Note not found" });

    note.isPinned = !note.isPinned;
    await note.save();

    res.json(note);
  } catch (error) {
    res.status(500).json({ message: "Server error" });
  }
});

export default router
