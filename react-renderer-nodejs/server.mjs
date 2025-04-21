import express from "express"
import React from "react"
import { renderToStaticMarkup } from 'react-dom/server';

const app = express()
const PORT = 3000

export default function startServer(){
    app.get("/render",renderReactHandler)

    app.listen(PORT, () => {

        console.log("express listening on port " + PORT)
    })

}

function renderReactHandler(req, res){
    const reqId = req.get("X-Request-ID")
    const component = React.createElement("div", {
        className:'rounded-xl p-4 border-2 justify-items-center border-black-50 shadow-xl',
        dangerouslySetInnerHTML: {__html: `hello from node microservice!\n Req ID= ${reqId}`}
    })
    const html = renderToStaticMarkup(component)

    res.json({
        markup: html
    })
}