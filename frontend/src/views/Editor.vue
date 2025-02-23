<template>
  <div class="editor-container">
    <div class="editor-header">
      <input
        v-model="title"
        type="text"
        class="title-input"
        placeholder="请输入标题"
      />
      <div class="actions">
        <button @click="saveArticle" class="save-btn">保存</button>
        <button @click="router.push('/articles')" class="cancel-btn">取消</button>
      </div>
    </div>

    <div class="editor-main">
      <div class="editor-section">
        <textarea
          v-model="content"
          class="content-input"
          placeholder="请输入文章内容 (Markdown格式)"
        ></textarea>
      </div>
      <div class="preview-section">
        <div class="markdown-body" v-html="htmlContent"></div>
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
const isEditing = ref(false)
const articleId = ref<number | null>(null)

const htmlContent = computed(() => {
  return marked(content.value)
})

const loadArticle = async (id: string) => {
  try {
    const response = await axios.get(`http://localhost:8080/api/articles/${id}`, {
      headers: {
        Authorization: `Bearer ${userStore.token}`
      }
    })
    const article = response.data
    title.value = article.title
    content.value = article.content
    articleId.value = article.id
    isEditing.value = true
  } catch (error) {
    console.error('加载文章失败:', error)
    router.push('/articles')
  }
}

const saveArticle = async () => {
  if (!title.value.trim() || !content.value.trim()) {
    alert('标题和内容不能为空')
    return
  }

  const article = {
    title: title.value,
    content: content.value
  }

  try {
    if (isEditing.value && articleId.value) {
      await axios.put(
        `http://localhost:8080/api/articles/${articleId.value}`,
        article,
        {
          headers: {
            Authorization: `Bearer ${userStore.token}`
          }
        }
      )
    } else {
      await axios.post('http://localhost:8080/api/articles', article, {
        headers: {
          Authorization: `Bearer ${userStore.token}`
        }
      })
    }
    router.push('/articles')
  } catch (error) {
    console.error('保存文章失败:', error)
    alert('保存失败，请重试')
  }
}

onMounted(() => {
  const id = route.query.id as string
  if (id) {
    loadArticle(id)
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

.content-input {
  width: 100%;
  height: 100%;
  border: none;
  resize: none;
  font-family: monospace;
  font-size: 14px;
  line-height: 1.6;
  padding: 10px;
}

.content-input:focus {
  outline: none;
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