import { hashToken } from "../utils/session";
import { UserSession } from "../models/User";

export default async function requireAuth(req, res, next) {
  try {

    const token = req.cookies.session;
    if (!token) return res.status(401).json({message: "Not logged in"});

    const tokenHash = hashToken(token);
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

