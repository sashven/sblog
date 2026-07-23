<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="app-shell">
    <aside class="site-nav" aria-label="Primary navigation">
      <div class="site-nav__head">
        <RouterLink class="site-brand" to="/">
          <!--<span class="site-brand__mark">S</span>-->
          <span>
            <strong>Sblog</strong>
            <small>Small blog :K</small>
          </span>
        </RouterLink>
      </div>

      <nav class="site-nav__links">
        <RouterLink to="/">Posts</RouterLink>
      </nav>
    </aside>

    <main id="main-content" class="doc-scrollview" tabindex="-1">
      <div class="doc-container">
        <div class="doc-content">
          <RouterView />
        </div>
      </div>
    </main>
  </div>
</template>
