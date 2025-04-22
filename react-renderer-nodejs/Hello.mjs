import React from "react";

export default function HelloFromReact({reqId}){
    return React.createElement("div", {
        className:'rounded-xl p-4 border-2 justify-items-center border-black-50 shadow-xl',
        dangerouslySetInnerHTML: {__html: `Hello from React NodeJS microservice! <br/>This is rendered on the server with nodejs and using gRPC!<br/>Req ID= ${reqId}`}
    })
}