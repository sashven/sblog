<script setup lang="ts">
import { onMounted, ref } from 'vue'

import { formatPublishedDate } from '@/date'

type PostSummary = {
  slug: string
  title: string
  summary: string
  tags: string[]
  published_at: string
}

const posts = ref<PostSummary[]>([])
const error = ref('')

onMounted(async () => {
  try {
    const response = await fetch('/api/v1/posts')
    if (!response.ok) {
      throw new Error(`load posts: ${response.status}`)
    }
    const data = (await response.json()) as { posts: PostSummary[] }
    posts.value = data.posts
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'load posts failed'
  }
})
</script>

<template>
  <section class="doc-node post-index">
    <div class="title">
      <h1>Sblog</h1>
      <p>Source controlled blog.</p>
    </div>

    <p v-if="error" class="notice notice--error">{{ error }}</p>
    <ol v-else class="post-list">
      <li v-for="post in posts" :key="post.slug" class="post-list__item">
        <RouterLink class="post-list__title" :to="`/posts/${post.slug}`">{{ post.title }}</RouterLink>
        <p class="post-list__summary">{{ post.summary }}</p>
        <div class="post-list__meta">
          <time :datetime="post.published_at">{{ formatPublishedDate(post.published_at) }}</time>
          <span v-for="tag in post.tags" :key="tag" class="tag">{{ tag }}</span>
        </div>
      </li>
    </ol>
  </section>
</template>
