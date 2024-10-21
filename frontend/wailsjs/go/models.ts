export namespace backend {
	
	export class Boxdims {
	    miktex: string;
	    texlive: string;
	
	    static createFrom(source: any = {}) {
	        return new Boxdims(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.miktex = source["miktex"];
	        this.texlive = source["texlive"];
	    }
	}
	export class Conf {
	    workspace: string;
	    pdflatexPaths: {[key: string]: string};
	    "last-distro": string;
	
	    static createFrom(source: any = {}) {
	        return new Conf(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.workspace = source["workspace"];
	        this.pdflatexPaths = source["pdflatexPaths"];
	        this["last-distro"] = source["last-distro"];
	    }
	}

}

