
import grpc from '@grpc/grpc-js';
import protoLoader from '@grpc/proto-loader';
import path from 'path';
import {renderToStaticMarkup} from "react-dom/server";
import HelloFromReact from "./Hello.mjs";


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


    const html = renderToStaticMarkup(HelloFromReact({reqId}))
    callback(null, {
        markup: html
    });
}

/**
 * Starts an RPC server that receives requests for the Greeter service at the
 * sample server port
 */
function main() {
    const server = new grpc.Server();
    server.addService(renderer_proto.RenderingEngine.service, {renderPage});
    server.bindAsync('0.0.0.0:9002', grpc.ServerCredentials.createInsecure(), (err, port) => {
        if (err != null) {
            return console.error(err);
        }
        console.log(`react node gRPC listening on ${port}`)
    });
}

main();