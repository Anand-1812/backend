import mongoose from "mongoose";

const userSchema = new mongoose.Schema(
  {
    name: {
      type: String,
      required: true,
      trim: true,
      maxlength: 50,
    },

    email: {
      type: String,
      required: true,
      unique: true,
      trim: true,
      index: true,
    },

    password: {
      type: String,
      required: true,
      minlength: 6,
      select: false,
    },
    salt: {
      type: String,
    }
  },
  { timestamps: true }
);

const userSessionSchema = new mongoose.Schema(
  {
    user: {
      type: mongoose.Schema.Types.ObjectId,
      ref: "User", // ✅ reference to User model
      required: true,
      index: true,
    },

    sessionToken: {
      type: String,
      required: true,
      unique: true,
    },

    expiresAt: {
      type: Date,
      required: true,
      index: true,
    },
  },
  { timestamps: true }
);

// Models (avoid overwrite error in Next.js)
export const User =
  mongoose.models.User || mongoose.model("User", userSchema);

export const UserSession =
  mongoose.models.UserSession || mongoose.model("UserSession", userSessionSchema);

