//// [tests/cases/compiler/declarationEmitCastReusesTypeNode3.ts] ////

//// [declarationEmitCastReusesTypeNode3.ts]
type P = { } & { name: string }

export let vLet = <P>null!
export const vConst = <P>null!

export function fn(p = <P>null!) {}

export function fnWithRequiredDefaultParam(p = <P>null!, req: number) {}

export class C {
    field = <P>null!
    optField? = <P>null!
    readonly roFiled = <P>null!;
    method(p = <P>null!) {}
    methodWithRequiredDefault(p = <P>null!, req: number) {}

    constructor(public ctorField = <P>null!) {}

    get x() { return <P>null! }
    set x(v) { }
}

export default <P>null!;

// allows `undefined` on the input side, thanks to the initializer
export function fnWithPartialAnnotationOnDefaultparam(x: P = <P>null!, b: number) {}

//// [declarationEmitCastReusesTypeNode3.js]
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.C = exports.vConst = exports.vLet = void 0;
exports.fn = fn;
exports.fnWithRequiredDefaultParam = fnWithRequiredDefaultParam;
exports.fnWithPartialAnnotationOnDefaultparam = fnWithPartialAnnotationOnDefaultparam;
exports.vLet = null;
exports.vConst = null;
function fn(p = null) { }
function fnWithRequiredDefaultParam(p = null, req) { }
class C {
    ctorField;
    field = null;
    optField = null;
    roFiled = null;
    method(p = null) { }
    methodWithRequiredDefault(p = null, req) { }
    constructor(ctorField = null) {
        this.ctorField = ctorField;
    }
    get x() { return null; }
    set x(v) { }
}
exports.C = C;
exports.default = null;
// allows `undefined` on the input side, thanks to the initializer
function fnWithPartialAnnotationOnDefaultparam(x = null, b) { }


//// [declarationEmitCastReusesTypeNode3.d.ts]
type P = {} & {
    name: string;
};
export declare let vLet: {
    name: string;
};
export declare const vConst: {
    name: string;
};
export declare function fn(p?: {
    name: string;
}): void;
export declare function fnWithRequiredDefaultParam(p: {
    name: string;
}, req: number): void;
export declare class C {
    ctorField: {
        name: string;
    } | undefined;
    field: {
        name: string;
    };
    optField?: {
        name: string;
    } | undefined;
    readonly roFiled: {
        name: string;
    };
    method(p?: {
        name: string;
    }): void;
    methodWithRequiredDefault(p: {
        name: string;
    }, req: number): void;
    constructor(ctorField?: {
        name: string;
    } | undefined);
    get x(): {
        name: string;
    };
    set x(v: {
        name: string;
    });
}
declare const _default: {
    name: string;
};
export default _default;
// allows `undefined` on the input side, thanks to the initializer
export declare function fnWithPartialAnnotationOnDefaultparam(x: P, b: number): void;
