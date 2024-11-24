import {h} from "vue"
import type {Theme} from "vitepress"
import DefaultTheme from "vitepress/theme"

import "./vars.css"
import "./rainbow.css"

export default {
  extends: DefaultTheme,
  Layout: () => {return h(DefaultTheme.Layout, null, {})
  },
  enhanceApp({ app, router, siteData }) {
  }
} satisfies Theme

function updateHomePageStyle(value: boolean) {
  if (value) {
    if (homePageStyle)
      return

    homePageStyle = document.createElement("style")
    homePageStyle.innerHTML = `
    :root {
      animation: rainbow 12s linear infinite;
    }`
    document.body.appendChild(homePageStyle)
  }
  else {
    if (!homePageStyle)
      return

    homePageStyle.remove()
    homePageStyle = undefined
  }
}
