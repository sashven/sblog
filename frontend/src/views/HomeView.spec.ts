import { mount, flushPromises, RouterLinkStub } from '@vue/test-utils'
import { afterAll, afterEach, beforeAll, describe, expect, it, vi } from 'vitest'

import HomeView from './HomeView.vue'

describe('HomeView', () => {
  beforeAll(() => {
    vi.stubEnv('TZ', 'America/Los_Angeles')
  })

  afterAll(() => {
    vi.unstubAllEnvs()
  })

  afterEach(() => {
    vi.unstubAllGlobals()
  })

  it('renders posts returned by the API', async () => {
    const publishedAt = '2026-07-13T00:00:00Z'
    const fetchMock = vi.fn<typeof fetch>().mockResolvedValue(
      new Response(
        JSON.stringify({
          posts: [
            {
              slug: 'hello-world',
              title: 'Hello World',
              summary: 'First post',
              tags: ['go'],
              published_at: publishedAt,
            },
          ],
        }),
      ),
    )
    vi.stubGlobal('fetch', fetchMock)

    const wrapper = mount(HomeView, {
      global: {
        stubs: {
          RouterLink: RouterLinkStub,
        },
      },
    })
    await flushPromises()

    expect(fetchMock).toHaveBeenCalledWith('/api/v1/posts')
    expect(wrapper.text()).toContain('Hello World')
    expect(wrapper.text()).toContain('First post')
    expect(wrapper.find('time').text()).toBe(
      new Date(publishedAt).toLocaleDateString(undefined, { timeZone: 'UTC' }),
    )
  })
})
