export class Mail {
    public sender: string;
    public receiver: string;
    public subject: string;
    public body: string;
    public time: number;

    constructor(sender: string, receiver: string, subject: string, body: string, time: number){
        this.sender = sender;
        this.receiver = receiver;
        this.subject = subject;
        this.body = body;
        this.time = time;
    }
}
