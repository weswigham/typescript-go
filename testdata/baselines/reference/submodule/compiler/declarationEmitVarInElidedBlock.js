//// [tests/cases/compiler/declarationEmitVarInElidedBlock.ts] ////

//// [declarationEmitVarInElidedBlock.ts]
{
    var a = "";
}
export let b: typeof a;

//// [declarationEmitVarInElidedBlock.js]
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.b = void 0;
{
    var a = "";
}


//// [declarationEmitVarInElidedBlock.d.ts]
export declare let b: typeof a;
