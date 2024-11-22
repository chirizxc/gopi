import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "GoPI",
  description: "REST API для работы с GIF.",
  base: "/gopi/",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
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
  }
})
