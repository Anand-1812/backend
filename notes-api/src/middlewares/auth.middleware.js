import { UserSession } from "../models/User.js";
import { generateHashToken } from "../utils/session.js";

const requireAuth = async (req, res, next) => {
  try {

    const token = req.cookies.session;
    if (!token) return res.status(401).json({message: "Not logged in"});

    const tokenHash = generateHashToken(token);
    const session = await UserSession.findOne({
      sessionToken: tokenHash,
      expiresAt: {$gt: new Date()},
    }).populate("user");

    req.user = session.user;
    req.session = session;

    next();

  } catch (error) {
    console.log(`Auth error = ${error}`);
    return res.status(500).json({message: "Auth error"});
  }
}

export default requireAuth;
