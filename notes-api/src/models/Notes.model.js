import mongoose from "mongoose";

const notesSchema = new mongoose.Schema(
  {
    user: {
      type: mongoose.Schema.Types.ObjectId,
      ref: "User",
      required: true,
    },
    title: {
      type: String,
      required: true,
      trim: true,
      maxlength: 50
    },
    content: {
      type: String,
      required: true,
      trim: true,
      maxlength: 1000
    },
    tag: {
      type: String,
      trim: true,
      default: "General"
    },
    color: {
      type: String,
      default: "border-white/10"
    },
    isPinned: {
      type: Boolean,
      default: false
    }
  },
  { timestamps: true }
);

export const Note = mongoose.model("Note", notesSchema);
