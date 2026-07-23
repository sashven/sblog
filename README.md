# Sblog

Sblog is a source controlled personal blog built with Go, Gin, Markdown and Vue.

## Content

Posts live in `content/posts`. (default)

Use filenames in this format:

```text
<preview-name>-MM-DD-YYYY.md
```

Example:

```text
hello-world-07-13-2026.md
```

Each post may include YAML front matter:

```markdown
---
title: Hello World
summary: Your Sblog post.
tags:
  - go
  - vue
---

# Hello World
```
