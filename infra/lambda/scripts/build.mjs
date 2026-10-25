import { execFileSync } from 'node:child_process'
import { mkdirSync, readFileSync, rmSync } from 'node:fs'
import { dirname, join } from 'node:path'
import { fileURLToPath } from 'node:url'
import { build } from 'esbuild'
import { resolveNodeRuntime } from './runtime-version.mjs'

const rootDir = dirname(dirname(fileURLToPath(import.meta.url)))
const distDir = join(rootDir, 'dist')
const handlers = ['warmup', 'health-check']
const nodeTarget = resolveNodeRuntime(
  readFileSync(join(rootDir, '../../.tool-versions'), 'utf8'),
  readFileSync(join(rootDir, '../opentofu/aws-runtime.tf'), 'utf8'),
)

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
    target: nodeTarget,
  })

  execFileSync('zip', ['-j', join(outputDir, 'archive.zip'), outputFile], { stdio: 'inherit' })
}
