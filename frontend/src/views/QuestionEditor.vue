<template>
  <div class="editor-container">
    <div class="editor-header">
      <input
        v-model="title"
        type="text"
        class="title-input"
        placeholder="请输入问题标题"
      />
      <div class="actions">
        <button @click="saveQuestion" class="save-btn">保存</button>
        <button @click="router.push('/questions')" class="cancel-btn">取消</button>
      </div>
    </div>

    <div class="editor-main">
      <div class="editor-section">
        <div class="section-header">问题描述</div>
        <textarea
          v-model="content"
          class="content-input"
          placeholder="请输入问题描述 (Markdown格式)"
        ></textarea>
        <div class="section-header">解答</div>
        <textarea
          v-model="answer"
          class="content-input"
          placeholder="请输入解答 (Markdown格式)"
        ></textarea>
      </div>
      <div class="preview-section">
        <div class="section-header">预览</div>
        <div class="preview-content">
          <h3>问题描述</h3>
          <div class="markdown-body" v-html="contentHtml"></div>
          <h3 class="answer-header">解答</h3>
          <div class="markdown-body" v-html="answerHtml"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '../stores/user'
import { marked } from 'marked'
import axios from 'axios'
import 'highlight.js/styles/github.css'
import hljs from 'highlight.js'

// 配置marked
marked.setOptions({
  highlight: function (code, lang) {
    if (lang && hljs.getLanguage(lang)) {
      return hljs.highlight(code, { language: lang }).value
    }
    return hljs.highlightAuto(code).value
  },
  breaks: true
})

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

const title = ref('')
const content = ref('')
const answer = ref('')
const isEditing = ref(false)
const questionId = ref<number | null>(null)

const contentHtml = computed(() => {
  return marked(content.value || '')
})

const answerHtml = computed(() => {
  return marked(answer.value || '')
})

const loadQuestion = async (id: string) => {
  try {
    const response = await axios.get(`http://localhost:8080/api/questions/${id}`, {
      headers: {
        Authorization: `Bearer ${userStore.token}`
      }
    })
    const question = response.data
    title.value = question.title
    content.value = question.content
    answer.value = question.answer
    questionId.value = question.id
    isEditing.value = true
  } catch (error) {
    console.error('加载问题失败:', error)
    router.push('/questions')
  }
}

const saveQuestion = async () => {
  if (!title.value.trim() || !content.value.trim()) {
    alert('标题和问题描述不能为空')
    return
  }

  const question = {
    title: title.value,
    content: content.value,
    answer: answer.value
  }

  try {
    if (isEditing.value && questionId.value) {
      await axios.put(
        `http://localhost:8080/api/questions/${questionId.value}`,
        question,
        {
          headers: {
            Authorization: `Bearer ${userStore.token}`
          }
        }
      )
    } else {
      await axios.post('http://localhost:8080/api/questions', question, {
        headers: {
          Authorization: `Bearer ${userStore.token}`
        }
      })
    }
    router.push('/questions')
  } catch (error) {
    console.error('保存问题失败:', error)
    alert('保存失败，请重试')
  }
}

onMounted(() => {
  const id = route.query.id as string
  if (id) {
    loadQuestion(id)
  }
})
</script>

<style scoped>
.editor-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.title-input {
  font-size: 1.5em;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  width: 60%;
}

.actions button {
  margin-left: 10px;
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.save-btn {
  background-color: #42b983;
  color: white;
}

.cancel-btn {
  background-color: #666;
  color: white;
}

.editor-main {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  height: calc(100vh - 150px);
}

.editor-section,
.preview-section {
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 10px;
  overflow-y: auto;
}

.section-header {
  font-size: 1.2em;
  font-weight: bold;
  margin: 10px 0;
  padding-bottom: 5px;
  border-bottom: 1px solid #eee;
}

.content-input {
  width: 100%;
  height: calc(50% - 60px);
  border: 1px solid #ddd;
  border-radius: 4px;
  resize: none;
  font-family: monospace;
  font-size: 14px;
  line-height: 1.6;
  padding: 10px;
  margin-bottom: 20px;
}

.content-input:focus {
  outline: none;
  border-color: #42b983;
}

.preview-content {
  padding: 10px;
}

.answer-header {
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #eee;
}

/* Markdown 预览样式 */
.markdown-body {
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
    sans-serif;
  font-size: 16px;
  line-height: 1.6;
  padding: 10px;
}

.markdown-body h1,
.markdown-body h2 {
  border-bottom: 1px solid #eaecef;
  padding-bottom: 0.3em;
}

.markdown-body pre {
  background-color: #f6f8fa;
  border-radius: 6px;
  padding: 16px;
  overflow: auto;
}

.markdown-body code {
  background-color: rgba(27, 31, 35, 0.05);
  border-radius: 3px;
  font-size: 85%;
  padding: 0.2em 0.4em;
}

.markdown-body pre code {
  background-color: transparent;
  padding: 0;
}
</style> 