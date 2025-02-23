import request from '../utils/request'
import type { Article } from '../types'

export const fetchArticles = async () => {
  const { data } = await request.get<Article[]>('/articles')
  return data
}

export const fetchArticle = async (id: number) => {
  const { data } = await request.get<Article>(`/articles/${id}`)
  return data
}

export const createArticle = async (article: Partial<Article>) => {
  const { data } = await request.post<Article>('/articles', article)
  return data
}

export const updateArticle = async (id: number, article: Partial<Article>) => {
  const { data } = await request.put<Article>(`/articles/${id}`, article)
  return data
}

export const deleteArticle = async (id: number) => {
  await request.delete(`/articles/${id}`)
} 