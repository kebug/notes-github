import request from '../utils/request'
import type { Question } from '../types'

export const fetchQuestions = async () => {
  const { data } = await request.get<Question[]>('/questions')
  return data
}

export const fetchQuestion = async (id: number) => {
  const { data } = await request.get<Question>(`/questions/${id}`)
  return data
}

export const createQuestion = async (question: Partial<Question>) => {
  const { data } = await request.post<Question>('/questions', question)
  return data
}

export const updateQuestion = async (id: number, question: Partial<Question>) => {
  const { data } = await request.put<Question>(`/questions/${id}`, question)
  return data
}

export const deleteQuestion = async (id: number) => {
  await request.delete(`/questions/${id}`)
} 