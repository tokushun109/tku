import assert from 'node:assert/strict'
import { test } from 'node:test'
import { resolveNodeRuntime } from './runtime-version.mjs'

test('matching versions select the Lambda build target', () => {
  assert.equal(resolveNodeRuntime('nodejs 24.22.1 # LTS\n', '  runtime = "nodejs24.x"\n'), 'node24')
  assert.equal(resolveNodeRuntime('nodejs 22.23.4\n', '  runtime = "nodejs22.x"\n'), 'node22')
})

test('a Node.js upgrade without a Lambda runtime upgrade is rejected', () => {
  assert.throws(() => resolveNodeRuntime('nodejs 24.22.1\n', 'runtime = "nodejs22.x"'), /must match/)
})

test('all Lambda runtimes must match', () => {
  assert.throws(() => resolveNodeRuntime('nodejs 24.22.1\n', 'runtime = "nodejs24.x"\nruntime = "nodejs22.x"'), /must match/)
})

test('missing or commented-out settings are rejected', () => {
  assert.throws(() => resolveNodeRuntime('# nodejs 24.22.1', 'runtime = "nodejs24.x"'), /missing/)
  assert.throws(() => resolveNodeRuntime('nodejs 24.22.1', '# runtime = "nodejs24.x"'), /missing/)
})
