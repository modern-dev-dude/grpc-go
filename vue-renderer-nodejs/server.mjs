import express from "express"
import cors from 'cors';
import { renderToString } from 'vue/server-renderer'
import {createApp} from "../packages/render-client/static/vue-app.mjs";
const app = express()
const PORT = 3000
app.use(cors());

export default function startServer(){
    app.get("/render",renderReactHandler)

    app.listen(PORT, () => {
        console.log("express listening on port " + PORT)
    })

}

function renderReactHandler(req, res){
    const reqId = req.get("X-Request-ID")


    const app = createApp(reqId)
    renderToString(app).then((html) => {
        res.send(`
            <!DOCTYPE html>
            <html>
              <head>
                <title>Vue SSR Example</title>
                <script type="importmap">
                    { 
                        "imports": {
                            "@mdd/hydrateVue.mjs": "http://localhost:8080/vue-client.mjs",
                            "vue": "https://unpkg.com/vue@3/dist/vue.esm-browser.js"
                        }
                    }
                </script>
              </head>
              <body>
                <div id="app">${html}</div>
                <script>
                import('@mdd/hydrateVue.mjs').then(({hydrateVue}) =>{
                    hydrateVue()
                }) 
                </script>
              </body>
            </html>`
        )
    })
}