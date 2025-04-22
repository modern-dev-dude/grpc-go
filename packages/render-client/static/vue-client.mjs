// client.js
import { createApp } from './vue-app.mjs'

export function hydrateVue(){
    createApp().mount('#vue-mfe')

    console.log("vue app hydrated")
}
