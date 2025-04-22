import {register} from "node:module"
import { pathToFileURL } from "node:url"
// register SSR import specifiers
register('./specifier-hook.mjs', pathToFileURL("./"))
// run the app
import ('./grpc-server.mjs')