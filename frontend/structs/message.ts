interface iMessage {
    id: number;
    text: string;
    sender: string;
    timestamp: number | string;
    user?: string
}