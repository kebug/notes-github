<template>
  <div class="articles">
    <div class="header">
      <h2>文章列表</h2>
      <NButton type="primary" @click="router.push('/editor')">
        新建文章
      </NButton>
    </div>

    <NDataTable :columns="columns" :data="articles" />
  </div>
</template>

<script setup lang="ts">
import { ref, h, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useMessage } from '../utils/message'
import { NButton, NDataTable, NSpace } from 'naive-ui'
import type { Article } from '../types'
import { fetchArticles, deleteArticle } from '../api/articles'
import type { DataTableColumns } from 'naive-ui'

const router = useRouter()
const message = useMessage()
const articles = ref<Article[]>([])

const columns: DataTableColumns<Article> = [
  {
    title: '标题',
    key: 'title'
  },
  {
    title: '创建时间',
    key: 'created_at',
    render: (row) => {
      return new Date(row.CreatedAt).toLocaleDateString()
    }
  },
  {
    title: '操作',
    key: 'actions',
    render: (row) => {
      return h(NSpace, null, {
        default: () => [
          h(
            NButton,
            { onClick: () => handleEdit(row.id) },
            { default: () => '编辑' }
          ),
          h(
            NButton,
            {
              type: 'error',
              onClick: () => handleDelete(row.id)
            },
            { default: () => '删除' }
          )
        ]
      })
    }
  }
]

const loadArticles = async () => {
  try {
    articles.value = await fetchArticles()
  } catch (error) {
    message.error('获取文章列表失败')
  }
}

const handleDelete = async (id: number) => {
  try {
    await deleteArticle(id)
    message.success('删除成功')
    await loadArticles()
  } catch (error) {
    message.error('删除失败')
  }
}

const handleEdit = (id: number) => {
  router.push(`/editor/${id}`)
}

onMounted(loadArticles)
</script>

<style scoped>
.articles {
  padding: 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}
</style> 