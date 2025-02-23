<template>
  <div class="questions">
    <div class="header">
      <h2>问题列表</h2>
      <NButton type="primary" @click="router.push('/question-editor')">
        新建问题
      </NButton>
    </div>

    <NDataTable :columns="columns" :data="questions" />
  </div>
</template>

<script setup lang="ts">
import { ref, h, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useMessage } from '../utils/message'
import { NButton, NDataTable, NSpace } from 'naive-ui'
import type { Question } from '../types'
import { fetchQuestions, deleteQuestion } from '../api/questions'
import type { DataTableColumns } from 'naive-ui'

const router = useRouter()
const message = useMessage()
const questions = ref<Question[]>([])

const columns: DataTableColumns<Question> = [
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

const loadQuestions = async () => {
  try {
    questions.value = await fetchQuestions()
  } catch (error) {
    message.error('获取问题列表失败')
  }
}

const handleDelete = async (id: number) => {
  try {
    await deleteQuestion(id)
    message.success('删除成功')
    await loadQuestions()
  } catch (error) {
    message.error('删除失败')
  }
}

const handleEdit = (id: number) => {
  router.push(`/question-editor/${id}`)
}

onMounted(loadQuestions)
</script>

<style scoped>
.questions {
  padding: 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}
</style> 