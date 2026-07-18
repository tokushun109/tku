import { execFileSync } from 'node:child_process'
import { mkdirSync, rmSync } from 'node:fs'
import { dirname, join } from 'node:path'
import { fileURLToPath } from 'node:url'
import { build } from 'esbuild'

const rootDir = dirname(dirname(fileURLToPath(import.meta.url)))
const distDir = join(rootDir, 'dist')
const handlers = ['warmup', 'health-check']

rmSync(distDir, { force: true, recursive: true })
mkdirSync(distDir, { recursive: true })

for (const handler of handlers) {
  const outputDir = join(distDir, handler)
  const outputFile = join(outputDir, 'index.js')

  mkdirSync(outputDir, { recursive: true })
  await build({
    bundle: true,
    entryPoints: [join(rootDir, 'handlers', handler, 'src', 'index.ts')],
    format: 'cjs',
    outfile: outputFile,
    platform: 'node',
    target: 'node22',
  })

  execFileSync('zip', ['-j', join(outputDir, 'archive.zip'), outputFile], { stdio: 'inherit' })
}
