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
    const component = React.createElement("div", {
        dangerouslySetInnerHTML: {__html: "hello from node microservice!"}
    })
    const html = renderToStaticMarkup(component)

    res.json({
        markup: html
    })
}