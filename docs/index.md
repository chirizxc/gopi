---
layout: home

hero:
  name: "GoPI"
  text: "API для работы с GIF"
  image:
    src: /logo.png
    style: "margin-top: 40px;
            filter: drop-shadow(0 0 50px rgba(255, 202, 202, 0.7));
            animation: glow 1s infinite alternate;"
  actions:
    - theme: brand
      text: Примеры работы с API
      link: /api-examples
---

<script setup>
import { VPTeamMembers } from 'vitepress/theme'

const members = [
  {
    avatar: 'https://www.github.com/chirizxc.png',
    name: 'chirizxc',
    links: [
      { icon: 'github', link: 'https://github.com/chirizxc' },
      { icon: 'telegram', link: 'https://t.me/autistic_kids' }
    ]
  },
  {
    avatar: 'https://www.github.com/FriedCerebrum.png',
    name: 'FriedCerebrum',
    links: [
      { icon: 'github', link: 'https://github.com/FriedCerebrum' },
    ]
  },
  {
    avatar: 'https://www.github.com/ChrisElli-dev.png',
    name: 'Christopher Elliot',
    links: [
      { icon: 'github', link: 'https://github.com/ChrisElli-dev' },
    ]
  },
  {
    avatar: 'https://www.github.com/Memory420.png',
    name: 'Memory420',
    links: [
      { icon: 'github', link: 'https://github.com/Memory420' },
    ]
  },
]
</script>

<h2 class="center-heading">Team</h2>

<VPTeamMembers size="small" :members="members"></VPTeamMembers>

<style scoped>
.center-heading {
  text-align: center;
  margin-bottom: 2rem; /* Add space between the heading and team members */
}
</style>
