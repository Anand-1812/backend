import crypto from "node:crypto"

// token for browser cookie
export function generateSessionToken() {
  return crypto.randomBytes(32).toString('hex');
}

// hash token to be stored in db
export function generateHashToken(token) {
  return crypto.createHash('sha256').update(token).digest('hex');
}
