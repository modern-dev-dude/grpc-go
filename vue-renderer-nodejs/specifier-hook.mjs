import fs from 'fs'
import path from 'path'

const { vueTeam } = JSON.parse(fs.readFileSync(path.join(import.meta.dirname,"../node-config.json"), "utf-8"))

export async function resolve(specifier, ctx, nextResolve){
    console.log(specifier)
    if(Object.hasOwn(vueTeam.specifiers, specifier)){
        const newImportPath = path.join(import.meta.url,vueTeam.specifiers[specifier] )
        return nextResolve(newImportPath, ctx)
    }
    return nextResolve(specifier, ctx)
}