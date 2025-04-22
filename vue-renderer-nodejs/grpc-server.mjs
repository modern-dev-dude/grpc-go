import grpc from '@grpc/grpc-js';
import protoLoader from '@grpc/proto-loader';
import fs from 'fs';
import path from 'path';
import { createSSRApp } from 'vue'
import { renderToString } from 'vue/server-renderer'
import {createApp} from "../packages/render-client/static/vue-app.mjs";

const CONFIG = JSON.parse(fs.readFileSync(path.join(import.meta.dirname,"../config.json"), "utf-8"))
const PROTO_PATH = path.join(import.meta.dirname,"../proto/go-renderer.proto")
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition)
const renderer_proto = protoDescriptor.go_renderer;

function renderPage(call, callback) {
    const reqId = call.request.metadata.reqId

    const app = createApp(reqId)
    renderToString(app).then((html) => {
        console.log(html)
        callback(null, {
            markup: html
        });
    })
}

function main() {
    const server = new grpc.Server();
    server.addService(renderer_proto.RenderingEngine.service, {renderPage});
    server.bindAsync(`:${CONFIG.vueTeam.port}`, grpc.ServerCredentials.createInsecure(), (err, port) => {
        if (err != null) {
            return console.error(err);
        }
        console.log(`react node gRPC listening on ${port}`)
    });
}

main();