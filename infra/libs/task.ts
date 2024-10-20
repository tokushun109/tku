import * as fs from 'fs'
interface TaskDefinition {
    containerDefinitions: string
}

export const importTaskDefinition = (path: string): string => {
    const { containerDefinitions }: TaskDefinition = JSON.parse(fs.readFileSync(path, 'utf-8'))
    return JSON.stringify(containerDefinitions)
}
