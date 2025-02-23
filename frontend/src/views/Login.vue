<template>
  <div class="login-container">
    <h2>登录</h2>
    <button @click="handleGithubLogin" class="github-login-btn">
      使用GitHub登录
    </button>
  </div>
</template>

<script setup lang="ts">
const handleGithubLogin = () => {
  window.location.href = 'http://localhost:8080/auth/github/login'
}
</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 40px auto;
  padding: 20px;
  text-align: center;
}

.github-login-btn {
  background-color: #24292e;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

.github-login-btn:hover {
  background-color: #2f363d;
}
</style> 