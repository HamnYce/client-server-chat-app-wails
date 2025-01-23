export namespace global {
	
	export class Message {
	    Sender: string;
	    Message: string;
	    Time: string;
	
	    static createFrom(source: any = {}) {
	        return new Message(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Sender = source["Sender"];
	        this.Message = source["Message"];
	        this.Time = source["Time"];
	    }
	}

}

