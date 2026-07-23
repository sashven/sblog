import { describe, it, expect } from 'vitest'

import { mount } from '@vue/test-utils'
import { createMemoryHistory, createRouter } from 'vue-router'
import App from '../App.vue'

describe('App', () => {
  it('renders the documentation-style shell around the active route', async () => {
    const router = createRouter({
      history: createMemoryHistory(),
      routes: [
        {
          path: '/',
          component: { template: '<main>Route content</main>' },
        },
      ],
    })
    await router.push('/')
    await router.isReady()

    const wrapper = mount(App, {
      global: {
        plugins: [router],
      },
    })
    expect(wrapper.find('.app-shell').exists()).toBe(true)
    expect(wrapper.find('.site-nav').exists()).toBe(true)
    expect(wrapper.find('.doc-scrollview').exists()).toBe(true)
    expect(wrapper.find('.doc-content').exists()).toBe(true)
    expect(wrapper.text()).toContain('Route content')
  })
})
