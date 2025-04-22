import { createSSRApp } from 'vue'

export function createApp(reqId) {
    return  createSSRApp({
        data: () => ({ count: 1, id: reqId }),
        template: `<p>Hello from Vue NodeJS microservice!<br/>This is rendered on the server with nodejs and using gRPC!<br/>Req ID= <span id="reqId">{{id}}</span></p><button class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600" @click="count++">Count clicked: {{ count }}</button>`,
        mounted() {
            console.log('Component mounted - Client rendered content:', {
                element: document.getElementById('vue-mfe').innerHTML,
                reqId: this.id,
                count: this.count
            })
        },
        beforeMount() {
            console.log('Before mount - Initial DOM:', {
                element: document.getElementById('vue-mfe')?.innerHTML,
                reqId: this.id
            })
        }
})

}