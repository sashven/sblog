<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

import { formatPublishedDate } from '@/date'

type PostDetail = {
  slug: string
  title: string
  summary: string
  tags: string[]
  published_at: string
  body_html: string
}

const route = useRoute()
const post = ref<PostDetail | null>(null)
const error = ref('')
const slug = computed(() => String(route.params.slug ?? ''))

onMounted(async () => {
  try {
    const response = await fetch(`/api/v1/posts/${slug.value}`)
    if (!response.ok) {
      throw new Error(`load post: ${response.status}`)
    }
    const data = (await response.json()) as { post: PostDetail }
    post.value = data.post
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'load post failed'
  }
})
</script>

<template>
  <section class="doc-node post-detail">
    <p v-if="error" class="notice notice--error">{{ error }}</p>
    <article v-else-if="post">
      <RouterLink class="back-link" to="/">Back to posts</RouterLink>
      <div class="title">
        <h1>{{ post.title }}</h1>
        <p>{{ post.summary }}</p>
      </div>
      <div class="post-detail__meta">
        <time :datetime="post.published_at">{{ formatPublishedDate(post.published_at) }}</time>
        <span v-for="tag in post.tags" :key="tag" class="tag">{{ tag }}</span>
      </div>
      <div class="markdown-body" v-html="post.body_html"></div>
    </article>
  </section>
</template>
