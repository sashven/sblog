import { mount, flushPromises } from '@vue/test-utils'
import { afterAll, afterEach, beforeAll, describe, expect, it, vi } from 'vitest'
import { createMemoryHistory, createRouter } from 'vue-router'

import PostView from './PostView.vue'

describe('PostView', () => {
  beforeAll(() => {
    vi.stubEnv('TZ', 'America/Los_Angeles')
  })

  afterAll(() => {
    vi.unstubAllEnvs()
  })

  afterEach(() => {
    vi.unstubAllGlobals()
  })

  it('renders the post returned by the API', async () => {
    const publishedAt = '2026-07-13T00:00:00Z'
    const fetchMock = vi.fn<typeof fetch>().mockResolvedValue(
      new Response(
        JSON.stringify({
          post: {
            slug: 'hello-world',
            title: 'Hello World',
            summary: 'First post',
            tags: ['go'],
            published_at: publishedAt,
            body_html: '<h1>Hello World</h1><p>From markdown.</p>',
          },
        }),
      ),
    )
    vi.stubGlobal('fetch', fetchMock)

    const router = createRouter({
      history: createMemoryHistory(),
      routes: [{ path: '/posts/:slug', component: PostView }],
    })
    await router.push('/posts/hello-world')
    await router.isReady()

    const wrapper = mount(PostView, {
      global: {
        plugins: [router],
      },
    })
    await flushPromises()

    expect(fetchMock).toHaveBeenCalledWith('/api/v1/posts/hello-world')
    expect(wrapper.text()).toContain('Hello World')
    expect(wrapper.find('time').text()).toBe(
      new Date(publishedAt).toLocaleDateString(undefined, { timeZone: 'UTC' }),
    )
    expect(wrapper.html()).toContain('<h1>Hello World</h1>')
    expect(wrapper.text()).toContain('From markdown.')
  })
})
