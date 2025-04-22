// client.js
import { createApp } from './vue-app.mjs'

export function hydrateVue(){
    createApp().mount('#app')

    console.log("vue app hydrated")
}
