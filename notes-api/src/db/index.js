import mongoose from "mongoose";

const connectDB = async () => {
  try {
    await mongoose.connect(process.env.MONGO_URL, {
      tls: true,
    });

  } catch (error) {
    console.error(`Error on db connection: ${error}`);
    process.exit();
  }
}

export default connectDB;
