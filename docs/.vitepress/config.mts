import { defineConfig } from 'vitepress'

export default defineConfig({
  title: "GoPI",
  description: "REST API для работы с GIF.",
  base: "/gopi/",
  head: [
    ['link', { href: '/gopi/favicon.png', rel: 'icon' }],
  ],
  themeConfig: {
    nav: [
      { text: 'Домой', link: '/' },
      { text: 'Примеры работы с API', link: '/api-examples' }
    ],

    sidebar: [
      {
        items: [
          { text: 'Примеры работы с API', link: '/api-examples' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/chirizxc/gopi' }
    ]
  },
})
