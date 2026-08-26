// These match, field for field, the JSON your Go backend actually returns — 
// that's intentional; TypeScript's whole value here is catching typos/mismatches against
//  what the API sends.
export interface User {
  id: number
  name: string
  email: string
}

export interface AuthResponse {
  user: User
  token: string
}

export interface LoginPayload {
  email: string
  password: string
}

export interface RegisterPayload {
  name: string
  email: string
  password: string
}