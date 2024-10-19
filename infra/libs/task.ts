import * as fs from 'fs'
import { TaskDefinition } from '../resources/ecs/types'

export const importTaskDefinition = (path: string): string => {
    const { containerDefinitions }: TaskDefinition = JSON.parse(fs.readFileSync(path, 'utf-8'))
    return JSON.stringify(containerDefinitions)
}
