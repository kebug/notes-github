export interface Article {
  id: number
  title: string
  content: string
  created_at: string
  updated_at: string
  user_id: number
  git_path: string
}

export interface Question {
  id: number
  title: string
  content: string
  answer: string
  created_at: string
  updated_at: string
  user_id: number
  issue_url: string
}

export interface User {
  id: number
  username: string
  github_id: number
  avatar_url: string
  created_at: string
  updated_at: string
} 