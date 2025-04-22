import { createSSRApp } from 'vue'

export function createApp(reqId) {
if(typeof window !== 'undefined' ){
    reqId=document.getElementById("reqId").innerText
    console.log(reqId)
}
return  createSSRApp({
        data: () => ({ count: 1 }),
        template: `<div>
                    <p>Hello from Vue NodeJS microservice! <br/>This is rendered on the server with nodejs and using gRPC!<br/>Req ID= <span id="reqId">${reqId}</span></p>
                        <button @click="count++">{{ count }}</button>
                    </div>`
    })

}