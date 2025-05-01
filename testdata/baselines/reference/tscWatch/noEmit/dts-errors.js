
currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::-w
//// [/home/src/workspaces/project/a.ts] new file
const a = class { private p = 10; };
//// [/home/src/workspaces/project/tsconfig.json] new file
{
	"compilerOptions": {
            "noEmit": true,
            "outFile": "../outFile.js",
            "declaration": true
	}
}



CompilerOptions::{
    "watch": true
}


Output::
a.ts(1,7): error TS4094: Property 'p' of exported anonymous class type may not be private or protected.


Found 1 error in a.ts[90m:1[0m

//// [/home/src/workspaces/project/a.ts] no change
//// [/home/src/workspaces/project/tsconfig.json] no change



Edit:: fix syntax error

Output::
//// [/home/src/workspaces/project/a.ts] modified. new content:
const a = "hello";
//// [/home/src/workspaces/project/tsconfig.json] no change



Edit:: emit after fixing error

Output::
//// [/home/src/workspaces/project/a.d.ts] new file
declare const a = "hello";

//// [/home/src/workspaces/project/a.js] new file
const a = "hello";

//// [/home/src/workspaces/project/a.ts] no change
//// [/home/src/workspaces/project/tsconfig.json] modified. new content:
{
	"compilerOptions": {
            "outFile": "../outFile.js",
            "declaration": true
	}
}



Edit:: no emit run after fixing error

Output::
//// [/home/src/workspaces/project/a.d.ts] no change
//// [/home/src/workspaces/project/a.js] no change
//// [/home/src/workspaces/project/a.ts] no change
//// [/home/src/workspaces/project/tsconfig.json] modified. new content:
{
	"compilerOptions": {
            "noEmit": true,
            "outFile": "../outFile.js",
            "declaration": true
	}
}



Edit:: introduce error

Output::
a.ts(1,7): error TS4094: Property 'p' of exported anonymous class type may not be private or protected.


Found 1 error in a.ts[90m:1[0m

//// [/home/src/workspaces/project/a.d.ts] no change
//// [/home/src/workspaces/project/a.js] no change
//// [/home/src/workspaces/project/a.ts] modified. new content:
const a = class { private p = 10; };
//// [/home/src/workspaces/project/tsconfig.json] no change



Edit:: emit when error

Output::
a.ts(1,7): error TS4094: Property 'p' of exported anonymous class type may not be private or protected.


Found 1 error in a.ts[90m:1[0m

//// [/home/src/workspaces/project/a.d.ts] modified. new content:
declare const a: {
    new (): {
        p: number;
    };
};

//// [/home/src/workspaces/project/a.js] modified. new content:
const a = class {
    p = 10;
};

//// [/home/src/workspaces/project/a.ts] no change
//// [/home/src/workspaces/project/tsconfig.json] modified. new content:
{
	"compilerOptions": {
            "outFile": "../outFile.js",
            "declaration": true
	}
}



Edit:: no emit run when error

Output::
a.ts(1,7): error TS4094: Property 'p' of exported anonymous class type may not be private or protected.


Found 1 error in a.ts[90m:1[0m

//// [/home/src/workspaces/project/a.d.ts] no change
//// [/home/src/workspaces/project/a.js] no change
//// [/home/src/workspaces/project/a.ts] no change
//// [/home/src/workspaces/project/tsconfig.json] modified. new content:
{
	"compilerOptions": {
            "noEmit": true,
            "outFile": "../outFile.js",
            "declaration": true
	}
}

