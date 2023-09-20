import { buildSync } from 'esbuild'
import { Construct } from 'constructs'
import { TerraformAsset, AssetType } from 'cdktf'
import * as path from 'path'

export interface compileForLambdaFunctionProps {
    path: string
    external?: string[]
}

const bundle = (workingDirectory: string, external: string[] = []) => {
    buildSync({
        entryPoints: ['src/index.ts'],
        platform: 'node',
        target: 'es2018',
        bundle: true,
        format: 'cjs',
        sourcemap: 'external',
        outdir: 'dist',
        absWorkingDir: workingDirectory,
        external,
    })

    return path.join(workingDirectory, 'dist')
}

export class compileForLambdaFunction extends Construct {
    public readonly asset: TerraformAsset

    constructor(
        scope: Construct,
        name: string,
        props: compileForLambdaFunctionProps
    ) {
        super(scope, name)

        const workingDirectory = path.resolve(props.path)
        const distPath = bundle(workingDirectory, props.external)

        this.asset = new TerraformAsset(this, 'lambda-asset', {
            path: distPath,
            type: AssetType.ARCHIVE, // if left empty it infers directory and file
        })
    }
}
