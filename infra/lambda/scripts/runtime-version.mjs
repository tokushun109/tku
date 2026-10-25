export function resolveNodeRuntime(toolVersions, infrastructureConfig) {
  const nodeVersion = toolVersions.match(/^nodejs\s+(\d+)\.\d+\.\d+(?:\s|$)/m)
  if (!nodeVersion) {
    throw new Error('Node.js version is missing from .tool-versions')
  }

  const runtimes = [...infrastructureConfig.matchAll(/^\s*runtime\s*=\s*"nodejs(\d+)\.x"/gm)]
  if (runtimes.length === 0) {
    throw new Error('Lambda Node.js runtime is missing from aws-runtime.tf')
  }
  if (runtimes.some((runtime) => runtime[1] !== nodeVersion[1])) {
    throw new Error('Lambda runtime must match the Node.js major version in .tool-versions')
  }

  return `node${nodeVersion[1]}`
}
